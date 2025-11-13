package interfacesawsec2


// A reference to a CapacityReservationFleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityReservationFleetReference := &CapacityReservationFleetReference{
//   	CapacityReservationFleetId: jsii.String("capacityReservationFleetId"),
//   }
//
type CapacityReservationFleetReference struct {
	// The CapacityReservationFleetId of the CapacityReservationFleet resource.
	CapacityReservationFleetId *string `field:"required" json:"capacityReservationFleetId" yaml:"capacityReservationFleetId"`
}

