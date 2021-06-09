package awselasticloadbalancingv2actions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		reflect.TypeOf((*AuthenticateCognitoAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "next", GoGetter: "Next"},
			_jsii_.MemberMethod{JsiiMethod: "renderActions", GoMethod: "RenderActions"},
			_jsii_.MemberMethod{JsiiMethod: "renumber", GoMethod: "Renumber"},
		},
		func() interface{} {
			j := jsiiProxy_AuthenticateCognitoAction{}
			_jsii_.InitJsiiProxy(&j.Type__awselasticloadbalancingv2ListenerAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2_actions.AuthenticateCognitoActionProps",
		reflect.TypeOf((*AuthenticateCognitoActionProps)(nil)).Elem(),
	)
}
