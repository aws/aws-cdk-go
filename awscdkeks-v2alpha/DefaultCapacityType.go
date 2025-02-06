package awscdkeks-v2alpha


// The default capacity type for the cluster.
// Experimental.
type DefaultCapacityType string

const (
	// managed node group.
	// Experimental.
	DefaultCapacityType_NODEGROUP DefaultCapacityType = "NODEGROUP"
	// EC2 autoscaling group.
	// Experimental.
	DefaultCapacityType_EC2 DefaultCapacityType = "EC2"
)

