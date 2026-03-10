package previewawssyntheticsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryStatusChange",
		reflect.TypeOf((*SyntheticsCanaryStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SyntheticsCanaryStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryStatusChange.ChangedConfig",
		reflect.TypeOf((*SyntheticsCanaryStatusChange_ChangedConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryStatusChange.ExecutionArn",
		reflect.TypeOf((*SyntheticsCanaryStatusChange_ExecutionArn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryStatusChange.SyntheticsCanaryStatusChangeProps",
		reflect.TypeOf((*SyntheticsCanaryStatusChange_SyntheticsCanaryStatusChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryStatusChange.TestCodeLayerVersionArn",
		reflect.TypeOf((*SyntheticsCanaryStatusChange_TestCodeLayerVersionArn)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunFailure",
		reflect.TypeOf((*SyntheticsCanaryTestRunFailure)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SyntheticsCanaryTestRunFailure{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunFailure.CanaryRunTimeline",
		reflect.TypeOf((*SyntheticsCanaryTestRunFailure_CanaryRunTimeline)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunFailure.SyntheticsCanaryTestRunFailureProps",
		reflect.TypeOf((*SyntheticsCanaryTestRunFailure_SyntheticsCanaryTestRunFailureProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunSuccessful",
		reflect.TypeOf((*SyntheticsCanaryTestRunSuccessful)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SyntheticsCanaryTestRunSuccessful{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunSuccessful.CanaryRunTimeline",
		reflect.TypeOf((*SyntheticsCanaryTestRunSuccessful_CanaryRunTimeline)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunSuccessful.SyntheticsCanaryTestRunSuccessfulProps",
		reflect.TypeOf((*SyntheticsCanaryTestRunSuccessful_SyntheticsCanaryTestRunSuccessfulProps)(nil)).Elem(),
	)
}
