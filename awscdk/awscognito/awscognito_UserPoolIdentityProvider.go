package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// User pool third-party identity providers.
type UserPoolIdentityProvider interface {
}

// The jsii proxy struct for UserPoolIdentityProvider
type jsiiProxy_UserPoolIdentityProvider struct {
	_ byte // padding
}

// Import an existing UserPoolIdentityProvider.
func UserPoolIdentityProvider_FromProviderName(scope constructs.Construct, id *string, providerName *string) IUserPoolIdentityProvider {
	_init_.Initialize()

	if err := validateUserPoolIdentityProvider_FromProviderNameParameters(scope, id, providerName); err != nil {
		panic(err)
	}
	var returns IUserPoolIdentityProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolIdentityProvider",
		"fromProviderName",
		[]interface{}{scope, id, providerName},
		&returns,
	)

	return returns
}

