package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// OAuth scopes that are allowed with this client.
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
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-idp-settings.html
//
type OAuthScope interface {
	// The name of this scope as recognized by CloudFormation.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-allowedoauthscopes
	//
	ScopeName() *string
}

// The jsii proxy struct for OAuthScope
type jsiiProxy_OAuthScope struct {
	_ byte // padding
}

func (j *jsiiProxy_OAuthScope) ScopeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeName",
		&returns,
	)
	return returns
}


// Custom scope is one that you define for your own resource server in the Resource Servers.
//
// The format is 'resource-server-identifier/scope'.
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-define-resource-servers.html
//
func OAuthScope_Custom(name *string) OAuthScope {
	_init_.Initialize()

	if err := validateOAuthScope_CustomParameters(name); err != nil {
		panic(err)
	}
	var returns OAuthScope

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.OAuthScope",
		"custom",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Adds a custom scope that's tied to a resource server in your stack.
func OAuthScope_ResourceServer(server IUserPoolResourceServer, scope ResourceServerScope) OAuthScope {
	_init_.Initialize()

	if err := validateOAuthScope_ResourceServerParameters(server, scope); err != nil {
		panic(err)
	}
	var returns OAuthScope

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.OAuthScope",
		"resourceServer",
		[]interface{}{server, scope},
		&returns,
	)

	return returns
}

func OAuthScope_COGNITO_ADMIN() OAuthScope {
	_init_.Initialize()
	var returns OAuthScope
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.OAuthScope",
		"COGNITO_ADMIN",
		&returns,
	)
	return returns
}

func OAuthScope_EMAIL() OAuthScope {
	_init_.Initialize()
	var returns OAuthScope
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.OAuthScope",
		"EMAIL",
		&returns,
	)
	return returns
}

func OAuthScope_OPENID() OAuthScope {
	_init_.Initialize()
	var returns OAuthScope
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.OAuthScope",
		"OPENID",
		&returns,
	)
	return returns
}

func OAuthScope_PHONE() OAuthScope {
	_init_.Initialize()
	var returns OAuthScope
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.OAuthScope",
		"PHONE",
		&returns,
	)
	return returns
}

func OAuthScope_PROFILE() OAuthScope {
	_init_.Initialize()
	var returns OAuthScope
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.OAuthScope",
		"PROFILE",
		&returns,
	)
	return returns
}

