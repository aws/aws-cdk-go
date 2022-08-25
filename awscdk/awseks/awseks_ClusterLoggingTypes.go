package awseks


// EKS cluster logging types.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	// ...
//   	version: eks.kubernetesVersion_V1_21(),
//   	clusterLogging: []clusterLoggingTypes{
//   		eks.*clusterLoggingTypes_API,
//   		eks.*clusterLoggingTypes_AUTHENTICATOR,
//   		eks.*clusterLoggingTypes_SCHEDULER,
//   	},
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

