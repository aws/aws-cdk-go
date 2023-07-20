package awsidentitystore


// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupProps := &CfnGroupProps{
//   	DisplayName: jsii.String("displayName"),
//   	IdentityStoreId: jsii.String("identityStoreId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-group.html
//
type CfnGroupProps struct {
	// A string containing the name of the group.
	//
	// This value is commonly displayed when the group is referenced.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-group.html#cfn-identitystore-group-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The globally unique identifier for the identity store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-group.html#cfn-identitystore-group-identitystoreid
	//
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// A string containing the description of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-group.html#cfn-identitystore-group-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

