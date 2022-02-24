/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"os"
	"strconv"
	"time"

	tutils "github.com/WASdev/websphere-traditional-operator/utils"
	oputils "github.com/application-stacks/runtime-component-operator/utils"
	"github.com/go-logr/logr"

	webspheretraditionalv1alpha1 "github.com/WASdev/websphere-traditional-operator/api/v1alpha1"
	"github.com/WASdev/websphere-traditional-operator/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ReconcileWebsphereTraditionalTrace reconciles a WebsphereTraditionalTrace object
type ReconcileWebsphereTraditionalTrace struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	Client     client.Client
	Scheme     *runtime.Scheme
	Recorder   record.EventRecorder
	RestConfig *rest.Config
	Log        logr.Logger
}

const traceFinalizer = "finalizer.webspheretraditionalTraces.apps.webspheretraditional.io.ibm"
const traceConfigFile = "/config/configDropins/overrides/add_trace.xml"
const serviceabilityDir = "/serviceability"

// +kubebuilder:rbac:groups=apps.webspheretraditional.io.ibm,resources=webspheretraditionaltraces;webspheretraditionaltraces/status;webspheretraditionaltraces/finalizers,verbs=*,namespace=websphere-traditional-operator
// +kubebuilder:rbac:groups=core,resources=pods;pods/exec,verbs=*,namespace=websphere-traditional-operator

// Reconcile reads that state of the cluster for a WebsphereTraditionalTrace object and makes changes based on the state read
// and what is in the WebsphereTraditionalTrace.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.

