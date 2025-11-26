package previewawscodeguruprofilermixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.mixins.CfnProfilingGroupMixinProps",
		reflect.TypeOf((*CfnProfilingGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.mixins.CfnProfilingGroupPropsMixin",
		reflect.TypeOf((*CfnProfilingGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProfilingGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.mixins.CfnProfilingGroupPropsMixin.AgentPermissionsProperty",
		reflect.TypeOf((*CfnProfilingGroupPropsMixin_AgentPermissionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.mixins.CfnProfilingGroupPropsMixin.ChannelProperty",
		reflect.TypeOf((*CfnProfilingGroupPropsMixin_ChannelProperty)(nil)).Elem(),
	)
}
