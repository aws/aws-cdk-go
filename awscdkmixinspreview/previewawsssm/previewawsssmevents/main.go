package previewawsssmevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssm.events.CalendarStateChange",
		reflect.TypeOf((*CalendarStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CalendarStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.events.CalendarStateChange.CalendarStateChangeProps",
		reflect.TypeOf((*CalendarStateChange_CalendarStateChangeProps)(nil)).Elem(),
	)
}
