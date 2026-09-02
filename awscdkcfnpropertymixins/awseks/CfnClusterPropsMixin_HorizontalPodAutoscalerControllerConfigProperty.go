package awseks


// The horizontal pod autoscaler controller configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   horizontalPodAutoscalerControllerConfigProperty := &HorizontalPodAutoscalerControllerConfigProperty{
//   	HorizontalPodAutoscalerSyncPeriod: jsii.String("horizontalPodAutoscalerSyncPeriod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-horizontalpodautoscalercontrollerconfig.html
//
type CfnClusterPropsMixin_HorizontalPodAutoscalerControllerConfigProperty struct {
	// The interval between each sync of the horizontal pod autoscaler (e.g., 15s, 1m).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-horizontalpodautoscalercontrollerconfig.html#cfn-eks-cluster-horizontalpodautoscalercontrollerconfig-horizontalpodautoscalersyncperiod
	//
	HorizontalPodAutoscalerSyncPeriod *string `field:"optional" json:"horizontalPodAutoscalerSyncPeriod" yaml:"horizontalPodAutoscalerSyncPeriod"`
}

