package awscdkapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Credentials used for AWS Service integrations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   integrationCredentials := apigatewayv2_alpha.IntegrationCredentials_FromRole(role)
//
// Deprecated.
type IntegrationCredentials interface {
	// The ARN of the credentials.
	// Deprecated.
	CredentialsArn() *string
}

// The jsii proxy struct for IntegrationCredentials
type jsiiProxy_IntegrationCredentials struct {
	_ byte // padding
}

func (j *jsiiProxy_IntegrationCredentials) CredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsArn",
		&returns,
	)
	return returns
}


// Deprecated.
func NewIntegrationCredentials_Override(i IntegrationCredentials) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.IntegrationCredentials",
		nil, // no parameters
		i,
	)
}

// Use the specified role for integration requests.
// Deprecated.
func IntegrationCredentials_FromRole(role awsiam.IRole) IntegrationCredentials {
	_init_.Initialize()

	if err := validateIntegrationCredentials_FromRoleParameters(role); err != nil {
		panic(err)
	}
	var returns IntegrationCredentials

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.IntegrationCredentials",
		"fromRole",
		[]interface{}{role},
		&returns,
	)

	return returns
}

// Use the calling user's identity to call the integration.
// Deprecated.
func IntegrationCredentials_UseCallerIdentity() IntegrationCredentials {
	_init_.Initialize()

	var returns IntegrationCredentials

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.IntegrationCredentials",
		"useCallerIdentity",
		nil, // no parameters
		&returns,
	)

	return returns
}

