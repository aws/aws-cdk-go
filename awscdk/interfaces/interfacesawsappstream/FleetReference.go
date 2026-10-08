package interfacesawsappstream


// A reference to a Fleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetReference := &FleetReference{
//   	FleetArn: jsii.String("fleetArn"),
//   	FleetName: jsii.String("fleetName"),
//   }
//
type FleetReference struct {
	// The ARN of the Fleet resource.
	FleetArn *string `field:"required" json:"fleetArn" yaml:"fleetArn"`
	// The Name of the Fleet resource.
	FleetName *string `field:"required" json:"fleetName" yaml:"fleetName"`
}

