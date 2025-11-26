package previewawseksmixins


// An object representing the scaling configuration details for the Amazon EC2 Auto Scaling group that is associated with your node group.
//
// When creating a node group, you must specify all or none of the properties. When updating a node group, you can specify any or none of the properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scalingConfigProperty := &ScalingConfigProperty{
//   	DesiredSize: jsii.Number(123),
//   	MaxSize: jsii.Number(123),
//   	MinSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-scalingconfig.html
//
type CfnNodegroupPropsMixin_ScalingConfigProperty struct {
	// The current number of nodes that the managed node group should maintain.
	//
	// > If you use the Kubernetes [Cluster Autoscaler](https://docs.aws.amazon.com/https://github.com/kubernetes/autoscaler#kubernetes-autoscaler) , you shouldn't change the `desiredSize` value directly, as this can cause the Cluster Autoscaler to suddenly scale up or scale down.
	//
	// Whenever this parameter changes, the number of worker nodes in the node group is updated to the specified size. If this parameter is given a value that is smaller than the current number of running worker nodes, the necessary number of worker nodes are terminated to match the given value. When using CloudFormation, no action occurs if you remove this parameter from your CFN template.
	//
	// This parameter can be different from `minSize` in some cases, such as when starting with extra hosts for testing. This parameter can also be different when you want to start with an estimated number of needed hosts, but let the Cluster Autoscaler reduce the number if there are too many. When the Cluster Autoscaler is used, the `desiredSize` parameter is altered by the Cluster Autoscaler (but can be out-of-date for short periods of time). the Cluster Autoscaler doesn't scale a managed node group lower than `minSize` or higher than `maxSize` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-scalingconfig.html#cfn-eks-nodegroup-scalingconfig-desiredsize
	//
	DesiredSize *float64 `field:"optional" json:"desiredSize" yaml:"desiredSize"`
	// The maximum number of nodes that the managed node group can scale out to.
	//
	// For information about the maximum number that you can specify, see [Amazon EKS service quotas](https://docs.aws.amazon.com/eks/latest/userguide/service-quotas.html) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-scalingconfig.html#cfn-eks-nodegroup-scalingconfig-maxsize
	//
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of nodes that the managed node group can scale in to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-scalingconfig.html#cfn-eks-nodegroup-scalingconfig-minsize
	//
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
}

