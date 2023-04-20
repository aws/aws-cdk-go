package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway/internal"
)

// Represents an APIGateway Stage.
// Experimental.
type IStage interface {
	awscdk.IResource
	// RestApi to which this stage is associated.
	// Experimental.
	RestApi() IRestApi
	// Name of this stage.
	// Experimental.
	StageName() *string
}

// The jsii proxy for IStage
type jsiiProxy_IStage struct {
	internal.Type__awscdkIResource
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

