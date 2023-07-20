package awseks


// Properties for defining a `CfnNodegroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNodegroupProps := &CfnNodegroupProps{
//   	ClusterName: jsii.String("clusterName"),
//   	NodeRole: jsii.String("nodeRole"),
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	AmiType: jsii.String("amiType"),
//   	CapacityType: jsii.String("capacityType"),
//   	DiskSize: jsii.Number(123),
//   	ForceUpdateEnabled: jsii.Boolean(false),
//   	InstanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	LaunchTemplate: &LaunchTemplateSpecificationProperty{
//   		Id: jsii.String("id"),
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	NodegroupName: jsii.String("nodegroupName"),
//   	ReleaseVersion: jsii.String("releaseVersion"),
//   	RemoteAccess: &RemoteAccessProperty{
//   		Ec2SshKey: jsii.String("ec2SshKey"),
//
//   		// the properties below are optional
//   		SourceSecurityGroups: []*string{
//   			jsii.String("sourceSecurityGroups"),
//   		},
//   	},
//   	ScalingConfig: &ScalingConfigProperty{
//   		DesiredSize: jsii.Number(123),
//   		MaxSize: jsii.Number(123),
//   		MinSize: jsii.Number(123),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Taints: []interface{}{
//   		&TaintProperty{
//   			Effect: jsii.String("effect"),
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UpdateConfig: &UpdateConfigProperty{
//   		MaxUnavailable: jsii.Number(123),
//   		MaxUnavailablePercentage: jsii.Number(123),
//   	},
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html
//
type CfnNodegroupProps struct {
	// The name of the cluster to create the node group in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	//
	// The Amazon EKS worker node `kubelet` daemon makes calls to AWS APIs on your behalf. Nodes receive permissions for these API calls through an IAM instance profile and associated policies. Before you can launch nodes and register them into a cluster, you must create an IAM role for those nodes to use when they are launched. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the **Amazon EKS User Guide** . If you specify `launchTemplate` , then don't specify [`IamInstanceProfile`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-noderole
	//
	NodeRole *string `field:"required" json:"nodeRole" yaml:"nodeRole"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// If you specify `launchTemplate` , then don't specify [`SubnetId`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-subnets
	//
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// The AMI type for your node group.
	//
	// If you specify `launchTemplate` , and your launch template uses a custom AMI, then don't specify `amiType` , or the node group deployment will fail. If your launch template uses a Windows custom AMI, then add `eks:kube-proxy-windows` to your Windows nodes `rolearn` in the `aws-auth` `ConfigMap` . For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-amitype
	//
	AmiType *string `field:"optional" json:"amiType" yaml:"amiType"`
	// The capacity type of your managed node group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-capacitytype
	//
	CapacityType *string `field:"optional" json:"capacityType" yaml:"capacityType"`
	// The root device disk size (in GiB) for your node group instances.
	//
	// The default disk size is 20 GiB for Linux and Bottlerocket. The default disk size is 50 GiB for Windows. If you specify `launchTemplate` , then don't specify `diskSize` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-disksize
	//
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old node whether or not any pods are running on the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-forceupdateenabled
	//
	ForceUpdateEnabled interface{} `field:"optional" json:"forceUpdateEnabled" yaml:"forceUpdateEnabled"`
	// Specify the instance types for a node group.
	//
	// If you specify a GPU instance type, make sure to also specify an applicable GPU AMI type with the `amiType` parameter. If you specify `launchTemplate` , then you can specify zero or one instance type in your launch template *or* you can specify 0-20 instance types for `instanceTypes` . If however, you specify an instance type in your launch template *and* specify any `instanceTypes` , the node group deployment will fail. If you don't specify an instance type in a launch template or for `instanceTypes` , then `t3.medium` is used, by default. If you specify `Spot` for `capacityType` , then we recommend specifying multiple values for `instanceTypes` . For more information, see [Managed node group capacity types](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types) and [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	//
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels applied to the nodes in the node group.
	//
	// > Only labels that are applied with the Amazon EKS API are shown here. There may be other Kubernetes labels applied to the nodes in this group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// An object representing a node group's launch template specification.
	//
	// If specified, then do not specify `instanceTypes` , `diskSize` , or `remoteAccess` and make sure that the launch template meets the requirements in `launchTemplateSpecification` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-launchtemplate
	//
	LaunchTemplate interface{} `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The unique name to give your node group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-nodegroupname
	//
	NodegroupName *string `field:"optional" json:"nodegroupName" yaml:"nodegroupName"`
	// The AMI version of the Amazon EKS optimized AMI to use with your node group (for example, `1.14.7- *YYYYMMDD*` ). By default, the latest available AMI version for the node group's current Kubernetes version is used. For more information, see [Amazon EKS optimized Linux AMI Versions](https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html) in the *Amazon EKS User Guide* .
	//
	// > Changing this value triggers an update of the node group if one is available. You can't update other properties at the same time as updating `Release Version` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-releaseversion
	//
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access configuration to use with your node group.
	//
	// For Linux, the protocol is SSH. For Windows, the protocol is RDP. If you specify `launchTemplate` , then don't specify `remoteAccess` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-remoteaccess
	//
	RemoteAccess interface{} `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-scalingconfig
	//
	ScalingConfig interface{} `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// The metadata applied to the node group to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Node group tags do not propagate to any other resources associated with the node group, such as the Amazon EC2 instances or subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	//
	// Effect is one of `No_Schedule` , `Prefer_No_Schedule` , or `No_Execute` . Kubernetes taints can be used together with tolerations to control how workloads are scheduled to your nodes. For more information, see [Node taints on managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-taints
	//
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// The node group update configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-updateconfig
	//
	UpdateConfig interface{} `field:"optional" json:"updateConfig" yaml:"updateConfig"`
	// The Kubernetes version to use for your managed nodes.
	//
	// By default, the Kubernetes version of the cluster is used, and this is the only accepted specified value. If you specify `launchTemplate` , and your launch template uses a custom AMI, then don't specify `version` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	//
	// > You can't update other properties at the same time as updating `Version` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

