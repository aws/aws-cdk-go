package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2"
)

// Represents a reference to an HTTP Stage.
type IHttpStageRef interface {
	interfacesawsapigatewayv2.IStageRef
	// Indicates that this is an HTTP Stage.
	//
	// Will always return true, but is necessary to prevent accidental structural
	// equality in TypeScript.
	IsHttpStage() *bool
}

// The jsii proxy for IHttpStageRef
type jsiiProxy_IHttpStageRef struct {
	internal.Type__interfacesawsapigatewayv2IStageRef
}

func (j *jsiiProxy_IHttpStageRef) IsHttpStage() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isHttpStage",
		&returns,
	)
	return returns
}

