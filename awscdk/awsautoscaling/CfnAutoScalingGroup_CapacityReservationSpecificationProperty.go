package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityReservationSpecificationProperty := &CapacityReservationSpecificationProperty{
//   	CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//
//   	// the properties below are optional
//   	CapacityReservationTarget: &CapacityReservationTargetProperty{
//   		CapacityReservationIds: []*string{
//   			jsii.String("capacityReservationIds"),
//   		},
//   		CapacityReservationResourceGroupArns: []*string{
//   			jsii.String("capacityReservationResourceGroupArns"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-capacityreservationspecification.html
//
type CfnAutoScalingGroup_CapacityReservationSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-capacityreservationspecification.html#cfn-autoscaling-autoscalinggroup-capacityreservationspecification-capacityreservationpreference
	//
	CapacityReservationPreference *string `field:"required" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-capacityreservationspecification.html#cfn-autoscaling-autoscalinggroup-capacityreservationspecification-capacityreservationtarget
	//
	CapacityReservationTarget interface{} `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

