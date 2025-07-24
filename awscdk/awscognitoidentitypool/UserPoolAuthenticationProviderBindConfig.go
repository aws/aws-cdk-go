package awscognitoidentitypool


// Represents a UserPoolAuthenticationProvider Bind Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolAuthenticationProviderBindConfig := &UserPoolAuthenticationProviderBindConfig{
//   	ClientId: jsii.String("clientId"),
//   	ProviderName: jsii.String("providerName"),
//   	ServerSideTokenCheck: jsii.Boolean(false),
//   }
//
type UserPoolAuthenticationProviderBindConfig struct {
	// Client Id of the Associated User Pool Client.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The identity providers associated with the UserPool.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Whether to enable the identity pool's server side token check.
	ServerSideTokenCheck *bool `field:"required" json:"serverSideTokenCheck" yaml:"serverSideTokenCheck"`
}

