package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
)

// API keys are alphanumeric string values that you distribute to app developer customers to grant access to your API.
type IApiKey interface {
	awscdk.IResource
	// The API key ARN.
	KeyArn() *string
	// The API key ID.
	KeyId() *string
}

// The jsii proxy for IApiKey
type jsiiProxy_IApiKey struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IApiKey) KeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKey) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

