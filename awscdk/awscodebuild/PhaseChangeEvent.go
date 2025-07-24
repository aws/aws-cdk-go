package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Event fields for the CodeBuild "phase change" event.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html#sample-build-notifications-ref
//
type PhaseChangeEvent interface {
}

// The jsii proxy struct for PhaseChangeEvent
type jsiiProxy_PhaseChangeEvent struct {
	_ byte // padding
}

func PhaseChangeEvent_BuildComplete() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"buildComplete",
		&returns,
	)
	return returns
}

func PhaseChangeEvent_BuildId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"buildId",
		&returns,
	)
	return returns
}

func PhaseChangeEvent_CompletedPhase() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"completedPhase",
		&returns,
	)
	return returns
}

func PhaseChangeEvent_CompletedPhaseDurationSeconds() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"completedPhaseDurationSeconds",
		&returns,
	)
	return returns
}

func PhaseChangeEvent_CompletedPhaseStatus() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"completedPhaseStatus",
		&returns,
	)
	return returns
}

func PhaseChangeEvent_ProjectName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"projectName",
		&returns,
	)
	return returns
}

