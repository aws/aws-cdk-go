package awscognito


// Properties to initialize UserPoolAmazonIdentityProvider.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   provider := cognito.NewUserPoolIdentityProviderAmazon(this, jsii.String("Amazon"), &userPoolIdentityProviderAmazonProps{
//   	userPool: pool,
//   	clientId: jsii.String("amzn-client-id"),
//   	clientSecret: jsii.String("amzn-client-secret"),
//   })
//
//   client := pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	// ...
//   	supportedIdentityProviders: []userPoolClientIdentityProvider{
//   		cognito.*userPoolClientIdentityProvider_AMAZON(),
//   	},
//   })
//
//   client.node.addDependency(provider)
//
type UserPoolIdentityProviderAmazonProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The client id recognized by 'Login with Amazon' APIs.
	// See: https://developer.amazon.com/docs/login-with-amazon/security-profile.html#client-identifier
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret to be accompanied with clientId for 'Login with Amazon' APIs to authenticate the client.
	// See: https://developer.amazon.com/docs/login-with-amazon/security-profile.html#client-identifier
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The types of user profile data to obtain for the Amazon profile.
	// See: https://developer.amazon.com/docs/login-with-amazon/customer-profile.html
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

