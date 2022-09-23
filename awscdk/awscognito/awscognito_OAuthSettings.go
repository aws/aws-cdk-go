package awscognito


// OAuth settings to configure the interaction between the app and this client.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("UserPool"), &userPoolProps{
//   })
//   client := userpool.addClient(jsii.String("Client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			implicitCodeGrant: jsii.Boolean(true),
//   		},
//   		callbackUrls: []*string{
//   			jsii.String("https://myapp.com/home"),
//   			jsii.String("https://myapp.com/users"),
//   		},
//   	},
//   })
//   domain := userpool.addDomain(jsii.String("Domain"), &userPoolDomainOptions{
//   })
//   signInUrl := domain.signInUrl(client, &signInUrlOptions{
//   	redirectUri: jsii.String("https://myapp.com/home"),
//   })
//
type OAuthSettings struct {
	// List of allowed redirect URLs for the identity providers.
	CallbackUrls *[]*string `field:"optional" json:"callbackUrls" yaml:"callbackUrls"`
	// OAuth flows that are allowed with this client.
	// See: - the 'Allowed OAuth Flows' section at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
	//
	Flows *OAuthFlows `field:"optional" json:"flows" yaml:"flows"`
	// List of allowed logout URLs for the identity providers.
	LogoutUrls *[]*string `field:"optional" json:"logoutUrls" yaml:"logoutUrls"`
	// OAuth scopes that are allowed with this client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
	//
	Scopes *[]OAuthScope `field:"optional" json:"scopes" yaml:"scopes"`
}

