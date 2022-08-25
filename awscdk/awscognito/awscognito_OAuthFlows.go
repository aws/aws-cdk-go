package awscognito


// Types of OAuth grant flows.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   pool.addClient(jsii.String("app-client"), &userPoolClientOptions{
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			authorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope_OPENID(),
//   		},
//   		callbackUrls: []*string{
//   			jsii.String("https://my-app-domain.com/welcome"),
//   		},
//   		logoutUrls: []*string{
//   			jsii.String("https://my-app-domain.com/signin"),
//   		},
//   	},
//   })
//
// See: - the 'Allowed OAuth Flows' section at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
//
type OAuthFlows struct {
	// Initiate an authorization code grant flow, which provides an authorization code as the response.
	AuthorizationCodeGrant *bool `field:"optional" json:"authorizationCodeGrant" yaml:"authorizationCodeGrant"`
	// Client should get the access token and ID token from the token endpoint using a combination of client and client_secret.
	ClientCredentials *bool `field:"optional" json:"clientCredentials" yaml:"clientCredentials"`
	// The client should get the access token and ID token directly.
	ImplicitCodeGrant *bool `field:"optional" json:"implicitCodeGrant" yaml:"implicitCodeGrant"`
}

