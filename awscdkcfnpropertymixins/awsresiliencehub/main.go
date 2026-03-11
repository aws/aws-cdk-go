package awsresiliencehub

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehub.CfnAppMixinProps",
		reflect.TypeOf((*CfnAppMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehub.CfnAppPropsMixin",
		reflect.TypeOf((*CfnAppPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehub.CfnAppPropsMixin.EventSubscriptionProperty",
		reflect.TypeOf((*CfnAppPropsMixin_EventSubscriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehub.CfnAppPropsMixin.PermissionModelProperty",
		reflect.TypeOf((*CfnAppPropsMixin_PermissionModelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehub.CfnAppPropsMixin.PhysicalResourceIdProperty",
		reflect.TypeOf((*CfnAppPropsMixin_PhysicalResourceIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehub.CfnAppPropsMixin.ResourceMappingProperty",
		reflect.TypeOf((*CfnAppPropsMixin_ResourceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehub.CfnResiliencyPolicyMixinProps",
		reflect.TypeOf((*CfnResiliencyPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehub.CfnResiliencyPolicyPropsMixin",
		reflect.TypeOf((*CfnResiliencyPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResiliencyPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehub.CfnResiliencyPolicyPropsMixin.FailurePolicyProperty",
		reflect.TypeOf((*CfnResiliencyPolicyPropsMixin_FailurePolicyProperty)(nil)).Elem(),
	)
}
