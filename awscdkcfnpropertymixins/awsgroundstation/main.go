package awsgroundstation

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigMixinProps",
		reflect.TypeOf((*CfnConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin",
		reflect.TypeOf((*CfnConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.AntennaDownlinkConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_AntennaDownlinkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.AntennaDownlinkDemodDecodeConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_AntennaDownlinkDemodDecodeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.AntennaUplinkConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_AntennaUplinkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.ConfigDataProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_ConfigDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.DataflowEndpointConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_DataflowEndpointConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.DecodeConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_DecodeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.DemodulationConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_DemodulationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.EirpProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_EirpProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.FrequencyBandwidthProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_FrequencyBandwidthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.FrequencyProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_FrequencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.KinesisDataStreamDataProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_KinesisDataStreamDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.S3RecordingConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_S3RecordingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.SpectrumConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_SpectrumConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.TelemetrySinkConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_TelemetrySinkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.TelemetrySinkDataProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_TelemetrySinkDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.TrackingConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_TrackingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.UplinkEchoConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_UplinkEchoConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnConfigPropsMixin.UplinkSpectrumConfigProperty",
		reflect.TypeOf((*CfnConfigPropsMixin_UplinkSpectrumConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupMixinProps",
		reflect.TypeOf((*CfnDataflowEndpointGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupPropsMixin",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataflowEndpointGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupPropsMixin.AwsGroundStationAgentEndpointProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_AwsGroundStationAgentEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupPropsMixin.ConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_ConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupPropsMixin.DataflowEndpointProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_DataflowEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupPropsMixin.EndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_EndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupPropsMixin.IntegerRangeProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_IntegerRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupPropsMixin.RangedConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_RangedConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupPropsMixin.RangedSocketAddressProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_RangedSocketAddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupPropsMixin.SecurityDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_SecurityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupPropsMixin.SocketAddressProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupPropsMixin_SocketAddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2MixinProps",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2MixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataflowEndpointGroupV2PropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.ConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_ConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.CreateEndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_CreateEndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.DownlinkAwsGroundStationAgentEndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_DownlinkAwsGroundStationAgentEndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.DownlinkAwsGroundStationAgentEndpointProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_DownlinkAwsGroundStationAgentEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.DownlinkConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_DownlinkConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.DownlinkDataflowDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_DownlinkDataflowDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.EndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_EndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.IntegerRangeProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_IntegerRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.RangedConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_RangedConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.RangedSocketAddressProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_RangedSocketAddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.SocketAddressProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_SocketAddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.UplinkAwsGroundStationAgentEndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_UplinkAwsGroundStationAgentEndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.UplinkAwsGroundStationAgentEndpointProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_UplinkAwsGroundStationAgentEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.UplinkConnectionDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_UplinkConnectionDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnDataflowEndpointGroupV2PropsMixin.UplinkDataflowDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroupV2PropsMixin_UplinkDataflowDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnMissionProfileMixinProps",
		reflect.TypeOf((*CfnMissionProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnMissionProfilePropsMixin",
		reflect.TypeOf((*CfnMissionProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMissionProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnMissionProfilePropsMixin.DataflowEdgeProperty",
		reflect.TypeOf((*CfnMissionProfilePropsMixin_DataflowEdgeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_groundstation.CfnMissionProfilePropsMixin.StreamsKmsKeyProperty",
		reflect.TypeOf((*CfnMissionProfilePropsMixin_StreamsKmsKeyProperty)(nil)).Elem(),
	)
}