func (r *ReconcileWebsphereTraditionalTrace) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	reqLogger := r.Log.WithValues("Namespace", request.Namespace, "Name", request.Name)
	reqLogger.Info("Reconciling WebsphereTraditionalTrace")

	// Fetch the WebsphereTraditionalTrace instance
	instance := &webspheretraditionalv1alpha1.WebsphereTraditionalTrace{}
	err := r.Client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("Not found. Return and don't requeue")
			return reconcile.Result{}, nil
		}
		reqLogger.Info("Error reading the object - requeue the request.")
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	instance.Initialize()

	//Pod is expected to be from the same namespace as the CR instance
	podNamespace := instance.Namespace
	podName := instance.Spec.PodName

	prevPodName := instance.GetStatus().GetOperatedResource().GetOperatedResourceName()
	prevTraceEnabled := instance.GetStatus().GetCondition(webspheretraditionalv1alpha1.OperationStatusConditionTypeEnabled).Status
	podChanged := prevPodName != podName

	// Check if the WebsphereTraditionalTrace instance is marked to be deleted, which is
	// indicated by the deletion timestamp being set.
	isInstanceMarkedToBeDeleted := instance.GetDeletionTimestamp() != nil
	if isInstanceMarkedToBeDeleted {
		if tutils.Contains(instance.GetFinalizers(), traceFinalizer) {
			// Run finalization logic for traceFinalizer. If the finalization logic fails, don't remove the
			// finalizer so that we can retry during the next reconciliation.
			if err := r.finalizeWebsphereTraditionalTrace(reqLogger, instance, prevTraceEnabled, prevPodName, podNamespace); err != nil {
				return reconcile.Result{}, err
			}

			// Remove traceFinalizer. Once all finalizers have been removed, the object will be deleted.
			instance.SetFinalizers(tutils.Remove(instance.GetFinalizers(), traceFinalizer))
			err := r.Client.Update(context.TODO(), instance)
			if err != nil {
				return reconcile.Result{}, err
			}
		}
		return reconcile.Result{}, nil
	}

	// Add finalizer for this CR
	if !tutils.Contains(instance.GetFinalizers(), traceFinalizer) {
		if err := r.addFinalizer(reqLogger, instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	//If pod name changed, then stop tracing on previous pod (if trace was enabled on it)
	if podChanged && (prevTraceEnabled == corev1.ConditionTrue) {
		r.disableTraceOnPrevPod(reqLogger, prevPodName, podNamespace)
	}

	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: podName, Namespace: podNamespace}, &corev1.Pod{})
	if err != nil && errors.IsNotFound(err) {
		//Pod is not found. Return and don't requeue
		reqLogger.Error(err, "Pod "+podName+" was not found in namespace "+podNamespace)
		return r.UpdateStatus(err, webspheretraditionalv1alpha1.OperationStatusConditionTypeEnabled, *instance, corev1.ConditionFalse, podName, podChanged)
	}

	if instance.Spec.Disable != nil && *instance.Spec.Disable {
		//Disable trace if trace was previously enabled on the same pod
		if !podChanged && prevTraceEnabled == corev1.ConditionTrue {
			_, err = utils.ExecuteCommandInContainer(r.RestConfig, podName, podNamespace, "app", []string{"/bin/sh", "-c", "rm -f " + traceConfigFile})
			if err != nil {
				reqLogger.Error(err, "Encountered error while disabling trace for pod "+podName+" in namespace "+podNamespace)
				return r.UpdateStatus(err, webspheretraditionalv1alpha1.OperationStatusConditionTypeEnabled, *instance, corev1.ConditionTrue, podName, podChanged)
			}
			reqLogger.Info("Disabled trace for pod " + podName + " in namespace " + podNamespace)
		}
		r.UpdateStatus(nil, webspheretraditionalv1alpha1.OperationStatusConditionTypeEnabled, *instance, corev1.ConditionFalse, podName, podChanged)
	} else {
		PROFILE_NAME := "AppSrv01"
		NODE_NAME := "DefaultNode01"
		SERVER_NAME := "server1"
		maxFileSize := 20
		maxFiles := 5
		traceOutputDir := serviceabilityDir + "/" + podNamespace + "/" + podName + "/" + "trace.log"
		if instance.Spec.MaxFileSize != nil {
			maxFileSize = (int(*instance.Spec.MaxFileSize))
		}
		if instance.Spec.MaxFiles != nil {
			maxFiles = (int(*instance.Spec.MaxFiles))
		}
		traceConfig := "wsadmin.sh -user" + "wsadmin" + "-password" + "PASSWORD" + "-c" + "\"" + "set server [" + "\\" + "$AdminConfig getid /Cell:" + PROFILE_NAME + "/Node:" + NODE_NAME + "/Server:" + SERVER_NAME + "/]" + "\"" + "-c" + "\"" + "set ts [" + "\\" + "$AdminControl completeObjectName WebSphere:type=TraceService,process=" + SERVER_NAME + ",*]" + "\"" + "-c" + "'" + "$AdminControl invoke $ts setTraceOutputToFile {" + traceOutputDir + "\"" + strconv.Itoa(maxFileSize) + "\"" + "\"" + strconv.Itoa(maxFiles) + "\"" + "basic}" + "'" + "-c" + "'" + "$AdminControl setAttribute $ts traceSpecification instance.Spec.TraceSpecification" + "'"

		_, err = utils.ExecuteCommandInContainer(r.RestConfig, podName, podNamespace, "app", []string{"/bin/sh", "-c", "mkdir -p " + traceOutputDir + " && echo '" + traceConfig + "' > " + traceConfigFile})
		if err != nil {
			reqLogger.Error(err, "Encountered error while setting up trace for pod "+podName+" in namespace "+podNamespace)
			return r.UpdateStatus(err, webspheretraditionalv1alpha1.OperationStatusConditionTypeEnabled, *instance, corev1.ConditionFalse, podName, podChanged)
		}

		if podChanged || prevTraceEnabled == corev1.ConditionFalse {
			reqLogger.Info("Enabled trace for pod " + podName + " in namespace " + podNamespace)
		} else {
			reqLogger.Info("Updated trace for pod " + podName + " in namespace " + podNamespace)
		}
		r.UpdateStatus(nil, webspheretraditionalv1alpha1.OperationStatusConditionTypeEnabled, *instance, corev1.ConditionTrue, podName, podChanged)
	}

	return reconcile.Result{}, nil
}

// UpdateStatus updates the status
func (r *ReconcileWebsphereTraditionalTrace) UpdateStatus(issue error, conditionType webspheretraditionalv1alpha1.OperationStatusConditionType, instance webspheretraditionalv1alpha1.WebsphereTraditionalTrace, newStatus corev1.ConditionStatus, podName string, podChanged bool) (reconcile.Result, error) {
	s := instance.GetStatus()

	s.SetOperatedResource(webspheretraditionalv1alpha1.OperatedResource{ResourceName: podName, ResourceType: "pod"})

	oldCondition := s.GetCondition(conditionType)
	// Keep the old `LastTransitionTime` when pod and status have not changed
	nowTime := metav1.Now()
	transitionTime := oldCondition.GetLastTransitionTime()
	if podChanged || oldCondition.GetStatus() != newStatus {
		transitionTime = &nowTime
	}

	statusCondition := s.NewCondition()
	statusCondition.SetLastTransitionTime(transitionTime)
	statusCondition.SetLastUpdateTime(nowTime)

	if issue != nil {
		statusCondition.SetReason("Error")
		statusCondition.SetMessage(issue.Error())
		r.Recorder.Event(&instance, "Warning", "ProcessingError", issue.Error())
	} else {
		statusCondition.SetReason("")
		statusCondition.SetMessage("")
	}

	statusCondition.SetStatus(newStatus)
	statusCondition.SetType(conditionType)

	s.SetCondition(statusCondition)

	err := r.Client.Status().Update(context.Background(), &instance)
	if err != nil {
		r.Log.Error(err, "Unable to update status")
		return reconcile.Result{
			RequeueAfter: time.Second,
			Requeue:      true,
		}, nil
	}

	return reconcile.Result{Requeue: false}, nil
}

