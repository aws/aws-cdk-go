package previewawscodepipelineevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineActionExecutionStateChange",
		reflect.TypeOf((*CodePipelineActionExecutionStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CodePipelineActionExecutionStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineActionExecutionStateChange.CodePipelineActionExecutionStateChangeProps",
		reflect.TypeOf((*CodePipelineActionExecutionStateChange_CodePipelineActionExecutionStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineActionExecutionStateChange.ExecutionResult",
		reflect.TypeOf((*CodePipelineActionExecutionStateChange_ExecutionResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineActionExecutionStateChange.InputArtifacts",
		reflect.TypeOf((*CodePipelineActionExecutionStateChange_InputArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineActionExecutionStateChange.OutputArtifacts",
		reflect.TypeOf((*CodePipelineActionExecutionStateChange_OutputArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineActionExecutionStateChange.S3Location",
		reflect.TypeOf((*CodePipelineActionExecutionStateChange_S3Location)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineActionExecutionStateChange.Type",
		reflect.TypeOf((*CodePipelineActionExecutionStateChange_Type)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelinePipelineExecutionStateChange",
		reflect.TypeOf((*CodePipelinePipelineExecutionStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CodePipelinePipelineExecutionStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelinePipelineExecutionStateChange.CodePipelinePipelineExecutionStateChangeProps",
		reflect.TypeOf((*CodePipelinePipelineExecutionStateChange_CodePipelinePipelineExecutionStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelinePipelineExecutionStateChange.ExecutionTrigger",
		reflect.TypeOf((*CodePipelinePipelineExecutionStateChange_ExecutionTrigger)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineStageExecutionStateChange",
		reflect.TypeOf((*CodePipelineStageExecutionStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CodePipelineStageExecutionStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.events.CodePipelineStageExecutionStateChange.CodePipelineStageExecutionStateChangeProps",
		reflect.TypeOf((*CodePipelineStageExecutionStateChange_CodePipelineStageExecutionStateChangeProps)(nil)).Elem(),
	)
}
