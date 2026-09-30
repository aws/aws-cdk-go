package awseventsv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnEventBusMixinProps",
		reflect.TypeOf((*CfnEventBusMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnEventBusPropsMixin",
		reflect.TypeOf((*CfnEventBusPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventBusPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnEventBusPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnEventBusPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnEventBusPropsMixin.StorageConfigurationProperty",
		reflect.TypeOf((*CfnEventBusPropsMixin_StorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnEventSourceMixinProps",
		reflect.TypeOf((*CfnEventSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnEventSourcePropsMixin",
		reflect.TypeOf((*CfnEventSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnEventSourcePropsMixin.AwsServiceEventsConfigurationProperty",
		reflect.TypeOf((*CfnEventSourcePropsMixin_AwsServiceEventsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnEventSourcePropsMixin.EventSourceConfigurationProperty",
		reflect.TypeOf((*CfnEventSourcePropsMixin_EventSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnEventSourcePropsMixin.OnFailureConfigurationProperty",
		reflect.TypeOf((*CfnEventSourcePropsMixin_OnFailureConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnEventSourcePropsMixin.PartnerEventsConfigurationProperty",
		reflect.TypeOf((*CfnEventSourcePropsMixin_PartnerEventsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnResourcePolicyMixinProps",
		reflect.TypeOf((*CfnResourcePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnResourcePolicyPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberMixinProps",
		reflect.TypeOf((*CfnSubscriberMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin",
		reflect.TypeOf((*CfnSubscriberPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubscriberPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin.BatchConfigurationProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_BatchConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin.FilterConfigurationProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_FilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin.JsonataConfigurationProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_JsonataConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin.OnFailureConfigurationProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_OnFailureConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin.PointInTimeConfigurationProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_PointInTimeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin.RetryPolicyProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_RetryPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin.TransformerProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_TransformerProperty)(nil)).Elem(),
	)
}
