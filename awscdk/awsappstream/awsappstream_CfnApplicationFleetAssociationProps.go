package awsappstream


// Properties for defining a `CfnApplicationFleetAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationFleetAssociationProps := &cfnApplicationFleetAssociationProps{
//   	applicationArn: jsii.String("applicationArn"),
//   	fleetName: jsii.String("fleetName"),
//   }
//
type CfnApplicationFleetAssociationProps struct {
	// The ARN of the application.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The name of the fleet.
	FleetName *string `field:"required" json:"fleetName" yaml:"fleetName"`
}

