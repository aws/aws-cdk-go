package previewawsmedialiveevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_medialive.events.ChannelEvents",
		reflect.TypeOf((*ChannelEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "mediaLiveChannelInputChangePattern", GoMethod: "MediaLiveChannelInputChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_ChannelEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelAlert",
		reflect.TypeOf((*MediaLiveChannelAlert)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MediaLiveChannelAlert{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelAlert.MediaLiveChannelAlertProps",
		reflect.TypeOf((*MediaLiveChannelAlert_MediaLiveChannelAlertProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelInputChange",
		reflect.TypeOf((*MediaLiveChannelInputChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MediaLiveChannelInputChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelInputChange.MediaLiveChannelInputChangeProps",
		reflect.TypeOf((*MediaLiveChannelInputChange_MediaLiveChannelInputChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelStateChange",
		reflect.TypeOf((*MediaLiveChannelStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MediaLiveChannelStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelStateChange.MediaLiveChannelStateChangeProps",
		reflect.TypeOf((*MediaLiveChannelStateChange_MediaLiveChannelStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveMultiplexAlert",
		reflect.TypeOf((*MediaLiveMultiplexAlert)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MediaLiveMultiplexAlert{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveMultiplexAlert.MediaLiveMultiplexAlertProps",
		reflect.TypeOf((*MediaLiveMultiplexAlert_MediaLiveMultiplexAlertProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveMultiplexStateChange",
		reflect.TypeOf((*MediaLiveMultiplexStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MediaLiveMultiplexStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveMultiplexStateChange.MediaLiveMultiplexStateChangeProps",
		reflect.TypeOf((*MediaLiveMultiplexStateChange_MediaLiveMultiplexStateChangeProps)(nil)).Elem(),
	)
}
