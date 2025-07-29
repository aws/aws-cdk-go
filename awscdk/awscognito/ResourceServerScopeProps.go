package awscognito


// Props to initialize ResourceServerScope.
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
type ResourceServerScopeProps struct {
	// A description of the scope.
	ScopeDescription *string `field:"required" json:"scopeDescription" yaml:"scopeDescription"`
	// The name of the scope.
	ScopeName *string `field:"required" json:"scopeName" yaml:"scopeName"`
}

