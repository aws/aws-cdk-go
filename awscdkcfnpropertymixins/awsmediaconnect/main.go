package awsmediaconnect

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeMixinProps",
		reflect.TypeOf((*CfnBridgeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeOutputMixinProps",
		reflect.TypeOf((*CfnBridgeOutputMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeOutputPropsMixin",
		reflect.TypeOf((*CfnBridgeOutputPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBridgeOutputPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeOutputPropsMixin.BridgeNetworkOutputProperty",
		reflect.TypeOf((*CfnBridgeOutputPropsMixin_BridgeNetworkOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin",
		reflect.TypeOf((*CfnBridgePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBridgePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.BridgeFlowSourceProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_BridgeFlowSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.BridgeNetworkOutputProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_BridgeNetworkOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.BridgeNetworkSourceProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_BridgeNetworkSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.BridgeOutputProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_BridgeOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.BridgeSourceProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_BridgeSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.EgressGatewayBridgeProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_EgressGatewayBridgeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.FailoverConfigProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_FailoverConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.IngressGatewayBridgeProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_IngressGatewayBridgeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.MulticastSourceSettingsProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_MulticastSourceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.SourcePriorityProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_SourcePriorityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgePropsMixin.VpcInterfaceAttachmentProperty",
		reflect.TypeOf((*CfnBridgePropsMixin_VpcInterfaceAttachmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeSourceMixinProps",
		reflect.TypeOf((*CfnBridgeSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeSourcePropsMixin",
		reflect.TypeOf((*CfnBridgeSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBridgeSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeSourcePropsMixin.BridgeFlowSourceProperty",
		reflect.TypeOf((*CfnBridgeSourcePropsMixin_BridgeFlowSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeSourcePropsMixin.BridgeNetworkSourceProperty",
		reflect.TypeOf((*CfnBridgeSourcePropsMixin_BridgeNetworkSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeSourcePropsMixin.MulticastSourceSettingsProperty",
		reflect.TypeOf((*CfnBridgeSourcePropsMixin_MulticastSourceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeSourcePropsMixin.VpcInterfaceAttachmentProperty",
		reflect.TypeOf((*CfnBridgeSourcePropsMixin_VpcInterfaceAttachmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointMixinProps",
		reflect.TypeOf((*CfnDestinationEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDestinationEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.DestinationEndpointAssociatedResourceDetailsProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_DestinationEndpointAssociatedResourceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.DestinationEndpointConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_DestinationEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.EndpointEncryptionSecretConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_EndpointEncryptionSecretConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.PublicDestinationEndpointNetworkingConfigurationAllowRuleProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_PublicDestinationEndpointNetworkingConfigurationAllowRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.PublicDestinationEndpointNetworkingConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_PublicDestinationEndpointNetworkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.RistStandardDestinationEndpointConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_RistStandardDestinationEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.RouterInputDestinationEndpointAssociatedResourceDetailsProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_RouterInputDestinationEndpointAssociatedResourceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.RouterInputDestinationEndpointConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_RouterInputDestinationEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.RtpStandardDestinationEndpointConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_RtpStandardDestinationEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.SrtCallerStandardDestinationEndpointConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_SrtCallerStandardDestinationEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.SrtListenerStandardDestinationEndpointConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_SrtListenerStandardDestinationEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.StandardDestinationEndpointConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_StandardDestinationEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.StandardDestinationEndpointNetworkingConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_StandardDestinationEndpointNetworkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.StandardDestinationEndpointProtocolConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_StandardDestinationEndpointProtocolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnDestinationEndpointPropsMixin.VpcDestinationEndpointNetworkingConfigurationProperty",
		reflect.TypeOf((*CfnDestinationEndpointPropsMixin_VpcDestinationEndpointNetworkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowEntitlementMixinProps",
		reflect.TypeOf((*CfnFlowEntitlementMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowEntitlementPropsMixin",
		reflect.TypeOf((*CfnFlowEntitlementPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowEntitlementPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowEntitlementPropsMixin.EncryptionProperty",
		reflect.TypeOf((*CfnFlowEntitlementPropsMixin_EncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowMediaStreamMixinProps",
		reflect.TypeOf((*CfnFlowMediaStreamMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowMediaStreamPropsMixin",
		reflect.TypeOf((*CfnFlowMediaStreamPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowMediaStreamPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowMediaStreamPropsMixin.FmtpProperty",
		reflect.TypeOf((*CfnFlowMediaStreamPropsMixin_FmtpProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowMediaStreamPropsMixin.MediaStreamAttributesProperty",
		reflect.TypeOf((*CfnFlowMediaStreamPropsMixin_MediaStreamAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowMixinProps",
		reflect.TypeOf((*CfnFlowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputMixinProps",
		reflect.TypeOf((*CfnFlowOutputMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin",
		reflect.TypeOf((*CfnFlowOutputPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowOutputPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.DestinationConfigurationProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.EncodingParametersProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_EncodingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.EncryptionProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_EncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.FlowTransitEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_FlowTransitEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.FlowTransitEncryptionProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_FlowTransitEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.InterfaceProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_InterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.MediaStreamOutputConfigurationProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_MediaStreamOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.RtmpPushSettingsProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_RtmpPushSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.SecretsManagerEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_SecretsManagerEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.TlsEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_TlsEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.TlsEncryptionProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_TlsEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowOutputPropsMixin.VpcInterfaceAttachmentProperty",
		reflect.TypeOf((*CfnFlowOutputPropsMixin_VpcInterfaceAttachmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin",
		reflect.TypeOf((*CfnFlowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.AudioMonitoringSettingProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_AudioMonitoringSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.BlackFramesProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_BlackFramesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.EncodingConfigProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_EncodingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.EncryptionProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_EncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.FailoverConfigProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_FailoverConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.FlowTransitEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_FlowTransitEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.FlowTransitEncryptionProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_FlowTransitEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.FmtpProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_FmtpProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.FrozenFramesProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_FrozenFramesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.GatewayBridgeSourceProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_GatewayBridgeSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.InputConfigurationProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_InputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.InterfaceProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_InterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.MaintenanceProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_MaintenanceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.ManualRtmpPushNamingConfigurationProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_ManualRtmpPushNamingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.MediaStreamAttributesProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_MediaStreamAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.MediaStreamProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_MediaStreamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.MediaStreamSourceConfigurationProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_MediaStreamSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.NdiConfigProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_NdiConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.NdiDiscoveryServerConfigProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_NdiDiscoveryServerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.NdiSourceSettingsProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_NdiSourceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.RtmpPushDetailsProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_RtmpPushDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.RtmpPushNamingConfigurationProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_RtmpPushNamingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.RtmpPushSettingsProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_RtmpPushSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.SecretsManagerEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_SecretsManagerEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.SilentAudioProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_SilentAudioProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.SourceMonitoringConfigProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_SourceMonitoringConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.SourcePriorityProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_SourcePriorityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.VideoMonitoringSettingProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_VideoMonitoringSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.VpcInterfaceAttachmentProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_VpcInterfaceAttachmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowPropsMixin.VpcInterfaceProperty",
		reflect.TypeOf((*CfnFlowPropsMixin_VpcInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowSourceMixinProps",
		reflect.TypeOf((*CfnFlowSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowSourcePropsMixin",
		reflect.TypeOf((*CfnFlowSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowSourcePropsMixin.EncryptionProperty",
		reflect.TypeOf((*CfnFlowSourcePropsMixin_EncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowSourcePropsMixin.GatewayBridgeSourceProperty",
		reflect.TypeOf((*CfnFlowSourcePropsMixin_GatewayBridgeSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowSourcePropsMixin.ManualRtmpPushNamingConfigurationProperty",
		reflect.TypeOf((*CfnFlowSourcePropsMixin_ManualRtmpPushNamingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowSourcePropsMixin.RtmpPushDetailsProperty",
		reflect.TypeOf((*CfnFlowSourcePropsMixin_RtmpPushDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowSourcePropsMixin.RtmpPushNamingConfigurationProperty",
		reflect.TypeOf((*CfnFlowSourcePropsMixin_RtmpPushNamingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowSourcePropsMixin.RtmpPushSettingsProperty",
		reflect.TypeOf((*CfnFlowSourcePropsMixin_RtmpPushSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowSourcePropsMixin.VpcInterfaceAttachmentProperty",
		reflect.TypeOf((*CfnFlowSourcePropsMixin_VpcInterfaceAttachmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowVpcInterfaceMixinProps",
		reflect.TypeOf((*CfnFlowVpcInterfaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnFlowVpcInterfacePropsMixin",
		reflect.TypeOf((*CfnFlowVpcInterfacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowVpcInterfacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnGatewayMixinProps",
		reflect.TypeOf((*CfnGatewayMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnGatewayPropsMixin",
		reflect.TypeOf((*CfnGatewayPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnGatewayPropsMixin.GatewayNetworkProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GatewayNetworkProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputMixinProps",
		reflect.TypeOf((*CfnRouterInputMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin",
		reflect.TypeOf((*CfnRouterInputPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRouterInputPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.BlackFramesConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_BlackFramesConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.ContentQualityAnalysisFeatureConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_ContentQualityAnalysisFeatureConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.DestinationEndpointRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_DestinationEndpointRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.EndpointEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_EndpointEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.EndpointEncryptionSecretConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_EndpointEncryptionSecretConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.FailoverRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_FailoverRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.FailoverRouterInputIndexedStreamDetailsProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_FailoverRouterInputIndexedStreamDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.FailoverRouterInputProtocolConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_FailoverRouterInputProtocolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.FailoverRouterInputStreamDetailsProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_FailoverRouterInputStreamDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.FlowTransitEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_FlowTransitEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.FlowTransitEncryptionProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_FlowTransitEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.FrozenFramesConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_FrozenFramesConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.MaintenanceConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_MaintenanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.ManualRtmpPushNamingConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_ManualRtmpPushNamingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.MediaConnectFlowRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_MediaConnectFlowRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.MediaLiveChannelRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_MediaLiveChannelRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.MediaLiveTransitEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_MediaLiveTransitEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.MediaLiveTransitEncryptionProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_MediaLiveTransitEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.MergeRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_MergeRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.MergeRouterInputIndexedStreamDetailsProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_MergeRouterInputIndexedStreamDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.MergeRouterInputProtocolConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_MergeRouterInputProtocolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.MergeRouterInputStreamDetailsProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_MergeRouterInputStreamDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.PreferredDayTimeMaintenanceConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_PreferredDayTimeMaintenanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RistRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RistRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RouterContentQualityAnalysisConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RouterContentQualityAnalysisConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RouterInputProtocolConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RouterInputProtocolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RouterInputStreamDetailsProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RouterInputStreamDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RouterInputTransitEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RouterInputTransitEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RouterInputTransitEncryptionProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RouterInputTransitEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RtmpPushNamingConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RtmpPushNamingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RtmpPushRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RtmpPushRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RtmpPushStreamDetailsProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RtmpPushStreamDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.RtpRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_RtpRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.SecretsManagerEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_SecretsManagerEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.SecretsManagerEndpointEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_SecretsManagerEndpointEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.SilentAudioConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_SilentAudioConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.SrtCallerRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_SrtCallerRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.SrtDecryptionConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_SrtDecryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.SrtListenerRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_SrtListenerRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.StandardRouterInputConfigurationProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_StandardRouterInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterInputPropsMixin.StandardRouterInputStreamDetailsProperty",
		reflect.TypeOf((*CfnRouterInputPropsMixin_StandardRouterInputStreamDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterNetworkInterfaceMixinProps",
		reflect.TypeOf((*CfnRouterNetworkInterfaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterNetworkInterfacePropsMixin",
		reflect.TypeOf((*CfnRouterNetworkInterfacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRouterNetworkInterfacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterNetworkInterfacePropsMixin.PublicRouterNetworkInterfaceConfigurationProperty",
		reflect.TypeOf((*CfnRouterNetworkInterfacePropsMixin_PublicRouterNetworkInterfaceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterNetworkInterfacePropsMixin.PublicRouterNetworkInterfaceRuleProperty",
		reflect.TypeOf((*CfnRouterNetworkInterfacePropsMixin_PublicRouterNetworkInterfaceRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterNetworkInterfacePropsMixin.RouterNetworkInterfaceConfigurationProperty",
		reflect.TypeOf((*CfnRouterNetworkInterfacePropsMixin_RouterNetworkInterfaceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterNetworkInterfacePropsMixin.VpcRouterNetworkInterfaceConfigurationProperty",
		reflect.TypeOf((*CfnRouterNetworkInterfacePropsMixin_VpcRouterNetworkInterfaceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputMixinProps",
		reflect.TypeOf((*CfnRouterOutputMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin",
		reflect.TypeOf((*CfnRouterOutputPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRouterOutputPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.EndpointEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_EndpointEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.EndpointEncryptionSecretConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_EndpointEncryptionSecretConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.FabricConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_FabricConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.FlowTransitEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_FlowTransitEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.FlowTransitEncryptionProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_FlowTransitEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.MaintenanceConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_MaintenanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.MediaConnectFlowRouterOutputConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_MediaConnectFlowRouterOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.MediaLiveInputRouterOutputConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_MediaLiveInputRouterOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.MediaLiveTransitEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_MediaLiveTransitEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.MediaLiveTransitEncryptionProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_MediaLiveTransitEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.PreferredDayTimeMaintenanceConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_PreferredDayTimeMaintenanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.RistRouterOutputConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_RistRouterOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.RouterOutputConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_RouterOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.RouterOutputProtocolConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_RouterOutputProtocolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.RtmpPushRouterOutputConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_RtmpPushRouterOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.RtpRouterOutputConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_RtpRouterOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.SecretsManagerEncryptionKeyConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_SecretsManagerEncryptionKeyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.SecretsManagerEndpointEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_SecretsManagerEndpointEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.SourceEndpointRouterOutputConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_SourceEndpointRouterOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.SrtCallerRouterOutputConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_SrtCallerRouterOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.SrtEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_SrtEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.SrtListenerRouterOutputConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_SrtListenerRouterOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.StandardRouterOutputConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_StandardRouterOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.TlsEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_TlsEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterOutputPropsMixin.TlsEncryptionProperty",
		reflect.TypeOf((*CfnRouterOutputPropsMixin_TlsEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointMixinProps",
		reflect.TypeOf((*CfnSourceEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSourceEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.EndpointEncryptionSecretConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_EndpointEncryptionSecretConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.PublicSourceEndpointNetworkingConfigurationAllowRuleProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_PublicSourceEndpointNetworkingConfigurationAllowRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.PublicSourceEndpointNetworkingConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_PublicSourceEndpointNetworkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.RistStandardSourceEndpointConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_RistStandardSourceEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.RouterOutputSourceEndpointAssociatedResourceDetailsProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_RouterOutputSourceEndpointAssociatedResourceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.RouterOutputSourceEndpointConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_RouterOutputSourceEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.RtpStandardSourceEndpointConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_RtpStandardSourceEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.SourceEndpointAssociatedResourceDetailsProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_SourceEndpointAssociatedResourceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.SourceEndpointConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_SourceEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.SrtCallerStandardSourceEndpointConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_SrtCallerStandardSourceEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.SrtListenerStandardSourceEndpointConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_SrtListenerStandardSourceEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.StandardSourceEndpointConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_StandardSourceEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.StandardSourceEndpointNetworkingConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_StandardSourceEndpointNetworkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.StandardSourceEndpointProtocolConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_StandardSourceEndpointProtocolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnSourceEndpointPropsMixin.VpcSourceEndpointNetworkingConfigurationProperty",
		reflect.TypeOf((*CfnSourceEndpointPropsMixin_VpcSourceEndpointNetworkingConfigurationProperty)(nil)).Elem(),
	)
}
