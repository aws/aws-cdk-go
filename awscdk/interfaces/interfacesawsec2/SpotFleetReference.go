package interfacesawsec2


// A reference to a SpotFleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotFleetReference := &SpotFleetReference{
//   	SpotFleetId: jsii.String("spotFleetId"),
//   }
//
type SpotFleetReference struct {
	// The Id of the SpotFleet resource.
	SpotFleetId *string `field:"required" json:"spotFleetId" yaml:"spotFleetId"`
}

