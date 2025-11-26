package previewawscloud9mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloud9.mixins.CfnEnvironmentEC2MixinProps",
		reflect.TypeOf((*CfnEnvironmentEC2MixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloud9.mixins.CfnEnvironmentEC2PropsMixin",
		reflect.TypeOf((*CfnEnvironmentEC2PropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentEC2PropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloud9.mixins.CfnEnvironmentEC2PropsMixin.RepositoryProperty",
		reflect.TypeOf((*CfnEnvironmentEC2PropsMixin_RepositoryProperty)(nil)).Elem(),
	)
}
