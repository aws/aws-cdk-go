package awsbedrockagentcore

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.AddApiGatewayTargetOptions",
		reflect.TypeOf((*AddApiGatewayTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.AddEndpointOptions",
		reflect.TypeOf((*AddEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.AddLambdaTargetOptions",
		reflect.TypeOf((*AddLambdaTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.AddMcpServerTargetOptions",
		reflect.TypeOf((*AddMcpServerTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.AddOpenApiTargetOptions",
		reflect.TypeOf((*AddOpenApiTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.AddSmithyTargetOptions",
		reflect.TypeOf((*AddSmithyTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.AgentCoreRuntime",
		reflect.TypeOf((*AgentCoreRuntime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_AgentCoreRuntime{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.AgentRuntimeArtifact",
		reflect.TypeOf((*AgentRuntimeArtifact)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_AgentRuntimeArtifact{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.AgentRuntimeAttributes",
		reflect.TypeOf((*AgentRuntimeAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayHttpMethod",
		reflect.TypeOf((*ApiGatewayHttpMethod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ApiGatewayHttpMethod{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayTargetConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayTargetConfigurationProps",
		reflect.TypeOf((*ApiGatewayTargetConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayToolConfiguration",
		reflect.TypeOf((*ApiGatewayToolConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayToolFilter",
		reflect.TypeOf((*ApiGatewayToolFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayToolOverride",
		reflect.TypeOf((*ApiGatewayToolOverride)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyAdditionalConfiguration",
		reflect.TypeOf((*ApiKeyAdditionalConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialLocation",
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
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialProvider",
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
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialProviderAttributes",
		reflect.TypeOf((*ApiKeyCredentialProviderAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialProviderIdentityPerms",
		reflect.TypeOf((*ApiKeyCredentialProviderIdentityPerms)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ApiKeyCredentialProviderIdentityPerms{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialProviderOptions",
		reflect.TypeOf((*ApiKeyCredentialProviderOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialProviderProps",
		reflect.TypeOf((*ApiKeyCredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.ApiSchema",
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
		"aws-cdk-lib.aws_bedrockagentcore.AssetApiSchema",
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
		"aws-cdk-lib.aws_bedrockagentcore.AssetToolSchema",
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
		"aws-cdk-lib.aws_bedrockagentcore.AtlassianOAuth2CredentialProviderProps",
		reflect.TypeOf((*AtlassianOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.BrowserCustom",
		reflect.TypeOf((*BrowserCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "browserArn", GoGetter: "BrowserArn"},
			_jsii_.MemberProperty{JsiiProperty: "browserCustomName", GoGetter: "BrowserCustomName"},
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
		"aws-cdk-lib.aws_bedrockagentcore.BrowserCustomAttributes",
		reflect.TypeOf((*BrowserCustomAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.BrowserCustomBase",
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
		"aws-cdk-lib.aws_bedrockagentcore.BrowserCustomProps",
		reflect.TypeOf((*BrowserCustomProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.BrowserNetworkConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.BrowserSigning",
		reflect.TypeOf((*BrowserSigning)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": BrowserSigning_ENABLED,
			"DISABLED": BrowserSigning_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.BuiltinEvaluator",
		reflect.TypeOf((*BuiltinEvaluator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_BuiltinEvaluator{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CategoricalRatingOption",
		reflect.TypeOf((*CategoricalRatingOption)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnApiKeyCredentialProvider",
		reflect.TypeOf((*CfnApiKeyCredentialProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiKey", GoGetter: "ApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyCredentialProviderRef", GoGetter: "ApiKeyCredentialProviderRef"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiKeySecretArn", GoGetter: "AttrApiKeySecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrCredentialProviderArn", GoGetter: "AttrCredentialProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApiKeyCredentialProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIApiKeyCredentialProviderRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnApiKeyCredentialProvider.ApiKeySecretArnProperty",
		reflect.TypeOf((*CfnApiKeyCredentialProvider_ApiKeySecretArnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnApiKeyCredentialProviderProps",
		reflect.TypeOf((*CfnApiKeyCredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserCustom",
		reflect.TypeOf((*CfnBrowserCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrBrowserArn", GoGetter: "AttrBrowserArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrBrowserId", GoGetter: "AttrBrowserId"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFailureReason", GoGetter: "AttrFailureReason"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedAt", GoGetter: "AttrLastUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "browserCustomRef", GoGetter: "BrowserCustomRef"},
			_jsii_.MemberProperty{JsiiProperty: "browserSigning", GoGetter: "BrowserSigning"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "certificates", GoGetter: "Certificates"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "enterprisePolicies", GoGetter: "EnterprisePolicies"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "recordingConfig", GoGetter: "RecordingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrowserCustom{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIBrowserCustomRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserCustom.BrowserEnterprisePolicyProperty",
		reflect.TypeOf((*CfnBrowserCustom_BrowserEnterprisePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserCustom.BrowserNetworkConfigurationProperty",
		reflect.TypeOf((*CfnBrowserCustom_BrowserNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserCustom.BrowserSigningProperty",
		reflect.TypeOf((*CfnBrowserCustom_BrowserSigningProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserCustom.CertificateLocationProperty",
		reflect.TypeOf((*CfnBrowserCustom_CertificateLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserCustom.CertificateProperty",
		reflect.TypeOf((*CfnBrowserCustom_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserCustom.RecordingConfigProperty",
		reflect.TypeOf((*CfnBrowserCustom_RecordingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserCustom.S3LocationProperty",
		reflect.TypeOf((*CfnBrowserCustom_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserCustom.VpcConfigProperty",
		reflect.TypeOf((*CfnBrowserCustom_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserCustomProps",
		reflect.TypeOf((*CfnBrowserCustomProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserProfile",
		reflect.TypeOf((*CfnBrowserProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastSavedAt", GoGetter: "AttrLastSavedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastSavedBrowserId", GoGetter: "AttrLastSavedBrowserId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastSavedBrowserSessionId", GoGetter: "AttrLastSavedBrowserSessionId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedAt", GoGetter: "AttrLastUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrProfileArn", GoGetter: "AttrProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrProfileId", GoGetter: "AttrProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "browserProfileRef", GoGetter: "BrowserProfileRef"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrowserProfile{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIBrowserProfileRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnBrowserProfileProps",
		reflect.TypeOf((*CfnBrowserProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnCodeInterpreterCustom",
		reflect.TypeOf((*CfnCodeInterpreterCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCodeInterpreterArn", GoGetter: "AttrCodeInterpreterArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCodeInterpreterId", GoGetter: "AttrCodeInterpreterId"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFailureReason", GoGetter: "AttrFailureReason"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedAt", GoGetter: "AttrLastUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "certificates", GoGetter: "Certificates"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterCustomRef", GoGetter: "CodeInterpreterCustomRef"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeInterpreterCustom{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreICodeInterpreterCustomRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnCodeInterpreterCustom.CertificateLocationProperty",
		reflect.TypeOf((*CfnCodeInterpreterCustom_CertificateLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnCodeInterpreterCustom.CertificateProperty",
		reflect.TypeOf((*CfnCodeInterpreterCustom_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnCodeInterpreterCustom.CodeInterpreterNetworkConfigurationProperty",
		reflect.TypeOf((*CfnCodeInterpreterCustom_CodeInterpreterNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnCodeInterpreterCustom.VpcConfigProperty",
		reflect.TypeOf((*CfnCodeInterpreterCustom_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnCodeInterpreterCustomProps",
		reflect.TypeOf((*CfnCodeInterpreterCustomProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator",
		reflect.TypeOf((*CfnEvaluator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrEvaluatorArn", GoGetter: "AttrEvaluatorArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrEvaluatorId", GoGetter: "AttrEvaluatorId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorConfig", GoGetter: "EvaluatorConfig"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorName", GoGetter: "EvaluatorName"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorRef", GoGetter: "EvaluatorRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyArn", GoGetter: "KmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEvaluator{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIEvaluatorRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator.BedrockEvaluatorModelConfigProperty",
		reflect.TypeOf((*CfnEvaluator_BedrockEvaluatorModelConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator.CategoricalScaleDefinitionProperty",
		reflect.TypeOf((*CfnEvaluator_CategoricalScaleDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator.CodeBasedEvaluatorConfigProperty",
		reflect.TypeOf((*CfnEvaluator_CodeBasedEvaluatorConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator.EvaluatorConfigProperty",
		reflect.TypeOf((*CfnEvaluator_EvaluatorConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator.EvaluatorModelConfigProperty",
		reflect.TypeOf((*CfnEvaluator_EvaluatorModelConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator.InferenceConfigurationProperty",
		reflect.TypeOf((*CfnEvaluator_InferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator.LambdaEvaluatorConfigProperty",
		reflect.TypeOf((*CfnEvaluator_LambdaEvaluatorConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator.LlmAsAJudgeEvaluatorConfigProperty",
		reflect.TypeOf((*CfnEvaluator_LlmAsAJudgeEvaluatorConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator.NumericalScaleDefinitionProperty",
		reflect.TypeOf((*CfnEvaluator_NumericalScaleDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluator.RatingScaleProperty",
		reflect.TypeOf((*CfnEvaluator_RatingScaleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnEvaluatorProps",
		reflect.TypeOf((*CfnEvaluatorProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway",
		reflect.TypeOf((*CfnGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrGatewayArn", GoGetter: "AttrGatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrGatewayIdentifier", GoGetter: "AttrGatewayIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "attrGatewayUrl", GoGetter: "AttrGatewayUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusReasons", GoGetter: "AttrStatusReasons"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkloadIdentityDetails", GoGetter: "AttrWorkloadIdentityDetails"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerConfiguration", GoGetter: "AuthorizerConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerType", GoGetter: "AuthorizerType"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "exceptionLevel", GoGetter: "ExceptionLevel"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRef", GoGetter: "GatewayRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "interceptorConfigurations", GoGetter: "InterceptorConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyArn", GoGetter: "KmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineConfiguration", GoGetter: "PolicyEngineConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "protocolConfiguration", GoGetter: "ProtocolConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "protocolType", GoGetter: "ProtocolType"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIGatewayRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.AuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnGateway_AuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.AuthorizingClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnGateway_AuthorizingClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.ClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnGateway_ClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.CustomClaimValidationTypeProperty",
		reflect.TypeOf((*CfnGateway_CustomClaimValidationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.CustomJWTAuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnGateway_CustomJWTAuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.GatewayInterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGateway_GatewayInterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.GatewayPolicyEngineConfigurationProperty",
		reflect.TypeOf((*CfnGateway_GatewayPolicyEngineConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.GatewayProtocolConfigurationProperty",
		reflect.TypeOf((*CfnGateway_GatewayProtocolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.InterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGateway_InterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.InterceptorInputConfigurationProperty",
		reflect.TypeOf((*CfnGateway_InterceptorInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.LambdaInterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGateway_LambdaInterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.MCPGatewayConfigurationProperty",
		reflect.TypeOf((*CfnGateway_MCPGatewayConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGateway.WorkloadIdentityDetailsProperty",
		reflect.TypeOf((*CfnGateway_WorkloadIdentityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayProps",
		reflect.TypeOf((*CfnGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget",
		reflect.TypeOf((*CfnGatewayTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrGatewayArn", GoGetter: "AttrGatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastSynchronizedAt", GoGetter: "AttrLastSynchronizedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusReasons", GoGetter: "AttrStatusReasons"},
			_jsii_.MemberProperty{JsiiProperty: "attrTargetId", GoGetter: "AttrTargetId"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderConfigurations", GoGetter: "CredentialProviderConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayIdentifier", GoGetter: "GatewayIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetRef", GoGetter: "GatewayTargetRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "metadataConfiguration", GoGetter: "MetadataConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "targetConfiguration", GoGetter: "TargetConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIGatewayTargetRef)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.ApiGatewayTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTarget_ApiGatewayTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.ApiGatewayToolConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTarget_ApiGatewayToolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.ApiGatewayToolFilterProperty",
		reflect.TypeOf((*CfnGatewayTarget_ApiGatewayToolFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.ApiGatewayToolOverrideProperty",
		reflect.TypeOf((*CfnGatewayTarget_ApiGatewayToolOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.ApiKeyCredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTarget_ApiKeyCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.ApiSchemaConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTarget_ApiSchemaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.CredentialProviderConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTarget_CredentialProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.CredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTarget_CredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.IamCredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTarget_IamCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.McpLambdaTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTarget_McpLambdaTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.McpServerTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTarget_McpServerTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.McpTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTarget_McpTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.MetadataConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTarget_MetadataConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.OAuthCredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTarget_OAuthCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.S3ConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTarget_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.SchemaDefinitionProperty",
		reflect.TypeOf((*CfnGatewayTarget_SchemaDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.TargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTarget_TargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.ToolDefinitionProperty",
		reflect.TypeOf((*CfnGatewayTarget_ToolDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTarget.ToolSchemaProperty",
		reflect.TypeOf((*CfnGatewayTarget_ToolSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnGatewayTargetProps",
		reflect.TypeOf((*CfnGatewayTargetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory",
		reflect.TypeOf((*CfnMemory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFailureReason", GoGetter: "AttrFailureReason"},
			_jsii_.MemberProperty{JsiiProperty: "attrMemoryArn", GoGetter: "AttrMemoryArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrMemoryId", GoGetter: "AttrMemoryId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKeyArn", GoGetter: "EncryptionKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "eventExpiryDuration", GoGetter: "EventExpiryDuration"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "indexedKeys", GoGetter: "IndexedKeys"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "memoryExecutionRoleArn", GoGetter: "MemoryExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "memoryRef", GoGetter: "MemoryRef"},
			_jsii_.MemberProperty{JsiiProperty: "memoryStrategies", GoGetter: "MemoryStrategies"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "streamDeliveryResources", GoGetter: "StreamDeliveryResources"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMemory{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIMemoryRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.ContentConfigurationProperty",
		reflect.TypeOf((*CfnMemory_ContentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.CustomConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_CustomConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.CustomMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemory_CustomMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.EpisodicMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemory_EpisodicMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.EpisodicOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_EpisodicOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.EpisodicOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_EpisodicOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.EpisodicOverrideProperty",
		reflect.TypeOf((*CfnMemory_EpisodicOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.EpisodicOverrideReflectionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_EpisodicOverrideReflectionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.EpisodicReflectionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_EpisodicReflectionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.ExtractionConfigProperty",
		reflect.TypeOf((*CfnMemory_ExtractionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.IndexedKeyProperty",
		reflect.TypeOf((*CfnMemory_IndexedKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.InvocationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_InvocationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.KinesisResourceProperty",
		reflect.TypeOf((*CfnMemory_KinesisResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.LlmExtractionConfigProperty",
		reflect.TypeOf((*CfnMemory_LlmExtractionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.MemoryRecordSchemaProperty",
		reflect.TypeOf((*CfnMemory_MemoryRecordSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.MemoryStrategyProperty",
		reflect.TypeOf((*CfnMemory_MemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.MessageBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemory_MessageBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.MetadataSchemaEntryProperty",
		reflect.TypeOf((*CfnMemory_MetadataSchemaEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.NumberValidationProperty",
		reflect.TypeOf((*CfnMemory_NumberValidationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.SelfManagedConfigurationProperty",
		reflect.TypeOf((*CfnMemory_SelfManagedConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.SemanticMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemory_SemanticMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.SemanticOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_SemanticOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.SemanticOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_SemanticOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.SemanticOverrideProperty",
		reflect.TypeOf((*CfnMemory_SemanticOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.StreamDeliveryResourceProperty",
		reflect.TypeOf((*CfnMemory_StreamDeliveryResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.StreamDeliveryResourcesProperty",
		reflect.TypeOf((*CfnMemory_StreamDeliveryResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.StringListValidationProperty",
		reflect.TypeOf((*CfnMemory_StringListValidationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.StringValidationProperty",
		reflect.TypeOf((*CfnMemory_StringValidationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.SummaryMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemory_SummaryMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.SummaryOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_SummaryOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.SummaryOverrideProperty",
		reflect.TypeOf((*CfnMemory_SummaryOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.TimeBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemory_TimeBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.TokenBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemory_TokenBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.TriggerConditionInputProperty",
		reflect.TypeOf((*CfnMemory_TriggerConditionInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.UserPreferenceMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemory_UserPreferenceMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.UserPreferenceOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_UserPreferenceOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.UserPreferenceOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemory_UserPreferenceOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.UserPreferenceOverrideProperty",
		reflect.TypeOf((*CfnMemory_UserPreferenceOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemory.ValidationProperty",
		reflect.TypeOf((*CfnMemory_ValidationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnMemoryProps",
		reflect.TypeOf((*CfnMemoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider",
		reflect.TypeOf((*CfnOAuth2CredentialProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCallbackUrl", GoGetter: "AttrCallbackUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrClientSecretArn", GoGetter: "AttrClientSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrCredentialProviderArn", GoGetter: "AttrCredentialProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrOauth2ProviderConfigOutput", GoGetter: "AttrOauth2ProviderConfigOutput"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderVendor", GoGetter: "CredentialProviderVendor"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oAuth2CredentialProviderRef", GoGetter: "OAuth2CredentialProviderRef"},
			_jsii_.MemberProperty{JsiiProperty: "oauth2ProviderConfigInput", GoGetter: "Oauth2ProviderConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOAuth2CredentialProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIOAuth2CredentialProviderRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.AtlassianOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_AtlassianOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.ClientSecretArnProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_ClientSecretArnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.CustomOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_CustomOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.GithubOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_GithubOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.GoogleOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_GoogleOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.IncludedOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_IncludedOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.LinkedinOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_LinkedinOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.MicrosoftOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_MicrosoftOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.Oauth2AuthorizationServerMetadataProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_Oauth2AuthorizationServerMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.Oauth2DiscoveryProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_Oauth2DiscoveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.Oauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_Oauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.Oauth2ProviderConfigOutputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_Oauth2ProviderConfigOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.OnBehalfOfTokenExchangeConfigProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_OnBehalfOfTokenExchangeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.SalesforceOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_SalesforceOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.SlackOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_SlackOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProvider.TokenExchangeGrantTypeConfigProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProvider_TokenExchangeGrantTypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOAuth2CredentialProviderProps",
		reflect.TypeOf((*CfnOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig",
		reflect.TypeOf((*CfnOnlineEvaluationConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrOnlineEvaluationConfigArn", GoGetter: "AttrOnlineEvaluationConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrOnlineEvaluationConfigId", GoGetter: "AttrOnlineEvaluationConfigId"},
			_jsii_.MemberProperty{JsiiProperty: "attrOutputConfig", GoGetter: "AttrOutputConfig"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceConfig", GoGetter: "DataSourceConfig"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationExecutionRoleArn", GoGetter: "EvaluationExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "evaluators", GoGetter: "Evaluators"},
			_jsii_.MemberProperty{JsiiProperty: "executionStatus", GoGetter: "ExecutionStatus"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigName", GoGetter: "OnlineEvaluationConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "onlineEvaluationConfigRef", GoGetter: "OnlineEvaluationConfigRef"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "rule", GoGetter: "Rule"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOnlineEvaluationConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIOnlineEvaluationConfigRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig.CloudWatchLogsInputConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfig_CloudWatchLogsInputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig.CloudWatchOutputConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfig_CloudWatchOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig.DataSourceConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfig_DataSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig.EvaluatorReferenceProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfig_EvaluatorReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig.FilterProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfig_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig.FilterValueProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfig_FilterValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig.OutputConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfig_OutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig.RuleProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfig_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig.SamplingConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfig_SamplingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfig.SessionConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfig_SessionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnOnlineEvaluationConfigProps",
		reflect.TypeOf((*CfnOnlineEvaluationConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnPolicy",
		reflect.TypeOf((*CfnPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrPolicyArn", GoGetter: "AttrPolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrPolicyId", GoGetter: "AttrPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusReasons", GoGetter: "AttrStatusReasons"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineId", GoGetter: "PolicyEngineId"},
			_jsii_.MemberProperty{JsiiProperty: "policyRef", GoGetter: "PolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "validationMode", GoGetter: "ValidationMode"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIPolicyRef)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnPolicy.CedarPolicyProperty",
		reflect.TypeOf((*CfnPolicy_CedarPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnPolicy.PolicyDefinitionProperty",
		reflect.TypeOf((*CfnPolicy_PolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnPolicyEngine",
		reflect.TypeOf((*CfnPolicyEngine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrPolicyEngineArn", GoGetter: "AttrPolicyEngineArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrPolicyEngineId", GoGetter: "AttrPolicyEngineId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusReasons", GoGetter: "AttrStatusReasons"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKeyArn", GoGetter: "EncryptionKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyEngineRef", GoGetter: "PolicyEngineRef"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyEngine{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIPolicyEngineRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnPolicyEngineProps",
		reflect.TypeOf((*CfnPolicyEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnPolicyProps",
		reflect.TypeOf((*CfnPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime",
		reflect.TypeOf((*CfnRuntime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeArtifact", GoGetter: "AgentRuntimeArtifact"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeName", GoGetter: "AgentRuntimeName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentRuntimeArn", GoGetter: "AttrAgentRuntimeArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentRuntimeId", GoGetter: "AttrAgentRuntimeId"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentRuntimeVersion", GoGetter: "AttrAgentRuntimeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFailureReason", GoGetter: "AttrFailureReason"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedAt", GoGetter: "AttrLastUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkloadIdentityDetails", GoGetter: "AttrWorkloadIdentityDetails"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerConfiguration", GoGetter: "AuthorizerConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environmentVariables", GoGetter: "EnvironmentVariables"},
			_jsii_.MemberProperty{JsiiProperty: "filesystemConfigurations", GoGetter: "FilesystemConfigurations"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfiguration", GoGetter: "LifecycleConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "protocolConfiguration", GoGetter: "ProtocolConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeaderConfiguration", GoGetter: "RequestHeaderConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeRef", GoGetter: "RuntimeRef"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuntime{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIRuntimeRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.AgentRuntimeArtifactProperty",
		reflect.TypeOf((*CfnRuntime_AgentRuntimeArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.AuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnRuntime_AuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.AuthorizingClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnRuntime_AuthorizingClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.ClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnRuntime_ClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.CodeConfigurationProperty",
		reflect.TypeOf((*CfnRuntime_CodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.CodeProperty",
		reflect.TypeOf((*CfnRuntime_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.ContainerConfigurationProperty",
		reflect.TypeOf((*CfnRuntime_ContainerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.CustomClaimValidationTypeProperty",
		reflect.TypeOf((*CfnRuntime_CustomClaimValidationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.CustomJWTAuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnRuntime_CustomJWTAuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.FilesystemConfigurationProperty",
		reflect.TypeOf((*CfnRuntime_FilesystemConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.LifecycleConfigurationProperty",
		reflect.TypeOf((*CfnRuntime_LifecycleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnRuntime_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.RequestHeaderConfigurationProperty",
		reflect.TypeOf((*CfnRuntime_RequestHeaderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.S3LocationProperty",
		reflect.TypeOf((*CfnRuntime_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.SessionStorageConfigurationProperty",
		reflect.TypeOf((*CfnRuntime_SessionStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.VpcConfigProperty",
		reflect.TypeOf((*CfnRuntime_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntime.WorkloadIdentityDetailsProperty",
		reflect.TypeOf((*CfnRuntime_WorkloadIdentityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntimeEndpoint",
		reflect.TypeOf((*CfnRuntimeEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeId", GoGetter: "AgentRuntimeId"},
			_jsii_.MemberProperty{JsiiProperty: "agentRuntimeVersion", GoGetter: "AgentRuntimeVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentRuntimeArn", GoGetter: "AttrAgentRuntimeArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentRuntimeEndpointArn", GoGetter: "AttrAgentRuntimeEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFailureReason", GoGetter: "AttrFailureReason"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedAt", GoGetter: "AttrLastUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrLiveVersion", GoGetter: "AttrLiveVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrTargetVersion", GoGetter: "AttrTargetVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeEndpointRef", GoGetter: "RuntimeEndpointRef"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuntimeEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIRuntimeEndpointRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntimeEndpointProps",
		reflect.TypeOf((*CfnRuntimeEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnRuntimeProps",
		reflect.TypeOf((*CfnRuntimeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CfnWorkloadIdentity",
		reflect.TypeOf((*CfnWorkloadIdentity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedResourceOauth2ReturnUrls", GoGetter: "AllowedResourceOauth2ReturnUrls"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkloadIdentityArn", GoGetter: "AttrWorkloadIdentityArn"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberMethod{JsiiMethod: "cfnPropertyName", GoMethod: "CfnPropertyName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnPropertyNames", GoGetter: "CfnPropertyNames"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "workloadIdentityRef", GoGetter: "WorkloadIdentityRef"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkloadIdentity{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsbedrockagentcoreIWorkloadIdentityRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CfnWorkloadIdentityProps",
		reflect.TypeOf((*CfnWorkloadIdentityProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CloudWatchLogsDataSourceConfig",
		reflect.TypeOf((*CloudWatchLogsDataSourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CodeAssetOptions",
		reflect.TypeOf((*CodeAssetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CodeBasedOptions",
		reflect.TypeOf((*CodeBasedOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CodeInterpreterCustom",
		reflect.TypeOf((*CodeInterpreterCustom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterArn", GoGetter: "CodeInterpreterArn"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterCustomName", GoGetter: "CodeInterpreterCustomName"},
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
		"aws-cdk-lib.aws_bedrockagentcore.CodeInterpreterCustomAttributes",
		reflect.TypeOf((*CodeInterpreterCustomAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CodeInterpreterCustomBase",
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
		"aws-cdk-lib.aws_bedrockagentcore.CodeInterpreterCustomProps",
		reflect.TypeOf((*CodeInterpreterCustomProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CodeInterpreterNetworkConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.CognitoAuthorizerProps",
		reflect.TypeOf((*CognitoAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_bedrockagentcore.CredentialProviderType",
		reflect.TypeOf((*CredentialProviderType)(nil)).Elem(),
		map[string]interface{}{
			"API_KEY": CredentialProviderType_API_KEY,
			"OAUTH": CredentialProviderType_OAUTH,
			"GATEWAY_IAM_ROLE": CredentialProviderType_GATEWAY_IAM_ROLE,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_bedrockagentcore.CustomClaimOperator",
		reflect.TypeOf((*CustomClaimOperator)(nil)).Elem(),
		map[string]interface{}{
			"EQUALS": CustomClaimOperator_EQUALS,
			"CONTAINS": CustomClaimOperator_CONTAINS,
			"CONTAINS_ANY": CustomClaimOperator_CONTAINS_ANY,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.CustomJwtAuthorizer",
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
		"aws-cdk-lib.aws_bedrockagentcore.CustomJwtConfiguration",
		reflect.TypeOf((*CustomJwtConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.CustomOAuth2CredentialProviderProps",
		reflect.TypeOf((*CustomOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.DataSourceConfig",
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
		"aws-cdk-lib.aws_bedrockagentcore.DataSourceConfigBindResult",
		reflect.TypeOf((*DataSourceConfigBindResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.DropboxOAuth2CredentialProviderProps",
		reflect.TypeOf((*DropboxOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.EpisodicReflectionConfiguration",
		reflect.TypeOf((*EpisodicReflectionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluationLevel",
		reflect.TypeOf((*EvaluationLevel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_EvaluationLevel{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.Evaluator",
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
		"aws-cdk-lib.aws_bedrockagentcore.EvaluatorAttributes",
		reflect.TypeOf((*EvaluatorAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluatorBase",
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
		"aws-cdk-lib.aws_bedrockagentcore.EvaluatorConfig",
		reflect.TypeOf((*EvaluatorConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
		},
		func() interface{} {
			return &jsiiProxy_EvaluatorConfig{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluatorInferenceConfig",
		reflect.TypeOf((*EvaluatorInferenceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluatorProps",
		reflect.TypeOf((*EvaluatorProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluatorRatingScale",
		reflect.TypeOf((*EvaluatorRatingScale)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EvaluatorRatingScale{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluatorSelector",
		reflect.TypeOf((*EvaluatorSelector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "evaluatorId", GoGetter: "EvaluatorId"},
		},
		func() interface{} {
			return &jsiiProxy_EvaluatorSelector{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluatorSelectorBindResult",
		reflect.TypeOf((*EvaluatorSelectorBindResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.ExecutionStatus",
		reflect.TypeOf((*ExecutionStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ExecutionStatus{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.FacebookOAuth2CredentialProviderProps",
		reflect.TypeOf((*FacebookOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.FilterConfig",
		reflect.TypeOf((*FilterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.FilterOperator",
		reflect.TypeOf((*FilterOperator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_FilterOperator{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.FilterValue",
		reflect.TypeOf((*FilterValue)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FilterValue{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.FromApiKeyIdentityOptions",
		reflect.TypeOf((*FromApiKeyIdentityOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.FromOauthIdentityOptions",
		reflect.TypeOf((*FromOauthIdentityOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.Gateway",
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
			_jsii_.MemberProperty{JsiiProperty: "gatewayName", GoGetter: "GatewayName"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oauthScopes", GoGetter: "OauthScopes"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
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
		"aws-cdk-lib.aws_bedrockagentcore.GatewayApiKeyIdentityBinding",
		reflect.TypeOf((*GatewayApiKeyIdentityBinding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayAttributes",
		reflect.TypeOf((*GatewayAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayAuthorizer",
		reflect.TypeOf((*GatewayAuthorizer)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GatewayAuthorizer{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayAuthorizerType",
		reflect.TypeOf((*GatewayAuthorizerType)(nil)).Elem(),
		map[string]interface{}{
			"CUSTOM_JWT": GatewayAuthorizerType_CUSTOM_JWT,
			"AWS_IAM": GatewayAuthorizerType_AWS_IAM,
			"NONE": GatewayAuthorizerType_NONE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayBase",
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
			_jsii_.MemberProperty{JsiiProperty: "gatewayName", GoGetter: "GatewayName"},
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
		"aws-cdk-lib.aws_bedrockagentcore.GatewayCredentialProvider",
		reflect.TypeOf((*GatewayCredentialProvider)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GatewayCredentialProvider{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayCustomClaim",
		reflect.TypeOf((*GatewayCustomClaim)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GatewayCustomClaim{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayExceptionLevel",
		reflect.TypeOf((*GatewayExceptionLevel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GatewayExceptionLevel{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayOAuth2IdentityBinding",
		reflect.TypeOf((*GatewayOAuth2IdentityBinding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayProps",
		reflect.TypeOf((*GatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayProtocol",
		reflect.TypeOf((*GatewayProtocol)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GatewayProtocol{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTarget",
		reflect.TypeOf((*GatewayTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderConfigurations", GoGetter: "CredentialProviderConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gateway", GoGetter: "Gateway"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetName", GoGetter: "GatewayTargetName"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetRef", GoGetter: "GatewayTargetRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantSync", GoMethod: "GrantSync"},
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
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTargetApiGatewayProps",
		reflect.TypeOf((*GatewayTargetApiGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTargetAttributes",
		reflect.TypeOf((*GatewayTargetAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTargetBase",
		reflect.TypeOf((*GatewayTargetBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderConfigurations", GoGetter: "CredentialProviderConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gateway", GoGetter: "Gateway"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetName", GoGetter: "GatewayTargetName"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetRef", GoGetter: "GatewayTargetRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
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
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTargetCommonProps",
		reflect.TypeOf((*GatewayTargetCommonProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTargetLambdaProps",
		reflect.TypeOf((*GatewayTargetLambdaProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTargetMcpServerProps",
		reflect.TypeOf((*GatewayTargetMcpServerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTargetOpenApiProps",
		reflect.TypeOf((*GatewayTargetOpenApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTargetProps",
		reflect.TypeOf((*GatewayTargetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTargetProtocolType",
		reflect.TypeOf((*GatewayTargetProtocolType)(nil)).Elem(),
		map[string]interface{}{
			"MCP": GatewayTargetProtocolType_MCP,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayTargetSmithyProps",
		reflect.TypeOf((*GatewayTargetSmithyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GithubOAuth2CredentialProviderProps",
		reflect.TypeOf((*GithubOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.GoogleOAuth2CredentialProviderProps",
		reflect.TypeOf((*GoogleOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.HubspotOAuth2CredentialProviderProps",
		reflect.TypeOf((*HubspotOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_bedrockagentcore.IApiKeyCredentialProvider",
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
		"aws-cdk-lib.aws_bedrockagentcore.IBedrockAgentRuntime",
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
		"aws-cdk-lib.aws_bedrockagentcore.IBrowserCustom",
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
		"aws-cdk-lib.aws_bedrockagentcore.ICodeInterpreterCustom",
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
		"aws-cdk-lib.aws_bedrockagentcore.ICredentialProviderConfig",
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
		"aws-cdk-lib.aws_bedrockagentcore.IEvaluator",
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
		"aws-cdk-lib.aws_bedrockagentcore.IGateway",
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
			_jsii_.MemberProperty{JsiiProperty: "gatewayName", GoGetter: "GatewayName"},
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
		"aws-cdk-lib.aws_bedrockagentcore.IGatewayAuthorizerConfig",
		reflect.TypeOf((*IGatewayAuthorizerConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizerType", GoGetter: "AuthorizerType"},
		},
		func() interface{} {
			return &jsiiProxy_IGatewayAuthorizerConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_bedrockagentcore.IGatewayProtocolConfig",
		reflect.TypeOf((*IGatewayProtocolConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "protocolType", GoGetter: "ProtocolType"},
		},
		func() interface{} {
			return &jsiiProxy_IGatewayProtocolConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_bedrockagentcore.IGatewayTarget",
		reflect.TypeOf((*IGatewayTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderConfigurations", GoGetter: "CredentialProviderConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gateway", GoGetter: "Gateway"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetName", GoGetter: "GatewayTargetName"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetRef", GoGetter: "GatewayTargetRef"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
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
		"aws-cdk-lib.aws_bedrockagentcore.IInterceptor",
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
		"aws-cdk-lib.aws_bedrockagentcore.IMcpGatewayTarget",
		reflect.TypeOf((*IMcpGatewayTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderConfigurations", GoGetter: "CredentialProviderConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gateway", GoGetter: "Gateway"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetName", GoGetter: "GatewayTargetName"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayTargetRef", GoGetter: "GatewayTargetRef"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantManage", GoMethod: "GrantManage"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
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
		"aws-cdk-lib.aws_bedrockagentcore.IMemory",
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
		"aws-cdk-lib.aws_bedrockagentcore.IMemoryStrategy",
		reflect.TypeOf((*IMemoryStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
			_jsii_.MemberProperty{JsiiProperty: "strategyName", GoGetter: "StrategyName"},
			_jsii_.MemberProperty{JsiiProperty: "strategyType", GoGetter: "StrategyType"},
		},
		func() interface{} {
			return &jsiiProxy_IMemoryStrategy{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_bedrockagentcore.IOAuth2CredentialProvider",
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
		"aws-cdk-lib.aws_bedrockagentcore.IOnlineEvaluationConfig",
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
		"aws-cdk-lib.aws_bedrockagentcore.IRuntimeEndpoint",
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
		"aws-cdk-lib.aws_bedrockagentcore.ITargetConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.IWorkloadIdentity",
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
		"aws-cdk-lib.aws_bedrockagentcore.IamAuthorizer",
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
		"aws-cdk-lib.aws_bedrockagentcore.IncludedOauth2TenantCredentialProviderProps",
		reflect.TypeOf((*IncludedOauth2TenantCredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.IncludedOauth2TenantEndpoints",
		reflect.TypeOf((*IncludedOauth2TenantEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.InlineApiSchema",
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
		"aws-cdk-lib.aws_bedrockagentcore.InlineToolSchema",
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
		"aws-cdk-lib.aws_bedrockagentcore.InterceptionPoint",
		reflect.TypeOf((*InterceptionPoint)(nil)).Elem(),
		map[string]interface{}{
			"REQUEST": InterceptionPoint_REQUEST,
			"RESPONSE": InterceptionPoint_RESPONSE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.InterceptorBindConfig",
		reflect.TypeOf((*InterceptorBindConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.InterceptorOptions",
		reflect.TypeOf((*InterceptorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.InvocationConfiguration",
		reflect.TypeOf((*InvocationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.LambdaInterceptor",
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
		"aws-cdk-lib.aws_bedrockagentcore.LambdaTargetConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.LifecycleConfiguration",
		reflect.TypeOf((*LifecycleConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.LinkedinOAuth2CredentialProviderProps",
		reflect.TypeOf((*LinkedinOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.LlmAsAJudgeOptions",
		reflect.TypeOf((*LlmAsAJudgeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.LogType",
		reflect.TypeOf((*LogType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_LogType{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.LoggingConfig",
		reflect.TypeOf((*LoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.LoggingDestination",
		reflect.TypeOf((*LoggingDestination)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_LoggingDestination{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.MCPProtocolVersion",
		reflect.TypeOf((*MCPProtocolVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MCPProtocolVersion{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.ManagedMemoryStrategy",
		reflect.TypeOf((*ManagedMemoryStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "consolidationOverride", GoGetter: "ConsolidationOverride"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "extractionOverride", GoGetter: "ExtractionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "namespaces", GoGetter: "Namespaces"},
			_jsii_.MemberProperty{JsiiProperty: "reflectionConfiguration", GoGetter: "ReflectionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
			_jsii_.MemberProperty{JsiiProperty: "strategyName", GoGetter: "StrategyName"},
			_jsii_.MemberProperty{JsiiProperty: "strategyType", GoGetter: "StrategyType"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedMemoryStrategy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMemoryStrategy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.ManagedStrategyProps",
		reflect.TypeOf((*ManagedStrategyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.McpConfiguration",
		reflect.TypeOf((*McpConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.McpGatewaySearchType",
		reflect.TypeOf((*McpGatewaySearchType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_McpGatewaySearchType{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.McpProtocolConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.McpServerTargetConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.McpTargetConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.McpTargetType",
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
		"aws-cdk-lib.aws_bedrockagentcore.Memory",
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
		"aws-cdk-lib.aws_bedrockagentcore.MemoryAttributes",
		reflect.TypeOf((*MemoryAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryBase",
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
		"aws-cdk-lib.aws_bedrockagentcore.MemoryProps",
		reflect.TypeOf((*MemoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		reflect.TypeOf((*MemoryStrategy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MemoryStrategy{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategyCommonProps",
		reflect.TypeOf((*MemoryStrategyCommonProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategyType",
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
		"aws-cdk-lib.aws_bedrockagentcore.MetadataConfiguration",
		reflect.TypeOf((*MetadataConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.MicrosoftOAuth2CredentialProviderProps",
		reflect.TypeOf((*MicrosoftOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.NetworkConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.NoAuthAuthorizer",
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
		"aws-cdk-lib.aws_bedrockagentcore.NotionOAuth2CredentialProviderProps",
		reflect.TypeOf((*NotionOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.NumericalRatingOption",
		reflect.TypeOf((*NumericalRatingOption)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2AuthorizationServerMetadata",
		reflect.TypeOf((*OAuth2AuthorizationServerMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2ClientCredentials",
		reflect.TypeOf((*OAuth2ClientCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProvider",
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
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderAttributes",
		reflect.TypeOf((*OAuth2CredentialProviderAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderBaseProps",
		reflect.TypeOf((*OAuth2CredentialProviderBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderFactoryBaseProps",
		reflect.TypeOf((*OAuth2CredentialProviderFactoryBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderIdentityPerms",
		reflect.TypeOf((*OAuth2CredentialProviderIdentityPerms)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_OAuth2CredentialProviderIdentityPerms{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderProps",
		reflect.TypeOf((*OAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		reflect.TypeOf((*OAuth2CredentialProviderVendor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_OAuth2CredentialProviderVendor{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.OAuthConfiguration",
		reflect.TypeOf((*OAuthConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.OnlineEvaluationBase",
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
		"aws-cdk-lib.aws_bedrockagentcore.OnlineEvaluationBaseProps",
		reflect.TypeOf((*OnlineEvaluationBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.OnlineEvaluationConfig",
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
		"aws-cdk-lib.aws_bedrockagentcore.OnlineEvaluationConfigAttributes",
		reflect.TypeOf((*OnlineEvaluationConfigAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.OnlineEvaluationConfigProps",
		reflect.TypeOf((*OnlineEvaluationConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.OpenApiTargetConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.OverrideConfig",
		reflect.TypeOf((*OverrideConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.ProtocolType",
		reflect.TypeOf((*ProtocolType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ProtocolType{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.RecordingConfig",
		reflect.TypeOf((*RecordingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.RedditOAuth2CredentialProviderProps",
		reflect.TypeOf((*RedditOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.RequestHeaderConfiguration",
		reflect.TypeOf((*RequestHeaderConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.Runtime",
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
		"aws-cdk-lib.aws_bedrockagentcore.RuntimeAuthorizerConfiguration",
		reflect.TypeOf((*RuntimeAuthorizerConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RuntimeAuthorizerConfiguration{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.RuntimeBase",
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
		"aws-cdk-lib.aws_bedrockagentcore.RuntimeCustomClaim",
		reflect.TypeOf((*RuntimeCustomClaim)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RuntimeCustomClaim{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.RuntimeEndpoint",
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
		"aws-cdk-lib.aws_bedrockagentcore.RuntimeEndpointAttributes",
		reflect.TypeOf((*RuntimeEndpointAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.RuntimeEndpointBase",
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
		"aws-cdk-lib.aws_bedrockagentcore.RuntimeEndpointProps",
		reflect.TypeOf((*RuntimeEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.RuntimeNetworkConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.RuntimeProps",
		reflect.TypeOf((*RuntimeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.S3ApiSchema",
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
		"aws-cdk-lib.aws_bedrockagentcore.S3ToolSchema",
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
		"aws-cdk-lib.aws_bedrockagentcore.SalesforceOAuth2CredentialProviderProps",
		reflect.TypeOf((*SalesforceOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.SchemaDefinition",
		reflect.TypeOf((*SchemaDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.SchemaDefinitionType",
		reflect.TypeOf((*SchemaDefinitionType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_SchemaDefinitionType{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.SelfManagedMemoryStrategy",
		reflect.TypeOf((*SelfManagedMemoryStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "historicalContextWindowSize", GoGetter: "HistoricalContextWindowSize"},
			_jsii_.MemberProperty{JsiiProperty: "invocationConfiguration", GoGetter: "InvocationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
			_jsii_.MemberProperty{JsiiProperty: "strategyName", GoGetter: "StrategyName"},
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
		"aws-cdk-lib.aws_bedrockagentcore.SelfManagedStrategyProps",
		reflect.TypeOf((*SelfManagedStrategyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.SlackOAuth2CredentialProviderProps",
		reflect.TypeOf((*SlackOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.SmithyTargetConfiguration",
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
		"aws-cdk-lib.aws_bedrockagentcore.SpotifyOAuth2CredentialProviderProps",
		reflect.TypeOf((*SpotifyOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.TargetConfigurationConfig",
		reflect.TypeOf((*TargetConfigurationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.ToolDefinition",
		reflect.TypeOf((*ToolDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.ToolSchema",
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
		"aws-cdk-lib.aws_bedrockagentcore.TriggerConditions",
		reflect.TypeOf((*TriggerConditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.TwitchOAuth2CredentialProviderProps",
		reflect.TypeOf((*TwitchOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.VpcConfigProps",
		reflect.TypeOf((*VpcConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.WorkloadIdentity",
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
		"aws-cdk-lib.aws_bedrockagentcore.WorkloadIdentityAttributes",
		reflect.TypeOf((*WorkloadIdentityAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrockagentcore.WorkloadIdentityPerms",
		reflect.TypeOf((*WorkloadIdentityPerms)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_WorkloadIdentityPerms{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.WorkloadIdentityProps",
		reflect.TypeOf((*WorkloadIdentityProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.XOAuth2CredentialProviderProps",
		reflect.TypeOf((*XOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.YandexOAuth2CredentialProviderProps",
		reflect.TypeOf((*YandexOAuth2CredentialProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrockagentcore.ZoomOAuth2CredentialProviderProps",
		reflect.TypeOf((*ZoomOAuth2CredentialProviderProps)(nil)).Elem(),
	)
}
