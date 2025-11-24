package mixinsawscodestar

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codestar.mixins.CfnGitHubRepositoryMixinProps",
		reflect.TypeOf((*CfnGitHubRepositoryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codestar.mixins.CfnGitHubRepositoryPropsMixin",
		reflect.TypeOf((*CfnGitHubRepositoryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGitHubRepositoryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codestar.mixins.CfnGitHubRepositoryPropsMixin.CodeProperty",
		reflect.TypeOf((*CfnGitHubRepositoryPropsMixin_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codestar.mixins.CfnGitHubRepositoryPropsMixin.S3Property",
		reflect.TypeOf((*CfnGitHubRepositoryPropsMixin_S3Property)(nil)).Elem(),
	)
}
