package previewawscodepipelineevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codepipeline@CodePipelinePipelineExecutionStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codePipelinePipelineExecutionStateChange := awscdkmixinspreview.Events.NewCodePipelinePipelineExecutionStateChange()
//
// Experimental.
type CodePipelinePipelineExecutionStateChange interface {
}

// The jsii proxy struct for CodePipelinePipelineExecutionStateChange
type jsiiProxy_CodePipelinePipelineExecutionStateChange struct {
	_ byte // padding
}

// Experimental.
func NewCodePipelinePipelineExecutionStateChange() CodePipelinePipelineExecutionStateChange {
	_init_.Initialize()

	j := jsiiProxy_CodePipelinePipelineExecutionStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelinePipelineExecutionStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodePipelinePipelineExecutionStateChange_Override(c CodePipelinePipelineExecutionStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelinePipelineExecutionStateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodePipeline Pipeline Execution State Change.
// Experimental.
func CodePipelinePipelineExecutionStateChange_EventPattern(options *CodePipelinePipelineExecutionStateChange_CodePipelinePipelineExecutionStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodePipelinePipelineExecutionStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelinePipelineExecutionStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

