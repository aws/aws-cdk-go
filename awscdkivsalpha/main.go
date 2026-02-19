// The CDK Construct Library for AWS::IVS
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
			"ADVANCED_SD": ChannelType_ADVANCED_SD,
			"ADVANCED_HD": ChannelType_ADVANCED_HD,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ivs-alpha.ContainerFormat",
		reflect.TypeOf((*ContainerFormat)(nil)).Elem(),
		map[string]interface{}{
			"TS": ContainerFormat_TS,
			"FRAGMENTED_MP4": ContainerFormat_FRAGMENTED_MP4,
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ivs-alpha.IChannel",
		reflect.TypeOf((*IChannel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStreamKey", GoMethod: "AddStreamKey"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "channelArn", GoGetter: "ChannelArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "playbackKeyPairArn", GoGetter: "PlaybackKeyPairArn"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IPlaybackKeyPair{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ivs-alpha.IRecordingConfiguration",
		reflect.TypeOf((*IRecordingConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "recordingConfigurationArn", GoGetter: "RecordingConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "recordingConfigurationId", GoGetter: "RecordingConfigurationId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IRecordingConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ivs-alpha.IStreamKey",
		reflect.TypeOf((*IStreamKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "streamKeyArn", GoGetter: "StreamKeyArn"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ivs-alpha.MaximumResolution",
		reflect.TypeOf((*MaximumResolution)(nil)).Elem(),
		map[string]interface{}{
			"FULL_HD": MaximumResolution_FULL_HD,
			"HD": MaximumResolution_HD,
			"SD": MaximumResolution_SD,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ivs-alpha.MultitrackInputConfiguration",
		reflect.TypeOf((*MultitrackInputConfiguration)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ivs-alpha.Policy",
		reflect.TypeOf((*Policy)(nil)).Elem(),
		map[string]interface{}{
			"ALLOW": Policy_ALLOW,
			"REQUIRE": Policy_REQUIRE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ivs-alpha.Preset",
		reflect.TypeOf((*Preset)(nil)).Elem(),
		map[string]interface{}{
			"CONSTRAINED_BANDWIDTH_DELIVERY": Preset_CONSTRAINED_BANDWIDTH_DELIVERY,
			"HIGHER_BANDWIDTH_DELIVERY": Preset_HIGHER_BANDWIDTH_DELIVERY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ivs-alpha.RecordingConfiguration",
		reflect.TypeOf((*RecordingConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "recordingConfigurationArn", GoGetter: "RecordingConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "recordingConfigurationId", GoGetter: "RecordingConfigurationId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_RecordingConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRecordingConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ivs-alpha.RecordingConfigurationProps",
		reflect.TypeOf((*RecordingConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ivs-alpha.RecordingMode",
		reflect.TypeOf((*RecordingMode)(nil)).Elem(),
		map[string]interface{}{
			"INTERVAL": RecordingMode_INTERVAL,
			"DISABLED": RecordingMode_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ivs-alpha.RenditionConfiguration",
		reflect.TypeOf((*RenditionConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "renditions", GoGetter: "Renditions"},
			_jsii_.MemberProperty{JsiiProperty: "renditionSelection", GoGetter: "RenditionSelection"},
		},
		func() interface{} {
			return &jsiiProxy_RenditionConfiguration{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ivs-alpha.RenditionSelection",
		reflect.TypeOf((*RenditionSelection)(nil)).Elem(),
		map[string]interface{}{
			"ALL": RenditionSelection_ALL,
			"NONE": RenditionSelection_NONE,
			"CUSTOM": RenditionSelection_CUSTOM,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ivs-alpha.Resolution",
		reflect.TypeOf((*Resolution)(nil)).Elem(),
		map[string]interface{}{
			"FULL_HD": Resolution_FULL_HD,
			"HD": Resolution_HD,
			"SD": Resolution_SD,
			"LOWEST_RESOLUTION": Resolution_LOWEST_RESOLUTION,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ivs-alpha.Storage",
		reflect.TypeOf((*Storage)(nil)).Elem(),
		map[string]interface{}{
			"SEQUENTIAL": Storage_SEQUENTIAL,
			"LATEST": Storage_LATEST,
		},
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ivs-alpha.ThumbnailConfiguration",
		reflect.TypeOf((*ThumbnailConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "recordingMode", GoGetter: "RecordingMode"},
			_jsii_.MemberProperty{JsiiProperty: "resolution", GoGetter: "Resolution"},
			_jsii_.MemberProperty{JsiiProperty: "storage", GoGetter: "Storage"},
			_jsii_.MemberProperty{JsiiProperty: "targetInterval", GoGetter: "TargetInterval"},
		},
		func() interface{} {
			return &jsiiProxy_ThumbnailConfiguration{}
		},
	)
}
