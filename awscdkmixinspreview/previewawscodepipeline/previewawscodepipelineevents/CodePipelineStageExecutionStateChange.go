package previewawscodepipelineevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codepipeline@CodePipelineStageExecutionStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codePipelineStageExecutionStateChange := awscdkmixinspreview.Events.NewCodePipelineStageExecutionStateChange()
//
// Experimental.
type CodePipelineStageExecutionStateChange interface {
}

// The jsii proxy struct for CodePipelineStageExecutionStateChange
type jsiiProxy_CodePipelineStageExecutionStateChange struct {
	_ byte // padding
}

// Experimental.
func NewCodePipelineStageExecutionStateChange() CodePipelineStageExecutionStateChange {
	_init_.Initialize()

	j := jsiiProxy_CodePipelineStageExecutionStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineStageExecutionStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodePipelineStageExecutionStateChange_Override(c CodePipelineStageExecutionStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineStageExecutionStateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodePipeline Stage Execution State Change.
// Experimental.
func CodePipelineStageExecutionStateChange_EventPattern(options *CodePipelineStageExecutionStateChange_CodePipelineStageExecutionStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodePipelineStageExecutionStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineStageExecutionStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

