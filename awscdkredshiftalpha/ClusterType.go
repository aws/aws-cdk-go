package awscdkredshiftalpha


// What cluster type to use.
//
// Used by `ClusterProps.clusterType`
// Experimental.
type ClusterType string

const (
	// single-node cluster, the `ClusterProps.numberOfNodes` parameter is not required.
	// Experimental.
	ClusterType_SINGLE_NODE ClusterType = "SINGLE_NODE"
	// multi-node cluster, set the amount of nodes using `ClusterProps.numberOfNodes` parameter.
	// Experimental.
	ClusterType_MULTI_NODE ClusterType = "MULTI_NODE"
)

