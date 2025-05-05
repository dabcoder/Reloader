package util

import (
	"strings"

	"github.com/stakater/Reloader/internal/pkg/constants"
	"github.com/stakater/Reloader/internal/pkg/options"
	v1 "k8s.io/api/core/v1"
)

// Config contains rolling upgrade configuration parameters
type Config struct {
	Namespace           string
	ResourceName        string
	ResourceAnnotations map[string]string
	Annotation          string
	TypedAutoAnnotation string
	SHAValue            string
	Type                string
}

// GetConfigmapConfig provides utility config for configmap
func GetConfigmapConfig(configmap *v1.ConfigMap) Config {
	return Config{
		Namespace:           configmap.Namespace,
		ResourceName:        configmap.Name,
		ResourceAnnotations: configmap.Annotations,
		Annotation:          options.ConfigmapUpdateOnChangeAnnotation,
		TypedAutoAnnotation: options.ConfigmapReloaderAutoAnnotation,
		SHAValue:            GetSHAfromConfigmap(configmap),
		Type:                constants.ConfigmapEnvVarPostfix,
	}
}

// GetSecretConfig provides utility config for secret
func GetSecretConfig(secret *v1.Secret) Config {
	return Config{
		Namespace:           secret.Namespace,
		ResourceName:        secret.Name,
		ResourceAnnotations: secret.Annotations,
		Annotation:          options.SecretUpdateOnChangeAnnotation,
		TypedAutoAnnotation: options.SecretReloaderAutoAnnotation,
		SHAValue:            GetSHAfromSecret(secret.Data),
		Type:                constants.SecretEnvVarPostfix,
	}
}

// ParseResourceName parses a resource reference that can be in the format "name" or "namespace/name"
// and returns the namespace and name separately
func ParseResourceName(resourceRef string) (namespace, name string) {
	parts := strings.SplitN(resourceRef, "/", 2)
	if len(parts) == 2 {
		// Format is "namespace/name"
		return parts[0], parts[1]
	}
	// Format is just "name" (no namespace specified)
	return "", resourceRef
}
