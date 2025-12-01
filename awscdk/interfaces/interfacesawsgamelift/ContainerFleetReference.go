package interfacesawsgamelift


// A reference to a ContainerFleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerFleetReference := &ContainerFleetReference{
//   	FleetArn: jsii.String("fleetArn"),
//   	FleetId: jsii.String("fleetId"),
//   }
//
type ContainerFleetReference struct {
	// The ARN of the ContainerFleet resource.
	FleetArn *string `field:"required" json:"fleetArn" yaml:"fleetArn"`
	// The FleetId of the ContainerFleet resource.
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
}

