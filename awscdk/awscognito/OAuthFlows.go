package awscognito


// Types of OAuth grant flows.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("UserPool"), &UserPoolProps{
//   })
//   client := userpool.addClient(jsii.String("Client"), &UserPoolClientOptions{
//   	// ...
//   	OAuth: &OAuthSettings{
//   		Flows: &OAuthFlows{
//   			ImplicitCodeGrant: jsii.Boolean(true),
//   		},
//   		CallbackUrls: []*string{
//   			jsii.String("https://myapp.com/home"),
//   			jsii.String("https://myapp.com/users"),
//   		},
//   	},
//   })
//   domain := userpool.addDomain(jsii.String("Domain"), &UserPoolDomainOptions{
//   })
//   signInUrl := domain.SignInUrl(client, &SignInUrlOptions{
//   	RedirectUri: jsii.String("https://myapp.com/home"),
//   })
//
// See:  - the 'Allowed OAuth Flows' section at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
//
type OAuthFlows struct {
	// Initiate an authorization code grant flow, which provides an authorization code as the response.
	// Default: false.
	//
	AuthorizationCodeGrant *bool `field:"optional" json:"authorizationCodeGrant" yaml:"authorizationCodeGrant"`
	// Client should get the access token and ID token from the token endpoint using a combination of client and client_secret.
	// Default: false.
	//
	ClientCredentials *bool `field:"optional" json:"clientCredentials" yaml:"clientCredentials"`
	// The client should get the access token and ID token directly.
	// Default: false.
	//
	ImplicitCodeGrant *bool `field:"optional" json:"implicitCodeGrant" yaml:"implicitCodeGrant"`
}

