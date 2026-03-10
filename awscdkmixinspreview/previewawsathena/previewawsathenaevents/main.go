package previewawsathenaevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_athena.events.AthenaQueryStateChange",
		reflect.TypeOf((*AthenaQueryStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AthenaQueryStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_athena.events.AthenaQueryStateChange.AthenaQueryStateChangeProps",
		reflect.TypeOf((*AthenaQueryStateChange_AthenaQueryStateChangeProps)(nil)).Elem(),
	)
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
}
