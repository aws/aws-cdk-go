package awsverifiedpermissions


// A reference to a IdentitySource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identitySourceReference := &IdentitySourceReference{
//   	IdentitySourceId: jsii.String("identitySourceId"),
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   }
//
type IdentitySourceReference struct {
	// The IdentitySourceId of the IdentitySource resource.
	IdentitySourceId *string `field:"required" json:"identitySourceId" yaml:"identitySourceId"`
	// The PolicyStoreId of the IdentitySource resource.
	PolicyStoreId *string `field:"required" json:"policyStoreId" yaml:"policyStoreId"`
}

