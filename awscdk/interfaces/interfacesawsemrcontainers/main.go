package interfacesawsemrcontainers

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_emrcontainers.IVirtualClusterRef",
		reflect.TypeOf((*IVirtualClusterRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "virtualClusterRef", GoGetter: "VirtualClusterRef"},
		},
		func() interface{} {
			j := jsiiProxy_IVirtualClusterRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_emrcontainers.VirtualClusterReference",
		reflect.TypeOf((*VirtualClusterReference)(nil)).Elem(),
	)
}
