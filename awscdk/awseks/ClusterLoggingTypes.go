package awseks


// EKS cluster logging types.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv33"
//
//
//   cluster := eks.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	// ...
//   	Version: eks.KubernetesVersion_V1_33(),
//   	ClusterLogging: []ClusterLoggingTypes{
//   		eks.ClusterLoggingTypes_API,
//   		eks.ClusterLoggingTypes_AUTHENTICATOR,
//   		eks.ClusterLoggingTypes_SCHEDULER,
//   	},
//   	KubectlLayer: kubectlv33.NewKubectlV33Layer(this, jsii.String("kubectl")),
//   })
//
type ClusterLoggingTypes string

const (
	// Logs pertaining to API requests to the cluster.
	ClusterLoggingTypes_API ClusterLoggingTypes = "API"
	// Logs pertaining to cluster access via the Kubernetes API.
	ClusterLoggingTypes_AUDIT ClusterLoggingTypes = "AUDIT"
	// Logs pertaining to authentication requests into the cluster.
	ClusterLoggingTypes_AUTHENTICATOR ClusterLoggingTypes = "AUTHENTICATOR"
	// Logs pertaining to state of cluster controllers.
	ClusterLoggingTypes_CONTROLLER_MANAGER ClusterLoggingTypes = "CONTROLLER_MANAGER"
	// Logs pertaining to scheduling decisions.
	ClusterLoggingTypes_SCHEDULER ClusterLoggingTypes = "SCHEDULER"
)

