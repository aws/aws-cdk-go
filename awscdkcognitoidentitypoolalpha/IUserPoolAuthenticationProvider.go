// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the concept of a User Pool Authentication Provider.
//
// You use user pool authentication providers to configure User Pools
// and User Pool Clients for use with Identity Pools.
// Experimental.
type IUserPoolAuthenticationProvider interface {
	// The method called when a given User Pool Authentication Provider is added (for the first time) to an Identity Pool.
	// Experimental.
	Bind(scope constructs.Construct, identityPool IIdentityPool, options *UserPoolAuthenticationProviderBindOptions) *UserPoolAuthenticationProviderBindConfig
}

// The jsii proxy for IUserPoolAuthenticationProvider
type jsiiProxy_IUserPoolAuthenticationProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_IUserPoolAuthenticationProvider) Bind(scope constructs.Construct, identityPool IIdentityPool, options *UserPoolAuthenticationProviderBindOptions) *UserPoolAuthenticationProviderBindConfig {
	if err := i.validateBindParameters(scope, identityPool, options); err != nil {
		panic(err)
	}
	var returns *UserPoolAuthenticationProviderBindConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, identityPool, options},
		&returns,
	)

	return returns
}

