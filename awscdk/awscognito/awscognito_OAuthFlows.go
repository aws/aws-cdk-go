package awscognito


// Types of OAuth grant flows.
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
// See: - the 'Allowed OAuth Flows' section at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
//
// Experimental.
type OAuthFlows struct {
	// Initiate an authorization code grant flow, which provides an authorization code as the response.
	// Experimental.
	AuthorizationCodeGrant *bool `field:"optional" json:"authorizationCodeGrant" yaml:"authorizationCodeGrant"`
	// Client should get the access token and ID token from the token endpoint using a combination of client and client_secret.
	// Experimental.
	ClientCredentials *bool `field:"optional" json:"clientCredentials" yaml:"clientCredentials"`
	// The client should get the access token and ID token directly.
	// Experimental.
	ImplicitCodeGrant *bool `field:"optional" json:"implicitCodeGrant" yaml:"implicitCodeGrant"`
}

