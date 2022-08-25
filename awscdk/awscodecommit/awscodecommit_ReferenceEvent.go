package awscodecommit

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Fields of CloudWatch Events that change references.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#codebuild_event_type
//
// Experimental.
type ReferenceEvent interface {
}

// The jsii proxy struct for ReferenceEvent
type jsiiProxy_ReferenceEvent struct {
	_ byte // padding
}

func ReferenceEvent_CommitId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"commitId",
		&returns,
	)
	return returns
}

func ReferenceEvent_EventType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"eventType",
		&returns,
	)
	return returns
}

func ReferenceEvent_ReferenceFullName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"referenceFullName",
		&returns,
	)
	return returns
}

func ReferenceEvent_ReferenceName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"referenceName",
		&returns,
	)
	return returns
}

func ReferenceEvent_ReferenceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"referenceType",
		&returns,
	)
	return returns
}

func ReferenceEvent_RepositoryId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"repositoryId",
		&returns,
	)
	return returns
}

func ReferenceEvent_RepositoryName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"repositoryName",
		&returns,
	)
	return returns
}

