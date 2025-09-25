package awssagemaker


// Specifies the autoscaling configuration for a HyperPod cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterAutoScalingConfigProperty := &ClusterAutoScalingConfigProperty{
//   	Mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	AutoScalerType: jsii.String("autoScalerType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterautoscalingconfig.html
//
type CfnCluster_ClusterAutoScalingConfigProperty struct {
	// Describes whether autoscaling is enabled or disabled for the cluster.
	//
	// Valid values are `Enable` and `Disable` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterautoscalingconfig.html#cfn-sagemaker-cluster-clusterautoscalingconfig-mode
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The type of autoscaler to use.
	//
	// Currently supported value is `Karpenter` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterautoscalingconfig.html#cfn-sagemaker-cluster-clusterautoscalingconfig-autoscalertype
	//
	// Default: - "Karpenter".
	//
	AutoScalerType *string `field:"optional" json:"autoScalerType" yaml:"autoScalerType"`
}

