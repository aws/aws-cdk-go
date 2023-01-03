// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha


// What cluster type to use.
//
// Used by {@link ClusterProps.clusterType}
// Experimental.
type ClusterType string

const (
	// single-node cluster, the {@link ClusterProps.numberOfNodes} parameter is not required.
	// Experimental.
	ClusterType_SINGLE_NODE ClusterType = "SINGLE_NODE"
	// multi-node cluster, set the amount of nodes using {@link ClusterProps.numberOfNodes} parameter.
	// Experimental.
	ClusterType_MULTI_NODE ClusterType = "MULTI_NODE"
)

