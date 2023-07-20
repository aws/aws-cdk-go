package awsec2


// Specifies an instance's Capacity Reservation targeting option. You can specify only one option at a time.
//
// `CapacityReservationSpecification` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-capacityreservationspecification.html
//
type CfnLaunchTemplate_CapacityReservationSpecificationProperty struct {
	// Indicates the instance's Capacity Reservation preferences. Possible preferences include:.
	//
	// - `open` - The instance can run in any `open` Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).
	// - `none` - The instance avoids running in a Capacity Reservation even if one is available. The instance runs in On-Demand capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-capacityreservationspecification.html#cfn-ec2-launchtemplate-capacityreservationspecification-capacityreservationpreference
	//
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Information about the target Capacity Reservation or Capacity Reservation group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-capacityreservationspecification.html#cfn-ec2-launchtemplate-capacityreservationspecification-capacityreservationtarget
	//
	CapacityReservationTarget interface{} `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

