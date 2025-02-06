package awscdkeks-v2alpha


// EKS cluster logging types.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	// ...
//   	Version: eks.KubernetesVersion_V1_31(),
//   	ClusterLogging: []clusterLoggingTypes{
//   		eks.*clusterLoggingTypes_API,
//   		eks.*clusterLoggingTypes_AUTHENTICATOR,
//   		eks.*clusterLoggingTypes_SCHEDULER,
//   	},
//   })
//
// Experimental.
type ClusterLoggingTypes string

const (
	// Logs pertaining to API requests to the cluster.
	// Experimental.
	ClusterLoggingTypes_API ClusterLoggingTypes = "API"
	// Logs pertaining to cluster access via the Kubernetes API.
	// Experimental.
	ClusterLoggingTypes_AUDIT ClusterLoggingTypes = "AUDIT"
	// Logs pertaining to authentication requests into the cluster.
	// Experimental.
	ClusterLoggingTypes_AUTHENTICATOR ClusterLoggingTypes = "AUTHENTICATOR"
	// Logs pertaining to state of cluster controllers.
	// Experimental.
	ClusterLoggingTypes_CONTROLLER_MANAGER ClusterLoggingTypes = "CONTROLLER_MANAGER"
	// Logs pertaining to scheduling decisions.
	// Experimental.
	ClusterLoggingTypes_SCHEDULER ClusterLoggingTypes = "SCHEDULER"
)

