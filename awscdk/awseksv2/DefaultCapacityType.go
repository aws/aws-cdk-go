package awseksv2


// The default capacity type for the cluster.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_34(),
//   	DefaultCapacityType: eks.DefaultCapacityType_NODEGROUP,
//   	DefaultCapacity: jsii.Number(0),
//   })
//
//   cluster.AddNodegroupCapacity(jsii.String("custom-node-group"), &NodegroupOptions{
//   	InstanceTypes: []InstanceType{
//   		ec2.NewInstanceType(jsii.String("m5.large")),
//   	},
//   	MinSize: jsii.Number(4),
//   	DiskSize: jsii.Number(100),
//   })
//
type DefaultCapacityType string

const (
	// managed node group.
	DefaultCapacityType_NODEGROUP DefaultCapacityType = "NODEGROUP"
	// EC2 autoscaling group.
	DefaultCapacityType_EC2 DefaultCapacityType = "EC2"
	// Auto Mode.
	DefaultCapacityType_AUTOMODE DefaultCapacityType = "AUTOMODE"
)

