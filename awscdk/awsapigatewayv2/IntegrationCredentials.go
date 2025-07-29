package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Credentials used for AWS Service integrations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   integrationCredentials := awscdk.Aws_apigatewayv2.IntegrationCredentials_FromRole(role)
//
type IntegrationCredentials interface {
	// The ARN of the credentials.
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


func NewIntegrationCredentials_Override(i IntegrationCredentials) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.IntegrationCredentials",
		nil, // no parameters
		i,
	)
}

// Use the specified role for integration requests.
func IntegrationCredentials_FromRole(role awsiam.IRole) IntegrationCredentials {
	_init_.Initialize()

	if err := validateIntegrationCredentials_FromRoleParameters(role); err != nil {
		panic(err)
	}
	var returns IntegrationCredentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.IntegrationCredentials",
		"fromRole",
		[]interface{}{role},
		&returns,
	)

	return returns
}

// Use the calling user's identity to call the integration.
func IntegrationCredentials_UseCallerIdentity() IntegrationCredentials {
	_init_.Initialize()

	var returns IntegrationCredentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.IntegrationCredentials",
		"useCallerIdentity",
		nil, // no parameters
		&returns,
	)

	return returns
}

