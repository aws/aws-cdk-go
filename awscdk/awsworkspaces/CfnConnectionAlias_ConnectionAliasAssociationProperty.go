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
//   connectionAliasAssociationProperty := &ConnectionAliasAssociationProperty{
//   	AssociatedAccountId: jsii.String("associatedAccountId"),
//   	AssociationStatus: jsii.String("associationStatus"),
//   	ConnectionIdentifier: jsii.String("connectionIdentifier"),
//   	ResourceId: jsii.String("resourceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-connectionalias-connectionaliasassociation.html
//
type CfnConnectionAlias_ConnectionAliasAssociationProperty struct {
	// The identifier of the AWS account that associated the connection alias with a directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-connectionalias-connectionaliasassociation.html#cfn-workspaces-connectionalias-connectionaliasassociation-associatedaccountid
	//
	AssociatedAccountId *string `field:"optional" json:"associatedAccountId" yaml:"associatedAccountId"`
	// The association status of the connection alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-connectionalias-connectionaliasassociation.html#cfn-workspaces-connectionalias-connectionaliasassociation-associationstatus
	//
	AssociationStatus *string `field:"optional" json:"associationStatus" yaml:"associationStatus"`
	// The identifier of the connection alias association.
	//
	// You use the connection identifier in the DNS TXT record when you're configuring your DNS routing policies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-connectionalias-connectionaliasassociation.html#cfn-workspaces-connectionalias-connectionaliasassociation-connectionidentifier
	//
	ConnectionIdentifier *string `field:"optional" json:"connectionIdentifier" yaml:"connectionIdentifier"`
	// The identifier of the directory associated with a connection alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-connectionalias-connectionaliasassociation.html#cfn-workspaces-connectionalias-connectionaliasassociation-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

