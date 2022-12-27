package awseks


// Properties for defining a `CfnNodegroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNodegroupProps := &cfnNodegroupProps{
//   	clusterName: jsii.String("clusterName"),
//   	nodeRole: jsii.String("nodeRole"),
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	amiType: jsii.String("amiType"),
//   	capacityType: jsii.String("capacityType"),
//   	diskSize: jsii.Number(123),
//   	forceUpdateEnabled: jsii.Boolean(false),
//   	instanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	launchTemplate: &launchTemplateSpecificationProperty{
//   		id: jsii.String("id"),
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   	nodegroupName: jsii.String("nodegroupName"),
//   	releaseVersion: jsii.String("releaseVersion"),
//   	remoteAccess: &remoteAccessProperty{
//   		ec2SshKey: jsii.String("ec2SshKey"),
//
//   		// the properties below are optional
//   		sourceSecurityGroups: []*string{
//   			jsii.String("sourceSecurityGroups"),
//   		},
//   	},
//   	scalingConfig: &scalingConfigProperty{
//   		desiredSize: jsii.Number(123),
//   		maxSize: jsii.Number(123),
//   		minSize: jsii.Number(123),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	taints: []interface{}{
//   		&taintProperty{
//   			effect: jsii.String("effect"),
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	updateConfig: &updateConfigProperty{
//   		maxUnavailable: jsii.Number(123),
//   		maxUnavailablePercentage: jsii.Number(123),
//   	},
//   	version: jsii.String("version"),
//   }
//
type CfnNodegroupProps struct {
	// The name of the cluster to create the node group in.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	//
	// The Amazon EKS worker node `kubelet` daemon makes calls to AWS APIs on your behalf. Nodes receive permissions for these API calls through an IAM instance profile and associated policies. Before you can launch nodes and register them into a cluster, you must create an IAM role for those nodes to use when they are launched. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the **Amazon EKS User Guide** . If you specify `launchTemplate` , then don't specify [`IamInstanceProfile`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	NodeRole *string `field:"required" json:"nodeRole" yaml:"nodeRole"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// If you specify `launchTemplate` , then don't specify [`SubnetId`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// The AMI type for your node group.
	//
	// GPU instance types should use the `AL2_x86_64_GPU` AMI type. Non-GPU instances should use the `AL2_x86_64` AMI type. Arm instances should use the `AL2_ARM_64` AMI type. All types use the Amazon EKS optimized Amazon Linux 2 AMI. If you specify `launchTemplate` , and your launch template uses a custom AMI, then don't specify `amiType` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	AmiType *string `field:"optional" json:"amiType" yaml:"amiType"`
	// The capacity type of your managed node group.
	CapacityType *string `field:"optional" json:"capacityType" yaml:"capacityType"`
	// The root device disk size (in GiB) for your node group instances.
	//
	// The default disk size is 20 GiB. If you specify `launchTemplate` , then don't specify `diskSize` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old node whether or not any pods are running on the node.
	ForceUpdateEnabled interface{} `field:"optional" json:"forceUpdateEnabled" yaml:"forceUpdateEnabled"`
	// Specify the instance types for a node group.
	//
	// If you specify a GPU instance type, be sure to specify `AL2_x86_64_GPU` with the `amiType` parameter. If you specify `launchTemplate` , then you can specify zero or one instance type in your launch template *or* you can specify 0-20 instance types for `instanceTypes` . If however, you specify an instance type in your launch template *and* specify any `instanceTypes` , the node group deployment will fail. If you don't specify an instance type in a launch template or for `instanceTypes` , then `t3.medium` is used, by default. If you specify `Spot` for `capacityType` , then we recommend specifying multiple values for `instanceTypes` . For more information, see [Managed node group capacity types](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types) and [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// An object representing a node group's launch template specification.
	//
	// If specified, then do not specify `instanceTypes` , `diskSize` , or `remoteAccess` and make sure that the launch template meets the requirements in `launchTemplateSpecification` .
	LaunchTemplate interface{} `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The unique name to give your node group.
	NodegroupName *string `field:"optional" json:"nodegroupName" yaml:"nodegroupName"`
	// The AMI version of the Amazon EKS optimized AMI to use with your node group (for example, `1.14.7- *YYYYMMDD*` ). By default, the latest available AMI version for the node group's current Kubernetes version is used. For more information, see [Amazon EKS optimized Linux AMI Versions](https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html) in the *Amazon EKS User Guide* .
	//
	// > Changing this value triggers an update of the node group if one is available. However, only the latest available AMI release version is valid as an input. You cannot roll back to a previous AMI release version.
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// If you specify `launchTemplate` , then don't specify `remoteAccess` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	RemoteAccess interface{} `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	ScalingConfig interface{} `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Node group tags do not propagate to any other resources associated with the node group, such as the Amazon EC2 instances or subnets.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	//
	// Effect is one of `No_Schedule` , `Prefer_No_Schedule` , or `No_Execute` . Kubernetes taints can be used together with tolerations to control how workloads are scheduled to your nodes. For more information, see [Node taints on managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html) .
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// The node group update configuration.
	UpdateConfig interface{} `field:"optional" json:"updateConfig" yaml:"updateConfig"`
	// The Kubernetes version to use for your managed nodes.
	//
	// By default, the Kubernetes version of the cluster is used, and this is the only accepted specified value. If you specify `launchTemplate` , and your launch template uses a custom AMI, then don't specify `version` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	Version *string `field:"optional" json:"version" yaml:"version"`
}

