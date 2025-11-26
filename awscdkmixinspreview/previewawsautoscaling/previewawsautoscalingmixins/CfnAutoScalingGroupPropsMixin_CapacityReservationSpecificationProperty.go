package previewawsautoscalingmixins


// Describes the Capacity Reservation preference and targeting options.
//
// If you specify `open` or `none` for `CapacityReservationPreference` , do not specify a `CapacityReservationTarget` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   capacityReservationSpecificationProperty := &CapacityReservationSpecificationProperty{
//   	CapacityReservationPreference: jsii.String("capacityReservationPreference"),
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
type CfnAutoScalingGroupPropsMixin_CapacityReservationSpecificationProperty struct {
	// The capacity reservation preference. The following options are available:.
	//
	// - `capacity-reservations-only` - Auto Scaling will only launch instances into a Capacity Reservation or Capacity Reservation resource group. If capacity isn't available, instances will fail to launch.
	// - `capacity-reservations-first` - Auto Scaling will try to launch instances into a Capacity Reservation or Capacity Reservation resource group first. If capacity isn't available, instances will run in On-Demand capacity.
	// - `none` - Auto Scaling will not launch instances into a Capacity Reservation. Instances will run in On-Demand capacity.
	// - `default` - Auto Scaling uses the Capacity Reservation preference from your launch template or an open Capacity Reservation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-capacityreservationspecification.html#cfn-autoscaling-autoscalinggroup-capacityreservationspecification-capacityreservationpreference
	//
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Describes a target Capacity Reservation or Capacity Reservation resource group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-capacityreservationspecification.html#cfn-autoscaling-autoscalinggroup-capacityreservationspecification-capacityreservationtarget
	//
	CapacityReservationTarget interface{} `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

