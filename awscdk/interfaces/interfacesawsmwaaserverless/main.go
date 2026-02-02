package interfacesawsmwaaserverless

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_mwaaserverless.IWorkflowRef",
		reflect.TypeOf((*IWorkflowRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "workflowRef", GoGetter: "WorkflowRef"},
		},
		func() interface{} {
			j := jsiiProxy_IWorkflowRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_mwaaserverless.WorkflowReference",
		reflect.TypeOf((*WorkflowReference)(nil)).Elem(),
	)
}
