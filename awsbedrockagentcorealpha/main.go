// The CDK Construct Library for Amazon Bedrock
package awsbedrockagentcorealpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AddEndpointOptions",
		reflect.TypeOf((*AddEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AgentRuntimeArtifact",
		reflect.TypeOf((*AgentRuntimeArtifact)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_AgentRuntimeArtifact{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AgentRuntimeAttributes",
		reflect.TypeOf((*AgentRuntimeAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserCustom",
		reflect.TypeOf((*BrowserCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "browserArn", GoGetter: "BrowserArn"},
			_jsii_.MemberProperty{JsiiProperty: "browserId", GoGetter: "BrowserId"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "failureReason", GoGetter: "FailureReason"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedAt", GoGetter: "LastUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrorsForApiOperation", GoMethod: "MetricErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricForApiOperation", GoMethod: "MetricForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsForApiOperation", GoMethod: "MetricInvocationsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatencyForApiOperation", GoMethod: "MetricLatencyForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionDuration", GoMethod: "MetricSessionDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrorsForApiOperation", GoMethod: "MetricSystemErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricTakeOverCount", GoMethod: "MetricTakeOverCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTakeOverDuration", GoMethod: "MetricTakeOverDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricTakeOverReleaseCount", GoMethod: "MetricTakeOverReleaseCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottlesForApiOperation", GoMethod: "MetricThrottlesForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrorsForApiOperation", GoMethod: "MetricUserErrorsForApiOperation"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "recordingConfig", GoGetter: "RecordingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BrowserCustom{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BrowserCustomBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserCustomAttributes",
		reflect.TypeOf((*BrowserCustomAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserCustomBase",
		reflect.TypeOf((*BrowserCustomBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "browserArn", GoGetter: "BrowserArn"},
			_jsii_.MemberProperty{JsiiProperty: "browserId", GoGetter: "BrowserId"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedAt", GoGetter: "LastUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrorsForApiOperation", GoMethod: "MetricErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricForApiOperation", GoMethod: "MetricForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsForApiOperation", GoMethod: "MetricInvocationsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatencyForApiOperation", GoMethod: "MetricLatencyForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionDuration", GoMethod: "MetricSessionDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrorsForApiOperation", GoMethod: "MetricSystemErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricTakeOverCount", GoMethod: "MetricTakeOverCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTakeOverDuration", GoMethod: "MetricTakeOverDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricTakeOverReleaseCount", GoMethod: "MetricTakeOverReleaseCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottlesForApiOperation", GoMethod: "MetricThrottlesForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrorsForApiOperation", GoMethod: "MetricUserErrorsForApiOperation"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BrowserCustomBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBrowserCustom)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserCustomProps",
		reflect.TypeOf((*BrowserCustomProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserNetworkConfiguration",
		reflect.TypeOf((*BrowserNetworkConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "networkMode", GoGetter: "NetworkMode"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			j := jsiiProxy_BrowserNetworkConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NetworkConfiguration)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterCustom",
		reflect.TypeOf((*CodeInterpreterCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterArn", GoGetter: "CodeInterpreterArn"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterId", GoGetter: "CodeInterpreterId"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "failureReason", GoGetter: "FailureReason"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedAt", GoGetter: "LastUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrorsForApiOperation", GoMethod: "MetricErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricForApiOperation", GoMethod: "MetricForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsForApiOperation", GoMethod: "MetricInvocationsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatencyForApiOperation", GoMethod: "MetricLatencyForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionDuration", GoMethod: "MetricSessionDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrorsForApiOperation", GoMethod: "MetricSystemErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottlesForApiOperation", GoMethod: "MetricThrottlesForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrorsForApiOperation", GoMethod: "MetricUserErrorsForApiOperation"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CodeInterpreterCustom{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CodeInterpreterCustomBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterCustomAttributes",
		reflect.TypeOf((*CodeInterpreterCustomAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterCustomBase",
		reflect.TypeOf((*CodeInterpreterCustomBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterArn", GoGetter: "CodeInterpreterArn"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterId", GoGetter: "CodeInterpreterId"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedAt", GoGetter: "LastUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrorsForApiOperation", GoMethod: "MetricErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricForApiOperation", GoMethod: "MetricForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsForApiOperation", GoMethod: "MetricInvocationsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatencyForApiOperation", GoMethod: "MetricLatencyForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionDuration", GoMethod: "MetricSessionDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrorsForApiOperation", GoMethod: "MetricSystemErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottlesForApiOperation", GoMethod: "MetricThrottlesForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrorsForApiOperation", GoMethod: "MetricUserErrorsForApiOperation"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CodeInterpreterCustomBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICodeInterpreterCustom)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterCustomProps",
		reflect.TypeOf((*CodeInterpreterCustomProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterNetworkConfiguration",
		reflect.TypeOf((*CodeInterpreterNetworkConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "networkMode", GoGetter: "NetworkMode"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			j := jsiiProxy_CodeInterpreterNetworkConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NetworkConfiguration)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IBedrockAgentRuntime",
		reflect.TypeOf((*IBedrockAgentRuntime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeArn", GoGetter: "AgentRuntimeArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeId", GoGetter: "AgentRuntimeId"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeName", GoGetter: "AgentRuntimeName"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeVersion", GoGetter: "AgentRuntimeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "agentStatus", GoGetter: "AgentStatus"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeRuntime", GoMethod: "GrantInvokeRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeRuntimeForUser", GoMethod: "GrantInvokeRuntimeForUser"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedAt", GoGetter: "LastUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsAggregated", GoMethod: "MetricInvocationsAggregated"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionCount", GoMethod: "MetricSessionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionsAggregated", GoMethod: "MetricSessionsAggregated"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrors", GoMethod: "MetricSystemErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottles", GoMethod: "MetricThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalErrors", GoMethod: "MetricTotalErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrors", GoMethod: "MetricUserErrors"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IBedrockAgentRuntime{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IBrowserCustom",
		reflect.TypeOf((*IBrowserCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "browserArn", GoGetter: "BrowserArn"},
			_jsii_.MemberProperty{JsiiProperty: "browserId", GoGetter: "BrowserId"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedAt", GoGetter: "LastUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrorsForApiOperation", GoMethod: "MetricErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricForApiOperation", GoMethod: "MetricForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsForApiOperation", GoMethod: "MetricInvocationsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatencyForApiOperation", GoMethod: "MetricLatencyForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionDuration", GoMethod: "MetricSessionDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrorsForApiOperation", GoMethod: "MetricSystemErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricTakeOverCount", GoMethod: "MetricTakeOverCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTakeOverDuration", GoMethod: "MetricTakeOverDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricTakeOverReleaseCount", GoMethod: "MetricTakeOverReleaseCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottlesForApiOperation", GoMethod: "MetricThrottlesForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrorsForApiOperation", GoMethod: "MetricUserErrorsForApiOperation"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
		},
		func() interface{} {
			j := jsiiProxy_IBrowserCustom{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ICodeInterpreterCustom",
		reflect.TypeOf((*ICodeInterpreterCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterArn", GoGetter: "CodeInterpreterArn"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterId", GoGetter: "CodeInterpreterId"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedAt", GoGetter: "LastUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrorsForApiOperation", GoMethod: "MetricErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricForApiOperation", GoMethod: "MetricForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsForApiOperation", GoMethod: "MetricInvocationsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatencyForApiOperation", GoMethod: "MetricLatencyForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionDuration", GoMethod: "MetricSessionDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrorsForApiOperation", GoMethod: "MetricSystemErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottlesForApiOperation", GoMethod: "MetricThrottlesForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrorsForApiOperation", GoMethod: "MetricUserErrorsForApiOperation"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
		},
		func() interface{} {
			j := jsiiProxy_ICodeInterpreterCustom{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IMemory",
		reflect.TypeOf((*IMemory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAdmin", GoMethod: "GrantAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantDeleteLongTermMemory", GoMethod: "GrantDeleteLongTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantDeleteShortTermMemory", GoMethod: "GrantDeleteShortTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadLongTermMemory", GoMethod: "GrantReadLongTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadShortTermMemory", GoMethod: "GrantReadShortTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "memoryArn", GoGetter: "MemoryArn"},
			_jsii_.MemberProperty{JsiiProperty: "memoryId", GoGetter: "MemoryId"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrorsForApiOperation", GoMethod: "MetricErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricEventCreationCount", GoMethod: "MetricEventCreationCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricForApiOperation", GoMethod: "MetricForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsForApiOperation", GoMethod: "MetricInvocationsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatencyForApiOperation", GoMethod: "MetricLatencyForApiOperation"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
		},
		func() interface{} {
			j := jsiiProxy_IMemory{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IMemoryStrategy",
		reflect.TypeOf((*IMemoryStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
			_jsii_.MemberProperty{JsiiProperty: "strategyType", GoGetter: "StrategyType"},
		},
		func() interface{} {
			return &jsiiProxy_IMemoryStrategy{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IRuntimeEndpoint",
		reflect.TypeOf((*IRuntimeEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeArn", GoGetter: "AgentRuntimeArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeEndpointArn", GoGetter: "AgentRuntimeEndpointArn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "endpointName", GoGetter: "EndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "liveVersion", GoGetter: "LiveVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "targetVersion", GoGetter: "TargetVersion"},
		},
		func() interface{} {
			j := jsiiProxy_IRuntimeEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InvocationConfiguration",
		reflect.TypeOf((*InvocationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ManagedMemoryStrategy",
		reflect.TypeOf((*ManagedMemoryStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "consolidationOverride", GoGetter: "ConsolidationOverride"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "extractionOverride", GoGetter: "ExtractionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespaces", GoGetter: "Namespaces"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
			_jsii_.MemberProperty{JsiiProperty: "strategyType", GoGetter: "StrategyType"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedMemoryStrategy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMemoryStrategy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ManagedStrategyProps",
		reflect.TypeOf((*ManagedStrategyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.Memory",
		reflect.TypeOf((*Memory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMemoryStrategy", GoMethod: "AddMemoryStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "expirationDuration", GoGetter: "ExpirationDuration"},
			_jsii_.MemberProperty{JsiiProperty: "failureReason", GoGetter: "FailureReason"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAdmin", GoMethod: "GrantAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantDeleteLongTermMemory", GoMethod: "GrantDeleteLongTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantDeleteShortTermMemory", GoMethod: "GrantDeleteShortTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadLongTermMemory", GoMethod: "GrantReadLongTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadShortTermMemory", GoMethod: "GrantReadShortTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "memoryArn", GoGetter: "MemoryArn"},
			_jsii_.MemberProperty{JsiiProperty: "memoryId", GoGetter: "MemoryId"},
			_jsii_.MemberProperty{JsiiProperty: "memoryName", GoGetter: "MemoryName"},
			_jsii_.MemberProperty{JsiiProperty: "memoryStrategies", GoGetter: "MemoryStrategies"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrorsForApiOperation", GoMethod: "MetricErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricEventCreationCount", GoMethod: "MetricEventCreationCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricForApiOperation", GoMethod: "MetricForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsForApiOperation", GoMethod: "MetricInvocationsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatencyForApiOperation", GoMethod: "MetricLatencyForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricMemoryRecordCreationCount", GoMethod: "MetricMemoryRecordCreationCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
		},
		func() interface{} {
			j := jsiiProxy_Memory{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_MemoryBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.MemoryAttributes",
		reflect.TypeOf((*MemoryAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.MemoryBase",
		reflect.TypeOf((*MemoryBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAdmin", GoMethod: "GrantAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantDeleteLongTermMemory", GoMethod: "GrantDeleteLongTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantDeleteShortTermMemory", GoMethod: "GrantDeleteShortTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadLongTermMemory", GoMethod: "GrantReadLongTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadShortTermMemory", GoMethod: "GrantReadShortTermMemory"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "memoryArn", GoGetter: "MemoryArn"},
			_jsii_.MemberProperty{JsiiProperty: "memoryId", GoGetter: "MemoryId"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrorsForApiOperation", GoMethod: "MetricErrorsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricEventCreationCount", GoMethod: "MetricEventCreationCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricForApiOperation", GoMethod: "MetricForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsForApiOperation", GoMethod: "MetricInvocationsForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatencyForApiOperation", GoMethod: "MetricLatencyForApiOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricMemoryRecordCreationCount", GoMethod: "MetricMemoryRecordCreationCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
		},
		func() interface{} {
			j := jsiiProxy_MemoryBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMemory)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.MemoryProps",
		reflect.TypeOf((*MemoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.MemoryStrategy",
		reflect.TypeOf((*MemoryStrategy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MemoryStrategy{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.MemoryStrategyCommonProps",
		reflect.TypeOf((*MemoryStrategyCommonProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.MemoryStrategyType",
		reflect.TypeOf((*MemoryStrategyType)(nil)).Elem(),
		map[string]interface{}{
			"SUMMARIZATION": MemoryStrategyType_SUMMARIZATION,
			"SEMANTIC": MemoryStrategyType_SEMANTIC,
			"USER_PREFERENCE": MemoryStrategyType_USER_PREFERENCE,
			"CUSTOM": MemoryStrategyType_CUSTOM,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.NetworkConfiguration",
		reflect.TypeOf((*NetworkConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "networkMode", GoGetter: "NetworkMode"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OverrideConfig",
		reflect.TypeOf((*OverrideConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ProtocolType",
		reflect.TypeOf((*ProtocolType)(nil)).Elem(),
		map[string]interface{}{
			"MCP": ProtocolType_MCP,
			"HTTP": ProtocolType_HTTP,
			"A2A": ProtocolType_A2A,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RecordingConfig",
		reflect.TypeOf((*RecordingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.Runtime",
		reflect.TypeOf((*Runtime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEndpoint", GoMethod: "AddEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeArn", GoGetter: "AgentRuntimeArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeArtifact", GoGetter: "AgentRuntimeArtifact"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeId", GoGetter: "AgentRuntimeId"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeName", GoGetter: "AgentRuntimeName"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeVersion", GoGetter: "AgentRuntimeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "agentStatus", GoGetter: "AgentStatus"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeRuntime", GoMethod: "GrantInvokeRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeRuntimeForUser", GoMethod: "GrantInvokeRuntimeForUser"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedAt", GoGetter: "LastUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsAggregated", GoMethod: "MetricInvocationsAggregated"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionCount", GoMethod: "MetricSessionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionsAggregated", GoMethod: "MetricSessionsAggregated"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrors", GoMethod: "MetricSystemErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottles", GoMethod: "MetricThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalErrors", GoMethod: "MetricTotalErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrors", GoMethod: "MetricUserErrors"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Runtime{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RuntimeBase)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeAuthorizerConfiguration",
		reflect.TypeOf((*RuntimeAuthorizerConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RuntimeAuthorizerConfiguration{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeBase",
		reflect.TypeOf((*RuntimeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeArn", GoGetter: "AgentRuntimeArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeId", GoGetter: "AgentRuntimeId"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeName", GoGetter: "AgentRuntimeName"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeVersion", GoGetter: "AgentRuntimeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "agentStatus", GoGetter: "AgentStatus"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeRuntime", GoMethod: "GrantInvokeRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeRuntimeForUser", GoMethod: "GrantInvokeRuntimeForUser"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedAt", GoGetter: "LastUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsAggregated", GoMethod: "MetricInvocationsAggregated"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionCount", GoMethod: "MetricSessionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricSessionsAggregated", GoMethod: "MetricSessionsAggregated"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrors", GoMethod: "MetricSystemErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottles", GoMethod: "MetricThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalErrors", GoMethod: "MetricTotalErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrors", GoMethod: "MetricUserErrors"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBedrockAgentRuntime)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeEndpoint",
		reflect.TypeOf((*RuntimeEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeArn", GoGetter: "AgentRuntimeArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeEndpointArn", GoGetter: "AgentRuntimeEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeId", GoGetter: "AgentRuntimeId"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeVersion", GoGetter: "AgentRuntimeVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "endpointId", GoGetter: "EndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "endpointName", GoGetter: "EndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedAt", GoGetter: "LastUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "liveVersion", GoGetter: "LiveVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "targetVersion", GoGetter: "TargetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RuntimeEndpointBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeEndpointAttributes",
		reflect.TypeOf((*RuntimeEndpointAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeEndpointBase",
		reflect.TypeOf((*RuntimeEndpointBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeArn", GoGetter: "AgentRuntimeArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeEndpointArn", GoGetter: "AgentRuntimeEndpointArn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "endpointName", GoGetter: "EndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "liveVersion", GoGetter: "LiveVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "targetVersion", GoGetter: "TargetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeEndpointBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRuntimeEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeEndpointProps",
		reflect.TypeOf((*RuntimeEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeNetworkConfiguration",
		reflect.TypeOf((*RuntimeNetworkConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "networkMode", GoGetter: "NetworkMode"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeNetworkConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NetworkConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeProps",
		reflect.TypeOf((*RuntimeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SelfManagedMemoryStrategy",
		reflect.TypeOf((*SelfManagedMemoryStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "historicalContextWindowSize", GoGetter: "HistoricalContextWindowSize"},
			_jsii_.MemberProperty{JsiiProperty: "invocationConfiguration", GoGetter: "InvocationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
			_jsii_.MemberProperty{JsiiProperty: "strategyType", GoGetter: "StrategyType"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConditions", GoGetter: "TriggerConditions"},
		},
		func() interface{} {
			j := jsiiProxy_SelfManagedMemoryStrategy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMemoryStrategy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SelfManagedStrategyProps",
		reflect.TypeOf((*SelfManagedStrategyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.TriggerConditions",
		reflect.TypeOf((*TriggerConditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.VpcConfigProps",
		reflect.TypeOf((*VpcConfigProps)(nil)).Elem(),
	)
}
