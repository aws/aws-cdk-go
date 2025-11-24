package mixinsawsmpa

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnApprovalTeamMixinProps",
		reflect.TypeOf((*CfnApprovalTeamMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnApprovalTeamPropsMixin",
		reflect.TypeOf((*CfnApprovalTeamPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApprovalTeamPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnApprovalTeamPropsMixin.ApprovalStrategyProperty",
		reflect.TypeOf((*CfnApprovalTeamPropsMixin_ApprovalStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnApprovalTeamPropsMixin.ApproverProperty",
		reflect.TypeOf((*CfnApprovalTeamPropsMixin_ApproverProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnApprovalTeamPropsMixin.MofNApprovalStrategyProperty",
		reflect.TypeOf((*CfnApprovalTeamPropsMixin_MofNApprovalStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnApprovalTeamPropsMixin.PolicyProperty",
		reflect.TypeOf((*CfnApprovalTeamPropsMixin_PolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnIdentitySourceMixinProps",
		reflect.TypeOf((*CfnIdentitySourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnIdentitySourcePropsMixin",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentitySourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnIdentitySourcePropsMixin.IamIdentityCenterProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_IamIdentityCenterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnIdentitySourcePropsMixin.IdentitySourceParametersProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_IdentitySourceParametersProperty)(nil)).Elem(),
	)
}
