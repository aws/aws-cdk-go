package awsfms

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnNotificationChannelMixinProps",
		reflect.TypeOf((*CfnNotificationChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnNotificationChannelPropsMixin",
		reflect.TypeOf((*CfnNotificationChannelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNotificationChannelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyMixinProps",
		reflect.TypeOf((*CfnPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin",
		reflect.TypeOf((*CfnPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.IEMapProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_IEMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.IcmpTypeCodeProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_IcmpTypeCodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.NetworkAclCommonPolicyProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_NetworkAclCommonPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.NetworkAclEntryProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_NetworkAclEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.NetworkAclEntrySetProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_NetworkAclEntrySetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.NetworkFirewallPolicyProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_NetworkFirewallPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.PolicyOptionProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_PolicyOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.PolicyTagProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_PolicyTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.PortRangeProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_PortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.ResourceTagProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_ResourceTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.SecurityServicePolicyDataProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_SecurityServicePolicyDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnPolicyPropsMixin.ThirdPartyFirewallPolicyProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_ThirdPartyFirewallPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnResourceSetMixinProps",
		reflect.TypeOf((*CfnResourceSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fms.CfnResourceSetPropsMixin",
		reflect.TypeOf((*CfnResourceSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourceSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
