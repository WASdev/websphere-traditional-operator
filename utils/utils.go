package utils

import (
	//"bytes"
	//	"context"
	//	"fmt"
	//	"sort"
	"strconv"
	"strings"

	webspheretraditionalv1alpha1 "github.com/WASdev/websphere-traditional-operator/api/v1alpha1"
	//	rcoutils "github.com/application-stacks/runtime-component-operator/utils"
	//	routev1 "github.com/openshift/api/route/v1"
	//	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	//	"k8s.io/apimachinery/pkg/api/resource"
	//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"k8s.io/apimachinery/pkg/types"
	//	"k8s.io/client-go/kubernetes"
	//	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	//	"k8s.io/client-go/tools/remotecommand"
	"sigs.k8s.io/controller-runtime/pkg/client"
	//	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// Utility methods specific to Open Liberty and its configuration

var log = logf.Log.WithName("webspheretraditional_utils")

// Validate if the WebsphereTraditionalApplication is valid
func Validate(olapp *webspheretraditionalv1alpha1.WebsphereTraditionalApplication) (bool, error) {
	log.Error(nil, "Need to implement function")
	return true, nil
}

func requiredFieldMessage(fieldPaths ...string) string {
	return "must set the field(s): " + strings.Join(fieldPaths, ",")
}

// ExecuteCommandInContainer Execute command inside a container in a pod through API
func ExecuteCommandInContainer(config *rest.Config, podName, podNamespace, containerName string, command []string) (string, error) {
	log.Error(nil, "Need to implement function")
	return "Need to implement function", nil
}

// CustomizeLibertyEnv adds configured env variables appending configured liberty settings
func CustomizeWebsphereTraditionalEnv(pts *corev1.PodTemplateSpec, wt *webspheretraditionalv1alpha1.WebsphereTraditionalApplication, client client.Client) error {
	log.Error(nil, "Need to implement function")
	return nil
}

func addSecretResourceVersionAsEnvVar(pts *corev1.PodTemplateSpec, wt *webspheretraditionalv1alpha1.WebsphereTraditionalApplication, client client.Client, secretName string, envNamePrefix string) error {
	log.Error(nil, "Need to implement function")
	return nil
}

func CustomizeWebsphereTraditionalAnnotations(pts *corev1.PodTemplateSpec, wt *webspheretraditionalv1alpha1.WebsphereTraditionalApplication) {
	log.Error(nil, "Need to implement function")
}

func findEnvVar(name string, envList []corev1.EnvVar) (*corev1.EnvVar, bool) {
	for i, val := range envList {
		if val.Name == name {
			return &envList[i], true
		}
	}
	return nil, false
}

func CreateServiceabilityPVC(instance *webspheretraditionalv1alpha1.WebsphereTraditionalApplication) *corev1.PersistentVolumeClaim {
	log.Error(nil, "Need to implement function")
	return nil
}

func ConfigureServiceability(pts *corev1.PodTemplateSpec, la *webspheretraditionalv1alpha1.WebsphereTraditionalApplication) {
	log.Error(nil, "Need to implement function")
}

func normalizeEnvVariableName(name string) string {
	return strings.NewReplacer("-", "_", ".", "_").Replace(strings.ToUpper(name))
}

// getValue returns value for string
func getValue(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case bool:
		return strconv.FormatBool(v.(bool))
	default:
		return ""
	}
}

// createEnvVarSSO creates an environment variable for SSO
func createEnvVarSSO(loginID string, envSuffix string, value interface{}) *corev1.EnvVar {
	log.Error(nil, "Need to implement function")
	return nil
}

func writeSSOSecretIfNeeded(client client.Client, ssoSecret *corev1.Secret, ssoSecretUpdates map[string][]byte) error {
	log.Error(nil, "Need to implement function")
	return nil
}

// CustomizeEnvSSO Process the configuration for SSO login providers
func CustomizeEnvSSO(pts *corev1.PodTemplateSpec, instance *webspheretraditionalv1alpha1.WebsphereTraditionalApplication, client client.Client, isOpenShift bool) error {
	log.Error(nil, "Need to implement function")
	return nil
}

func Contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func Remove(list []string, s string) []string {
	for i, v := range list {
		if v == s {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list
}
