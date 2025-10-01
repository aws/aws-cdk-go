package awsappstream


// A reference to a ApplicationFleetAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationFleetAssociationReference := &ApplicationFleetAssociationReference{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	FleetName: jsii.String("fleetName"),
//   }
//
type ApplicationFleetAssociationReference struct {
	// The ApplicationArn of the ApplicationFleetAssociation resource.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The FleetName of the ApplicationFleetAssociation resource.
	FleetName *string `field:"required" json:"fleetName" yaml:"fleetName"`
}

