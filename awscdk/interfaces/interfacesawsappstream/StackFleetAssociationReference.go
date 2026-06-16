package interfacesawsappstream


// A reference to a StackFleetAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackFleetAssociationReference := &StackFleetAssociationReference{
//   	FleetName: jsii.String("fleetName"),
//   	StackName: jsii.String("stackName"),
//   }
//
type StackFleetAssociationReference struct {
	// The FleetName of the StackFleetAssociation resource.
	FleetName *string `field:"required" json:"fleetName" yaml:"fleetName"`
	// The StackName of the StackFleetAssociation resource.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
}

