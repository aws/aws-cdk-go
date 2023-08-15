package awscognito


// OAuth settings to configure the interaction between the app and this client.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   readOnlyScope := cognito.NewResourceServerScope(&ResourceServerScopeProps{
//   	ScopeName: jsii.String("read"),
//   	ScopeDescription: jsii.String("Read-only access"),
//   })
//   fullAccessScope := cognito.NewResourceServerScope(&ResourceServerScopeProps{
//   	ScopeName: jsii.String("*"),
//   	ScopeDescription: jsii.String("Full access"),
//   })
//
//   userServer := pool.addResourceServer(jsii.String("ResourceServer"), &UserPoolResourceServerOptions{
//   	Identifier: jsii.String("users"),
//   	Scopes: []resourceServerScope{
//   		readOnlyScope,
//   		fullAccessScope,
//   	},
//   })
//
//   readOnlyClient := pool.addClient(jsii.String("read-only-client"), &UserPoolClientOptions{
//   	// ...
//   	OAuth: &OAuthSettings{
//   		// ...
//   		Scopes: []oAuthScope{
//   			cognito.*oAuthScope_ResourceServer(userServer, readOnlyScope),
//   		},
//   	},
//   })
//
//   fullAccessClient := pool.addClient(jsii.String("full-access-client"), &UserPoolClientOptions{
//   	// ...
//   	OAuth: &OAuthSettings{
//   		// ...
//   		Scopes: []*oAuthScope{
//   			cognito.*oAuthScope_*ResourceServer(userServer, fullAccessScope),
//   		},
//   	},
//   })
//
type OAuthSettings struct {
	// List of allowed redirect URLs for the identity providers.
	// Default: - ['https://example.com'] if either authorizationCodeGrant or implicitCodeGrant flows are enabled, no callback URLs otherwise.
	//
	CallbackUrls *[]*string `field:"optional" json:"callbackUrls" yaml:"callbackUrls"`
	// OAuth flows that are allowed with this client.
	// See:  - the 'Allowed OAuth Flows' section at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
	//
	// Default: {authorizationCodeGrant:true,implicitCodeGrant:true}.
	//
	Flows *OAuthFlows `field:"optional" json:"flows" yaml:"flows"`
	// List of allowed logout URLs for the identity providers.
	// Default: - no logout URLs.
	//
	LogoutUrls *[]*string `field:"optional" json:"logoutUrls" yaml:"logoutUrls"`
	// OAuth scopes that are allowed with this client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
	//
	// Default: [OAuthScope.PHONE,OAuthScope.EMAIL,OAuthScope.OPENID,OAuthScope.PROFILE,OAuthScope.COGNITO_ADMIN]
	//
	Scopes *[]OAuthScope `field:"optional" json:"scopes" yaml:"scopes"`
}

