package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   distributionSegmentProperty := &DistributionSegmentProperty{
//   	TargetCapacityTypes: []*string{
//   		jsii.String("targetCapacityTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-distributionsegment.html
//
type CfnAutoScalingGroupPropsMixin_DistributionSegmentProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-distributionsegment.html#cfn-autoscaling-autoscalinggroup-distributionsegment-targetcapacitytypes
	//
	TargetCapacityTypes *[]*string `field:"optional" json:"targetCapacityTypes" yaml:"targetCapacityTypes"`
}

