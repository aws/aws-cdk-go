package awscognito


// A reference to a UserPoolIdentityProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolIdentityProviderReference := &UserPoolIdentityProviderReference{
//   	ProviderName: jsii.String("providerName"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolIdentityProviderReference struct {
	// The ProviderName of the UserPoolIdentityProvider resource.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The UserPoolId of the UserPoolIdentityProvider resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

