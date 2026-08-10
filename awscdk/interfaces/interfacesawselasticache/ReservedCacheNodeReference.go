package interfacesawselasticache


// A reference to a ReservedCacheNode resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reservedCacheNodeReference := &ReservedCacheNodeReference{
//   	ReservationArn: jsii.String("reservationArn"),
//   }
//
type ReservedCacheNodeReference struct {
	// The ReservationARN of the ReservedCacheNode resource.
	ReservationArn *string `field:"required" json:"reservationArn" yaml:"reservationArn"`
}

