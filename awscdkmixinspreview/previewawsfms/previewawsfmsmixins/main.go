package previewawsfmsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnNotificationChannelMixinProps",
		reflect.TypeOf((*CfnNotificationChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnNotificationChannelPropsMixin",
		reflect.TypeOf((*CfnNotificationChannelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNotificationChannelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyMixinProps",
		reflect.TypeOf((*CfnPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin",
		reflect.TypeOf((*CfnPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.IEMapProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_IEMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.IcmpTypeCodeProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_IcmpTypeCodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.NetworkAclCommonPolicyProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_NetworkAclCommonPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.NetworkAclEntryProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_NetworkAclEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.NetworkAclEntrySetProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_NetworkAclEntrySetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.NetworkFirewallPolicyProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_NetworkFirewallPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.PolicyOptionProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_PolicyOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.PolicyTagProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_PolicyTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.PortRangeProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_PortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.ResourceTagProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_ResourceTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.SecurityServicePolicyDataProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_SecurityServicePolicyDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnPolicyPropsMixin.ThirdPartyFirewallPolicyProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_ThirdPartyFirewallPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnResourceSetMixinProps",
		reflect.TypeOf((*CfnResourceSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_fms.mixins.CfnResourceSetPropsMixin",
		reflect.TypeOf((*CfnResourceSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourceSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
