package interfaces

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.IEnvironmentAware",
		reflect.TypeOf((*IEnvironmentAware)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
		},
		func() interface{} {
			return &jsiiProxy_IEnvironmentAware{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.ResourceEnvironment",
		reflect.TypeOf((*ResourceEnvironment)(nil)).Elem(),
	)
}
