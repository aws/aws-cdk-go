package mixinsawsidentitystore


// Properties for CfnGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGroupMixinProps := &CfnGroupMixinProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	IdentityStoreId: jsii.String("identityStoreId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-group.html
//
type CfnGroupMixinProps struct {
	// A string containing the description of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-group.html#cfn-identitystore-group-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name value for the group.
	//
	// The length limit is 1,024 characters. This value can consist of letters, accented characters, symbols, numbers, punctuation, tab, new line, carriage return, space, and nonbreaking space in this attribute. This value is specified at the time the group is created and stored as an attribute of the group object in the identity store.
	//
	// Prefix search supports a maximum of 1,000 characters for the string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-group.html#cfn-identitystore-group-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The globally unique identifier for the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-group.html#cfn-identitystore-group-identitystoreid
	//
	IdentityStoreId *string `field:"optional" json:"identityStoreId" yaml:"identityStoreId"`
}

