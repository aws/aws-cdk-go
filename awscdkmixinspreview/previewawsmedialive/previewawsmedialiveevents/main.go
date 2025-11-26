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
		"@aws-cdk/mixins-preview.aws_medialive.events.ChannelEvents.MediaLiveChannelInputChange",
		reflect.TypeOf((*ChannelEvents_MediaLiveChannelInputChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ChannelEvents_MediaLiveChannelInputChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_medialive.events.ChannelEvents.MediaLiveChannelInputChange.MediaLiveChannelInputChangeProps",
		reflect.TypeOf((*ChannelEvents_MediaLiveChannelInputChange_MediaLiveChannelInputChangeProps)(nil)).Elem(),
	)
}
