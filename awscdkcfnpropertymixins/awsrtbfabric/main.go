package awsrtbfabric

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnInboundExternalLinkMixinProps",
		reflect.TypeOf((*CfnInboundExternalLinkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnInboundExternalLinkPropsMixin",
		reflect.TypeOf((*CfnInboundExternalLinkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInboundExternalLinkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnInboundExternalLinkPropsMixin.ApplicationLogsProperty",
		reflect.TypeOf((*CfnInboundExternalLinkPropsMixin_ApplicationLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnInboundExternalLinkPropsMixin.LinkApplicationLogSamplingProperty",
		reflect.TypeOf((*CfnInboundExternalLinkPropsMixin_LinkApplicationLogSamplingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnInboundExternalLinkPropsMixin.LinkAttributesProperty",
		reflect.TypeOf((*CfnInboundExternalLinkPropsMixin_LinkAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnInboundExternalLinkPropsMixin.LinkLogSettingsProperty",
		reflect.TypeOf((*CfnInboundExternalLinkPropsMixin_LinkLogSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnInboundExternalLinkPropsMixin.ResponderErrorMaskingForHttpCodeProperty",
		reflect.TypeOf((*CfnInboundExternalLinkPropsMixin_ResponderErrorMaskingForHttpCodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkMixinProps",
		reflect.TypeOf((*CfnLinkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin",
		reflect.TypeOf((*CfnLinkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLinkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.ApplicationLogsProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_ApplicationLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.FilterCriterionProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_FilterCriterionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.HeaderTagActionProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_HeaderTagActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.LinkApplicationLogSamplingProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_LinkApplicationLogSamplingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.LinkAttributesProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_LinkAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.LinkLogSettingsProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_LinkLogSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.ModuleConfigurationProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_ModuleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.ModuleParametersProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_ModuleParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.NoBidActionProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_NoBidActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.NoBidModuleParametersProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_NoBidModuleParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.OpenRtbAttributeModuleParametersProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_OpenRtbAttributeModuleParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkPropsMixin.ResponderErrorMaskingForHttpCodeProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_ResponderErrorMaskingForHttpCodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnOutboundExternalLinkMixinProps",
		reflect.TypeOf((*CfnOutboundExternalLinkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnOutboundExternalLinkPropsMixin",
		reflect.TypeOf((*CfnOutboundExternalLinkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOutboundExternalLinkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnOutboundExternalLinkPropsMixin.ApplicationLogsProperty",
		reflect.TypeOf((*CfnOutboundExternalLinkPropsMixin_ApplicationLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnOutboundExternalLinkPropsMixin.LinkApplicationLogSamplingProperty",
		reflect.TypeOf((*CfnOutboundExternalLinkPropsMixin_LinkApplicationLogSamplingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnOutboundExternalLinkPropsMixin.LinkAttributesProperty",
		reflect.TypeOf((*CfnOutboundExternalLinkPropsMixin_LinkAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnOutboundExternalLinkPropsMixin.LinkLogSettingsProperty",
		reflect.TypeOf((*CfnOutboundExternalLinkPropsMixin_LinkLogSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnOutboundExternalLinkPropsMixin.ResponderErrorMaskingForHttpCodeProperty",
		reflect.TypeOf((*CfnOutboundExternalLinkPropsMixin_ResponderErrorMaskingForHttpCodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnRequesterGatewayMixinProps",
		reflect.TypeOf((*CfnRequesterGatewayMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnRequesterGatewayPropsMixin",
		reflect.TypeOf((*CfnRequesterGatewayPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRequesterGatewayPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnResponderGatewayMixinProps",
		reflect.TypeOf((*CfnResponderGatewayMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnResponderGatewayPropsMixin",
		reflect.TypeOf((*CfnResponderGatewayPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResponderGatewayPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnResponderGatewayPropsMixin.AutoScalingGroupsConfigurationProperty",
		reflect.TypeOf((*CfnResponderGatewayPropsMixin_AutoScalingGroupsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnResponderGatewayPropsMixin.EksEndpointsConfigurationProperty",
		reflect.TypeOf((*CfnResponderGatewayPropsMixin_EksEndpointsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnResponderGatewayPropsMixin.HealthCheckConfigProperty",
		reflect.TypeOf((*CfnResponderGatewayPropsMixin_HealthCheckConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnResponderGatewayPropsMixin.ManagedEndpointConfigurationProperty",
		reflect.TypeOf((*CfnResponderGatewayPropsMixin_ManagedEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnResponderGatewayPropsMixin.TrustStoreConfigurationProperty",
		reflect.TypeOf((*CfnResponderGatewayPropsMixin_TrustStoreConfigurationProperty)(nil)).Elem(),
	)
}
