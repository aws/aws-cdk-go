package awscognitoidentitypool


// Represents a UserPoolAuthenticationProvider Bind Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolAuthenticationProviderBindConfig := &userPoolAuthenticationProviderBindConfig{
//   	clientId: jsii.String("clientId"),
//   	providerName: jsii.String("providerName"),
//   	serverSideTokenCheck: jsii.Boolean(false),
//   }
//
// Experimental.
type UserPoolAuthenticationProviderBindConfig struct {
	// Client Id of the Associated User Pool Client.
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The identity providers associated with the UserPool.
	// Experimental.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Whether to enable the identity pool's server side token check.
	// Experimental.
	ServerSideTokenCheck *bool `field:"required" json:"serverSideTokenCheck" yaml:"serverSideTokenCheck"`
}

