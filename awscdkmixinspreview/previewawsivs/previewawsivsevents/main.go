package previewawsivsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ivs.events.LimitBreach",
		reflect.TypeOf((*LimitBreach)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_LimitBreach{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivs.events.LimitBreach.LimitBreachProps",
		reflect.TypeOf((*LimitBreach_LimitBreachProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ivs.events.RecordingStateChange",
		reflect.TypeOf((*RecordingStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RecordingStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivs.events.RecordingStateChange.RecordingStateChangeProps",
		reflect.TypeOf((*RecordingStateChange_RecordingStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ivs.events.StreamHealthChange",
		reflect.TypeOf((*StreamHealthChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StreamHealthChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivs.events.StreamHealthChange.StreamHealthChangeProps",
		reflect.TypeOf((*StreamHealthChange_StreamHealthChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ivs.events.StreamStateChange",
		reflect.TypeOf((*StreamStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StreamStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivs.events.StreamStateChange.StreamStateChangeProps",
		reflect.TypeOf((*StreamStateChange_StreamStateChangeProps)(nil)).Elem(),
	)
}
