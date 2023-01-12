package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
)

// Represents an APIGateway Stage.
type IStage interface {
	awscdk.IResource
	// Add an ApiKey to this Stage.
	AddApiKey(id *string, options *ApiKeyOptions) IApiKey
	// RestApi to which this stage is associated.
	RestApi() IRestApi
	// Name of this stage.
	StageName() *string
}

// The jsii proxy for IStage
type jsiiProxy_IStage struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IStage) AddApiKey(id *string, options *ApiKeyOptions) IApiKey {
	if err := i.validateAddApiKeyParameters(id, options); err != nil {
		panic(err)
	}
	var returns IApiKey

	_jsii_.Invoke(
		i,
		"addApiKey",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IStage) RestApi() IRestApi {
	var returns IRestApi
	_jsii_.Get(
		j,
		"restApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

