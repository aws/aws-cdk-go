package awscognito


// A reference to a IdentityPoolPrincipalTag resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityPoolPrincipalTagReference := &IdentityPoolPrincipalTagReference{
//   	IdentityPoolId: jsii.String("identityPoolId"),
//   	IdentityProviderName: jsii.String("identityProviderName"),
//   }
//
type IdentityPoolPrincipalTagReference struct {
	// The IdentityPoolId of the IdentityPoolPrincipalTag resource.
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// The IdentityProviderName of the IdentityPoolPrincipalTag resource.
	IdentityProviderName *string `field:"required" json:"identityProviderName" yaml:"identityProviderName"`
}

