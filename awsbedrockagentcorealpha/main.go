// The CDK Construct Library for Amazon Bedrock
package awsbedrockagentcorealpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AddApiGatewayTargetOptions",
		reflect.TypeOf((*AddApiGatewayTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AddEndpointOptions",
		reflect.TypeOf((*AddEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AddLambdaTargetOptions",
		reflect.TypeOf((*AddLambdaTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AddMcpServerTargetOptions",
		reflect.TypeOf((*AddMcpServerTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AddOpenApiTargetOptions",
		reflect.TypeOf((*AddOpenApiTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AddPolicyOptions",
		reflect.TypeOf((*AddPolicyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AddSmithyTargetOptions",
		reflect.TypeOf((*AddSmithyTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AgentCoreRuntime",
		reflect.TypeOf((*AgentCoreRuntime)(nil)).Elem(),
		map[string]interface{}{
			"PYTHON_3_10": AgentCoreRuntime_PYTHON_3_10,
			"PYTHON_3_11": AgentCoreRuntime_PYTHON_3_11,
			"PYTHON_3_12": AgentCoreRuntime_PYTHON_3_12,
			"PYTHON_3_13": AgentCoreRuntime_PYTHON_3_13,
			"PYTHON_3_14": AgentCoreRuntime_PYTHON_3_14,
			"NODE_22": AgentCoreRuntime_NODE_22,
		},
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
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiGatewayHttpMethod",
		reflect.TypeOf((*ApiGatewayHttpMethod)(nil)).Elem(),
		map[string]interface{}{
			"GET": ApiGatewayHttpMethod_GET,
			"POST": ApiGatewayHttpMethod_POST,
			"PUT": ApiGatewayHttpMethod_PUT,
			"DELETE": ApiGatewayHttpMethod_DELETE,
			"PATCH": ApiGatewayHttpMethod_PATCH,
			"HEAD": ApiGatewayHttpMethod_HEAD,
			"OPTIONS": ApiGatewayHttpMethod_OPTIONS,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiGatewayTargetConfiguration",
		reflect.TypeOf((*ApiGatewayTargetConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayToolConfiguration", GoGetter: "ApiGatewayToolConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "metadataConfiguration", GoGetter: "MetadataConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "renderMcpConfiguration", GoMethod: "RenderMcpConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "restApiId", GoGetter: "RestApiId"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayTargetConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_McpTargetConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiGatewayTargetConfigurationProps",
		reflect.TypeOf((*ApiGatewayTargetConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiGatewayToolConfiguration",
		reflect.TypeOf((*ApiGatewayToolConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiGatewayToolFilter",
		reflect.TypeOf((*ApiGatewayToolFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiGatewayToolOverride",
		reflect.TypeOf((*ApiGatewayToolOverride)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiKeyAdditionalConfiguration",
		reflect.TypeOf((*ApiKeyAdditionalConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiKeyCredentialLocation",
		reflect.TypeOf((*ApiKeyCredentialLocation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "credentialLocationType", GoGetter: "CredentialLocationType"},
			_jsii_.MemberProperty{JsiiProperty: "credentialParameterName", GoGetter: "CredentialParameterName"},
			_jsii_.MemberProperty{JsiiProperty: "credentialPrefix", GoGetter: "CredentialPrefix"},
		},
		func() interface{} {
			return &jsiiProxy_ApiKeyCredentialLocation{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiKeyCredentialProvider",
		reflect.TypeOf((*ApiKeyCredentialProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiKeyCredentialProviderName", GoGetter: "ApiKeyCredentialProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyCredentialProviderRef", GoGetter: "ApiKeyCredentialProviderRef"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeySecretArn", GoGetter: "ApiKeySecretArn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindForGatewayApiKeyTarget", GoMethod: "BindForGatewayApiKeyTarget"},
			_jsii_.MemberProperty{JsiiProperty: "createdTime", GoGetter: "CreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderArn", GoGetter: "CredentialProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAdmin", GoMethod: "GrantAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedTime", GoGetter: "LastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ApiKeyCredentialProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApiKeyCredentialProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiKeyCredentialProviderAttributes",
		reflect.TypeOf((*ApiKeyCredentialProviderAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiKeyCredentialProviderIdentityPerms",
		reflect.TypeOf((*ApiKeyCredentialProviderIdentityPerms)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ApiKeyCredentialProviderIdentityPerms{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiKeyCredentialProviderProps",
		reflect.TypeOf((*ApiKeyCredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiKeyCredentialProviderResourceProps",
		reflect.TypeOf((*ApiKeyCredentialProviderResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiSchema",
		reflect.TypeOf((*ApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucketOwnerAccountId", GoGetter: "BucketOwnerAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "grantPermissionsToRole", GoMethod: "GrantPermissionsToRole"},
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			return &jsiiProxy_ApiSchema{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AssetApiSchema",
		reflect.TypeOf((*AssetApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucketOwnerAccountId", GoGetter: "BucketOwnerAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "grantPermissionsToRole", GoMethod: "GrantPermissionsToRole"},
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_AssetApiSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiSchema)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AssetToolSchema",
		reflect.TypeOf((*AssetToolSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucketOwnerAccountId", GoGetter: "BucketOwnerAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "grantPermissionsToRole", GoMethod: "GrantPermissionsToRole"},
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_AssetToolSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ToolSchema)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AtlassianOAuth2CredentialProviderProps",
		reflect.TypeOf((*AtlassianOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AttributeAccessor",
		reflect.TypeOf((*AttributeAccessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "contains", GoMethod: "Contains"},
			_jsii_.MemberMethod{JsiiMethod: "equalTo", GoMethod: "EqualTo"},
			_jsii_.MemberMethod{JsiiMethod: "greaterThan", GoMethod: "GreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "greaterThanOrEqualTo", GoMethod: "GreaterThanOrEqualTo"},
			_jsii_.MemberMethod{JsiiMethod: "isIn", GoMethod: "IsIn"},
			_jsii_.MemberMethod{JsiiMethod: "isInRange", GoMethod: "IsInRange"},
			_jsii_.MemberMethod{JsiiMethod: "lessThan", GoMethod: "LessThan"},
			_jsii_.MemberMethod{JsiiMethod: "lessThanOrEqualTo", GoMethod: "LessThanOrEqualTo"},
			_jsii_.MemberMethod{JsiiMethod: "notEqualTo", GoMethod: "NotEqualTo"},
		},
		func() interface{} {
			return &jsiiProxy_AttributeAccessor{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserCustom",
		reflect.TypeOf((*BrowserCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "browserArn", GoGetter: "BrowserArn"},
			_jsii_.MemberProperty{JsiiProperty: "browserCustomRef", GoGetter: "BrowserCustomRef"},
			_jsii_.MemberProperty{JsiiProperty: "browserId", GoGetter: "BrowserId"},
			_jsii_.MemberProperty{JsiiProperty: "browserSigning", GoGetter: "BrowserSigning"},
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
			_jsii_.MemberProperty{JsiiProperty: "browserCustomRef", GoGetter: "BrowserCustomRef"},
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserSigning",
		reflect.TypeOf((*BrowserSigning)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": BrowserSigning_ENABLED,
			"DISABLED": BrowserSigning_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		reflect.TypeOf((*BuiltinEvaluator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_BuiltinEvaluator{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CategoricalRatingOption",
		reflect.TypeOf((*CategoricalRatingOption)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CloudWatchLogsDataSourceConfig",
		reflect.TypeOf((*CloudWatchLogsDataSourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeAssetOptions",
		reflect.TypeOf((*CodeAssetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeBasedOptions",
		reflect.TypeOf((*CodeBasedOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterCustom",
		reflect.TypeOf((*CodeInterpreterCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterArn", GoGetter: "CodeInterpreterArn"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterCustomRef", GoGetter: "CodeInterpreterCustomRef"},
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterCustomRef", GoGetter: "CodeInterpreterCustomRef"},
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CognitoAuthorizerProps",
		reflect.TypeOf((*CognitoAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ConditionBuilder",
		reflect.TypeOf((*ConditionBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "and", GoMethod: "And"},
			_jsii_.MemberMethod{JsiiMethod: "contextAttribute", GoMethod: "ContextAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "or", GoMethod: "Or"},
			_jsii_.MemberMethod{JsiiMethod: "principalAttribute", GoMethod: "PrincipalAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resourceAttribute", GoMethod: "ResourceAttribute"},
		},
		func() interface{} {
			return &jsiiProxy_ConditionBuilder{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ConditionalAttributeAccessor",
		reflect.TypeOf((*ConditionalAttributeAccessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "contains", GoMethod: "Contains"},
			_jsii_.MemberMethod{JsiiMethod: "equalTo", GoMethod: "EqualTo"},
			_jsii_.MemberMethod{JsiiMethod: "greaterThan", GoMethod: "GreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "greaterThanOrEqualTo", GoMethod: "GreaterThanOrEqualTo"},
			_jsii_.MemberMethod{JsiiMethod: "isIn", GoMethod: "IsIn"},
			_jsii_.MemberMethod{JsiiMethod: "isInRange", GoMethod: "IsInRange"},
			_jsii_.MemberMethod{JsiiMethod: "lessThan", GoMethod: "LessThan"},
			_jsii_.MemberMethod{JsiiMethod: "lessThanOrEqualTo", GoMethod: "LessThanOrEqualTo"},
			_jsii_.MemberMethod{JsiiMethod: "notEqualTo", GoMethod: "NotEqualTo"},
		},
		func() interface{} {
			return &jsiiProxy_ConditionalAttributeAccessor{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ConditionalPolicyStatement",
		reflect.TypeOf((*ConditionalPolicyStatement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "and", GoMethod: "And"},
			_jsii_.MemberMethod{JsiiMethod: "contextAttribute", GoMethod: "ContextAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "done", GoMethod: "Done"},
			_jsii_.MemberMethod{JsiiMethod: "or", GoMethod: "Or"},
			_jsii_.MemberMethod{JsiiMethod: "principalAttribute", GoMethod: "PrincipalAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resourceAttribute", GoMethod: "ResourceAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "unless", GoMethod: "Unless"},
		},
		func() interface{} {
			return &jsiiProxy_ConditionalPolicyStatement{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CredentialProviderType",
		reflect.TypeOf((*CredentialProviderType)(nil)).Elem(),
		map[string]interface{}{
			"API_KEY": CredentialProviderType_API_KEY,
			"OAUTH": CredentialProviderType_OAUTH,
			"GATEWAY_IAM_ROLE": CredentialProviderType_GATEWAY_IAM_ROLE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CustomClaimOperator",
		reflect.TypeOf((*CustomClaimOperator)(nil)).Elem(),
		map[string]interface{}{
			"EQUALS": CustomClaimOperator_EQUALS,
			"CONTAINS": CustomClaimOperator_CONTAINS,
			"CONTAINS_ANY": CustomClaimOperator_CONTAINS_ANY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CustomJwtAuthorizer",
		reflect.TypeOf((*CustomJwtAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizerType", GoGetter: "AuthorizerType"},
		},
		func() interface{} {
			j := jsiiProxy_CustomJwtAuthorizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGatewayAuthorizerConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CustomJwtConfiguration",
		reflect.TypeOf((*CustomJwtConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CustomOAuth2CredentialProviderProps",
		reflect.TypeOf((*CustomOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.DataSourceConfig",
		reflect.TypeOf((*DataSourceConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "cloudWatchLogsConfig", GoGetter: "CloudWatchLogsConfig"},
		},
		func() interface{} {
			return &jsiiProxy_DataSourceConfig{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.DataSourceConfigBindResult",
		reflect.TypeOf((*DataSourceConfigBindResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.DropboxOAuth2CredentialProviderProps",
		reflect.TypeOf((*DropboxOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EpisodicReflectionConfiguration",
		reflect.TypeOf((*EpisodicReflectionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluationLevel",
		reflect.TypeOf((*EvaluationLevel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_EvaluationLevel{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.Evaluator",
		reflect.TypeOf((*Evaluator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorArn", GoGetter: "EvaluatorArn"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorId", GoGetter: "EvaluatorId"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorName", GoGetter: "EvaluatorName"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorRef", GoGetter: "EvaluatorRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Evaluator{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EvaluatorBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorAttributes",
		reflect.TypeOf((*EvaluatorAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorBase",
		reflect.TypeOf((*EvaluatorBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorArn", GoGetter: "EvaluatorArn"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorId", GoGetter: "EvaluatorId"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorName", GoGetter: "EvaluatorName"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorRef", GoGetter: "EvaluatorRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_EvaluatorBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEvaluator)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorConfig",
		reflect.TypeOf((*EvaluatorConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
		},
		func() interface{} {
			return &jsiiProxy_EvaluatorConfig{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorInferenceConfig",
		reflect.TypeOf((*EvaluatorInferenceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorProps",
		reflect.TypeOf((*EvaluatorProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorRatingScale",
		reflect.TypeOf((*EvaluatorRatingScale)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EvaluatorRatingScale{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorReference",
		reflect.TypeOf((*EvaluatorReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorId", GoGetter: "EvaluatorId"},
		},
		func() interface{} {
			return &jsiiProxy_EvaluatorReference{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorReferenceBindResult",
		reflect.TypeOf((*EvaluatorReferenceBindResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ExecutionStatus",
		reflect.TypeOf((*ExecutionStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ExecutionStatus{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FacebookOAuth2CredentialProviderProps",
		reflect.TypeOf((*FacebookOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterConfig",
		reflect.TypeOf((*FilterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		reflect.TypeOf((*FilterOperator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_FilterOperator{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterValue",
		reflect.TypeOf((*FilterValue)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FilterValue{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FromApiKeyIdentityOptions",
		reflect.TypeOf((*FromApiKeyIdentityOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FromOauthIdentityOptions",
		reflect.TypeOf((*FromOauthIdentityOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.Gateway",
		reflect.TypeOf((*Gateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addApiGatewayTarget", GoMethod: "AddApiGatewayTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addInterceptor", GoMethod: "AddInterceptor"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaTarget", GoMethod: "AddLambdaTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addMcpServerTarget", GoMethod: "AddMcpServerTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOpenApiTarget", GoMethod: "AddOpenApiTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addSmithyTarget", GoMethod: "AddSmithyTarget"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerConfiguration", GoGetter: "AuthorizerConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "exceptionLevel", GoGetter: "ExceptionLevel"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayArn", GoGetter: "GatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayId", GoGetter: "GatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRef", GoGetter: "GatewayRef"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayUrl", GoGetter: "GatewayUrl"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricDuration", GoMethod: "MetricDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrors", GoMethod: "MetricSystemErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetExecutionTime", GoMethod: "MetricTargetExecutionTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetType", GoMethod: "MetricTargetType"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottles", GoMethod: "MetricThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrors", GoMethod: "MetricUserErrors"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oauthScopes", GoGetter: "OauthScopes"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineConfiguration", GoGetter: "PolicyEngineConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "protocolConfiguration", GoGetter: "ProtocolConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "resourceServer", GoGetter: "ResourceServer"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusReason", GoGetter: "StatusReason"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpointUrl", GoGetter: "TokenEndpointUrl"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "userPool", GoGetter: "UserPool"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolClient", GoGetter: "UserPoolClient"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolDomain", GoGetter: "UserPoolDomain"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Gateway{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GatewayBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayApiKeyIdentityBinding",
		reflect.TypeOf((*GatewayApiKeyIdentityBinding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAttributes",
		reflect.TypeOf((*GatewayAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizer",
		reflect.TypeOf((*GatewayAuthorizer)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GatewayAuthorizer{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizerType",
		reflect.TypeOf((*GatewayAuthorizerType)(nil)).Elem(),
		map[string]interface{}{
			"CUSTOM_JWT": GatewayAuthorizerType_CUSTOM_JWT,
			"AWS_IAM": GatewayAuthorizerType_AWS_IAM,
			"NONE": GatewayAuthorizerType_NONE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayBase",
		reflect.TypeOf((*GatewayBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerConfiguration", GoGetter: "AuthorizerConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "exceptionLevel", GoGetter: "ExceptionLevel"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayArn", GoGetter: "GatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayId", GoGetter: "GatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRef", GoGetter: "GatewayRef"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayUrl", GoGetter: "GatewayUrl"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricDuration", GoMethod: "MetricDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrors", GoMethod: "MetricSystemErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetExecutionTime", GoMethod: "MetricTargetExecutionTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetType", GoMethod: "MetricTargetType"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottles", GoMethod: "MetricThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrors", GoMethod: "MetricUserErrors"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "protocolConfiguration", GoGetter: "ProtocolConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusReason", GoGetter: "StatusReason"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_GatewayBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGateway)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayCredentialProvider",
		reflect.TypeOf((*GatewayCredentialProvider)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GatewayCredentialProvider{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayCustomClaim",
		reflect.TypeOf((*GatewayCustomClaim)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GatewayCustomClaim{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayExceptionLevel",
		reflect.TypeOf((*GatewayExceptionLevel)(nil)).Elem(),
		map[string]interface{}{
			"DEBUG": GatewayExceptionLevel_DEBUG,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayOAuth2IdentityBinding",
		reflect.TypeOf((*GatewayOAuth2IdentityBinding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayPolicyEngineConfig",
		reflect.TypeOf((*GatewayPolicyEngineConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayProps",
		reflect.TypeOf((*GatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayProtocol",
		reflect.TypeOf((*GatewayProtocol)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GatewayProtocol{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTarget",
		reflect.TypeOf((*GatewayTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderConfigurations", GoGetter: "CredentialProviderConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gateway", GoGetter: "Gateway"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetRef", GoGetter: "GatewayTargetRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantSync", GoMethod: "GrantSync"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusReasons", GoGetter: "StatusReasons"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetId", GoGetter: "TargetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetProtocolType", GoGetter: "TargetProtocolType"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_GatewayTarget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GatewayTargetBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMcpGatewayTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTargetApiGatewayProps",
		reflect.TypeOf((*GatewayTargetApiGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTargetAttributes",
		reflect.TypeOf((*GatewayTargetAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTargetBase",
		reflect.TypeOf((*GatewayTargetBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderConfigurations", GoGetter: "CredentialProviderConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gateway", GoGetter: "Gateway"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetRef", GoGetter: "GatewayTargetRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusReasons", GoGetter: "StatusReasons"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetId", GoGetter: "TargetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetProtocolType", GoGetter: "TargetProtocolType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_GatewayTargetBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGatewayTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTargetCommonProps",
		reflect.TypeOf((*GatewayTargetCommonProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTargetLambdaProps",
		reflect.TypeOf((*GatewayTargetLambdaProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTargetMcpServerProps",
		reflect.TypeOf((*GatewayTargetMcpServerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTargetOpenApiProps",
		reflect.TypeOf((*GatewayTargetOpenApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTargetProps",
		reflect.TypeOf((*GatewayTargetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTargetProtocolType",
		reflect.TypeOf((*GatewayTargetProtocolType)(nil)).Elem(),
		map[string]interface{}{
			"MCP": GatewayTargetProtocolType_MCP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayTargetSmithyProps",
		reflect.TypeOf((*GatewayTargetSmithyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GithubOAuth2CredentialProviderProps",
		reflect.TypeOf((*GithubOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GoogleOAuth2CredentialProviderProps",
		reflect.TypeOf((*GoogleOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.HubspotOAuth2CredentialProviderProps",
		reflect.TypeOf((*HubspotOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IApiKeyCredentialProvider",
		reflect.TypeOf((*IApiKeyCredentialProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiKeyCredentialProviderRef", GoGetter: "ApiKeyCredentialProviderRef"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeySecretArn", GoGetter: "ApiKeySecretArn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindForGatewayApiKeyTarget", GoMethod: "BindForGatewayApiKeyTarget"},
			_jsii_.MemberProperty{JsiiProperty: "createdTime", GoGetter: "CreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderArn", GoGetter: "CredentialProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAdmin", GoMethod: "GrantAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedTime", GoGetter: "LastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IApiKeyCredentialProvider{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIApiKeyCredentialProviderRef)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
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
			_jsii_.MemberProperty{JsiiProperty: "runtimeRef", GoGetter: "RuntimeRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IBedrockAgentRuntime{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIRuntimeRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IBrowserCustom",
		reflect.TypeOf((*IBrowserCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "browserArn", GoGetter: "BrowserArn"},
			_jsii_.MemberProperty{JsiiProperty: "browserCustomRef", GoGetter: "BrowserCustomRef"},
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IBrowserCustom{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIBrowserCustomRef)
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
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterCustomRef", GoGetter: "CodeInterpreterCustomRef"},
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ICodeInterpreterCustom{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreICodeInterpreterCustomRef)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ICredentialProviderConfig",
		reflect.TypeOf((*ICredentialProviderConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderType", GoGetter: "CredentialProviderType"},
			_jsii_.MemberMethod{JsiiMethod: "grantNeededPermissionsToRole", GoMethod: "GrantNeededPermissionsToRole"},
		},
		func() interface{} {
			return &jsiiProxy_ICredentialProviderConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IEvaluator",
		reflect.TypeOf((*IEvaluator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorArn", GoGetter: "EvaluatorArn"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorId", GoGetter: "EvaluatorId"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorName", GoGetter: "EvaluatorName"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorRef", GoGetter: "EvaluatorRef"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IEvaluator{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIEvaluatorRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IGateway",
		reflect.TypeOf((*IGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerConfiguration", GoGetter: "AuthorizerConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "exceptionLevel", GoGetter: "ExceptionLevel"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayArn", GoGetter: "GatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayId", GoGetter: "GatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRef", GoGetter: "GatewayRef"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayUrl", GoGetter: "GatewayUrl"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricDuration", GoMethod: "MetricDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrors", GoMethod: "MetricSystemErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetExecutionTime", GoMethod: "MetricTargetExecutionTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetType", GoMethod: "MetricTargetType"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottles", GoMethod: "MetricThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrors", GoMethod: "MetricUserErrors"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "protocolConfiguration", GoGetter: "ProtocolConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusReason", GoGetter: "StatusReason"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IGateway{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIGatewayRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IGatewayAuthorizerConfig",
		reflect.TypeOf((*IGatewayAuthorizerConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizerType", GoGetter: "AuthorizerType"},
		},
		func() interface{} {
			return &jsiiProxy_IGatewayAuthorizerConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IGatewayProtocolConfig",
		reflect.TypeOf((*IGatewayProtocolConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "protocolType", GoGetter: "ProtocolType"},
		},
		func() interface{} {
			return &jsiiProxy_IGatewayProtocolConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IGatewayTarget",
		reflect.TypeOf((*IGatewayTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderConfigurations", GoGetter: "CredentialProviderConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gateway", GoGetter: "Gateway"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetRef", GoGetter: "GatewayTargetRef"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusReasons", GoGetter: "StatusReasons"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetId", GoGetter: "TargetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetProtocolType", GoGetter: "TargetProtocolType"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IGatewayTarget{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIGatewayTargetRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IInterceptor",
		reflect.TypeOf((*IInterceptor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "interceptionPoint", GoGetter: "InterceptionPoint"},
		},
		func() interface{} {
			return &jsiiProxy_IInterceptor{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IMcpGatewayTarget",
		reflect.TypeOf((*IMcpGatewayTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderConfigurations", GoGetter: "CredentialProviderConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gateway", GoGetter: "Gateway"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetRef", GoGetter: "GatewayTargetRef"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusReasons", GoGetter: "StatusReasons"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetId", GoGetter: "TargetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetProtocolType", GoGetter: "TargetProtocolType"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IMcpGatewayTarget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGatewayTarget)
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
			_jsii_.MemberProperty{JsiiProperty: "memoryRef", GoGetter: "MemoryRef"},
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IMemory{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIMemoryRef)
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
		"@aws-cdk/aws-bedrock-agentcore-alpha.IOAuth2CredentialProvider",
		reflect.TypeOf((*IOAuth2CredentialProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindForGatewayOAuthTarget", GoMethod: "BindForGatewayOAuthTarget"},
			_jsii_.MemberProperty{JsiiProperty: "callbackUrl", GoGetter: "CallbackUrl"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretArn", GoGetter: "ClientSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "createdTime", GoGetter: "CreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderArn", GoGetter: "CredentialProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderVendor", GoGetter: "CredentialProviderVendor"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAdmin", GoMethod: "GrantAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedTime", GoGetter: "LastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oAuth2CredentialProviderRef", GoGetter: "OAuth2CredentialProviderRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IOAuth2CredentialProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIOAuth2CredentialProviderRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IOnlineEvaluationConfig",
		reflect.TypeOf((*IOnlineEvaluationConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "executionStatus", GoGetter: "ExecutionStatus"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigArn", GoGetter: "OnlineEvaluationConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigId", GoGetter: "OnlineEvaluationConfigId"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigName", GoGetter: "OnlineEvaluationConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigRef", GoGetter: "OnlineEvaluationConfigRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IOnlineEvaluationConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIOnlineEvaluationConfigRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IPolicy",
		reflect.TypeOf((*IPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEvaluationLatency", GoMethod: "MetricEvaluationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricEvaluations", GoMethod: "MetricEvaluations"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policyArn", GoGetter: "PolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngine", GoGetter: "PolicyEngine"},
			_jsii_.MemberProperty{JsiiProperty: "policyId", GoGetter: "PolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "policyName", GoGetter: "PolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "policyRef", GoGetter: "PolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "validationMode", GoGetter: "ValidationMode"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIPolicyRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IPolicyEngine",
		reflect.TypeOf((*IPolicyEngine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantEvaluate", GoMethod: "GrantEvaluate"},
			_jsii_.MemberMethod{JsiiMethod: "grantEvaluateForGateway", GoMethod: "GrantEvaluateForGateway"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAuthorizationLatency", GoMethod: "MetricAuthorizationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricDeniedRequests", GoMethod: "MetricDeniedRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrors", GoMethod: "MetricErrors"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineArn", GoGetter: "PolicyEngineArn"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineId", GoGetter: "PolicyEngineId"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineName", GoGetter: "PolicyEngineName"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineRef", GoGetter: "PolicyEngineRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IPolicyEngine{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIPolicyEngineRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
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
			_jsii_.MemberProperty{JsiiProperty: "runtimeEndpointRef", GoGetter: "RuntimeEndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "targetVersion", GoGetter: "TargetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IRuntimeEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIRuntimeEndpointRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ITargetConfiguration",
		reflect.TypeOf((*ITargetConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
		},
		func() interface{} {
			return &jsiiProxy_ITargetConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IWorkloadIdentity",
		reflect.TypeOf((*IWorkloadIdentity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdTime", GoGetter: "CreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAdmin", GoMethod: "GrantAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedTime", GoGetter: "LastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "workloadIdentityArn", GoGetter: "WorkloadIdentityArn"},
			_jsii_.MemberProperty{JsiiProperty: "workloadIdentityName", GoGetter: "WorkloadIdentityName"},
			_jsii_.MemberProperty{JsiiProperty: "workloadIdentityRef", GoGetter: "WorkloadIdentityRef"},
		},
		func() interface{} {
			j := jsiiProxy_IWorkloadIdentity{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIWorkloadIdentityRef)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IamAuthorizer",
		reflect.TypeOf((*IamAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizerType", GoGetter: "AuthorizerType"},
		},
		func() interface{} {
			j := jsiiProxy_IamAuthorizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGatewayAuthorizerConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IncludedOauth2TenantCredentialProviderProps",
		reflect.TypeOf((*IncludedOauth2TenantCredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IncludedOauth2TenantEndpoints",
		reflect.TypeOf((*IncludedOauth2TenantEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineApiSchema",
		reflect.TypeOf((*InlineApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucketOwnerAccountId", GoGetter: "BucketOwnerAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "grantPermissionsToRole", GoMethod: "GrantPermissionsToRole"},
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_InlineApiSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiSchema)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineToolSchema",
		reflect.TypeOf((*InlineToolSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucketOwnerAccountId", GoGetter: "BucketOwnerAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "grantPermissionsToRole", GoMethod: "GrantPermissionsToRole"},
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_InlineToolSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ToolSchema)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InterceptionPoint",
		reflect.TypeOf((*InterceptionPoint)(nil)).Elem(),
		map[string]interface{}{
			"REQUEST": InterceptionPoint_REQUEST,
			"RESPONSE": InterceptionPoint_RESPONSE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InterceptorBindConfig",
		reflect.TypeOf((*InterceptorBindConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InterceptorOptions",
		reflect.TypeOf((*InterceptorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InvocationConfiguration",
		reflect.TypeOf((*InvocationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LambdaInterceptor",
		reflect.TypeOf((*LambdaInterceptor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "interceptionPoint", GoGetter: "InterceptionPoint"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaInterceptor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInterceptor)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LambdaTargetConfiguration",
		reflect.TypeOf((*LambdaTargetConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberMethod{JsiiMethod: "renderMcpConfiguration", GoMethod: "RenderMcpConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
			_jsii_.MemberProperty{JsiiProperty: "toolSchema", GoGetter: "ToolSchema"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaTargetConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_McpTargetConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LifecycleConfiguration",
		reflect.TypeOf((*LifecycleConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LinkedinOAuth2CredentialProviderProps",
		reflect.TypeOf((*LinkedinOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LlmAsAJudgeOptions",
		reflect.TypeOf((*LlmAsAJudgeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LogType",
		reflect.TypeOf((*LogType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_LogType{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LoggingConfig",
		reflect.TypeOf((*LoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LoggingDestination",
		reflect.TypeOf((*LoggingDestination)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_LoggingDestination{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.MCPProtocolVersion",
		reflect.TypeOf((*MCPProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"MCP_2025_06_18": MCPProtocolVersion_MCP_2025_06_18,
			"MCP_2025_03_26": MCPProtocolVersion_MCP_2025_03_26,
		},
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
			_jsii_.MemberProperty{JsiiProperty: "reflectionConfiguration", GoGetter: "ReflectionConfiguration"},
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
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpConfiguration",
		reflect.TypeOf((*McpConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpGatewaySearchType",
		reflect.TypeOf((*McpGatewaySearchType)(nil)).Elem(),
		map[string]interface{}{
			"SEMANTIC": McpGatewaySearchType_SEMANTIC,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpProtocolConfiguration",
		reflect.TypeOf((*McpProtocolConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "instructions", GoGetter: "Instructions"},
			_jsii_.MemberProperty{JsiiProperty: "protocolType", GoGetter: "ProtocolType"},
			_jsii_.MemberProperty{JsiiProperty: "searchType", GoGetter: "SearchType"},
			_jsii_.MemberProperty{JsiiProperty: "supportedVersions", GoGetter: "SupportedVersions"},
		},
		func() interface{} {
			j := jsiiProxy_McpProtocolConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGatewayProtocolConfig)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpServerTargetConfiguration",
		reflect.TypeOf((*McpServerTargetConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberMethod{JsiiMethod: "renderMcpConfiguration", GoMethod: "RenderMcpConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
		},
		func() interface{} {
			j := jsiiProxy_McpServerTargetConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_McpTargetConfiguration)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpTargetConfiguration",
		reflect.TypeOf((*McpTargetConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderMcpConfiguration", GoMethod: "RenderMcpConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
		},
		func() interface{} {
			j := jsiiProxy_McpTargetConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITargetConfiguration)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpTargetType",
		reflect.TypeOf((*McpTargetType)(nil)).Elem(),
		map[string]interface{}{
			"OPENAPI_SCHEMA": McpTargetType_OPENAPI_SCHEMA,
			"SMITHY_MODEL": McpTargetType_SMITHY_MODEL,
			"LAMBDA": McpTargetType_LAMBDA,
			"MCP_SERVER": McpTargetType_MCP_SERVER,
			"API_GATEWAY": McpTargetType_API_GATEWAY,
		},
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
			_jsii_.MemberProperty{JsiiProperty: "memoryRef", GoGetter: "MemoryRef"},
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
			_jsii_.MemberProperty{JsiiProperty: "memoryRef", GoGetter: "MemoryRef"},
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
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
			"EPISODIC": MemoryStrategyType_EPISODIC,
			"CUSTOM": MemoryStrategyType_CUSTOM,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.MetadataConfiguration",
		reflect.TypeOf((*MetadataConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.MicrosoftOAuth2CredentialProviderProps",
		reflect.TypeOf((*MicrosoftOAuth2CredentialProviderProps)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.NoAuthAuthorizer",
		reflect.TypeOf((*NoAuthAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizerType", GoGetter: "AuthorizerType"},
		},
		func() interface{} {
			j := jsiiProxy_NoAuthAuthorizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGatewayAuthorizerConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.NotionOAuth2CredentialProviderProps",
		reflect.TypeOf((*NotionOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.NumericalRatingOption",
		reflect.TypeOf((*NumericalRatingOption)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2AuthorizationServerMetadata",
		reflect.TypeOf((*OAuth2AuthorizationServerMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2ClientCredentials",
		reflect.TypeOf((*OAuth2ClientCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		reflect.TypeOf((*OAuth2CredentialProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindForGatewayOAuthTarget", GoMethod: "BindForGatewayOAuthTarget"},
			_jsii_.MemberProperty{JsiiProperty: "callbackUrl", GoGetter: "CallbackUrl"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretArn", GoGetter: "ClientSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "createdTime", GoGetter: "CreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderArn", GoGetter: "CredentialProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderVendor", GoGetter: "CredentialProviderVendor"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAdmin", GoMethod: "GrantAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedTime", GoGetter: "LastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oAuth2CredentialProviderName", GoGetter: "OAuth2CredentialProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "oAuth2CredentialProviderRef", GoGetter: "OAuth2CredentialProviderRef"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_OAuth2CredentialProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOAuth2CredentialProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProviderAttributes",
		reflect.TypeOf((*OAuth2CredentialProviderAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProviderBaseProps",
		reflect.TypeOf((*OAuth2CredentialProviderBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProviderFactoryBaseProps",
		reflect.TypeOf((*OAuth2CredentialProviderFactoryBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProviderIdentityPerms",
		reflect.TypeOf((*OAuth2CredentialProviderIdentityPerms)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_OAuth2CredentialProviderIdentityPerms{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProviderProps",
		reflect.TypeOf((*OAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProviderVendor",
		reflect.TypeOf((*OAuth2CredentialProviderVendor)(nil)).Elem(),
		map[string]interface{}{
			"GOOGLE": OAuth2CredentialProviderVendor_GOOGLE,
			"GITHUB": OAuth2CredentialProviderVendor_GITHUB,
			"SLACK": OAuth2CredentialProviderVendor_SLACK,
			"SALESFORCE": OAuth2CredentialProviderVendor_SALESFORCE,
			"MICROSOFT": OAuth2CredentialProviderVendor_MICROSOFT,
			"CUSTOM": OAuth2CredentialProviderVendor_CUSTOM,
			"ATLASSIAN": OAuth2CredentialProviderVendor_ATLASSIAN,
			"LINKEDIN": OAuth2CredentialProviderVendor_LINKEDIN,
			"X": OAuth2CredentialProviderVendor_X,
			"OKTA": OAuth2CredentialProviderVendor_OKTA,
			"ONE_LOGIN": OAuth2CredentialProviderVendor_ONE_LOGIN,
			"PING_ONE": OAuth2CredentialProviderVendor_PING_ONE,
			"FACEBOOK": OAuth2CredentialProviderVendor_FACEBOOK,
			"YANDEX": OAuth2CredentialProviderVendor_YANDEX,
			"REDDIT": OAuth2CredentialProviderVendor_REDDIT,
			"ZOOM": OAuth2CredentialProviderVendor_ZOOM,
			"TWITCH": OAuth2CredentialProviderVendor_TWITCH,
			"SPOTIFY": OAuth2CredentialProviderVendor_SPOTIFY,
			"DROPBOX": OAuth2CredentialProviderVendor_DROPBOX,
			"NOTION": OAuth2CredentialProviderVendor_NOTION,
			"HUBSPOT": OAuth2CredentialProviderVendor_HUBSPOT,
			"CYBER_ARK": OAuth2CredentialProviderVendor_CYBER_ARK,
			"FUSION_AUTH": OAuth2CredentialProviderVendor_FUSION_AUTH,
			"AUTH0": OAuth2CredentialProviderVendor_AUTH0,
			"COGNITO": OAuth2CredentialProviderVendor_COGNITO,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuthConfiguration",
		reflect.TypeOf((*OAuthConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OnlineEvaluationBase",
		reflect.TypeOf((*OnlineEvaluationBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "executionStatus", GoGetter: "ExecutionStatus"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigArn", GoGetter: "OnlineEvaluationConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigId", GoGetter: "OnlineEvaluationConfigId"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigName", GoGetter: "OnlineEvaluationConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigRef", GoGetter: "OnlineEvaluationConfigRef"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_OnlineEvaluationBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOnlineEvaluationConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OnlineEvaluationBaseProps",
		reflect.TypeOf((*OnlineEvaluationBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OnlineEvaluationConfig",
		reflect.TypeOf((*OnlineEvaluationConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "executionStatus", GoGetter: "ExecutionStatus"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigArn", GoGetter: "OnlineEvaluationConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigId", GoGetter: "OnlineEvaluationConfigId"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigName", GoGetter: "OnlineEvaluationConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigRef", GoGetter: "OnlineEvaluationConfigRef"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_OnlineEvaluationConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_OnlineEvaluationBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OnlineEvaluationConfigAttributes",
		reflect.TypeOf((*OnlineEvaluationConfigAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OnlineEvaluationConfigProps",
		reflect.TypeOf((*OnlineEvaluationConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OpenApiTargetConfiguration",
		reflect.TypeOf((*OpenApiTargetConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiSchema", GoGetter: "ApiSchema"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderMcpConfiguration", GoMethod: "RenderMcpConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
		},
		func() interface{} {
			j := jsiiProxy_OpenApiTargetConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_McpTargetConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OverrideConfig",
		reflect.TypeOf((*OverrideConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.Policy",
		reflect.TypeOf((*Policy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEvaluationLatency", GoMethod: "MetricEvaluationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricEvaluations", GoMethod: "MetricEvaluations"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyArn", GoGetter: "PolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngine", GoGetter: "PolicyEngine"},
			_jsii_.MemberProperty{JsiiProperty: "policyId", GoGetter: "PolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "policyName", GoGetter: "PolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "policyRef", GoGetter: "PolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "validationMode", GoGetter: "ValidationMode"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Policy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PolicyBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyAttributes",
		reflect.TypeOf((*PolicyAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyBase",
		reflect.TypeOf((*PolicyBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEvaluationLatency", GoMethod: "MetricEvaluationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricEvaluations", GoMethod: "MetricEvaluations"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyArn", GoGetter: "PolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngine", GoGetter: "PolicyEngine"},
			_jsii_.MemberProperty{JsiiProperty: "policyId", GoGetter: "PolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "policyName", GoGetter: "PolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "policyRef", GoGetter: "PolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "validationMode", GoGetter: "ValidationMode"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPolicy)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyEngine",
		reflect.TypeOf((*PolicyEngine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPolicy", GoMethod: "AddPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantEvaluate", GoMethod: "GrantEvaluate"},
			_jsii_.MemberMethod{JsiiMethod: "grantEvaluateForGateway", GoMethod: "GrantEvaluateForGateway"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAuthorizationLatency", GoMethod: "MetricAuthorizationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricAuthorizations", GoMethod: "MetricAuthorizations"},
			_jsii_.MemberMethod{JsiiMethod: "metricDeniedRequests", GoMethod: "MetricDeniedRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrors", GoMethod: "MetricErrors"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policies", GoGetter: "Policies"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineArn", GoGetter: "PolicyEngineArn"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineId", GoGetter: "PolicyEngineId"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineName", GoGetter: "PolicyEngineName"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineRef", GoGetter: "PolicyEngineRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyEngine{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PolicyEngineBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyEngineAttributes",
		reflect.TypeOf((*PolicyEngineAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyEngineBase",
		reflect.TypeOf((*PolicyEngineBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantEvaluate", GoMethod: "GrantEvaluate"},
			_jsii_.MemberMethod{JsiiMethod: "grantEvaluateForGateway", GoMethod: "GrantEvaluateForGateway"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAuthorizationLatency", GoMethod: "MetricAuthorizationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricAuthorizations", GoMethod: "MetricAuthorizations"},
			_jsii_.MemberMethod{JsiiMethod: "metricDeniedRequests", GoMethod: "MetricDeniedRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrors", GoMethod: "MetricErrors"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineArn", GoGetter: "PolicyEngineArn"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineId", GoGetter: "PolicyEngineId"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineName", GoGetter: "PolicyEngineName"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineRef", GoGetter: "PolicyEngineRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyEngineBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPolicyEngine)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyEngineMode",
		reflect.TypeOf((*PolicyEngineMode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PolicyEngineMode{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyEngineProps",
		reflect.TypeOf((*PolicyEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyProps",
		reflect.TypeOf((*PolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyStatement",
		reflect.TypeOf((*PolicyStatement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "forAllPrincipals", GoMethod: "ForAllPrincipals"},
			_jsii_.MemberMethod{JsiiMethod: "forPrincipal", GoMethod: "ForPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "forPrincipalInGroup", GoMethod: "ForPrincipalInGroup"},
			_jsii_.MemberMethod{JsiiMethod: "onAction", GoMethod: "OnAction"},
			_jsii_.MemberMethod{JsiiMethod: "onActions", GoMethod: "OnActions"},
			_jsii_.MemberMethod{JsiiMethod: "onAllActions", GoMethod: "OnAllActions"},
			_jsii_.MemberMethod{JsiiMethod: "onAllResources", GoMethod: "OnAllResources"},
			_jsii_.MemberMethod{JsiiMethod: "onResource", GoMethod: "OnResource"},
			_jsii_.MemberMethod{JsiiMethod: "onResourceType", GoMethod: "OnResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toCedar", GoMethod: "ToCedar"},
			_jsii_.MemberMethod{JsiiMethod: "unless", GoMethod: "Unless"},
			_jsii_.MemberMethod{JsiiMethod: "when", GoMethod: "When"},
		},
		func() interface{} {
			return &jsiiProxy_PolicyStatement{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyValidationMode",
		reflect.TypeOf((*PolicyValidationMode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PolicyValidationMode{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ProtocolType",
		reflect.TypeOf((*ProtocolType)(nil)).Elem(),
		map[string]interface{}{
			"MCP": ProtocolType_MCP,
			"HTTP": ProtocolType_HTTP,
			"A2A": ProtocolType_A2A,
			"AGUI": ProtocolType_AGUI,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RecordingConfig",
		reflect.TypeOf((*RecordingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RedditOAuth2CredentialProviderProps",
		reflect.TypeOf((*RedditOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RequestHeaderConfiguration",
		reflect.TypeOf((*RequestHeaderConfiguration)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "runtimeRef", GoGetter: "RuntimeRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
			_jsii_.MemberProperty{JsiiProperty: "runtimeRef", GoGetter: "RuntimeRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_RuntimeBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBedrockAgentRuntime)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeCustomClaim",
		reflect.TypeOf((*RuntimeCustomClaim)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RuntimeCustomClaim{}
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
			_jsii_.MemberProperty{JsiiProperty: "runtimeEndpointRef", GoGetter: "RuntimeEndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "targetVersion", GoGetter: "TargetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
			_jsii_.MemberProperty{JsiiProperty: "runtimeEndpointRef", GoGetter: "RuntimeEndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "targetVersion", GoGetter: "TargetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
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
		"@aws-cdk/aws-bedrock-agentcore-alpha.S3ApiSchema",
		reflect.TypeOf((*S3ApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucketOwnerAccountId", GoGetter: "BucketOwnerAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "grantPermissionsToRole", GoMethod: "GrantPermissionsToRole"},
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_S3ApiSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiSchema)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.S3ToolSchema",
		reflect.TypeOf((*S3ToolSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucketOwnerAccountId", GoGetter: "BucketOwnerAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "grantPermissionsToRole", GoMethod: "GrantPermissionsToRole"},
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_S3ToolSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ToolSchema)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SalesforceOAuth2CredentialProviderProps",
		reflect.TypeOf((*SalesforceOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SchemaDefinition",
		reflect.TypeOf((*SchemaDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SchemaDefinitionType",
		reflect.TypeOf((*SchemaDefinitionType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": SchemaDefinitionType_STRING,
			"NUMBER": SchemaDefinitionType_NUMBER,
			"OBJECT": SchemaDefinitionType_OBJECT,
			"ARRAY": SchemaDefinitionType_ARRAY,
			"BOOLEAN": SchemaDefinitionType_BOOLEAN,
			"INTEGER": SchemaDefinitionType_INTEGER,
		},
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
		"@aws-cdk/aws-bedrock-agentcore-alpha.SlackOAuth2CredentialProviderProps",
		reflect.TypeOf((*SlackOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SmithyTargetConfiguration",
		reflect.TypeOf((*SmithyTargetConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderMcpConfiguration", GoMethod: "RenderMcpConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "smithyModel", GoGetter: "SmithyModel"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
		},
		func() interface{} {
			j := jsiiProxy_SmithyTargetConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_McpTargetConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SpotifyOAuth2CredentialProviderProps",
		reflect.TypeOf((*SpotifyOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.TargetConfigurationConfig",
		reflect.TypeOf((*TargetConfigurationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ToolDefinition",
		reflect.TypeOf((*ToolDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ToolSchema",
		reflect.TypeOf((*ToolSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucketOwnerAccountId", GoGetter: "BucketOwnerAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "grantPermissionsToRole", GoMethod: "GrantPermissionsToRole"},
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			return &jsiiProxy_ToolSchema{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.TriggerConditions",
		reflect.TypeOf((*TriggerConditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.TwitchOAuth2CredentialProviderProps",
		reflect.TypeOf((*TwitchOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.VpcConfigProps",
		reflect.TypeOf((*VpcConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.WorkloadIdentity",
		reflect.TypeOf((*WorkloadIdentity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdTime", GoGetter: "CreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAdmin", GoMethod: "GrantAdmin"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUse", GoMethod: "GrantUse"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedTime", GoGetter: "LastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "workloadIdentityArn", GoGetter: "WorkloadIdentityArn"},
			_jsii_.MemberProperty{JsiiProperty: "workloadIdentityName", GoGetter: "WorkloadIdentityName"},
			_jsii_.MemberProperty{JsiiProperty: "workloadIdentityRef", GoGetter: "WorkloadIdentityRef"},
		},
		func() interface{} {
			j := jsiiProxy_WorkloadIdentity{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWorkloadIdentity)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.WorkloadIdentityAttributes",
		reflect.TypeOf((*WorkloadIdentityAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-agentcore-alpha.WorkloadIdentityPerms",
		reflect.TypeOf((*WorkloadIdentityPerms)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_WorkloadIdentityPerms{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.WorkloadIdentityProps",
		reflect.TypeOf((*WorkloadIdentityProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.XOAuth2CredentialProviderProps",
		reflect.TypeOf((*XOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.YandexOAuth2CredentialProviderProps",
		reflect.TypeOf((*YandexOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ZoomOAuth2CredentialProviderProps",
		reflect.TypeOf((*ZoomOAuth2CredentialProviderProps)(nil)).Elem(),
	)
}
