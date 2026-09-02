package awsbedrockagentcore


// The Capacity Reservation targeting option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityReservationSpecificationProperty := &CapacityReservationSpecificationProperty{
//   	CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   	CapacityReservationTarget: &CapacityReservationTargetProperty{
//   		CapacityReservationId: jsii.String("capacityReservationId"),
//   		CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-capacityreservationspecification.html
//
type CfnCapacityProvider_CapacityReservationSpecificationProperty struct {
	// Indicates the instance's Capacity Reservation preferences.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-capacityreservationspecification.html#cfn-bedrockagentcore-capacityprovider-capacityreservationspecification-capacityreservationpreference
	//
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Information about the target Capacity Reservation or Capacity Reservation group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-capacityreservationspecification.html#cfn-bedrockagentcore-capacityprovider-capacityreservationspecification-capacityreservationtarget
	//
	CapacityReservationTarget interface{} `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

