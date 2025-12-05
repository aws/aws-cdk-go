package core

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		reflect.TypeOf((*ConstructSelector)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConstructSelector{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/mixins-preview.core.IConstructSelector",
		reflect.TypeOf((*IConstructSelector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
		},
		func() interface{} {
			return &jsiiProxy_IConstructSelector{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/mixins-preview.core.IMixin",
		reflect.TypeOf((*IMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			return &jsiiProxy_IMixin{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.core.Mixin",
		reflect.TypeOf((*Mixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_Mixin{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.core.MixinApplicator",
		reflect.TypeOf((*MixinApplicator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "apply", GoMethod: "Apply"},
			_jsii_.MemberMethod{JsiiMethod: "mustApply", GoMethod: "MustApply"},
		},
		func() interface{} {
			return &jsiiProxy_MixinApplicator{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.core.Mixins",
		reflect.TypeOf((*Mixins)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Mixins{}
		},
	)
}
