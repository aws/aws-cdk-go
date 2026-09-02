package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   capacityReservationsProperty := &CapacityReservationsProperty{
//   	ReservationGroupArn: jsii.String("reservationGroupArn"),
//   	ReservationPreference: jsii.String("reservationPreference"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-capacityreservations.html
//
type CfnComputeEnvironmentPropsMixin_CapacityReservationsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-capacityreservations.html#cfn-batch-computeenvironment-capacityreservations-reservationgrouparn
	//
	ReservationGroupArn *string `field:"optional" json:"reservationGroupArn" yaml:"reservationGroupArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-capacityreservations.html#cfn-batch-computeenvironment-capacityreservations-reservationpreference
	//
	ReservationPreference *string `field:"optional" json:"reservationPreference" yaml:"reservationPreference"`
}

