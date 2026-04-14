package interfacesawsbedrockagentcore

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.ApiKeyCredentialProviderReference",
		reflect.TypeOf((*ApiKeyCredentialProviderReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.BrowserCustomReference",
		reflect.TypeOf((*BrowserCustomReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.BrowserProfileReference",
		reflect.TypeOf((*BrowserProfileReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.CodeInterpreterCustomReference",
		reflect.TypeOf((*CodeInterpreterCustomReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.EvaluatorReference",
		reflect.TypeOf((*EvaluatorReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.GatewayReference",
		reflect.TypeOf((*GatewayReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.GatewayTargetReference",
		reflect.TypeOf((*GatewayTargetReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IApiKeyCredentialProviderRef",
		reflect.TypeOf((*IApiKeyCredentialProviderRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiKeyCredentialProviderRef", GoGetter: "ApiKeyCredentialProviderRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IApiKeyCredentialProviderRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IBrowserCustomRef",
		reflect.TypeOf((*IBrowserCustomRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "browserCustomRef", GoGetter: "BrowserCustomRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IBrowserCustomRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IBrowserProfileRef",
		reflect.TypeOf((*IBrowserProfileRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "browserProfileRef", GoGetter: "BrowserProfileRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IBrowserProfileRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.ICodeInterpreterCustomRef",
		reflect.TypeOf((*ICodeInterpreterCustomRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterCustomRef", GoGetter: "CodeInterpreterCustomRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ICodeInterpreterCustomRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IEvaluatorRef",
		reflect.TypeOf((*IEvaluatorRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorRef", GoGetter: "EvaluatorRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IEvaluatorRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IGatewayRef",
		reflect.TypeOf((*IGatewayRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRef", GoGetter: "GatewayRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IGatewayRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IGatewayTargetRef",
		reflect.TypeOf((*IGatewayTargetRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetRef", GoGetter: "GatewayTargetRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IGatewayTargetRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IMemoryRef",
		reflect.TypeOf((*IMemoryRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "memoryRef", GoGetter: "MemoryRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IMemoryRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IOAuth2CredentialProviderRef",
		reflect.TypeOf((*IOAuth2CredentialProviderRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oAuth2CredentialProviderRef", GoGetter: "OAuth2CredentialProviderRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IOAuth2CredentialProviderRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IOnlineEvaluationConfigRef",
		reflect.TypeOf((*IOnlineEvaluationConfigRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigRef", GoGetter: "OnlineEvaluationConfigRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IOnlineEvaluationConfigRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IPolicyEngineRef",
		reflect.TypeOf((*IPolicyEngineRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineRef", GoGetter: "PolicyEngineRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IPolicyEngineRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IPolicyRef",
		reflect.TypeOf((*IPolicyRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policyRef", GoGetter: "PolicyRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IPolicyRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IRuntimeEndpointRef",
		reflect.TypeOf((*IRuntimeEndpointRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeEndpointRef", GoGetter: "RuntimeEndpointRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IRuntimeEndpointRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IRuntimeRef",
		reflect.TypeOf((*IRuntimeRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeRef", GoGetter: "RuntimeRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IRuntimeRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.IWorkloadIdentityRef",
		reflect.TypeOf((*IWorkloadIdentityRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "workloadIdentityRef", GoGetter: "WorkloadIdentityRef"},
		},
		func() interface{} {
			j := jsiiProxy_IWorkloadIdentityRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.MemoryReference",
		reflect.TypeOf((*MemoryReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.OAuth2CredentialProviderReference",
		reflect.TypeOf((*OAuth2CredentialProviderReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.OnlineEvaluationConfigReference",
		reflect.TypeOf((*OnlineEvaluationConfigReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.PolicyEngineReference",
		reflect.TypeOf((*PolicyEngineReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.PolicyReference",
		reflect.TypeOf((*PolicyReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.RuntimeEndpointReference",
		reflect.TypeOf((*RuntimeEndpointReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.RuntimeReference",
		reflect.TypeOf((*RuntimeReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_bedrockagentcore.WorkloadIdentityReference",
		reflect.TypeOf((*WorkloadIdentityReference)(nil)).Elem(),
	)
}
