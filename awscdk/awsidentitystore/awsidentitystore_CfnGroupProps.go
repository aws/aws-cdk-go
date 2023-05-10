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
type CfnGroupProps struct {
	// `AWS::IdentityStore::Group.DisplayName`.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// `AWS::IdentityStore::Group.IdentityStoreId`.
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// A string containing the description of the group.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

