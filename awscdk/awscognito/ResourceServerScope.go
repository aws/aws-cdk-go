package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A scope for ResourceServer.
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
type ResourceServerScope interface {
	// A description of the scope.
	ScopeDescription() *string
	// The name of the scope.
	ScopeName() *string
}

// The jsii proxy struct for ResourceServerScope
type jsiiProxy_ResourceServerScope struct {
	_ byte // padding
}

func (j *jsiiProxy_ResourceServerScope) ScopeDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceServerScope) ScopeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeName",
		&returns,
	)
	return returns
}


func NewResourceServerScope(props *ResourceServerScopeProps) ResourceServerScope {
	_init_.Initialize()

	if err := validateNewResourceServerScopeParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourceServerScope{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.ResourceServerScope",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewResourceServerScope_Override(r ResourceServerScope, props *ResourceServerScopeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.ResourceServerScope",
		[]interface{}{props},
		r,
	)
}

