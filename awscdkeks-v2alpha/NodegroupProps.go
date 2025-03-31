package awscdkeks-v2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// NodeGroup properties interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkeks-v2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var instanceType instanceType
//   var role role
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//
//   nodegroupProps := &NodegroupProps{
//   	Cluster: cluster,
//
//   	// the properties below are optional
//   	AmiType: eks_v2_alpha.NodegroupAmiType_AL2_X86_64,
//   	CapacityType: eks_v2_alpha.CapacityType_SPOT,
//   	DesiredSize: jsii.Number(123),
//   	DiskSize: jsii.Number(123),
//   	EnableNodeAutoRepair: jsii.Boolean(false),
//   	ForceUpdate: jsii.Boolean(false),
//   	InstanceType: instanceType,
//   	InstanceTypes: []*instanceType{
//   		instanceType,
//   	},
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	LaunchTemplateSpec: &LaunchTemplateSpec{
//   		Id: jsii.String("id"),
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//   	MaxSize: jsii.Number(123),
//   	MaxUnavailable: jsii.Number(123),
//   	MaxUnavailablePercentage: jsii.Number(123),
//   	MinSize: jsii.Number(123),
//   	NodegroupName: jsii.String("nodegroupName"),
//   	NodeRole: role,
//   	ReleaseVersion: jsii.String("releaseVersion"),
//   	RemoteAccess: &NodegroupRemoteAccess{
//   		SshKeyName: jsii.String("sshKeyName"),
//
//   		// the properties below are optional
//   		SourceSecurityGroups: []iSecurityGroup{
//   			securityGroup,
//   		},
//   	},
//   	Subnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Taints: []taintSpec{
//   		&taintSpec{
//   			Effect: eks_v2_alpha.TaintEffect_NO_SCHEDULE,
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// Experimental.
type NodegroupProps struct {
	// The AMI type for your node group.
	//
	// If you explicitly specify the launchTemplate with custom AMI, do not specify this property, or
	// the node group deployment will fail. In other cases, you will need to specify correct amiType for the nodegroup.
	// Default: - auto-determined from the instanceTypes property when launchTemplateSpec property is not specified.
	//
	// Experimental.
	AmiType NodegroupAmiType `field:"optional" json:"amiType" yaml:"amiType"`
	// The capacity type of the nodegroup.
	// Default: - ON_DEMAND.
	//
	// Experimental.
	CapacityType CapacityType `field:"optional" json:"capacityType" yaml:"capacityType"`
	// The current number of worker nodes that the managed node group should maintain.
	//
	// If not specified,
	// the nodewgroup will initially create `minSize` instances.
	// Default: 2.
	//
	// Experimental.
	DesiredSize *float64 `field:"optional" json:"desiredSize" yaml:"desiredSize"`
	// The root device disk size (in GiB) for your node group instances.
	// Default: 20.
	//
	// Experimental.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Specifies whether to enable node auto repair for the node group.
	//
	// Node auto repair is disabled by default.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/node-health.html#node-auto-repair
	//
	// Default: - disabled.
	//
	// Experimental.
	EnableNodeAutoRepair *bool `field:"optional" json:"enableNodeAutoRepair" yaml:"enableNodeAutoRepair"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old
	// node whether or not any pods are
	// running on the node.
	// Default: true.
	//
	// Experimental.
	ForceUpdate *bool `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// The instance type to use for your node group.
	//
	// Currently, you can specify a single instance type for a node group.
	// The default value for this parameter is `t3.medium`. If you choose a GPU instance type, be sure to specify the
	// `AL2_x86_64_GPU`, `BOTTLEROCKET_ARM_64_NVIDIA`, or `BOTTLEROCKET_x86_64_NVIDIA` with the amiType parameter.
	// Default: t3.medium
	//
	// Deprecated: Use `instanceTypes` instead.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The instance types to use for your node group.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	//
	// Default: t3.medium will be used according to the cloudformation document.
	//
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	// Default: - None.
	//
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Launch template specification used for the nodegroup.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html
	//
	// Default: - no launch template.
	//
	// Experimental.
	LaunchTemplateSpec *LaunchTemplateSpec `field:"optional" json:"launchTemplateSpec" yaml:"launchTemplateSpec"`
	// The maximum number of worker nodes that the managed node group can scale out to.
	//
	// Managed node groups can support up to 100 nodes by default.
	// Default: - desiredSize.
	//
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The maximum number of nodes unavailable at once during a version update.
	//
	// Nodes will be updated in parallel. The maximum number is 100.
	//
	// This value or `maxUnavailablePercentage` is required to have a value for custom update configurations to be applied.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-updateconfig.html#cfn-eks-nodegroup-updateconfig-maxunavailable
	//
	// Default: 1.
	//
	// Experimental.
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// The maximum percentage of nodes unavailable during a version update.
	//
	// This percentage of nodes will be updated in parallel, up to 100 nodes at once.
	//
	// This value or `maxUnavailable` is required to have a value for custom update configurations to be applied.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-updateconfig.html#cfn-eks-nodegroup-updateconfig-maxunavailablepercentage
	//
	// Default: undefined - node groups will update instances one at a time.
	//
	// Experimental.
	MaxUnavailablePercentage *float64 `field:"optional" json:"maxUnavailablePercentage" yaml:"maxUnavailablePercentage"`
	// The minimum number of worker nodes that the managed node group can scale in to.
	//
	// This number must be greater than or equal to zero.
	// Default: 1.
	//
	// Experimental.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Name of the Nodegroup.
	// Default: - resource ID.
	//
	// Experimental.
	NodegroupName *string `field:"optional" json:"nodegroupName" yaml:"nodegroupName"`
	// The IAM role to associate with your node group.
	//
	// The Amazon EKS worker node kubelet daemon
	// makes calls to AWS APIs on your behalf. Worker nodes receive permissions for these API calls through
	// an IAM instance profile and associated policies. Before you can launch worker nodes and register them
	// into a cluster, you must create an IAM role for those worker nodes to use when they are launched.
	// Default: - None. Auto-generated if not specified.
	//
	// Experimental.
	NodeRole awsiam.IRole `field:"optional" json:"nodeRole" yaml:"nodeRole"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group (for example, `1.14.7-YYYYMMDD`).
	// Default: - The latest available AMI version for the node group's current Kubernetes version is used.
	//
	// Experimental.
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// Disabled by default, however, if you
	// specify an Amazon EC2 SSH key but do not specify a source security group when you create a managed node group,
	// then port 22 on the worker nodes is opened to the internet (0.0.0.0/0)
	// Default: - disabled.
	//
	// Experimental.
	RemoteAccess *NodegroupRemoteAccess `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// By specifying the
	// SubnetSelection, the selected subnets will automatically apply required tags i.e.
	// `kubernetes.io/cluster/CLUSTER_NAME` with a value of `shared`, where `CLUSTER_NAME` is replaced with
	// the name of your cluster.
	// Default: - private subnets.
	//
	// Experimental.
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of
	// a key and an optional value, both of which you define. Node group tags do not propagate to any other resources
	// associated with the node group, such as the Amazon EC2 instances or subnets.
	// Default: - None.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	// Default: - None.
	//
	// Experimental.
	Taints *[]*TaintSpec `field:"optional" json:"taints" yaml:"taints"`
	// Cluster resource.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

