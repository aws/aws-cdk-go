package awscdkivsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ivs-alpha.Channel",
		reflect.TypeOf((*Channel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStreamKey", GoMethod: "AddStreamKey"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "channelArn", GoGetter: "ChannelArn"},
			_jsii_.MemberProperty{JsiiProperty: "channelIngestEndpoint", GoGetter: "ChannelIngestEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "channelPlaybackUrl", GoGetter: "ChannelPlaybackUrl"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Channel{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IChannel)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ivs-alpha.ChannelProps",
		reflect.TypeOf((*ChannelProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ivs-alpha.ChannelType",
		reflect.TypeOf((*ChannelType)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": ChannelType_STANDARD,
			"BASIC": ChannelType_BASIC,
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ivs-alpha.IChannel",
		reflect.TypeOf((*IChannel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStreamKey", GoMethod: "AddStreamKey"},
			_jsii_.MemberProperty{JsiiProperty: "channelArn", GoGetter: "ChannelArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IChannel{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ivs-alpha.IPlaybackKeyPair",
		reflect.TypeOf((*IPlaybackKeyPair)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "playbackKeyPairArn", GoGetter: "PlaybackKeyPairArn"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IPlaybackKeyPair{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ivs-alpha.IStreamKey",
		reflect.TypeOf((*IStreamKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "streamKeyArn", GoGetter: "StreamKeyArn"},
		},
		func() interface{} {
			j := jsiiProxy_IStreamKey{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ivs-alpha.LatencyMode",
		reflect.TypeOf((*LatencyMode)(nil)).Elem(),
		map[string]interface{}{
			"LOW": LatencyMode_LOW,
			"NORMAL": LatencyMode_NORMAL,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ivs-alpha.PlaybackKeyPair",
		reflect.TypeOf((*PlaybackKeyPair)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "playbackKeyPairArn", GoGetter: "PlaybackKeyPairArn"},
			_jsii_.MemberProperty{JsiiProperty: "playbackKeyPairFingerprint", GoGetter: "PlaybackKeyPairFingerprint"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PlaybackKeyPair{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPlaybackKeyPair)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ivs-alpha.PlaybackKeyPairProps",
		reflect.TypeOf((*PlaybackKeyPairProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ivs-alpha.StreamKey",
		reflect.TypeOf((*StreamKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "streamKeyArn", GoGetter: "StreamKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "streamKeyValue", GoGetter: "StreamKeyValue"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StreamKey{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStreamKey)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ivs-alpha.StreamKeyProps",
		reflect.TypeOf((*StreamKeyProps)(nil)).Elem(),
	)
}
