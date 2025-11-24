package mixinsawsmanagedblockchain

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnAccessorMixinProps",
		reflect.TypeOf((*CfnAccessorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnAccessorPropsMixin",
		reflect.TypeOf((*CfnAccessorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberMixinProps",
		reflect.TypeOf((*CfnMemberMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin",
		reflect.TypeOf((*CfnMemberPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMemberPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin.ApprovalThresholdPolicyProperty",
		reflect.TypeOf((*CfnMemberPropsMixin_ApprovalThresholdPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin.MemberConfigurationProperty",
		reflect.TypeOf((*CfnMemberPropsMixin_MemberConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin.MemberFabricConfigurationProperty",
		reflect.TypeOf((*CfnMemberPropsMixin_MemberFabricConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin.MemberFrameworkConfigurationProperty",
		reflect.TypeOf((*CfnMemberPropsMixin_MemberFrameworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnMemberPropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin.NetworkFabricConfigurationProperty",
		reflect.TypeOf((*CfnMemberPropsMixin_NetworkFabricConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin.NetworkFrameworkConfigurationProperty",
		reflect.TypeOf((*CfnMemberPropsMixin_NetworkFrameworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin.VotingPolicyProperty",
		reflect.TypeOf((*CfnMemberPropsMixin_VotingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnNodeMixinProps",
		reflect.TypeOf((*CfnNodeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnNodePropsMixin",
		reflect.TypeOf((*CfnNodePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNodePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnNodePropsMixin.NodeConfigurationProperty",
		reflect.TypeOf((*CfnNodePropsMixin_NodeConfigurationProperty)(nil)).Elem(),
	)
}
