package awss3objectlambda

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3objectlambda.CfnAccessPointMixinProps",
		reflect.TypeOf((*CfnAccessPointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3objectlambda.CfnAccessPointPolicyMixinProps",
		reflect.TypeOf((*CfnAccessPointPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_s3objectlambda.CfnAccessPointPolicyPropsMixin",
		reflect.TypeOf((*CfnAccessPointPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessPointPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_s3objectlambda.CfnAccessPointPropsMixin",
		reflect.TypeOf((*CfnAccessPointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessPointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3objectlambda.CfnAccessPointPropsMixin.AliasProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_AliasProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3objectlambda.CfnAccessPointPropsMixin.ObjectLambdaConfigurationProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_ObjectLambdaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3objectlambda.CfnAccessPointPropsMixin.PolicyStatusProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_PolicyStatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3objectlambda.CfnAccessPointPropsMixin.PublicAccessBlockConfigurationProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_PublicAccessBlockConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3objectlambda.CfnAccessPointPropsMixin.TransformationConfigurationProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_TransformationConfigurationProperty)(nil)).Elem(),
	)
}
