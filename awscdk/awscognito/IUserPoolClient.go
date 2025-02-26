package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
)

// Represents a Cognito user pool client.
type IUserPoolClient interface {
	awscdk.IResource
	// Name of the application client.
	UserPoolClientId() *string
	// The generated client secret.
	//
	// Only available if the "generateSecret" props is set to true.
	UserPoolClientSecret() awscdk.SecretValue
}

// The jsii proxy for IUserPoolClient
type jsiiProxy_IUserPoolClient struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IUserPoolClient) UserPoolClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolClient) UserPoolClientSecret() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"userPoolClientSecret",
		&returns,
	)
	return returns
}