func (r *ReconcileWebsphereTraditionalTrace) disableTraceOnPrevPod(reqLogger logr.Logger, prevPodName string, podNamespace string) {
	err := r.Client.Get(context.TODO(), types.NamespacedName{Name: prevPodName, Namespace: podNamespace}, &corev1.Pod{})
	if err != nil && errors.IsNotFound(err) {
		//Previous Pod is not found. No-op
		reqLogger.Info("Previous pod " + prevPodName + " was not found in namespace " + podNamespace)
	} else {
		//Stop tracing on previous Pod
		_, err = utils.ExecuteCommandInContainer(r.RestConfig, prevPodName, podNamespace, "app", []string{"/bin/sh", "-c", "rm -f " + traceConfigFile})
		if err == nil {
			reqLogger.Info("Disabled trace on previous pod " + prevPodName + " in namespace " + podNamespace)
		} else {
			reqLogger.Error(err, "Encountered error while disabling trace on previous pod "+prevPodName+" in namespace "+podNamespace)
		}
	}
}

func (r *ReconcileWebsphereTraditionalTrace) finalizeWebsphereTraditionalTrace(reqLogger logr.Logger, olt *webspheretraditionalv1alpha1.WebsphereTraditionalTrace, prevTraceEnabled corev1.ConditionStatus, prevPodName string, podNamespace string) error {
	if prevTraceEnabled == corev1.ConditionTrue {
		r.disableTraceOnPrevPod(reqLogger, prevPodName, podNamespace)
	}
	return nil
}

func (r *ReconcileWebsphereTraditionalTrace) addFinalizer(reqLogger logr.Logger, olt *webspheretraditionalv1alpha1.WebsphereTraditionalTrace) error {
	reqLogger.Info("Adding Finalizer for WebsphereTraditionalTrace")
	olt.SetFinalizers(append(olt.GetFinalizers(), traceFinalizer))

	// Update CR
	err := r.Client.Update(context.TODO(), olt)
	if err != nil {
		reqLogger.Error(err, "Failed to update WebsphereTraditionalTrace with finalizer")
		return err
	}

	return nil
}

func (r *ReconcileWebsphereTraditionalTrace) SetupWithManager(mgr ctrl.Manager) error {

	watchNamespaces, err := oputils.GetWatchNamespaces()
	if err != nil {
		r.Log.Error(err, "Failed to get watch namespace")
		os.Exit(1)
	}

	watchNamespacesMap := make(map[string]bool)
	for _, ns := range watchNamespaces {
		watchNamespacesMap[ns] = true
	}
	isClusterWide := len(watchNamespacesMap) == 1 && watchNamespacesMap[""]

	r.Log.V(1).Info("Adding a new controller", "watchNamespaces", watchNamespaces, "isClusterWide", isClusterWide)

	pred := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			// Ignore updates to CR status in which case metadata.Generation does not change
			return e.ObjectOld.GetGeneration() != e.ObjectNew.GetGeneration() && (isClusterWide || watchNamespacesMap[e.ObjectOld.GetNamespace()])
		},
		CreateFunc: func(e event.CreateEvent) bool {
			return isClusterWide || watchNamespacesMap[e.Object.GetNamespace()]
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			return isClusterWide || watchNamespacesMap[e.Object.GetNamespace()]
		},
		GenericFunc: func(e event.GenericEvent) bool {
			return isClusterWide || watchNamespacesMap[e.Object.GetNamespace()]
		},
	}
	return ctrl.NewControllerManagedBy(mgr).For(&webspheretraditionalv1alpha1.WebsphereTraditionalTrace{}, builder.WithPredicates(pred)).Complete(r)
}
