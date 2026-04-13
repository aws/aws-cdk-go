package interfacesawsnovaact

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_novaact.IWorkflowDefinitionRef",
		reflect.TypeOf((*IWorkflowDefinitionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "workflowDefinitionRef", GoGetter: "WorkflowDefinitionRef"},
		},
		func() interface{} {
			j := jsiiProxy_IWorkflowDefinitionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_novaact.WorkflowDefinitionReference",
		reflect.TypeOf((*WorkflowDefinitionReference)(nil)).Elem(),
	)
}
