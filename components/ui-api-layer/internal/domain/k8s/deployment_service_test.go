package k8s_test

import (
	"testing"
	"time"

	"github.com/kyma-project/kyma/components/ui-api-layer/internal/domain/k8s"
	testingUtils "github.com/kyma-project/kyma/components/ui-api-layer/internal/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
)

func TestDeploymentService_List(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		deployment1 := fixDeployment("one", "ns1", "deployment")
		deployment2 := fixDeployment("two", "ns1", "function")
		deployment3 := fixDeployment("three", "ns2", "deployment")
		deployment4 := fixDeployment("four", "ns2", "function")

		informer := fixDeploymentInformer(deployment1, deployment2, deployment3, deployment4)
		svc, err := k8s.NewDeploymentService(informer)
		require.NoError(t, err)
		testingUtils.WaitForInformerStartAtMost(t, time.Second, informer)

		result, err := svc.List("ns1")

		require.NoError(t, err)
		assert.Equal(t, []*v1beta2.Deployment{
			deployment1, deployment2,
		}, result)
	})

	t.Run("Not found", func(t *testing.T) {
		var expected []*v1beta2.Deployment

		informer := fixDeploymentInformer()
		svc, err := k8s.NewDeploymentService(informer)
		require.NoError(t, err)
		testingUtils.WaitForInformerStartAtMost(t, time.Second, informer)

		result, err := svc.List("ns1")

		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}

func TestDeploymentService_ListWithoutFunctions(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		deployment1 := fixDeployment("one", "ns1", "deployment")
		deployment2 := fixDeployment("two", "ns1", "function")
		deployment3 := fixDeployment("three", "ns2", "deployment")
		deployment4 := fixDeployment("four", "ns2", "function")

		informer := fixDeploymentInformer(deployment1, deployment2, deployment3, deployment4)
		svc, err := k8s.NewDeploymentService(informer)
		require.NoError(t, err)
		testingUtils.WaitForInformerStartAtMost(t, time.Second, informer)

		result, err := svc.ListWithoutFunctions("ns1")

		require.NoError(t, err)
		assert.Equal(t, []*v1beta2.Deployment{
			deployment1,
		}, result)
	})

	t.Run("Not found", func(t *testing.T) {
		var expected []*v1beta2.Deployment

		informer := fixDeploymentInformer()
		svc, err := k8s.NewDeploymentService(informer)
		require.NoError(t, err)
		testingUtils.WaitForInformerStartAtMost(t, time.Second, informer)

		result, err := svc.ListWithoutFunctions("ns1")

		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}

func fixDeployment(name, namespace, kind string) *v1beta2.Deployment {
	return &v1beta2.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				kind: "",
			},
		},
	}
}

func fixDeploymentInformer(objects ...runtime.Object) cache.SharedIndexInformer {
	client := fake.NewSimpleClientset(objects...)
	informerFactory := informers.NewSharedInformerFactory(client, 0)

	return informerFactory.Apps().V1beta2().Deployments().Informer()
}
