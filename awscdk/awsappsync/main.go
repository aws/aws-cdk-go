package awsappsync

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.ApiBase",
		reflect.TypeOf((*ApiBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiArn", GoGetter: "ApiArn"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApiBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApi)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ApiKeyConfig",
		reflect.TypeOf((*ApiKeyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncApiKeyConfig",
		reflect.TypeOf((*AppSyncApiKeyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncAuthProvider",
		reflect.TypeOf((*AppSyncAuthProvider)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.AppSyncAuthorizationType",
		reflect.TypeOf((*AppSyncAuthorizationType)(nil)).Elem(),
		map[string]interface{}{
			"API_KEY": AppSyncAuthorizationType_API_KEY,
			"IAM": AppSyncAuthorizationType_IAM,
			"USER_POOL": AppSyncAuthorizationType_USER_POOL,
			"OIDC": AppSyncAuthorizationType_OIDC,
			"LAMBDA": AppSyncAuthorizationType_LAMBDA,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncAwsIamConfig",
		reflect.TypeOf((*AppSyncAwsIamConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AppSyncBackedDataSource",
		reflect.TypeOf((*AppSyncBackedDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppSyncBackedDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AppSyncBaseDataSource)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncBackedDataSourceProps",
		reflect.TypeOf((*AppSyncBackedDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AppSyncBaseDataSource",
		reflect.TypeOf((*AppSyncBaseDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppSyncBaseDataSource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncBaseDataSourceProps",
		reflect.TypeOf((*AppSyncBaseDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncCognitoConfig",
		reflect.TypeOf((*AppSyncCognitoConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncDataSourceOptions",
		reflect.TypeOf((*AppSyncDataSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.AppSyncDataSourceType",
		reflect.TypeOf((*AppSyncDataSourceType)(nil)).Elem(),
		map[string]interface{}{
			"LAMBDA": AppSyncDataSourceType_LAMBDA,
			"DYNAMODB": AppSyncDataSourceType_DYNAMODB,
			"EVENTBRIDGE": AppSyncDataSourceType_EVENTBRIDGE,
			"OPENSEARCH_SERVICE": AppSyncDataSourceType_OPENSEARCH_SERVICE,
			"HTTP": AppSyncDataSourceType_HTTP,
			"RELATIONAL_DATABASE": AppSyncDataSourceType_RELATIONAL_DATABASE,
			"BEDROCK": AppSyncDataSourceType_BEDROCK,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncDomainOptions",
		reflect.TypeOf((*AppSyncDomainOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AppSyncDynamoDbDataSource",
		reflect.TypeOf((*AppSyncDynamoDbDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppSyncDynamoDbDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AppSyncBackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncDynamoDbDataSourceProps",
		reflect.TypeOf((*AppSyncDynamoDbDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AppSyncEventBridgeDataSource",
		reflect.TypeOf((*AppSyncEventBridgeDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppSyncEventBridgeDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AppSyncBackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncEventBridgeDataSourceProps",
		reflect.TypeOf((*AppSyncEventBridgeDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AppSyncEventResource",
		reflect.TypeOf((*AppSyncEventResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resourceArns", GoMethod: "ResourceArns"},
		},
		func() interface{} {
			return &jsiiProxy_AppSyncEventResource{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncExtendedDataSourceProps",
		reflect.TypeOf((*AppSyncExtendedDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.AppSyncFieldLogLevel",
		reflect.TypeOf((*AppSyncFieldLogLevel)(nil)).Elem(),
		map[string]interface{}{
			"NONE": AppSyncFieldLogLevel_NONE,
			"ERROR": AppSyncFieldLogLevel_ERROR,
			"INFO": AppSyncFieldLogLevel_INFO,
			"DEBUG": AppSyncFieldLogLevel_DEBUG,
			"ALL": AppSyncFieldLogLevel_ALL,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AppSyncHttpDataSource",
		reflect.TypeOf((*AppSyncHttpDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppSyncHttpDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AppSyncBackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncHttpDataSourceOptions",
		reflect.TypeOf((*AppSyncHttpDataSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncHttpDataSourceProps",
		reflect.TypeOf((*AppSyncHttpDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncLambdaAuthorizerConfig",
		reflect.TypeOf((*AppSyncLambdaAuthorizerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AppSyncLambdaDataSource",
		reflect.TypeOf((*AppSyncLambdaDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppSyncLambdaDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AppSyncBackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncLambdaDataSourceProps",
		reflect.TypeOf((*AppSyncLambdaDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncLogConfig",
		reflect.TypeOf((*AppSyncLogConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncOpenIdConnectConfig",
		reflect.TypeOf((*AppSyncOpenIdConnectConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AppSyncOpenSearchDataSource",
		reflect.TypeOf((*AppSyncOpenSearchDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppSyncOpenSearchDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AppSyncBackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncOpenSearchDataSourceProps",
		reflect.TypeOf((*AppSyncOpenSearchDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AppSyncRdsDataSource",
		reflect.TypeOf((*AppSyncRdsDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppSyncRdsDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AppSyncBackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncRdsDataSourceProps",
		reflect.TypeOf((*AppSyncRdsDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppSyncRdsDataSourcePropsV2",
		reflect.TypeOf((*AppSyncRdsDataSourcePropsV2)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AppsyncFunction",
		reflect.TypeOf((*AppsyncFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dataSource", GoGetter: "DataSource"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionId", GoGetter: "FunctionId"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppsyncFunction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAppsyncFunction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppsyncFunctionAttributes",
		reflect.TypeOf((*AppsyncFunctionAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AppsyncFunctionProps",
		reflect.TypeOf((*AppsyncFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AssetCode",
		reflect.TypeOf((*AssetCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
		},
		func() interface{} {
			j := jsiiProxy_AssetCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.Assign",
		reflect.TypeOf((*Assign)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "putInMap", GoMethod: "PutInMap"},
			_jsii_.MemberMethod{JsiiMethod: "renderAsAssignment", GoMethod: "RenderAsAssignment"},
		},
		func() interface{} {
			return &jsiiProxy_Assign{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AttributeValues",
		reflect.TypeOf((*AttributeValues)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "attribute", GoMethod: "Attribute"},
			_jsii_.MemberMethod{JsiiMethod: "renderTemplate", GoMethod: "RenderTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "renderVariables", GoMethod: "RenderVariables"},
		},
		func() interface{} {
			return &jsiiProxy_AttributeValues{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.AttributeValuesStep",
		reflect.TypeOf((*AttributeValuesStep)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "is", GoMethod: "Is"},
		},
		func() interface{} {
			return &jsiiProxy_AttributeValuesStep{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AuthorizationConfig",
		reflect.TypeOf((*AuthorizationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AuthorizationMode",
		reflect.TypeOf((*AuthorizationMode)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.AuthorizationType",
		reflect.TypeOf((*AuthorizationType)(nil)).Elem(),
		map[string]interface{}{
			"API_KEY": AuthorizationType_API_KEY,
			"IAM": AuthorizationType_IAM,
			"USER_POOL": AuthorizationType_USER_POOL,
			"OIDC": AuthorizationType_OIDC,
			"LAMBDA": AuthorizationType_LAMBDA,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AwsIamConfig",
		reflect.TypeOf((*AwsIamConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.BackedDataSource",
		reflect.TypeOf((*BackedDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "createFunction", GoMethod: "CreateFunction"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "ds", GoGetter: "Ds"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BackedDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseDataSource)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.BackedDataSourceProps",
		reflect.TypeOf((*BackedDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.BaseAppsyncFunctionProps",
		reflect.TypeOf((*BaseAppsyncFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.BaseChannelNamespaceProps",
		reflect.TypeOf((*BaseChannelNamespaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.BaseDataSource",
		reflect.TypeOf((*BaseDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "createFunction", GoMethod: "CreateFunction"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "ds", GoGetter: "Ds"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BaseDataSource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.BaseDataSourceProps",
		reflect.TypeOf((*BaseDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.BaseResolverProps",
		reflect.TypeOf((*BaseResolverProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CachingConfig",
		reflect.TypeOf((*CachingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnApi",
		reflect.TypeOf((*CfnApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiArn", GoGetter: "AttrApiArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiId", GoGetter: "AttrApiId"},
			_jsii_.MemberProperty{JsiiProperty: "attrDns", GoGetter: "AttrDns"},
			_jsii_.MemberProperty{JsiiProperty: "attrDnsHttp", GoGetter: "AttrDnsHttp"},
			_jsii_.MemberProperty{JsiiProperty: "attrDnsRealtime", GoGetter: "AttrDnsRealtime"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventConfig", GoGetter: "EventConfig"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ownerContact", GoGetter: "OwnerContact"},
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
		},
		func() interface{} {
			j := jsiiProxy_CfnApi{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApi.AuthModeProperty",
		reflect.TypeOf((*CfnApi_AuthModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApi.AuthProviderProperty",
		reflect.TypeOf((*CfnApi_AuthProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApi.CognitoConfigProperty",
		reflect.TypeOf((*CfnApi_CognitoConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApi.DnsMapProperty",
		reflect.TypeOf((*CfnApi_DnsMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApi.EventConfigProperty",
		reflect.TypeOf((*CfnApi_EventConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApi.EventLogConfigProperty",
		reflect.TypeOf((*CfnApi_EventLogConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApi.LambdaAuthorizerConfigProperty",
		reflect.TypeOf((*CfnApi_LambdaAuthorizerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApi.OpenIDConnectConfigProperty",
		reflect.TypeOf((*CfnApi_OpenIDConnectConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnApiCache",
		reflect.TypeOf((*CfnApiCache)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiCachingBehavior", GoGetter: "ApiCachingBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "atRestEncryptionEnabled", GoGetter: "AtRestEncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "healthMetricsConfig", GoGetter: "HealthMetricsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
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
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitEncryptionEnabled", GoGetter: "TransitEncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ttl", GoGetter: "Ttl"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApiCache{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApiCacheProps",
		reflect.TypeOf((*CfnApiCacheProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnApiKey",
		reflect.TypeOf((*CfnApiKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiKey", GoGetter: "AttrApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiKeyId", GoGetter: "AttrApiKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "expires", GoGetter: "Expires"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
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
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApiKey{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApiKeyProps",
		reflect.TypeOf((*CfnApiKeyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnApiProps",
		reflect.TypeOf((*CfnApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnChannelNamespace",
		reflect.TypeOf((*CfnChannelNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrChannelNamespaceArn", GoGetter: "AttrChannelNamespaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "codeHandlers", GoGetter: "CodeHandlers"},
			_jsii_.MemberProperty{JsiiProperty: "codeS3Location", GoGetter: "CodeS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "handlerConfigs", GoGetter: "HandlerConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "publishAuthModes", GoGetter: "PublishAuthModes"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subscribeAuthModes", GoGetter: "SubscribeAuthModes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelNamespace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnChannelNamespace.AuthModeProperty",
		reflect.TypeOf((*CfnChannelNamespace_AuthModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnChannelNamespace.HandlerConfigProperty",
		reflect.TypeOf((*CfnChannelNamespace_HandlerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnChannelNamespace.HandlerConfigsProperty",
		reflect.TypeOf((*CfnChannelNamespace_HandlerConfigsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnChannelNamespace.IntegrationProperty",
		reflect.TypeOf((*CfnChannelNamespace_IntegrationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnChannelNamespace.LambdaConfigProperty",
		reflect.TypeOf((*CfnChannelNamespace_LambdaConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnChannelNamespaceProps",
		reflect.TypeOf((*CfnChannelNamespaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnDataSource",
		reflect.TypeOf((*CfnDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDataSourceArn", GoGetter: "AttrDataSourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrName", GoGetter: "AttrName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "dynamoDbConfig", GoGetter: "DynamoDbConfig"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchConfig", GoGetter: "ElasticsearchConfig"},
			_jsii_.MemberProperty{JsiiProperty: "eventBridgeConfig", GoGetter: "EventBridgeConfig"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "httpConfig", GoGetter: "HttpConfig"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaConfig", GoGetter: "LambdaConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "metricsConfig", GoGetter: "MetricsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "openSearchServiceConfig", GoGetter: "OpenSearchServiceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "relationalDatabaseConfig", GoGetter: "RelationalDatabaseConfig"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleArn", GoGetter: "ServiceRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.AuthorizationConfigProperty",
		reflect.TypeOf((*CfnDataSource_AuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.AwsIamConfigProperty",
		reflect.TypeOf((*CfnDataSource_AwsIamConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.DeltaSyncConfigProperty",
		reflect.TypeOf((*CfnDataSource_DeltaSyncConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.DynamoDBConfigProperty",
		reflect.TypeOf((*CfnDataSource_DynamoDBConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.ElasticsearchConfigProperty",
		reflect.TypeOf((*CfnDataSource_ElasticsearchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.EventBridgeConfigProperty",
		reflect.TypeOf((*CfnDataSource_EventBridgeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.HttpConfigProperty",
		reflect.TypeOf((*CfnDataSource_HttpConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.LambdaConfigProperty",
		reflect.TypeOf((*CfnDataSource_LambdaConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.OpenSearchServiceConfigProperty",
		reflect.TypeOf((*CfnDataSource_OpenSearchServiceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.RdsHttpEndpointConfigProperty",
		reflect.TypeOf((*CfnDataSource_RdsHttpEndpointConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.RelationalDatabaseConfigProperty",
		reflect.TypeOf((*CfnDataSource_RelationalDatabaseConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSourceProps",
		reflect.TypeOf((*CfnDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnDomainName",
		reflect.TypeOf((*CfnDomainName)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAppSyncDomainName", GoGetter: "AttrAppSyncDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomainName", GoGetter: "AttrDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomainNameArn", GoGetter: "AttrDomainNameArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrHostedZoneId", GoGetter: "AttrHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "certificateArn", GoGetter: "CertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
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
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainName{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnDomainNameApiAssociation",
		reflect.TypeOf((*CfnDomainNameApiAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiAssociationIdentifier", GoGetter: "AttrApiAssociationIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
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
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainNameApiAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDomainNameApiAssociationProps",
		reflect.TypeOf((*CfnDomainNameApiAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDomainNameProps",
		reflect.TypeOf((*CfnDomainNameProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration",
		reflect.TypeOf((*CfnFunctionConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDataSourceName", GoGetter: "AttrDataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "attrFunctionArn", GoGetter: "AttrFunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrFunctionId", GoGetter: "AttrFunctionId"},
			_jsii_.MemberProperty{JsiiProperty: "attrName", GoGetter: "AttrName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "code", GoGetter: "Code"},
			_jsii_.MemberProperty{JsiiProperty: "codeS3Location", GoGetter: "CodeS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "functionVersion", GoGetter: "FunctionVersion"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maxBatchSize", GoGetter: "MaxBatchSize"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "requestMappingTemplate", GoGetter: "RequestMappingTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "requestMappingTemplateS3Location", GoGetter: "RequestMappingTemplateS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "responseMappingTemplate", GoGetter: "ResponseMappingTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "responseMappingTemplateS3Location", GoGetter: "ResponseMappingTemplateS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "runtime", GoGetter: "Runtime"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "syncConfig", GoGetter: "SyncConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFunctionConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration.AppSyncRuntimeProperty",
		reflect.TypeOf((*CfnFunctionConfiguration_AppSyncRuntimeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration.LambdaConflictHandlerConfigProperty",
		reflect.TypeOf((*CfnFunctionConfiguration_LambdaConflictHandlerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration.SyncConfigProperty",
		reflect.TypeOf((*CfnFunctionConfiguration_SyncConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfigurationProps",
		reflect.TypeOf((*CfnFunctionConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		reflect.TypeOf((*CfnGraphQLApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "additionalAuthenticationProviders", GoGetter: "AdditionalAuthenticationProviders"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiType", GoGetter: "ApiType"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiId", GoGetter: "AttrApiId"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrGraphQlDns", GoGetter: "AttrGraphQlDns"},
			_jsii_.MemberProperty{JsiiProperty: "attrGraphQlEndpointArn", GoGetter: "AttrGraphQlEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrGraphQlUrl", GoGetter: "AttrGraphQlUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrRealtimeDns", GoGetter: "AttrRealtimeDns"},
			_jsii_.MemberProperty{JsiiProperty: "attrRealtimeUrl", GoGetter: "AttrRealtimeUrl"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationType", GoGetter: "AuthenticationType"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enhancedMetricsConfig", GoGetter: "EnhancedMetricsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "environmentVariables", GoGetter: "EnvironmentVariables"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "introspectionConfig", GoGetter: "IntrospectionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaAuthorizerConfig", GoGetter: "LambdaAuthorizerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logConfig", GoGetter: "LogConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "mergedApiExecutionRoleArn", GoGetter: "MergedApiExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectConfig", GoGetter: "OpenIdConnectConfig"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ownerContact", GoGetter: "OwnerContact"},
			_jsii_.MemberProperty{JsiiProperty: "queryDepthLimit", GoGetter: "QueryDepthLimit"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "resolverCountLimit", GoGetter: "ResolverCountLimit"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolConfig", GoGetter: "UserPoolConfig"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "visibility", GoGetter: "Visibility"},
			_jsii_.MemberProperty{JsiiProperty: "xrayEnabled", GoGetter: "XrayEnabled"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGraphQLApi{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi.AdditionalAuthenticationProviderProperty",
		reflect.TypeOf((*CfnGraphQLApi_AdditionalAuthenticationProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi.CognitoUserPoolConfigProperty",
		reflect.TypeOf((*CfnGraphQLApi_CognitoUserPoolConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi.EnhancedMetricsConfigProperty",
		reflect.TypeOf((*CfnGraphQLApi_EnhancedMetricsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi.LambdaAuthorizerConfigProperty",
		reflect.TypeOf((*CfnGraphQLApi_LambdaAuthorizerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi.LogConfigProperty",
		reflect.TypeOf((*CfnGraphQLApi_LogConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi.OpenIDConnectConfigProperty",
		reflect.TypeOf((*CfnGraphQLApi_OpenIDConnectConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi.UserPoolConfigProperty",
		reflect.TypeOf((*CfnGraphQLApi_UserPoolConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApiProps",
		reflect.TypeOf((*CfnGraphQLApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnGraphQLSchema",
		reflect.TypeOf((*CfnGraphQLSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "definitionS3Location", GoGetter: "DefinitionS3Location"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
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
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGraphQLSchema{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnGraphQLSchemaProps",
		reflect.TypeOf((*CfnGraphQLSchemaProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnResolver",
		reflect.TypeOf((*CfnResolver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrFieldName", GoGetter: "AttrFieldName"},
			_jsii_.MemberProperty{JsiiProperty: "attrResolverArn", GoGetter: "AttrResolverArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrTypeName", GoGetter: "AttrTypeName"},
			_jsii_.MemberProperty{JsiiProperty: "cachingConfig", GoGetter: "CachingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "code", GoGetter: "Code"},
			_jsii_.MemberProperty{JsiiProperty: "codeS3Location", GoGetter: "CodeS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "fieldName", GoGetter: "FieldName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maxBatchSize", GoGetter: "MaxBatchSize"},
			_jsii_.MemberProperty{JsiiProperty: "metricsConfig", GoGetter: "MetricsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineConfig", GoGetter: "PipelineConfig"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "requestMappingTemplate", GoGetter: "RequestMappingTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "requestMappingTemplateS3Location", GoGetter: "RequestMappingTemplateS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "responseMappingTemplate", GoGetter: "ResponseMappingTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "responseMappingTemplateS3Location", GoGetter: "ResponseMappingTemplateS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "runtime", GoGetter: "Runtime"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "syncConfig", GoGetter: "SyncConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResolver{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnResolver.AppSyncRuntimeProperty",
		reflect.TypeOf((*CfnResolver_AppSyncRuntimeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnResolver.CachingConfigProperty",
		reflect.TypeOf((*CfnResolver_CachingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnResolver.LambdaConflictHandlerConfigProperty",
		reflect.TypeOf((*CfnResolver_LambdaConflictHandlerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnResolver.PipelineConfigProperty",
		reflect.TypeOf((*CfnResolver_PipelineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnResolver.SyncConfigProperty",
		reflect.TypeOf((*CfnResolver_SyncConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnResolverProps",
		reflect.TypeOf((*CfnResolverProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnSourceApiAssociation",
		reflect.TypeOf((*CfnSourceApiAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAssociationArn", GoGetter: "AttrAssociationArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrAssociationId", GoGetter: "AttrAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastSuccessfulMergeDate", GoGetter: "AttrLastSuccessfulMergeDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrMergedApiArn", GoGetter: "AttrMergedApiArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrMergedApiId", GoGetter: "AttrMergedApiId"},
			_jsii_.MemberProperty{JsiiProperty: "attrSourceApiArn", GoGetter: "AttrSourceApiArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrSourceApiAssociationStatus", GoGetter: "AttrSourceApiAssociationStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrSourceApiAssociationStatusDetail", GoGetter: "AttrSourceApiAssociationStatusDetail"},
			_jsii_.MemberProperty{JsiiProperty: "attrSourceApiId", GoGetter: "AttrSourceApiId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "mergedApiIdentifier", GoGetter: "MergedApiIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceApiAssociationConfig", GoGetter: "SourceApiAssociationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "sourceApiIdentifier", GoGetter: "SourceApiIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSourceApiAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnSourceApiAssociation.SourceApiAssociationConfigProperty",
		reflect.TypeOf((*CfnSourceApiAssociation_SourceApiAssociationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnSourceApiAssociationProps",
		reflect.TypeOf((*CfnSourceApiAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.ChannelNamespace",
		reflect.TypeOf((*ChannelNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "channelNamespaceArn", GoGetter: "ChannelNamespaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublish", GoMethod: "GrantPublish"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublishAndSubscribe", GoMethod: "GrantPublishAndSubscribe"},
			_jsii_.MemberMethod{JsiiMethod: "grantSubscribe", GoMethod: "GrantSubscribe"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ChannelNamespace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IChannelNamespace)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ChannelNamespaceOptions",
		reflect.TypeOf((*ChannelNamespaceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ChannelNamespaceProps",
		reflect.TypeOf((*ChannelNamespaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.Code",
		reflect.TypeOf((*Code)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_Code{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CodeConfig",
		reflect.TypeOf((*CodeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.DataSourceOptions",
		reflect.TypeOf((*DataSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.Definition",
		reflect.TypeOf((*Definition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "sourceApiOptions", GoGetter: "SourceApiOptions"},
		},
		func() interface{} {
			return &jsiiProxy_Definition{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.DomainOptions",
		reflect.TypeOf((*DomainOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.DynamoDbDataSource",
		reflect.TypeOf((*DynamoDbDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "createFunction", GoMethod: "CreateFunction"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "ds", GoGetter: "Ds"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DynamoDbDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.DynamoDbDataSourceProps",
		reflect.TypeOf((*DynamoDbDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.ElasticsearchDataSource",
		reflect.TypeOf((*ElasticsearchDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "createFunction", GoMethod: "CreateFunction"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "ds", GoGetter: "Ds"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticsearchDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ElasticsearchDataSourceProps",
		reflect.TypeOf((*ElasticsearchDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.EventApi",
		reflect.TypeOf((*EventApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addChannelNamespace", GoMethod: "AddChannelNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addEventBridgeDataSource", GoMethod: "AddEventBridgeDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addOpenSearchDataSource", GoMethod: "AddOpenSearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberProperty{JsiiProperty: "apiArn", GoGetter: "ApiArn"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeys", GoGetter: "ApiKeys"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "appSyncDomainName", GoGetter: "AppSyncDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "authProviderTypes", GoGetter: "AuthProviderTypes"},
			_jsii_.MemberProperty{JsiiProperty: "connectionModeTypes", GoGetter: "ConnectionModeTypes"},
			_jsii_.MemberProperty{JsiiProperty: "customHttpEndpoint", GoGetter: "CustomHttpEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "customRealtimeEndpoint", GoGetter: "CustomRealtimeEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPublishModeTypes", GoGetter: "DefaultPublishModeTypes"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSubscribeModeTypes", GoGetter: "DefaultSubscribeModeTypes"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublish", GoMethod: "GrantPublish"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublishAndSubscribe", GoMethod: "GrantPublishAndSubscribe"},
			_jsii_.MemberMethod{JsiiMethod: "grantSubscribe", GoMethod: "GrantSubscribe"},
			_jsii_.MemberProperty{JsiiProperty: "httpDns", GoGetter: "HttpDns"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeDns", GoGetter: "RealtimeDns"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventApi{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EventApiBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.EventApiAttributes",
		reflect.TypeOf((*EventApiAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.EventApiAuthConfig",
		reflect.TypeOf((*EventApiAuthConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.EventApiBase",
		reflect.TypeOf((*EventApiBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addChannelNamespace", GoMethod: "AddChannelNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addEventBridgeDataSource", GoMethod: "AddEventBridgeDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addOpenSearchDataSource", GoMethod: "AddOpenSearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberProperty{JsiiProperty: "apiArn", GoGetter: "ApiArn"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authProviderTypes", GoGetter: "AuthProviderTypes"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublish", GoMethod: "GrantPublish"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublishAndSubscribe", GoMethod: "GrantPublishAndSubscribe"},
			_jsii_.MemberMethod{JsiiMethod: "grantSubscribe", GoMethod: "GrantSubscribe"},
			_jsii_.MemberProperty{JsiiProperty: "httpDns", GoGetter: "HttpDns"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeDns", GoGetter: "RealtimeDns"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventApiBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEventApi)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.EventApiProps",
		reflect.TypeOf((*EventApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.EventBridgeDataSource",
		reflect.TypeOf((*EventBridgeDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "createFunction", GoMethod: "CreateFunction"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "ds", GoGetter: "Ds"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventBridgeDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.EventBridgeDataSourceProps",
		reflect.TypeOf((*EventBridgeDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ExtendedDataSourceProps",
		reflect.TypeOf((*ExtendedDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ExtendedResolverProps",
		reflect.TypeOf((*ExtendedResolverProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.FieldLogLevel",
		reflect.TypeOf((*FieldLogLevel)(nil)).Elem(),
		map[string]interface{}{
			"NONE": FieldLogLevel_NONE,
			"ERROR": FieldLogLevel_ERROR,
			"INFO": FieldLogLevel_INFO,
			"DEBUG": FieldLogLevel_DEBUG,
			"ALL": FieldLogLevel_ALL,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.FunctionRuntime",
		reflect.TypeOf((*FunctionRuntime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "toProperties", GoMethod: "ToProperties"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_FunctionRuntime{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.FunctionRuntimeFamily",
		reflect.TypeOf((*FunctionRuntimeFamily)(nil)).Elem(),
		map[string]interface{}{
			"JS": FunctionRuntimeFamily_JS,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.GraphqlApi",
		reflect.TypeOf((*GraphqlApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addElasticsearchDataSource", GoMethod: "AddElasticsearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addEnvironmentVariable", GoMethod: "AddEnvironmentVariable"},
			_jsii_.MemberMethod{JsiiMethod: "addEventBridgeDataSource", GoMethod: "AddEventBridgeDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addNoneDataSource", GoMethod: "AddNoneDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addOpenSearchDataSource", GoMethod: "AddOpenSearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSourceV2", GoMethod: "AddRdsDataSourceV2"},
			_jsii_.MemberMethod{JsiiMethod: "addSchemaDependency", GoMethod: "AddSchemaDependency"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "apiKey", GoGetter: "ApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "appSyncDomainName", GoGetter: "AppSyncDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantMutation", GoMethod: "GrantMutation"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantSubscription", GoMethod: "GrantSubscription"},
			_jsii_.MemberProperty{JsiiProperty: "graphQLEndpointArn", GoGetter: "GraphQLEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "graphqlUrl", GoGetter: "GraphqlUrl"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "visibility", GoGetter: "Visibility"},
		},
		func() interface{} {
			j := jsiiProxy_GraphqlApi{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GraphqlApiBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.GraphqlApiAttributes",
		reflect.TypeOf((*GraphqlApiAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.GraphqlApiBase",
		reflect.TypeOf((*GraphqlApiBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addElasticsearchDataSource", GoMethod: "AddElasticsearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addEventBridgeDataSource", GoMethod: "AddEventBridgeDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addNoneDataSource", GoMethod: "AddNoneDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addOpenSearchDataSource", GoMethod: "AddOpenSearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSourceV2", GoMethod: "AddRdsDataSourceV2"},
			_jsii_.MemberMethod{JsiiMethod: "addSchemaDependency", GoMethod: "AddSchemaDependency"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantMutation", GoMethod: "GrantMutation"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantSubscription", GoMethod: "GrantSubscription"},
			_jsii_.MemberProperty{JsiiProperty: "graphQLEndpointArn", GoGetter: "GraphQLEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "visibility", GoGetter: "Visibility"},
		},
		func() interface{} {
			j := jsiiProxy_GraphqlApiBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGraphqlApi)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.GraphqlApiProps",
		reflect.TypeOf((*GraphqlApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.HandlerBehavior",
		reflect.TypeOf((*HandlerBehavior)(nil)).Elem(),
		map[string]interface{}{
			"CODE": HandlerBehavior_CODE,
			"DIRECT": HandlerBehavior_DIRECT,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.HandlerConfig",
		reflect.TypeOf((*HandlerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.HttpDataSource",
		reflect.TypeOf((*HttpDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "createFunction", GoMethod: "CreateFunction"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "ds", GoGetter: "Ds"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HttpDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.HttpDataSourceOptions",
		reflect.TypeOf((*HttpDataSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.HttpDataSourceProps",
		reflect.TypeOf((*HttpDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.IApi",
		reflect.TypeOf((*IApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiArn", GoGetter: "ApiArn"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IApi{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.IAppSyncAuthConfig",
		reflect.TypeOf((*IAppSyncAuthConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "setupCognitoConfig", GoMethod: "SetupCognitoConfig"},
			_jsii_.MemberMethod{JsiiMethod: "setupLambdaAuthorizerConfig", GoMethod: "SetupLambdaAuthorizerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "setupOpenIdConnectConfig", GoMethod: "SetupOpenIdConnectConfig"},
		},
		func() interface{} {
			return &jsiiProxy_IAppSyncAuthConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.IAppsyncFunction",
		reflect.TypeOf((*IAppsyncFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionId", GoGetter: "FunctionId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAppsyncFunction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.IChannelNamespace",
		reflect.TypeOf((*IChannelNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "channelNamespaceArn", GoGetter: "ChannelNamespaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IChannelNamespace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.IEventApi",
		reflect.TypeOf((*IEventApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addChannelNamespace", GoMethod: "AddChannelNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addEventBridgeDataSource", GoMethod: "AddEventBridgeDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addOpenSearchDataSource", GoMethod: "AddOpenSearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberProperty{JsiiProperty: "apiArn", GoGetter: "ApiArn"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authProviderTypes", GoGetter: "AuthProviderTypes"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublish", GoMethod: "GrantPublish"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublishAndSubscribe", GoMethod: "GrantPublishAndSubscribe"},
			_jsii_.MemberMethod{JsiiMethod: "grantSubscribe", GoMethod: "GrantSubscribe"},
			_jsii_.MemberProperty{JsiiProperty: "httpDns", GoGetter: "HttpDns"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeDns", GoGetter: "RealtimeDns"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IEventApi{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApi)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.IGraphqlApi",
		reflect.TypeOf((*IGraphqlApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addElasticsearchDataSource", GoMethod: "AddElasticsearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addEventBridgeDataSource", GoMethod: "AddEventBridgeDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addNoneDataSource", GoMethod: "AddNoneDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addOpenSearchDataSource", GoMethod: "AddOpenSearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSourceV2", GoMethod: "AddRdsDataSourceV2"},
			_jsii_.MemberMethod{JsiiMethod: "addSchemaDependency", GoMethod: "AddSchemaDependency"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantMutation", GoMethod: "GrantMutation"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantSubscription", GoMethod: "GrantSubscription"},
			_jsii_.MemberProperty{JsiiProperty: "graphQLEndpointArn", GoGetter: "GraphQLEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "visibility", GoGetter: "Visibility"},
		},
		func() interface{} {
			j := jsiiProxy_IGraphqlApi{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.ISchema",
		reflect.TypeOf((*ISchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ISchema{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.ISchemaConfig",
		reflect.TypeOf((*ISchemaConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
		},
		func() interface{} {
			return &jsiiProxy_ISchemaConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.ISourceApiAssociation",
		reflect.TypeOf((*ISourceApiAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "associationArn", GoGetter: "AssociationArn"},
			_jsii_.MemberProperty{JsiiProperty: "associationId", GoGetter: "AssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "mergedApi", GoGetter: "MergedApi"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "sourceApi", GoGetter: "SourceApi"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ISourceApiAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.IamResource",
		reflect.TypeOf((*IamResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resourceArns", GoMethod: "ResourceArns"},
		},
		func() interface{} {
			return &jsiiProxy_IamResource{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.InlineCode",
		reflect.TypeOf((*InlineCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_InlineCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.IntrospectionConfig",
		reflect.TypeOf((*IntrospectionConfig)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": IntrospectionConfig_ENABLED,
			"DISABLED": IntrospectionConfig_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.KeyCondition",
		reflect.TypeOf((*KeyCondition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "and", GoMethod: "And"},
			_jsii_.MemberMethod{JsiiMethod: "renderTemplate", GoMethod: "RenderTemplate"},
		},
		func() interface{} {
			return &jsiiProxy_KeyCondition{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.LambdaAuthorizerConfig",
		reflect.TypeOf((*LambdaAuthorizerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.LambdaDataSource",
		reflect.TypeOf((*LambdaDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "createFunction", GoMethod: "CreateFunction"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "ds", GoGetter: "Ds"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.LambdaDataSourceProps",
		reflect.TypeOf((*LambdaDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.LambdaInvokeType",
		reflect.TypeOf((*LambdaInvokeType)(nil)).Elem(),
		map[string]interface{}{
			"EVENT": LambdaInvokeType_EVENT,
			"REQUEST_RESPONSE": LambdaInvokeType_REQUEST_RESPONSE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.LogConfig",
		reflect.TypeOf((*LogConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.MappingTemplate",
		reflect.TypeOf((*MappingTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "renderTemplate", GoMethod: "RenderTemplate"},
		},
		func() interface{} {
			return &jsiiProxy_MappingTemplate{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.MergeType",
		reflect.TypeOf((*MergeType)(nil)).Elem(),
		map[string]interface{}{
			"MANUAL_MERGE": MergeType_MANUAL_MERGE,
			"AUTO_MERGE": MergeType_AUTO_MERGE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.NamespaceAuthConfig",
		reflect.TypeOf((*NamespaceAuthConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.NoneDataSource",
		reflect.TypeOf((*NoneDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "createFunction", GoMethod: "CreateFunction"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "ds", GoGetter: "Ds"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NoneDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.NoneDataSourceProps",
		reflect.TypeOf((*NoneDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.OpenIdConnectConfig",
		reflect.TypeOf((*OpenIdConnectConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.OpenSearchDataSource",
		reflect.TypeOf((*OpenSearchDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "createFunction", GoMethod: "CreateFunction"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "ds", GoGetter: "Ds"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OpenSearchDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.OpenSearchDataSourceProps",
		reflect.TypeOf((*OpenSearchDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.PartitionKey",
		reflect.TypeOf((*PartitionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "pkey", GoGetter: "Pkey"},
			_jsii_.MemberMethod{JsiiMethod: "renderTemplate", GoMethod: "RenderTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "sort", GoMethod: "Sort"},
		},
		func() interface{} {
			j := jsiiProxy_PartitionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PrimaryKey)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.PartitionKeyStep",
		reflect.TypeOf((*PartitionKeyStep)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "auto", GoMethod: "Auto"},
			_jsii_.MemberMethod{JsiiMethod: "is", GoMethod: "Is"},
		},
		func() interface{} {
			return &jsiiProxy_PartitionKeyStep{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.PrimaryKey",
		reflect.TypeOf((*PrimaryKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "pkey", GoGetter: "Pkey"},
			_jsii_.MemberMethod{JsiiMethod: "renderTemplate", GoMethod: "RenderTemplate"},
		},
		func() interface{} {
			return &jsiiProxy_PrimaryKey{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.RdsDataSource",
		reflect.TypeOf((*RdsDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "createFunction", GoMethod: "CreateFunction"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "ds", GoGetter: "Ds"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RdsDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BackedDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.RdsDataSourceProps",
		reflect.TypeOf((*RdsDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.RdsDataSourcePropsV2",
		reflect.TypeOf((*RdsDataSourcePropsV2)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.Resolver",
		reflect.TypeOf((*Resolver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Resolver{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ResolverProps",
		reflect.TypeOf((*ResolverProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.RuntimeConfig",
		reflect.TypeOf((*RuntimeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.SchemaBindOptions",
		reflect.TypeOf((*SchemaBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.SchemaFile",
		reflect.TypeOf((*SchemaFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
		},
		func() interface{} {
			j := jsiiProxy_SchemaFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISchema)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.SchemaProps",
		reflect.TypeOf((*SchemaProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.SortKeyStep",
		reflect.TypeOf((*SortKeyStep)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "auto", GoMethod: "Auto"},
			_jsii_.MemberMethod{JsiiMethod: "is", GoMethod: "Is"},
		},
		func() interface{} {
			return &jsiiProxy_SortKeyStep{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.SourceApi",
		reflect.TypeOf((*SourceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.SourceApiAssociation",
		reflect.TypeOf((*SourceApiAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "association", GoGetter: "Association"},
			_jsii_.MemberProperty{JsiiProperty: "associationArn", GoGetter: "AssociationArn"},
			_jsii_.MemberProperty{JsiiProperty: "associationId", GoGetter: "AssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mergedApi", GoGetter: "MergedApi"},
			_jsii_.MemberProperty{JsiiProperty: "mergeType", GoGetter: "MergeType"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceApi", GoGetter: "SourceApi"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SourceApiAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISourceApiAssociation)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.SourceApiAssociationAttributes",
		reflect.TypeOf((*SourceApiAssociationAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.SourceApiAssociationProps",
		reflect.TypeOf((*SourceApiAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.SourceApiOptions",
		reflect.TypeOf((*SourceApiOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.UserPoolConfig",
		reflect.TypeOf((*UserPoolConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.UserPoolDefaultAction",
		reflect.TypeOf((*UserPoolDefaultAction)(nil)).Elem(),
		map[string]interface{}{
			"ALLOW": UserPoolDefaultAction_ALLOW,
			"DENY": UserPoolDefaultAction_DENY,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.Values",
		reflect.TypeOf((*Values)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Values{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.Visibility",
		reflect.TypeOf((*Visibility)(nil)).Elem(),
		map[string]interface{}{
			"GLOBAL": Visibility_GLOBAL,
			"PRIVATE": Visibility_PRIVATE,
		},
	)
}
