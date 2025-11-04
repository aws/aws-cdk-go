package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Stage.
// Experimental.
type IStageRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Stage resource.
	// Experimental.
	StageRef() *StageReference
}

// The jsii proxy for IStageRef
type jsiiProxy_IStageRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IStageRef) StageRef() *StageReference {
	var returns *StageReference
	_jsii_.Get(
		j,
		"stageRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStageRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStageRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

