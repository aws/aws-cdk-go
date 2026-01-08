package previewawsgroundstationmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigMixinProps",
		reflect.TypeOf((*CfnConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin",
		reflect.TypeOf((*CfnConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.AntennaDownlinkConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_AntennaDownlinkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.AntennaDownlinkDemodDecodeConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_AntennaDownlinkDemodDecodeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.AntennaUplinkConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_AntennaUplinkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.ConfigDataProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_ConfigDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.DataflowEndpointConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_DataflowEndpointConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.DecodeConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_DecodeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.DemodulationConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_DemodulationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.EirpProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_EirpProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.FrequencyBandwidthProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_FrequencyBandwidthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.FrequencyProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_FrequencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.S3RecordingConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_S3RecordingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.SpectrumConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_SpectrumConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.TrackingConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_TrackingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.UplinkEchoConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_UplinkEchoConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin.UplinkSpectrumConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_UplinkSpectrumConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupMixinProps",
		reflect.TypeOf((*CfnDataflowEndpointGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataflowEndpointGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin.AwsGroundStationAgentEndpointProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_AwsGroundStationAgentEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin.ConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_ConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin.DataflowEndpointProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_DataflowEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin.EndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_EndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin.IntegerRangeProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_IntegerRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin.RangedConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_RangedConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin.RangedSocketAddressProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_RangedSocketAddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin.SecurityDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_SecurityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin.SocketAddressProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_SocketAddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2MixinProps",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2MixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataflowEndpointGroupV2PropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.ConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_ConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.CreateEndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_CreateEndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.DownlinkAwsGroundStationAgentEndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_DownlinkAwsGroundStationAgentEndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.DownlinkAwsGroundStationAgentEndpointProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_DownlinkAwsGroundStationAgentEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.DownlinkConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_DownlinkConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.DownlinkDataflowDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_DownlinkDataflowDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.EndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_EndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.IntegerRangeProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_IntegerRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.RangedConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_RangedConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.RangedSocketAddressProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_RangedSocketAddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.SocketAddressProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_SocketAddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.UplinkAwsGroundStationAgentEndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_UplinkAwsGroundStationAgentEndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.UplinkAwsGroundStationAgentEndpointProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_UplinkAwsGroundStationAgentEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.UplinkConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_UplinkConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin.UplinkDataflowDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_UplinkDataflowDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnMissionProfileMixinProps",
		reflect.TypeOf((*CfnMissionProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnMissionProfilePropsMixin",
		reflect.TypeOf((*CfnMissionProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMissionProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnMissionProfilePropsMixin.DataflowEdgeProperty",
		reflect.TypeOf((*CfnMissionProfilePropsMixin_DataflowEdgeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnMissionProfilePropsMixin.StreamsKmsKeyProperty",
		reflect.TypeOf((*CfnMissionProfilePropsMixin_StreamsKmsKeyProperty)(nil)).Elem(),
	)
}
