package awsecrmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ecr.mixins.RepositoryAutoDeleteImages",
		reflect.TypeOf((*RepositoryAutoDeleteImages)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_RepositoryAutoDeleteImages{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			return &j
		},
	)
}
