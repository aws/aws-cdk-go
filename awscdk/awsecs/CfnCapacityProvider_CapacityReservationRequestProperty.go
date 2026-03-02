package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityReservationRequestProperty := &CapacityReservationRequestProperty{
//   	ReservationGroupArn: jsii.String("reservationGroupArn"),
//   	ReservationPreference: jsii.String("reservationPreference"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-capacityreservationrequest.html
//
type CfnCapacityProvider_CapacityReservationRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-capacityreservationrequest.html#cfn-ecs-capacityprovider-capacityreservationrequest-reservationgrouparn
	//
	ReservationGroupArn *string `field:"optional" json:"reservationGroupArn" yaml:"reservationGroupArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-capacityreservationrequest.html#cfn-ecs-capacityprovider-capacityreservationrequest-reservationpreference
	//
	ReservationPreference *string `field:"optional" json:"reservationPreference" yaml:"reservationPreference"`
}

