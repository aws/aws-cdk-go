// The CDK Construct Library for Amazon EventBridge Pipes Targets
package awscdkpipestargetsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-targets-alpha.SqsTarget",
		reflect.TypeOf((*SqsTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_SqsTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaITarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-targets-alpha.SqsTargetParameters",
		reflect.TypeOf((*SqsTargetParameters)(nil)).Elem(),
	)
}
