package awscognitoidentitypool

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

// Defines a User Pool Authentication Provider.
//
// Example:
//   var identityPool identityPool
//
//   userPool := cognito.NewUserPool(this, jsii.String("Pool"))
//   identityPool.AddUserPoolAuthentication(awscdk.NewUserPoolAuthenticationProvider(&UserPoolAuthenticationProviderProps{
//   	UserPool: UserPool,
//   	DisableServerSideTokenCheck: jsii.Boolean(true),
//   }))
//
// Experimental.
type UserPoolAuthenticationProvider interface {
	IUserPoolAuthenticationProvider
	// The method called when a given User Pool Authentication Provider is added (for the first time) to an Identity Pool.
	// Experimental.
	Bind(scope constructs.Construct, identityPool IIdentityPool, _options *UserPoolAuthenticationProviderBindOptions) *UserPoolAuthenticationProviderBindConfig
}

// The jsii proxy struct for UserPoolAuthenticationProvider
type jsiiProxy_UserPoolAuthenticationProvider struct {
	jsiiProxy_IUserPoolAuthenticationProvider
}

// Experimental.
func NewUserPoolAuthenticationProvider(props *UserPoolAuthenticationProviderProps) UserPoolAuthenticationProvider {
	_init_.Initialize()

	if err := validateNewUserPoolAuthenticationProviderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserPoolAuthenticationProvider{}

	_jsii_.Create(
		"monocdk.aws_cognito_identitypool.UserPoolAuthenticationProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPoolAuthenticationProvider_Override(u UserPoolAuthenticationProvider, props *UserPoolAuthenticationProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cognito_identitypool.UserPoolAuthenticationProvider",
		[]interface{}{props},
		u,
	)
}

func (u *jsiiProxy_UserPoolAuthenticationProvider) Bind(scope constructs.Construct, identityPool IIdentityPool, _options *UserPoolAuthenticationProviderBindOptions) *UserPoolAuthenticationProviderBindConfig {
	if err := u.validateBindParameters(scope, identityPool, _options); err != nil {
		panic(err)
	}
	var returns *UserPoolAuthenticationProviderBindConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{scope, identityPool, _options},
		&returns,
	)

	return returns
}

