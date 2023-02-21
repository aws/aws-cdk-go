package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Event fields for the CodeBuild "state change" event.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html#sample-build-notifications-ref
//
type StateChangeEvent interface {
}

// The jsii proxy struct for StateChangeEvent
type jsiiProxy_StateChangeEvent struct {
	_ byte // padding
}

func StateChangeEvent_BuildId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.StateChangeEvent",
		"buildId",
		&returns,
	)
	return returns
}

func StateChangeEvent_BuildStatus() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.StateChangeEvent",
		"buildStatus",
		&returns,
	)
	return returns
}

func StateChangeEvent_CurrentPhase() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.StateChangeEvent",
		"currentPhase",
		&returns,
	)
	return returns
}

func StateChangeEvent_ProjectName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.StateChangeEvent",
		"projectName",
		&returns,
	)
	return returns
}

