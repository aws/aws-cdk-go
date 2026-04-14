package awss3express

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnAccessPointMixinProps",
		reflect.TypeOf((*CfnAccessPointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnAccessPointPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnAccessPointPropsMixin.PublicAccessBlockConfigurationProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_PublicAccessBlockConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnAccessPointPropsMixin.ScopeProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_ScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnAccessPointPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnBucketPolicyMixinProps",
		reflect.TypeOf((*CfnBucketPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnBucketPolicyPropsMixin",
		reflect.TypeOf((*CfnBucketPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBucketPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnDirectoryBucketMixinProps",
		reflect.TypeOf((*CfnDirectoryBucketMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnDirectoryBucketPropsMixin",
		reflect.TypeOf((*CfnDirectoryBucketPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDirectoryBucketPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnDirectoryBucketPropsMixin.AbortIncompleteMultipartUploadProperty",
		reflect.TypeOf((*CfnDirectoryBucketPropsMixin_AbortIncompleteMultipartUploadProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnDirectoryBucketPropsMixin.BucketEncryptionProperty",
		reflect.TypeOf((*CfnDirectoryBucketPropsMixin_BucketEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnDirectoryBucketPropsMixin.LifecycleConfigurationProperty",
		reflect.TypeOf((*CfnDirectoryBucketPropsMixin_LifecycleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnDirectoryBucketPropsMixin.MetricsConfigurationProperty",
		reflect.TypeOf((*CfnDirectoryBucketPropsMixin_MetricsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnDirectoryBucketPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnDirectoryBucketPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnDirectoryBucketPropsMixin.ServerSideEncryptionByDefaultProperty",
		reflect.TypeOf((*CfnDirectoryBucketPropsMixin_ServerSideEncryptionByDefaultProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_s3express.CfnDirectoryBucketPropsMixin.ServerSideEncryptionRuleProperty",
		reflect.TypeOf((*CfnDirectoryBucketPropsMixin_ServerSideEncryptionRuleProperty)(nil)).Elem(),
	)
}
