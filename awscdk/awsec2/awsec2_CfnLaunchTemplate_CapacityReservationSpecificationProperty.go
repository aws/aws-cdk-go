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
//   capacityReservationSpecificationProperty := &capacityReservationSpecificationProperty{
//   	capacityReservationPreference: jsii.String("capacityReservationPreference"),
//   	capacityReservationTarget: &capacityReservationTargetProperty{
//   		capacityReservationId: jsii.String("capacityReservationId"),
//   		capacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   	},
//   }
//
type CfnLaunchTemplate_CapacityReservationSpecificationProperty struct {
	// Indicates the instance's Capacity Reservation preferences. Possible preferences include:.
	//
	// - `open` - The instance can run in any `open` Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).
	// - `none` - The instance avoids running in a Capacity Reservation even if one is available. The instance runs in On-Demand capacity.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Information about the target Capacity Reservation or Capacity Reservation group.
	CapacityReservationTarget interface{} `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

