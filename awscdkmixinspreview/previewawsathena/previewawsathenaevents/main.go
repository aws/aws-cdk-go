package previewawsathenaevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_athena.events.WorkGroupEvents",
		reflect.TypeOf((*WorkGroupEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "athenaQueryStateChangePattern", GoMethod: "AthenaQueryStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_WorkGroupEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_athena.events.WorkGroupEvents.AthenaQueryStateChange",
		reflect.TypeOf((*WorkGroupEvents_AthenaQueryStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_WorkGroupEvents_AthenaQueryStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_athena.events.WorkGroupEvents.AthenaQueryStateChange.AthenaQueryStateChangeProps",
		reflect.TypeOf((*WorkGroupEvents_AthenaQueryStateChange_AthenaQueryStateChangeProps)(nil)).Elem(),
	)
}
