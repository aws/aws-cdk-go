package awscognito


// OAuth settings to configure the interaction between the app and this client.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   readOnlyScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("read"),
//   	scopeDescription: jsii.String("Read-only access"),
//   })
//   fullAccessScope := cognito.NewResourceServerScope(&resourceServerScopeProps{
//   	scopeName: jsii.String("*"),
//   	scopeDescription: jsii.String("Full access"),
//   })
//
//   userServer := pool.addResourceServer(jsii.String("ResourceServer"), &userPoolResourceServerOptions{
//   	identifier: jsii.String("users"),
//   	scopes: []resourceServerScope{
//   		readOnlyScope,
//   		fullAccessScope,
//   	},
//   })
//
//   readOnlyClient := pool.addClient(jsii.String("read-only-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, readOnlyScope),
//   		},
//   	},
//   })
//
//   fullAccessClient := pool.addClient(jsii.String("full-access-client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		// ...
//   		scopes: []*oAuthScope{
//   			cognito.*oAuthScope.resourceServer(userServer, fullAccessScope),
//   		},
//   	},
//   })
//
// Experimental.
type OAuthSettings struct {
	// List of allowed redirect URLs for the identity providers.
	// Experimental.
	CallbackUrls *[]*string `field:"optional" json:"callbackUrls" yaml:"callbackUrls"`
	// OAuth flows that are allowed with this client.
	// See: - the 'Allowed OAuth Flows' section at https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
	//
	// Experimental.
	Flows *OAuthFlows `field:"optional" json:"flows" yaml:"flows"`
	// List of allowed logout URLs for the identity providers.
	// Experimental.
	LogoutUrls *[]*string `field:"optional" json:"logoutUrls" yaml:"logoutUrls"`
	// OAuth scopes that are allowed with this client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
	//
	// Experimental.
	Scopes *[]OAuthScope `field:"optional" json:"scopes" yaml:"scopes"`
}

