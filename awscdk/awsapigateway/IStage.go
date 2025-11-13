package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an APIGateway Stage.
type IStage interface {
	awscdk.IResource
	interfacesawsapigateway.IStageRef
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
	internal.Type__interfacesawsapigatewayIStageRef
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

func (i *jsiiProxy_IStage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IStage) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) StageRef() *interfacesawsapigateway.StageReference {
	var returns *interfacesawsapigateway.StageReference
	_jsii_.Get(
		j,
		"stageRef",
		&returns,
	)
	return returns
}

