package mixinsawssagemaker


// Specifies the autoscaling configuration for a HyperPod cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterAutoScalingConfigProperty := &ClusterAutoScalingConfigProperty{
//   	AutoScalerType: jsii.String("autoScalerType"),
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterautoscalingconfig.html
//
type CfnClusterPropsMixin_ClusterAutoScalingConfigProperty struct {
	// The type of autoscaler to use.
	//
	// Currently supported value is `Karpenter` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterautoscalingconfig.html#cfn-sagemaker-cluster-clusterautoscalingconfig-autoscalertype
	//
	// Default: - "Karpenter".
	//
	AutoScalerType *string `field:"optional" json:"autoScalerType" yaml:"autoScalerType"`
	// Describes whether autoscaling is enabled or disabled for the cluster.
	//
	// Valid values are `Enable` and `Disable` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterautoscalingconfig.html#cfn-sagemaker-cluster-clusterautoscalingconfig-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

