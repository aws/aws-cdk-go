package awscdkcognitoidentitypoolalpha


// Represents a UserPoolAuthenticationProvider Bind Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cognito_identitypool_alpha "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
//
//   userPoolAuthenticationProviderBindConfig := &UserPoolAuthenticationProviderBindConfig{
//   	ClientId: jsii.String("clientId"),
//   	ProviderName: jsii.String("providerName"),
//   	ServerSideTokenCheck: jsii.Boolean(false),
//   }
//
// Deprecated.
type UserPoolAuthenticationProviderBindConfig struct {
	// Client Id of the Associated User Pool Client.
	// Deprecated.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The identity providers associated with the UserPool.
	// Deprecated.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Whether to enable the identity pool's server side token check.
	// Deprecated.
	ServerSideTokenCheck *bool `field:"required" json:"serverSideTokenCheck" yaml:"serverSideTokenCheck"`
}

