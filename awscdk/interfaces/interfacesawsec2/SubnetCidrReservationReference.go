package interfacesawsec2


// A reference to a SubnetCidrReservation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetCidrReservationReference := &SubnetCidrReservationReference{
//   	SubnetCidrReservationArn: jsii.String("subnetCidrReservationArn"),
//   	SubnetCidrReservationId: jsii.String("subnetCidrReservationId"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
type SubnetCidrReservationReference struct {
	// The ARN of the SubnetCidrReservation resource.
	SubnetCidrReservationArn *string `field:"required" json:"subnetCidrReservationArn" yaml:"subnetCidrReservationArn"`
	// The SubnetCidrReservationId of the SubnetCidrReservation resource.
	SubnetCidrReservationId *string `field:"required" json:"subnetCidrReservationId" yaml:"subnetCidrReservationId"`
	// The SubnetId of the SubnetCidrReservation resource.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

