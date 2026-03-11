package awscodecommit

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codecommit.CfnRepositoryMixinProps",
		reflect.TypeOf((*CfnRepositoryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codecommit.CfnRepositoryPropsMixin",
		reflect.TypeOf((*CfnRepositoryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRepositoryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codecommit.CfnRepositoryPropsMixin.CodeProperty",
		reflect.TypeOf((*CfnRepositoryPropsMixin_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codecommit.CfnRepositoryPropsMixin.RepositoryTriggerProperty",
		reflect.TypeOf((*CfnRepositoryPropsMixin_RepositoryTriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codecommit.CfnRepositoryPropsMixin.S3Property",
		reflect.TypeOf((*CfnRepositoryPropsMixin_S3Property)(nil)).Elem(),
	)
}
