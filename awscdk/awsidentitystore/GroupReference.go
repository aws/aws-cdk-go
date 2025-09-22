package awsidentitystore


// A reference to a Group resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupReference := &GroupReference{
//   	GroupId: jsii.String("groupId"),
//   	IdentityStoreId: jsii.String("identityStoreId"),
//   }
//
type GroupReference struct {
	// The GroupId of the Group resource.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The IdentityStoreId of the Group resource.
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
}

