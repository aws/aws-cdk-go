package interfacesawsnetworksecuritymanager

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_networksecuritymanager.DeploymentReference",
		reflect.TypeOf((*DeploymentReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_networksecuritymanager.IDeploymentRef",
		reflect.TypeOf((*IDeploymentRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentRef", GoGetter: "DeploymentRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IDeploymentRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_networksecuritymanager.IRuleRef",
		reflect.TypeOf((*IRuleRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ruleRef", GoGetter: "RuleRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IRuleRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_networksecuritymanager.IScopeRef",
		reflect.TypeOf((*IScopeRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "scopeRef", GoGetter: "ScopeRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IScopeRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_networksecuritymanager.ITemplateRef",
		reflect.TypeOf((*ITemplateRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "templateRef", GoGetter: "TemplateRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ITemplateRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_networksecuritymanager.RuleReference",
		reflect.TypeOf((*RuleReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_networksecuritymanager.ScopeReference",
		reflect.TypeOf((*ScopeReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_networksecuritymanager.TemplateReference",
		reflect.TypeOf((*TemplateReference)(nil)).Elem(),
	)
}
