// Code generated by reverse-kube-resource. DO NOT EDIT.

package scmigration

import v1unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

var (
	// Unstructured "helm-broker-cleanup"
	helmBrokerCleanupUnstructuredJob = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "batch/v1",
			"kind":       "Job",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-cleanup",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-addons-ui"
	helmBrokerAddonsUiUnstructuredPodSecurityPolicy = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "policy/v1beta1",
			"kind":       "PodSecurityPolicy",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-addons-ui",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-addons-ui"
	helmBrokerAddonsUiUnstructuredServiceAccount = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-addons-ui",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-etcd-stateful-etcd-certs"
	helmBrokerEtcdStatefulEtcdCertsUnstructuredServiceAccount = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-etcd-stateful-etcd-certs",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker"
	helmBrokerUnstructuredServiceAccount = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "helm-broker",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-secret"
	helmSecretUnstructuredSecret = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      "helm-secret",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-webhook-cert"
	helmBrokerWebhookCertUnstructuredSecret = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-webhook-cert",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "addons-ui"
	addonsUiUnstructuredConfigMap = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "addons-ui",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-dashboard"
	helmBrokerDashboardUnstructuredConfigMap = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-dashboard",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-config-map"
	helmConfigMapUnstructuredConfigMap = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "helm-config-map",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "ssh-cfg"
	sshCfgUnstructuredConfigMap = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "ssh-cfg",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-etcd-stateful-etcd-certs"
	helmBrokerEtcdStatefulEtcdCertsUnstructuredClusterRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-etcd-stateful-etcd-certs",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-h3"
	helmBrokerH3UnstructuredClusterRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-h3",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-etcd-stateful-etcd-certs"
	helmBrokerEtcdStatefulEtcdCertsUnstructuredClusterRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-etcd-stateful-etcd-certs",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-h3"
	helmBrokerH3UnstructuredClusterRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-h3",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-addons-ui"
	helmBrokerAddonsUiUnstructuredRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-addons-ui",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-addons-ui"
	helmBrokerAddonsUiUnstructuredRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-addons-ui",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-addons-ui"
	helmBrokerAddonsUiUnstructuredService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-addons-ui",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-etcd-stateful"
	helmBrokerEtcdStatefulUnstructuredService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-etcd-stateful",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-etcd-stateful-client"
	helmBrokerEtcdStatefulClientUnstructuredService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-etcd-stateful-client",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-metrics"
	helmBrokerMetricsUnstructuredService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-metrics",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "addon-controller-metrics"
	addonControllerMetricsUnstructuredService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "addon-controller-metrics",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker"
	helmBrokerUnstructuredService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "helm-broker",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-webhook"
	helmBrokerWebhookUnstructuredService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-webhook",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-addons-ui"
	helmBrokerAddonsUiUnstructuredDeployment = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-addons-ui",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker"
	helmBrokerUnstructuredDeployment = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "helm-broker",
				"namespace": "kyma-system",
			},
			"spec": map[string]interface{}{
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "helm-broker",
					},
				},
			},
		},
	}

	// Unstructured "helm-broker-webhook"
	helmBrokerWebhookUnstructuredDeployment = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-webhook",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-etcd-stateful"
	helmBrokerEtcdStatefulUnstructuredStatefulSet = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "StatefulSet",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-etcd-stateful",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker"
	helmBrokerUnstructuredAuthorizationPolicy = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "security.istio.io/v1beta1",
			"kind":       "AuthorizationPolicy",
			"metadata": map[string]interface{}{
				"name":      "helm-broker",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-repos-urls"
	helmReposUrlsUnstructuredClusterAddonsConfiguration = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "addons.kyma-project.io/v1alpha1",
			"kind":       "ClusterAddonsConfiguration",
			"metadata": map[string]interface{}{
				"name":      "helm-repos-urls",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "addonsclustermicrofrontend"
	addonsclustermicrofrontendUnstructuredClusterMicroFrontend = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "ui.kyma-project.io/v1alpha1",
			"kind":       "ClusterMicroFrontend",
			"metadata": map[string]interface{}{
				"name":      "addonsclustermicrofrontend",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "addonsmicrofrontend"
	addonsmicrofrontendUnstructuredClusterMicroFrontend = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "ui.kyma-project.io/v1alpha1",
			"kind":       "ClusterMicroFrontend",
			"metadata": map[string]interface{}{
				"name":      "addonsmicrofrontend",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-addons-ui"
	helmBrokerAddonsUiUnstructuredDestinationRule = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.istio.io/v1alpha3",
			"kind":       "DestinationRule",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-addons-ui",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-etcd-stateful-client"
	helmBrokerEtcdStatefulClientUnstructuredDestinationRule = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.istio.io/v1alpha3",
			"kind":       "DestinationRule",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-etcd-stateful-client",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-mutating-webhook"
	helmBrokerMutatingWebhookUnstructuredMutatingWebhookConfiguration = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "MutatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-mutating-webhook",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker"
	helmBrokerUnstructuredPeerAuthentication = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "security.istio.io/v1beta1",
			"kind":       "PeerAuthentication",
			"metadata": map[string]interface{}{
				"name":      "helm-broker",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-etcd-stateful"
	helmBrokerEtcdStatefulUnstructuredServiceMonitor = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "monitoring.coreos.com/v1",
			"kind":       "ServiceMonitor",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-etcd-stateful",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker"
	helmBrokerUnstructuredServiceMonitor = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "monitoring.coreos.com/v1",
			"kind":       "ServiceMonitor",
			"metadata": map[string]interface{}{
				"name":      "helm-broker",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-addon-controller"
	helmBrokerAddonControllerUnstructuredServiceMonitor = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "monitoring.coreos.com/v1",
			"kind":       "ServiceMonitor",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-addon-controller",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "helm-broker-addons-ui"
	helmBrokerAddonsUiUnstructuredVirtualService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.istio.io/v1alpha3",
			"kind":       "VirtualService",
			"metadata": map[string]interface{}{
				"name":      "helm-broker-addons-ui",
				"namespace": "kyma-system",
			},
		},
	}
)