package awsworkspaces


// Describes a connection alias association that is used for cross-Region redirection.
//
// For more information, see [Cross-Region Redirection for Amazon WorkSpaces](https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionAliasAssociationProperty := &connectionAliasAssociationProperty{
//   	associatedAccountId: jsii.String("associatedAccountId"),
//   	associationStatus: jsii.String("associationStatus"),
//   	connectionIdentifier: jsii.String("connectionIdentifier"),
//   	resourceId: jsii.String("resourceId"),
//   }
//
type CfnConnectionAlias_ConnectionAliasAssociationProperty struct {
	// The identifier of the AWS account that associated the connection alias with a directory.
	AssociatedAccountId *string `field:"optional" json:"associatedAccountId" yaml:"associatedAccountId"`
	// The association status of the connection alias.
	AssociationStatus *string `field:"optional" json:"associationStatus" yaml:"associationStatus"`
	// The identifier of the connection alias association.
	//
	// You use the connection identifier in the DNS TXT record when you're configuring your DNS routing policies.
	ConnectionIdentifier *string `field:"optional" json:"connectionIdentifier" yaml:"connectionIdentifier"`
	// The identifier of the directory associated with a connection alias.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

