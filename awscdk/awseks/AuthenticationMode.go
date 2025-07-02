package awseks


// Represents the authentication mode for an Amazon EKS cluster.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv33"
//   var vpc vpc
//
//
//   eks.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   	Version: eks.KubernetesVersion_V1_33(),
//   	KubectlLayer: kubectlv33.NewKubectlV33Layer(this, jsii.String("KubectlLayer")),
//   	AuthenticationMode: eks.AuthenticationMode_API_AND_CONFIG_MAP,
//   })
//
type AuthenticationMode string

const (
	// Authenticates using a Kubernetes ConfigMap.
	AuthenticationMode_CONFIG_MAP AuthenticationMode = "CONFIG_MAP"
	// Authenticates using both the Kubernetes API server and a ConfigMap.
	AuthenticationMode_API_AND_CONFIG_MAP AuthenticationMode = "API_AND_CONFIG_MAP"
	// Authenticates using the Kubernetes API server.
	AuthenticationMode_API AuthenticationMode = "API"
)

