package awsworkspaces


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionAliasAssociationProperty := &ConnectionAliasAssociationProperty{
//   	AssociatedAccountId: jsii.String("associatedAccountId"),
//   	AssociationStatus: jsii.String("associationStatus"),
//   	ConnectionIdentifier: jsii.String("connectionIdentifier"),
//   	ResourceId: jsii.String("resourceId"),
//   }
//
type CfnConnectionAlias_ConnectionAliasAssociationProperty struct {
	// `CfnConnectionAlias.ConnectionAliasAssociationProperty.AssociatedAccountId`.
	AssociatedAccountId *string `field:"optional" json:"associatedAccountId" yaml:"associatedAccountId"`
	// `CfnConnectionAlias.ConnectionAliasAssociationProperty.AssociationStatus`.
	AssociationStatus *string `field:"optional" json:"associationStatus" yaml:"associationStatus"`
	// `CfnConnectionAlias.ConnectionAliasAssociationProperty.ConnectionIdentifier`.
	ConnectionIdentifier *string `field:"optional" json:"connectionIdentifier" yaml:"connectionIdentifier"`
	// `CfnConnectionAlias.ConnectionAliasAssociationProperty.ResourceId`.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

