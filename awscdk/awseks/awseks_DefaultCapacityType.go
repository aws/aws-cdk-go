package awseks


// The default capacity type for the cluster.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	defaultCapacityType: eks.defaultCapacityType_EC2,
//   })
//
type DefaultCapacityType string

const (
	// managed node group.
	DefaultCapacityType_NODEGROUP DefaultCapacityType = "NODEGROUP"
	// EC2 autoscaling group.
	DefaultCapacityType_EC2 DefaultCapacityType = "EC2"
)

