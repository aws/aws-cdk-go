package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// The Nodegroup Options for addNodeGroup() method.
//
// Example:
//   var cluster cluster
//
//   cluster.addNodegroupCapacity(jsii.String("extra-ng-spot"), &nodegroupOptions{
//   	instanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("c5.large")),
//   		ec2.NewInstanceType(jsii.String("c5a.large")),
//   		ec2.NewInstanceType(jsii.String("c5d.large")),
//   	},
//   	minSize: jsii.Number(3),
//   	capacityType: eks.capacityType_SPOT,
//   })
//
// Experimental.
type NodegroupOptions struct {
	// The AMI type for your node group.
	//
	// If you explicitly specify the launchTemplate with custom AMI, do not specify this property, or
	// the node group deployment will fail. In other cases, you will need to specify correct amiType for the nodegroup.
	// Experimental.
	AmiType NodegroupAmiType `field:"optional" json:"amiType" yaml:"amiType"`
	// The capacity type of the nodegroup.
	// Experimental.
	CapacityType CapacityType `field:"optional" json:"capacityType" yaml:"capacityType"`
	// The current number of worker nodes that the managed node group should maintain.
	//
	// If not specified,
	// the nodewgroup will initially create `minSize` instances.
	// Experimental.
	DesiredSize *float64 `field:"optional" json:"desiredSize" yaml:"desiredSize"`
	// The root device disk size (in GiB) for your node group instances.
	// Experimental.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old
	// node whether or not any pods are
	// running on the node.
	// Experimental.
	ForceUpdate *bool `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// The instance type to use for your node group.
	//
	// Currently, you can specify a single instance type for a node group.
	// The default value for this parameter is `t3.medium`. If you choose a GPU instance type, be sure to specify the
	// `AL2_x86_64_GPU` with the amiType parameter.
	// Deprecated: Use `instanceTypes` instead.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The instance types to use for your node group.
	// See: - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	//
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Launch template specification used for the nodegroup.
	// See: - https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html
	//
	// Experimental.
	LaunchTemplateSpec *LaunchTemplateSpec `field:"optional" json:"launchTemplateSpec" yaml:"launchTemplateSpec"`
	// The maximum number of worker nodes that the managed node group can scale out to.
	//
	// Managed node groups can support up to 100 nodes by default.
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of worker nodes that the managed node group can scale in to.
	//
	// This number must be greater than or equal to zero.
	// Experimental.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Name of the Nodegroup.
	// Experimental.
	NodegroupName *string `field:"optional" json:"nodegroupName" yaml:"nodegroupName"`
	// The IAM role to associate with your node group.
	//
	// The Amazon EKS worker node kubelet daemon
	// makes calls to AWS APIs on your behalf. Worker nodes receive permissions for these API calls through
	// an IAM instance profile and associated policies. Before you can launch worker nodes and register them
	// into a cluster, you must create an IAM role for those worker nodes to use when they are launched.
	// Experimental.
	NodeRole awsiam.IRole `field:"optional" json:"nodeRole" yaml:"nodeRole"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group (for example, `1.14.7-YYYYMMDD`).
	// Experimental.
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// Disabled by default, however, if you
	// specify an Amazon EC2 SSH key but do not specify a source security group when you create a managed node group,
	// then port 22 on the worker nodes is opened to the internet (0.0.0.0/0)
	// Experimental.
	RemoteAccess *NodegroupRemoteAccess `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// By specifying the
	// SubnetSelection, the selected subnets will automatically apply required tags i.e.
	// `kubernetes.io/cluster/CLUSTER_NAME` with a value of `shared`, where `CLUSTER_NAME` is replaced with
	// the name of your cluster.
	// Experimental.
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of
	// a key and an optional value, both of which you define. Node group tags do not propagate to any other resources
	// associated with the node group, such as the Amazon EC2 instances or subnets.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	// Experimental.
	Taints *[]*TaintSpec `field:"optional" json:"taints" yaml:"taints"`
}

