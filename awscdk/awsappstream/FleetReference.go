package awsappstream


// A reference to a Fleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetReference := &FleetReference{
//   	FleetId: jsii.String("fleetId"),
//   }
//
type FleetReference struct {
	// The Id of the Fleet resource.
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
}

