package awsec2


// Specifies a target Capacity Reservation.
//
// `CapacityReservationTarget` is a property of the [Amazon EC2 LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityReservationTargetProperty := &capacityReservationTargetProperty{
//   	capacityReservationId: jsii.String("capacityReservationId"),
//   	capacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   }
//
type CfnLaunchTemplate_CapacityReservationTargetProperty struct {
	// The ID of the Capacity Reservation in which to run the instance.
	CapacityReservationId *string `field:"optional" json:"capacityReservationId" yaml:"capacityReservationId"`
	// The ARN of the Capacity Reservation resource group in which to run the instance.
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
}

