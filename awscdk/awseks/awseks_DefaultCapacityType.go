package awseks


// The default capacity type for the cluster.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	defaultCapacityType: eks.defaultCapacityType_EC2,
//   })
//
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

