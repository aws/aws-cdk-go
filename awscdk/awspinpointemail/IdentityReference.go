package awspinpointemail


// A reference to a Identity resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityReference := &IdentityReference{
//   	IdentityId: jsii.String("identityId"),
//   }
//
type IdentityReference struct {
	// The Id of the Identity resource.
	IdentityId *string `field:"required" json:"identityId" yaml:"identityId"`
}

