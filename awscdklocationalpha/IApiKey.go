package awscdklocationalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdklocationalpha/v2/internal"
)

// An API Key.
// Experimental.
type IApiKey interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the api key resource.
	// Experimental.
	ApiKeyArn() *string
	// The name of the api key.
	// Experimental.
	ApiKeyName() *string
}

// The jsii proxy for IApiKey
type jsiiProxy_IApiKey struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IApiKey) ApiKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKey) ApiKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyName",
		&returns,
	)
	return returns
}

