package awscognito


// Options to create a UserPoolResourceServer.
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
type UserPoolResourceServerOptions struct {
	// A unique resource server identifier for the resource server.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Oauth scopes.
	Scopes *[]ResourceServerScope `field:"optional" json:"scopes" yaml:"scopes"`
	// A friendly name for the resource server.
	UserPoolResourceServerName *string `field:"optional" json:"userPoolResourceServerName" yaml:"userPoolResourceServerName"`
}

