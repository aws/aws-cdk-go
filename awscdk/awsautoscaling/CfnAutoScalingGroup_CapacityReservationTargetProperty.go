package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityReservationTargetProperty := &CapacityReservationTargetProperty{
//   	CapacityReservationIds: []*string{
//   		jsii.String("capacityReservationIds"),
//   	},
//   	CapacityReservationResourceGroupArns: []*string{
//   		jsii.String("capacityReservationResourceGroupArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-capacityreservationtarget.html
//
type CfnAutoScalingGroup_CapacityReservationTargetProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-capacityreservationtarget.html#cfn-autoscaling-autoscalinggroup-capacityreservationtarget-capacityreservationids
	//
	CapacityReservationIds *[]*string `field:"optional" json:"capacityReservationIds" yaml:"capacityReservationIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-capacityreservationtarget.html#cfn-autoscaling-autoscalinggroup-capacityreservationtarget-capacityreservationresourcegrouparns
	//
	CapacityReservationResourceGroupArns *[]*string `field:"optional" json:"capacityReservationResourceGroupArns" yaml:"capacityReservationResourceGroupArns"`
}

