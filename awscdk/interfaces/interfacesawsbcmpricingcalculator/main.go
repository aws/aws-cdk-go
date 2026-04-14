package interfacesawsbcmpricingcalculator

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bcmpricingcalculator.BillScenarioReference",
		reflect.TypeOf((*BillScenarioReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bcmpricingcalculator.IBillScenarioRef",
		reflect.TypeOf((*IBillScenarioRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "billScenarioRef", GoGetter: "BillScenarioRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IBillScenarioRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
}
