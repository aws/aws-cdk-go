package awscdkintegtestsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.IntegTest",
		reflect.TypeOf((*IntegTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IntegTest{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.IntegTestCase",
		reflect.TypeOf((*IntegTestCase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IntegTestCase{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseProps",
		reflect.TypeOf((*IntegTestCaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.IntegTestProps",
		reflect.TypeOf((*IntegTestProps)(nil)).Elem(),
	)
}
