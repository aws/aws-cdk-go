package awscdkeksv2alpha


// EKS cluster logging types.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	// ...
//   	Version: eks.KubernetesVersion_V1_34(),
//   	ClusterLogging: []ClusterLoggingTypes{
//   		eks.ClusterLoggingTypes_API,
//   		eks.ClusterLoggingTypes_AUTHENTICATOR,
//   		eks.ClusterLoggingTypes_SCHEDULER,
//   	},
//   })
//
// Deprecated.
type ClusterLoggingTypes string

const (
	// Logs pertaining to API requests to the cluster.
	// Deprecated.
	ClusterLoggingTypes_API ClusterLoggingTypes = "API"
	// Logs pertaining to cluster access via the Kubernetes API.
	// Deprecated.
	ClusterLoggingTypes_AUDIT ClusterLoggingTypes = "AUDIT"
	// Logs pertaining to authentication requests into the cluster.
	// Deprecated.
	ClusterLoggingTypes_AUTHENTICATOR ClusterLoggingTypes = "AUTHENTICATOR"
	// Logs pertaining to state of cluster controllers.
	// Deprecated.
	ClusterLoggingTypes_CONTROLLER_MANAGER ClusterLoggingTypes = "CONTROLLER_MANAGER"
	// Logs pertaining to scheduling decisions.
	// Deprecated.
	ClusterLoggingTypes_SCHEDULER ClusterLoggingTypes = "SCHEDULER"
)

