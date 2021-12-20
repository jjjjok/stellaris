package namespace_mapping

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"harmonycloud.cn/stellaris/pkg/apis/multicluster/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/utils/pointer"
	"math/rand"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"testing"
	"time"
)

var cfg *rest.Config
var k8sClient client.Client
var testEnv *envtest.Environment
var testScheme = runtime.NewScheme()
var reconciler *NamespaceMappingReconciler
var controllerDone context.CancelFunc
var mgr ctrl.Manager

func TestNamespaceMapping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Namespace-Mapping Suite")
}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.New(zap.UseDevMode(true), zap.WriteTo(GinkgoWriter)))
	rand.Seed(time.Now().UnixNano())
	By("bootstrapping test environment")

	// TODO fix test env
	yamlPath := ""
	testEnv = &envtest.Environment{
		ControlPlaneStartTimeout: time.Minute,
		ControlPlaneStopTimeout:  time.Minute,
		UseExistingCluster:       pointer.BoolPtr(false),
		CRDDirectoryPaths:        []string{yamlPath},
	}

	var err error
	cfg, err = testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	err = v1alpha1.SchemeBuilder.AddToScheme(testScheme)
	Expect(err).NotTo(HaveOccurred())

	err = scheme.AddToScheme(testScheme)
	Expect(err).NotTo(HaveOccurred())

	k8sClient, err = client.New(cfg, client.Options{Scheme: testScheme})
	Expect(err).ToNot(HaveOccurred())
	Expect(k8sClient).ToNot(BeNil())

	// setup the controller manager since we need the component handler to run in the background
	mgr, err = ctrl.NewManager(cfg, ctrl.Options{
		Scheme:                  testScheme,
		MetricsBindAddress:      "0",
		LeaderElection:          false,
		LeaderElectionNamespace: "default",
		LeaderElectionID:        "test",
	})
	Expect(err).NotTo(HaveOccurred())

	reconciler = &NamespaceMappingReconciler{
		Scheme: testScheme,
		Client: k8sClient,
		log:    logf.Log.WithName("namespace_mapping_controller"),
	}

	var ctx context.Context
	ctx, controllerDone = context.WithCancel(context.Background())
	// start the controller in the background so that new componentRevisions are created
	go func() {
		err = mgr.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	}()
	close(done)

}, 120)

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	controllerDone()
	err := testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
})
