package awseks


// An object representing the scaling configuration details for the Auto Scaling group that is associated with your node group.
//
// When creating a node group, you must specify all or none of the properties. When updating a node group, you can specify any or none of the properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingConfigProperty := &scalingConfigProperty{
//   	desiredSize: jsii.Number(123),
//   	maxSize: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   }
//
type CfnNodegroup_ScalingConfigProperty struct {
	// The current number of nodes that the managed node group should maintain.
	//
	// > If you use Cluster Autoscaler, you shouldn't change the desiredSize value directly, as this can cause the Cluster Autoscaler to suddenly scale up or scale down.
	//
	// Whenever this parameter changes, the number of worker nodes in the node group is updated to the specified size. If this parameter is given a value that is smaller than the current number of running worker nodes, the necessary number of worker nodes are terminated to match the given value. When using CloudFormation, no action occurs if you remove this parameter from your CFN template.
	//
	// This parameter can be different from minSize in some cases, such as when starting with extra hosts for testing. This parameter can also be different when you want to start with an estimated number of needed hosts, but let Cluster Autoscaler reduce the number if there are too many. When Cluster Autoscaler is used, the desiredSize parameter is altered by Cluster Autoscaler (but can be out-of-date for short periods of time). Cluster Autoscaler doesn't scale a managed node group lower than minSize or higher than maxSize.
	DesiredSize *float64 `field:"optional" json:"desiredSize" yaml:"desiredSize"`
	// The maximum number of nodes that the managed node group can scale out to.
	//
	// For information about the maximum number that you can specify, see [Amazon EKS service quotas](https://docs.aws.amazon.com/eks/latest/userguide/service-quotas.html) in the *Amazon EKS User Guide* .
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of nodes that the managed node group can scale in to.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
}

