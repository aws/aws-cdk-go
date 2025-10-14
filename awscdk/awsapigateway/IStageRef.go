package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Stage.
// Experimental.
type IStageRef interface {
	constructs.IConstruct
	// A reference to a Stage resource.
	// Experimental.
	StageRef() *StageReference
}

// The jsii proxy for IStageRef
type jsiiProxy_IStageRef struct {
	internal.Type__constructsIConstruct
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

