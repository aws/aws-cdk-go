package awscloudtrail

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnChannelMixinProps",
		reflect.TypeOf((*CfnChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnChannelPropsMixin",
		reflect.TypeOf((*CfnChannelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnChannelPropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnDashboardMixinProps",
		reflect.TypeOf((*CfnDashboardMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnDashboardPropsMixin",
		reflect.TypeOf((*CfnDashboardPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDashboardPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnDashboardPropsMixin.FrequencyProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_FrequencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnDashboardPropsMixin.RefreshScheduleProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_RefreshScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnDashboardPropsMixin.WidgetProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_WidgetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnEventDataStoreMixinProps",
		reflect.TypeOf((*CfnEventDataStoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnEventDataStorePropsMixin",
		reflect.TypeOf((*CfnEventDataStorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventDataStorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnEventDataStorePropsMixin.AdvancedEventSelectorProperty",
		reflect.TypeOf((*CfnEventDataStorePropsMixin_AdvancedEventSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnEventDataStorePropsMixin.AdvancedFieldSelectorProperty",
		reflect.TypeOf((*CfnEventDataStorePropsMixin_AdvancedFieldSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnEventDataStorePropsMixin.ContextKeySelectorProperty",
		reflect.TypeOf((*CfnEventDataStorePropsMixin_ContextKeySelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnEventDataStorePropsMixin.InsightSelectorProperty",
		reflect.TypeOf((*CfnEventDataStorePropsMixin_InsightSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnResourcePolicyMixinProps",
		reflect.TypeOf((*CfnResourcePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnResourcePolicyPropsMixin",
		reflect.TypeOf((*CfnResourcePolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourcePolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnTrailMixinProps",
		reflect.TypeOf((*CfnTrailMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnTrailPropsMixin",
		reflect.TypeOf((*CfnTrailPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrailPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnTrailPropsMixin.AdvancedEventSelectorProperty",
		reflect.TypeOf((*CfnTrailPropsMixin_AdvancedEventSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnTrailPropsMixin.AdvancedFieldSelectorProperty",
		reflect.TypeOf((*CfnTrailPropsMixin_AdvancedFieldSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnTrailPropsMixin.AggregationConfigurationProperty",
		reflect.TypeOf((*CfnTrailPropsMixin_AggregationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnTrailPropsMixin.DataResourceProperty",
		reflect.TypeOf((*CfnTrailPropsMixin_DataResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnTrailPropsMixin.EventSelectorProperty",
		reflect.TypeOf((*CfnTrailPropsMixin_EventSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudtrail.CfnTrailPropsMixin.InsightSelectorProperty",
		reflect.TypeOf((*CfnTrailPropsMixin_InsightSelectorProperty)(nil)).Elem(),
	)
}
