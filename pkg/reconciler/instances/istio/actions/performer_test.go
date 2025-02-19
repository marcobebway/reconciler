package actions

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/kyma-incubator/reconciler/pkg/reconciler/chart"
	workspacemocks "github.com/kyma-incubator/reconciler/pkg/reconciler/chart/mocks"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/kyma-incubator/reconciler/pkg/logger"
	clientsetmocks "github.com/kyma-incubator/reconciler/pkg/reconciler/instances/istio/clientset/mocks"
	"github.com/kyma-incubator/reconciler/pkg/reconciler/instances/istio/istioctl"
	istioctlmocks "github.com/kyma-incubator/reconciler/pkg/reconciler/instances/istio/istioctl/mocks"
	proxymocks "github.com/kyma-incubator/reconciler/pkg/reconciler/instances/istio/reset/proxy/mocks"
	"github.com/kyma-incubator/reconciler/pkg/reconciler/kubernetes/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	istioManifest = `
apiVersion: version/v1
kind: Kind1
metadata:
  namespace: namespace
  name: name
---
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: namespace
  name: name
---
apiVersion: version/v2
kind: Kind2
metadata:
  namespace: namespace
  name: name
`
	istioctlMockSimpleVersion = `no running Istio pods in "istio-system"
{
	"clientVersion": {
	"version": "1.11.2",
	"revision": "revision",
	"golang_version": "go1.16.0",
	"status": "Clean",
	"tag": "1.11.2"
	}
}`

	istioctlMockCompleteVersion = `{
		"clientVersion": {
		  "version": "1.11.1",
		  "revision": "revision",
		  "golang_version": "go1.16.7",
		  "status": "Clean",
		  "tag": "1.11.1"
		},
		"meshVersion": [
		  {
			"Component": "pilot",
			"Info": {
			  "version": "1.11.1",
			  "revision": "revision",
			  "golang_version": "",
			  "status": "Clean",
			  "tag": "1.11.1"
			}
		  }
		],
		"dataPlaneVersion": [
		  {
			"ID": "id",
			"IstioVersion": "1.11.1"
		  }
		]
	  }`
)

func Test_DefaultIstioPerformer_Install(t *testing.T) {

	kubeConfig := "kubeConfig"
	log := logger.NewLogger(false)

	t.Run("should not install when istio version could not be resolved", func(t *testing.T) {
		// given
		cmdResolver := TestCommanderResolver{err: errors.New("istioctl not found")}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Install(kubeConfig, "", "1.2.3", log)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "Istio Operator definition could not be found")
	})

	t.Run("should not install when istio operator could not be found in manifest", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmder.On("Install", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(errors.New("istioctl error"))
		cmdResolver := TestCommanderResolver{cmder: &cmder}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Install(kubeConfig, "", "1.2.3", log)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "Istio Operator definition could not be found")
		cmder.AssertNotCalled(t, "Install", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger"))
	})

	t.Run("should not install Istio when istioctl returned an error", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmder.On("Install", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(errors.New("istioctl error"))

		cmdResolver := TestCommanderResolver{cmder: &cmder}
		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Install(kubeConfig, istioManifest, "1.2.3", log)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "istioctl error")
		cmder.AssertCalled(t, "Install", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger"))
	})

	t.Run("should install Istio when istioctl command was successful", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmder.On("Install", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(nil)
		cmdResolver := TestCommanderResolver{cmder: &cmder}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Install(kubeConfig, istioManifest, "1.2.3", log)

		// then
		require.NoError(t, err)
		cmder.AssertCalled(t, "Install", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger"))
	})

}

func Test_DefaultIstioPerformer_Uninstall(t *testing.T) {
	kc := &mocks.Client{}
	kc.On("Kubeconfig").Return("kubeconfig")
	kc.On("Clientset").Return(fake.NewSimpleClientset(&corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: "istio-system"},
	}), nil)
	log := logger.NewLogger(false)

	t.Run("should not uninstall when istio version could not be resolved", func(t *testing.T) {
		// given
		cmdResolver := TestCommanderResolver{err: errors.New("istioctl not found")}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		var wrapper IstioPerformer = NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Uninstall(kc, "1.2.3", log)

		// then
		require.Error(t, err)
		require.Equal(t, "istioctl not found", err.Error())
	})

	t.Run("should not uninstall Istio when istioctl returned an error", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmder.On("Uninstall", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(errors.New("istioctl error"))
		cmdResolver := TestCommanderResolver{cmder: &cmder}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		var wrapper IstioPerformer = NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Uninstall(kc, "1.2.3", log)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "istioctl error")
		cmder.AssertCalled(t, "Uninstall", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger"))
	})

	t.Run("should uninstall Istio when istioctl command was successful", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmder.On("Uninstall", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(nil)
		cmdResolver := TestCommanderResolver{cmder: &cmder}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Uninstall(kc, "1.2.3", log)

		// then
		require.NoError(t, err)
		cmder.AssertCalled(t, "Uninstall", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger"))
	})

}

func Test_DefaultIstioPerformer_LabelNamespaces(t *testing.T) {

	log := logger.NewLogger(false)

	t.Run("should not label namespaces when sidecar migration value doesn't exist", func(t *testing.T) {
		// given
		namespace := "test"
		kubeClient := mocks.Client{}
		clientset := fake.NewSimpleClientset(createNamespace(namespace))
		kubeClient.On("Clientset").Return(clientset, nil)
		wrapper := NewDefaultIstioPerformer(nil, nil, nil)
		istioChart := "istio-sidecar-empty"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		err := wrapper.LabelNamespaces(context.TODO(), &kubeClient, factory, "", istioChart, log)
		require.NoError(t, err)

		// then
		got, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
		require.NoError(t, err)
		require.NotContains(t, got.Labels, "istio-injection")
	})

	t.Run("should not label namespaces when sidecar migration is disabled", func(t *testing.T) {
		// given
		namespace := "test"
		kubeClient := mocks.Client{}
		clientset := fake.NewSimpleClientset(createNamespace(namespace))
		kubeClient.On("Clientset").Return(clientset, nil)
		wrapper := NewDefaultIstioPerformer(nil, nil, nil)
		istioChart := "istio-sidecar-disabled"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		err := wrapper.LabelNamespaces(context.TODO(), &kubeClient, factory, "", istioChart, log)
		require.NoError(t, err)

		// then
		got, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
		require.NoError(t, err)
		require.NotContains(t, got.Labels, "istio-injection")
	})

	t.Run("should label namespaces when sidecar migration is enabled", func(t *testing.T) {
		// given
		namespace := "test"
		kubeClient := mocks.Client{}
		clientset := fake.NewSimpleClientset(createNamespace(namespace))
		kubeClient.On("Clientset").Return(clientset, nil)
		wrapper := NewDefaultIstioPerformer(nil, nil, nil)
		istioChart := "istio-sidecar-enabled"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		err := wrapper.LabelNamespaces(context.TODO(), &kubeClient, factory, "", istioChart, log)
		require.NoError(t, err)

		// then
		got, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
		require.NoError(t, err)
		require.Contains(t, got.Labels, "istio-injection")
		require.Equal(t, "enabled", got.Labels["istio-injection"])
	})

	t.Run("should not label kube-system namespace when sidecar migration is enabled", func(t *testing.T) {
		// given
		namespace := "kube-system"
		kubeClient := mocks.Client{}
		clientset := fake.NewSimpleClientset(createNamespace(namespace))
		kubeClient.On("Clientset").Return(clientset, nil)
		wrapper := NewDefaultIstioPerformer(nil, nil, nil)
		istioChart := "istio-sidecar-enabled"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		err := wrapper.LabelNamespaces(context.TODO(), &kubeClient, factory, "", istioChart, log)
		require.NoError(t, err)

		// then
		got, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
		require.NoError(t, err)
		require.NotContains(t, got.Labels, "istio-injection")
	})

	t.Run("should not label namespace with user created label when sidecar migration is enabled", func(t *testing.T) {
		// given
		namespace := "user-ns"
		kubeClient := mocks.Client{}
		clientset := fake.NewSimpleClientset(createNamespaceWithLabel(namespace, map[string]string{"istio-injection": "disabled"}))
		kubeClient.On("Clientset").Return(clientset, nil)
		wrapper := NewDefaultIstioPerformer(nil, nil, nil)
		istioChart := "istio-sidecar-enabled"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		err := wrapper.LabelNamespaces(context.TODO(), &kubeClient, factory, "", istioChart, log)
		require.NoError(t, err)

		// then
		got, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
		require.NoError(t, err)
		require.Contains(t, got.Labels, "istio-injection")
		require.Equal(t, "disabled", got.Labels["istio-injection"])
	})
}

func createNamespace(namespace string) *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: namespace},
	}
}

func createNamespaceWithLabel(name string, labels map[string]string) *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: name, Labels: labels},
	}
}

func Test_DefaultIstioPerformer_Update(t *testing.T) {

	kubeConfig := "kubeConfig"
	log := logger.NewLogger(false)

	t.Run("should not update when istio version could not be resolved", func(t *testing.T) {
		// given
		cmdResolver := TestCommanderResolver{err: errors.New("istioctl not found")}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Update(kubeConfig, "", "1.2.3", log)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "Istio Operator definition could not be found")
	})

	t.Run("should not update when istio operator could not be found in manifest", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmder.On("Upgrade", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(errors.New("istioctl error"))
		cmdResolver := TestCommanderResolver{cmder: &cmder}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Update(kubeConfig, "", "1.2.3", log)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "Istio Operator definition could not be found")
		cmder.AssertNotCalled(t, "Update", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger"))
	})

	t.Run("should not update Istio when istioctl returned an error", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmder.On("Upgrade", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(errors.New("istioctl error"))
		cmdResolver := TestCommanderResolver{cmder: &cmder}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Update(kubeConfig, istioManifest, "1.2.3", log)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "istioctl error")
		cmder.AssertCalled(t, "Upgrade", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger"))
	})

	t.Run("should update Istio when istioctl command was successful", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmder.On("Upgrade", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(nil)
		cmdResolver := TestCommanderResolver{cmder: &cmder}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		err := wrapper.Update(kubeConfig, istioManifest, "1.2.3", log)

		// then
		require.NoError(t, err)
		cmder.AssertCalled(t, "Upgrade", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger"))
	})

}

func Test_DefaultIstioPerformer_ResetProxy(t *testing.T) {

	kubeConfig := "kubeconfig"
	log := logger.NewLogger(false)
	ctx := context.Background()
	defer ctx.Done()
	t.Run("should return error when kubeclient could not be retrieved", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmdResolver := TestCommanderResolver{cmder: &cmder}
		canUpdate := true

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		provider.On("RetrieveFrom", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(nil, errors.New("Kubeclient error"))

		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)
		proxyImageVersion := "1.2.0"
		istioChart := "istio-sidecar-disabled"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		err := wrapper.ResetProxy(ctx, kubeConfig, factory, "", istioChart, proxyImageVersion, "", log, canUpdate)
		// ResetProxy(context context.Context, kubeConfig string, workspace chart.Factory, branchVersion string, istioChart string, proxyImageVersion string, proxyImagePrefix string, logger *zap.SugaredLogger, canUpdate bool) error {

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "Kubeclient error")
	})

	t.Run("should return error when istio proxy reset returned an error", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmdResolver := TestCommanderResolver{cmder: &cmder}
		canUpdate := true

		proxy := proxymocks.IstioProxyReset{}
		proxy.On("Run", mock.Anything).Return(errors.New("Proxy reset error"))
		provider := clientsetmocks.Provider{}
		provider.On("RetrieveFrom", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(fake.NewSimpleClientset(), nil)

		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)
		proxyImageVersion := "1.2.0"
		proxyImagePrefix := "anything"
		istioChart := "istio-sidecar-disabled"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		err := wrapper.ResetProxy(ctx, kubeConfig, factory, "", istioChart, proxyImageVersion, proxyImagePrefix, log, canUpdate)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "Proxy reset error")
	})

	t.Run("should return no error when istio proxy reset was successful", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		cmdResolver := TestCommanderResolver{cmder: &cmder}
		canUpdate := true

		proxy := proxymocks.IstioProxyReset{}
		proxy.On("Run", mock.Anything).Return(nil)
		provider := clientsetmocks.Provider{}
		provider.On("RetrieveFrom", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return(fake.NewSimpleClientset(), nil)

		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)
		proxyImageVersion := "1.2.0"
		proxyImagePrefix := "anything"
		istioChart := "istio-sidecar-disabled"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		err := wrapper.ResetProxy(ctx, kubeConfig, factory, "", istioChart, proxyImageVersion, proxyImagePrefix, log, canUpdate)

		// then
		require.NoError(t, err)
	})

}

func Test_DefaultIstioPerformer_Version(t *testing.T) {

	kubeConfig := "kubeConfig"
	log := logger.NewLogger(false)

	t.Run("should not proceed if the istio version could not be resolved", func(t *testing.T) {
		// given
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)
		cmdResolver := TestCommanderResolver{err: errors.New("istioctl not found")}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		ver, err := wrapper.Version(factory, "version", "istio-test", kubeConfig, log)

		// then
		require.Empty(t, ver)
		require.Error(t, err)
		require.Equal(t, "istioctl not found", err.Error())
	})

	t.Run("should not proceed if the version command output returns an empty string", func(t *testing.T) {
		// given
		cmder := istioctlmocks.Commander{}
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)
		cmder.On("Version", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return([]byte(""), nil)
		cmdResolver := TestCommanderResolver{cmder: &cmder}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		ver, err := wrapper.Version(factory, "version", "istio-test", kubeConfig, log)

		// then
		require.Empty(t, ver)
		require.Error(t, err)
		require.Contains(t, err.Error(), "command is empty")
	})

	t.Run("should not proceed if the targetVersion is not found", func(t *testing.T) {
		// given
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{}, nil)
		cmder := istioctlmocks.Commander{}
		cmder.On("Version", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return([]byte(""), nil)
		cmdResolver := TestCommanderResolver{cmder: &cmder}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		ver, err := wrapper.Version(factory, "version", "istio-test", kubeConfig, log)

		// then
		require.Empty(t, ver)
		require.Error(t, err)
		require.Contains(t, err.Error(), "Target Version could not be found")
	})

	t.Run("should get only the client version when istio is not yet installed on the cluster", func(t *testing.T) {
		// given
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)
		cmder := istioctlmocks.Commander{}
		cmder.On("Version", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return([]byte(istioctlMockSimpleVersion), nil)
		cmdResolver := TestCommanderResolver{cmder: &cmder}

		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		ver, err := wrapper.Version(factory, "version", "istio-test", kubeConfig, log)

		// then
		require.EqualValues(t, IstioStatus{ClientVersion: "1.11.2", TargetVersion: "1.2.3-solo-fips-distroless", TargetPrefix: "anything/anything"}, ver)
		require.NoError(t, err)
		cmder.AssertCalled(t, "Version", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger"))
		cmder.AssertNumberOfCalls(t, "Version", 1)
	})

	t.Run("should get all the expected versions when istio installed on the cluster", func(t *testing.T) {
		// given
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)
		cmder := istioctlmocks.Commander{}
		cmder.On("Version", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger")).Return([]byte(istioctlMockCompleteVersion), nil)
		cmdResolver := TestCommanderResolver{cmder: &cmder}
		proxy := proxymocks.IstioProxyReset{}
		provider := clientsetmocks.Provider{}
		wrapper := NewDefaultIstioPerformer(cmdResolver, &proxy, &provider)

		// when
		ver, err := wrapper.Version(factory, "version", "istio-test", kubeConfig, log)

		// then
		require.EqualValues(t, IstioStatus{ClientVersion: "1.11.1", TargetVersion: "1.2.3-solo-fips-distroless", TargetPrefix: "anything/anything", PilotVersion: "1.11.1", DataPlaneVersion: "1.11.1"}, ver)
		require.NoError(t, err)
		cmder.AssertCalled(t, "Version", mock.AnythingOfType("string"), mock.AnythingOfType("*zap.SugaredLogger"))
		cmder.AssertNumberOfCalls(t, "Version", 1)
	})
}

func Test_getTargetProxyV2PrefixFromIstioChart(t *testing.T) {
	branch := "branch"
	log := logger.NewLogger(false)

	t.Run("should correctly parse the prefix from istio helm chart", func(t *testing.T) {
		// given
		istioChart := "istio"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files/path-tests"}, nil)

		// when
		targetPrefix, err := getTargetProxyV2PrefixFromIstioChart(factory, branch, istioChart, log)

		// then
		expectedPrefix := "istio-proxy-path/istio-proxy-dir"
		require.NoError(t, err)
		require.EqualValues(t, expectedPrefix, targetPrefix)
	})
}

func Test_getTargetVersionFromIstioChart(t *testing.T) {
	branch := "branch"
	log := logger.NewLogger(false)

	t.Run("should not get target version when the istio Chart does not exist", func(t *testing.T) {
		// given
		istioChart := "not-existing-chart"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		targetVersion, err := getTargetVersionFromIstioChart(factory, branch, istioChart, log)

		// then
		require.Empty(t, targetVersion)
		require.Error(t, err)
		require.Contains(t, err.Error(), "no such file or directory")
	})

	t.Run("should return pilot version from values when version was found in values", func(t *testing.T) {
		// given
		istioChart := "istio-values-appversion"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		targetVersion, err := getTargetVersionFromIstioChart(factory, branch, istioChart, log)

		// then
		require.NoError(t, err)
		require.EqualValues(t, "1.2.3-solo-fips-distroless", targetVersion)
	})

	t.Run("should fallback to chart appVersion when version is not found in values", func(t *testing.T) {
		// given
		istioChart := "istio-incorrect-values-only-appversion"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		targetVersion, err := getTargetVersionFromIstioChart(factory, branch, istioChart, log)

		// then
		require.NoError(t, err)
		require.EqualValues(t, "1.2.3-distroless", targetVersion)
	})

	t.Run("should fallback to chart appVersion when values.yaml is not present in the chart", func(t *testing.T) {
		// given
		istioChart := "istio-no-values-only-appversion"
		factory := &workspacemocks.Factory{}
		factory.On("Get", mock.AnythingOfType("string")).Return(&chart.KymaWorkspace{ResourceDir: "../test_files"}, nil)

		// when
		targetVersion, err := getTargetVersionFromIstioChart(factory, branch, istioChart, log)

		// then
		require.NoError(t, err)
		require.EqualValues(t, "1.2.3-distroless", targetVersion)
	})
}

func TestMapVersionToStruct(t *testing.T) {

	t.Run("Empty byte array for version command returns an error", func(t *testing.T) {
		// given
		versionOutput := []byte("")
		targetVersion := "targetVersion"
		targetDirectory := "targetDirectory"

		// when
		_, err := mapVersionToStruct(versionOutput, targetVersion, targetDirectory)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "command is empty")
	})

	t.Run("If unmarshalled properly, the byte array must be converted to struct", func(t *testing.T) {
		// given
		versionOutput := []byte(istioctlMockCompleteVersion)
		targetVersion := "targetVersion"
		targetPrefix := "anything/anything"
		expectedStruct := IstioStatus{
			ClientVersion:    "1.11.1",
			TargetVersion:    targetVersion,
			TargetPrefix:     targetPrefix,
			PilotVersion:     "1.11.1",
			DataPlaneVersion: "1.11.1",
		}

		// when
		gotStruct, err := mapVersionToStruct(versionOutput, targetVersion, targetPrefix)

		// then
		require.NoError(t, err)
		require.EqualValues(t, expectedStruct, gotStruct)
	})

}

func TestGetVersionFromJSON(t *testing.T) {
	t.Run("should get all the expected versions when istio installed on the cluster", func(t *testing.T) {
		// given
		var version IstioVersionOutput
		err := json.Unmarshal([]byte(istioctlMockCompleteVersion), &version)

		// when
		gotClient := getVersionFromJSON("client", version)
		gotPilot := getVersionFromJSON("pilot", version)
		gotDataPlane := getVersionFromJSON("dataPlane", version)
		gotNothing := getVersionFromJSON("", version)

		// then
		require.NoError(t, err)
		require.Equal(t, "1.11.1", gotClient)
		require.Equal(t, "1.11.1", gotPilot)
		require.Equal(t, "1.11.1", gotDataPlane)
		require.Equal(t, "", gotNothing)

	})
}

type TestCommanderResolver struct {
	err   error
	cmder istioctl.Commander
}

func (tcr TestCommanderResolver) GetCommander(_ istioctl.Version) (istioctl.Commander, error) {
	if tcr.err != nil {
		return nil, tcr.err
	}
	return tcr.cmder, nil
}
