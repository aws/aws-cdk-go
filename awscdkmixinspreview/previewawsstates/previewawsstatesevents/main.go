package previewawsstatesevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_states.events.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_states.events.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_states.events.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_states.events.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_states.events.AWSAPICallViaCloudTrail.ResponseElements",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ResponseElements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_states.events.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_states.events.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_states.events.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_states.events.ExecutionEvents",
		reflect.TypeOf((*ExecutionEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "stepFunctionsExecutionStatusChangePattern", GoMethod: "StepFunctionsExecutionStatusChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_ExecutionEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_states.events.StepFunctionsExecutionStatusChange",
		reflect.TypeOf((*StepFunctionsExecutionStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StepFunctionsExecutionStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_states.events.StepFunctionsExecutionStatusChange.StepFunctionsExecutionStatusChangeProps",
		reflect.TypeOf((*StepFunctionsExecutionStatusChange_StepFunctionsExecutionStatusChangeProps)(nil)).Elem(),
	)
}
