package awscognito


// Options to create a UserPoolResourceServer.
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
type UserPoolResourceServerOptions struct {
	// A unique resource server identifier for the resource server.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Oauth scopes.
	// Default: - No scopes will be added.
	//
	Scopes *[]ResourceServerScope `field:"optional" json:"scopes" yaml:"scopes"`
	// A friendly name for the resource server.
	// Default: - same as `identifier`.
	//
	UserPoolResourceServerName *string `field:"optional" json:"userPoolResourceServerName" yaml:"userPoolResourceServerName"`
}

