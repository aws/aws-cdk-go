package interfacesawscodebuild


// A reference to a Fleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetReference := &FleetReference{
//   	FleetArn: jsii.String("fleetArn"),
//   }
//
type FleetReference struct {
	// The Arn of the Fleet resource.
	FleetArn *string `field:"required" json:"fleetArn" yaml:"fleetArn"`
}

