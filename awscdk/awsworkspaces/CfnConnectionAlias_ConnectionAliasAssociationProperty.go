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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-connectionalias-connectionaliasassociation.html
//
type CfnConnectionAlias_ConnectionAliasAssociationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-connectionalias-connectionaliasassociation.html#cfn-workspaces-connectionalias-connectionaliasassociation-associatedaccountid
	//
	AssociatedAccountId *string `field:"optional" json:"associatedAccountId" yaml:"associatedAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-connectionalias-connectionaliasassociation.html#cfn-workspaces-connectionalias-connectionaliasassociation-associationstatus
	//
	AssociationStatus *string `field:"optional" json:"associationStatus" yaml:"associationStatus"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-connectionalias-connectionaliasassociation.html#cfn-workspaces-connectionalias-connectionaliasassociation-connectionidentifier
	//
	ConnectionIdentifier *string `field:"optional" json:"connectionIdentifier" yaml:"connectionIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-connectionalias-connectionaliasassociation.html#cfn-workspaces-connectionalias-connectionaliasassociation-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

