package awsbedrockagentcore


// Information about the target Capacity Reservation or Capacity Reservation group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityReservationTargetProperty := &CapacityReservationTargetProperty{
//   	CapacityReservationId: jsii.String("capacityReservationId"),
//   	CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-capacityreservationtarget.html
//
type CfnCapacityProvider_CapacityReservationTargetProperty struct {
	// The ID of the Capacity Reservation in which to run the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-capacityreservationtarget.html#cfn-bedrockagentcore-capacityprovider-capacityreservationtarget-capacityreservationid
	//
	CapacityReservationId *string `field:"optional" json:"capacityReservationId" yaml:"capacityReservationId"`
	// The ARN of the Capacity Reservation resource group in which to run the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-capacityreservationtarget.html#cfn-bedrockagentcore-capacityprovider-capacityreservationtarget-capacityreservationresourcegrouparn
	//
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
}

