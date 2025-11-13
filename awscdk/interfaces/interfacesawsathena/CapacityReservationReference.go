package interfacesawsathena


// A reference to a CapacityReservation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityReservationReference := &CapacityReservationReference{
//   	CapacityReservationArn: jsii.String("capacityReservationArn"),
//   }
//
type CapacityReservationReference struct {
	// The Arn of the CapacityReservation resource.
	CapacityReservationArn *string `field:"required" json:"capacityReservationArn" yaml:"capacityReservationArn"`
}

