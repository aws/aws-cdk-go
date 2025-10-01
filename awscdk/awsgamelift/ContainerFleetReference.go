package awsgamelift


// A reference to a ContainerFleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerFleetReference := &ContainerFleetReference{
//   	FleetId: jsii.String("fleetId"),
//   }
//
type ContainerFleetReference struct {
	// The FleetId of the ContainerFleet resource.
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
}

