package interfacesawsiotfleetwise


// A reference to a Fleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetReference := &FleetReference{
//   	FleetArn: jsii.String("fleetArn"),
//   	FleetId: jsii.String("fleetId"),
//   }
//
type FleetReference struct {
	// The ARN of the Fleet resource.
	FleetArn *string `field:"required" json:"fleetArn" yaml:"fleetArn"`
	// The Id of the Fleet resource.
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
}

