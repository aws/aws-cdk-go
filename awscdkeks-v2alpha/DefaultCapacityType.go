package awscdkeks-v2alpha


// The default capacity type for the cluster.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_33(),
//   	DefaultCapacityType: eks.DefaultCapacityType_NODEGROUP,
//   	DefaultCapacity: jsii.Number(0),
//   })
//
//   cluster.AddNodegroupCapacity(jsii.String("custom-node-group"), &NodegroupOptions{
//   	InstanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("m5.large")),
//   	},
//   	MinSize: jsii.Number(4),
//   	DiskSize: jsii.Number(100),
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
	// Auto Mode.
	// Experimental.
	DefaultCapacityType_AUTOMODE DefaultCapacityType = "AUTOMODE"
)

