package interfacesawscomprehend

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_comprehend.DocumentClassifierEndpointReference",
		reflect.TypeOf((*DocumentClassifierEndpointReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_comprehend.DocumentClassifierReference",
		reflect.TypeOf((*DocumentClassifierReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_comprehend.EntityRecognizerEndpointReference",
		reflect.TypeOf((*EntityRecognizerEndpointReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_comprehend.EntityRecognizerReference",
		reflect.TypeOf((*EntityRecognizerReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_comprehend.FlywheelReference",
		reflect.TypeOf((*FlywheelReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_comprehend.IDocumentClassifierEndpointRef",
		reflect.TypeOf((*IDocumentClassifierEndpointRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "documentClassifierEndpointRef", GoGetter: "DocumentClassifierEndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IDocumentClassifierEndpointRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_comprehend.IDocumentClassifierRef",
		reflect.TypeOf((*IDocumentClassifierRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "documentClassifierRef", GoGetter: "DocumentClassifierRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IDocumentClassifierRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_comprehend.IEntityRecognizerEndpointRef",
		reflect.TypeOf((*IEntityRecognizerEndpointRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "entityRecognizerEndpointRef", GoGetter: "EntityRecognizerEndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IEntityRecognizerEndpointRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_comprehend.IEntityRecognizerRef",
		reflect.TypeOf((*IEntityRecognizerRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "entityRecognizerRef", GoGetter: "EntityRecognizerRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IEntityRecognizerRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_comprehend.IFlywheelRef",
		reflect.TypeOf((*IFlywheelRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flywheelRef", GoGetter: "FlywheelRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IFlywheelRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
}
