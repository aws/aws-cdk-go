package previewawscodepipelineevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codepipeline@CodePipelineActionExecutionStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codePipelineActionExecutionStateChange := awscdkmixinspreview.Events.NewCodePipelineActionExecutionStateChange()
//
// Experimental.
type CodePipelineActionExecutionStateChange interface {
}

// The jsii proxy struct for CodePipelineActionExecutionStateChange
type jsiiProxy_CodePipelineActionExecutionStateChange struct {
	_ byte // padding
}

// Experimental.
func NewCodePipelineActionExecutionStateChange() CodePipelineActionExecutionStateChange {
	_init_.Initialize()

	j := jsiiProxy_CodePipelineActionExecutionStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineActionExecutionStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodePipelineActionExecutionStateChange_Override(c CodePipelineActionExecutionStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineActionExecutionStateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodePipeline Action Execution State Change.
// Experimental.
func CodePipelineActionExecutionStateChange_EventPattern(options *CodePipelineActionExecutionStateChange_CodePipelineActionExecutionStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodePipelineActionExecutionStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineActionExecutionStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

