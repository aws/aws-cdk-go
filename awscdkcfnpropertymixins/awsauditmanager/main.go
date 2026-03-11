package awsauditmanager

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_auditmanager.CfnAssessmentMixinProps",
		reflect.TypeOf((*CfnAssessmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_auditmanager.CfnAssessmentPropsMixin",
		reflect.TypeOf((*CfnAssessmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssessmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_auditmanager.CfnAssessmentPropsMixin.AWSAccountProperty",
		reflect.TypeOf((*CfnAssessmentPropsMixin_AWSAccountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_auditmanager.CfnAssessmentPropsMixin.AWSServiceProperty",
		reflect.TypeOf((*CfnAssessmentPropsMixin_AWSServiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_auditmanager.CfnAssessmentPropsMixin.AssessmentReportsDestinationProperty",
		reflect.TypeOf((*CfnAssessmentPropsMixin_AssessmentReportsDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_auditmanager.CfnAssessmentPropsMixin.DelegationProperty",
		reflect.TypeOf((*CfnAssessmentPropsMixin_DelegationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_auditmanager.CfnAssessmentPropsMixin.RoleProperty",
		reflect.TypeOf((*CfnAssessmentPropsMixin_RoleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_auditmanager.CfnAssessmentPropsMixin.ScopeProperty",
		reflect.TypeOf((*CfnAssessmentPropsMixin_ScopeProperty)(nil)).Elem(),
	)
}
