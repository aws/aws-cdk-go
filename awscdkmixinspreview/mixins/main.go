package mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.mixins.CfnPropertyMixinOptions",
		reflect.TypeOf((*CfnPropertyMixinOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/mixins-preview.mixins.IMergeStrategy",
		reflect.TypeOf((*IMergeStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "apply", GoMethod: "Apply"},
		},
		func() interface{} {
			return &jsiiProxy_IMergeStrategy{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.mixins.PropertyMergeStrategy",
		reflect.TypeOf((*PropertyMergeStrategy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_PropertyMergeStrategy{}
		},
	)
}
