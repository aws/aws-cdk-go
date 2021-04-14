package awsappsync

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.AddFieldOptions",
		reflect.TypeOf((*AddFieldOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ApiKeyConfig",
		reflect.TypeOf((*ApiKeyConfig)(nil)).Elem(),
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
		"aws-cdk-lib.aws_appsync.BaseTypeOptions",
		reflect.TypeOf((*BaseTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnApiCache",
		reflect.TypeOf((*CfnApiCache)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiCachingBehavior", GoGetter: "ApiCachingBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "atRestEncryptionEnabled", GoGetter: "AtRestEncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitEncryptionEnabled", GoGetter: "TransitEncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ttl", GoGetter: "Ttl"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyId", GoGetter: "ApiKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiKey", GoGetter: "AttrApiKey"},
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.CfnDataSource",
		reflect.TypeOf((*CfnDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
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
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "httpConfig", GoGetter: "HttpConfig"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaConfig", GoGetter: "LambdaConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "relationalDatabaseConfig", GoGetter: "RelationalDatabaseConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleArn", GoGetter: "ServiceRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
		"aws-cdk-lib.aws_appsync.CfnDataSource.HttpConfigProperty",
		reflect.TypeOf((*CfnDataSource_HttpConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.CfnDataSource.LambdaConfigProperty",
		reflect.TypeOf((*CfnDataSource_LambdaConfigProperty)(nil)).Elem(),
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
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration",
		reflect.TypeOf((*CfnFunctionConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
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
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "functionVersion", GoGetter: "FunctionVersion"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "requestMappingTemplate", GoGetter: "RequestMappingTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "requestMappingTemplateS3Location", GoGetter: "RequestMappingTemplateS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "responseMappingTemplate", GoGetter: "ResponseMappingTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "responseMappingTemplateS3Location", GoGetter: "ResponseMappingTemplateS3Location"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "syncConfig", GoGetter: "SyncConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "additionalAuthenticationProviders", GoGetter: "AdditionalAuthenticationProviders"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiId", GoGetter: "AttrApiId"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrGraphQlUrl", GoGetter: "AttrGraphQlUrl"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationType", GoGetter: "AuthenticationType"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaAuthorizerConfig", GoGetter: "LambdaAuthorizerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logConfig", GoGetter: "LogConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectConfig", GoGetter: "OpenIdConnectConfig"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolConfig", GoGetter: "UserPoolConfig"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "xrayEnabled", GoGetter: "XrayEnabled"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGraphQLApi{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
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
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "fieldName", GoGetter: "FieldName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineConfig", GoGetter: "PipelineConfig"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "requestMappingTemplate", GoGetter: "RequestMappingTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "requestMappingTemplateS3Location", GoGetter: "RequestMappingTemplateS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "responseMappingTemplate", GoGetter: "ResponseMappingTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "responseMappingTemplateS3Location", GoGetter: "ResponseMappingTemplateS3Location"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "syncConfig", GoGetter: "SyncConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
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
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.DataSourceOptions",
		reflect.TypeOf((*DataSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.Directive",
		reflect.TypeOf((*Directive)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "mutationFields", GoGetter: "MutationFields"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_Directive{}
		},
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
		"aws-cdk-lib.aws_appsync.EnumType",
		reflect.TypeOf((*EnumType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addField", GoMethod: "AddField"},
			_jsii_.MemberMethod{JsiiMethod: "attribute", GoMethod: "Attribute"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EnumType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIntermediateType)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.EnumTypeOptions",
		reflect.TypeOf((*EnumTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ExtendedDataSourceProps",
		reflect.TypeOf((*ExtendedDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ExtendedResolverProps",
		reflect.TypeOf((*ExtendedResolverProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.Field",
		reflect.TypeOf((*Field)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "argsToString", GoMethod: "ArgsToString"},
			_jsii_.MemberMethod{JsiiMethod: "directivesToString", GoMethod: "DirectivesToString"},
			_jsii_.MemberProperty{JsiiProperty: "fieldOptions", GoGetter: "FieldOptions"},
			_jsii_.MemberProperty{JsiiProperty: "intermediateType", GoGetter: "IntermediateType"},
			_jsii_.MemberProperty{JsiiProperty: "isList", GoGetter: "IsList"},
			_jsii_.MemberProperty{JsiiProperty: "isRequired", GoGetter: "IsRequired"},
			_jsii_.MemberProperty{JsiiProperty: "isRequiredList", GoGetter: "IsRequiredList"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_Field{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GraphqlType)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IField)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.FieldLogLevel",
		reflect.TypeOf((*FieldLogLevel)(nil)).Elem(),
		map[string]interface{}{
			"NONE": FieldLogLevel_NONE,
			"ERROR": FieldLogLevel_ERROR,
			"ALL": FieldLogLevel_ALL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.FieldOptions",
		reflect.TypeOf((*FieldOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.GraphqlApi",
		reflect.TypeOf((*GraphqlApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addMutation", GoMethod: "AddMutation"},
			_jsii_.MemberMethod{JsiiMethod: "addNoneDataSource", GoMethod: "AddNoneDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addQuery", GoMethod: "AddQuery"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSchemaDependency", GoMethod: "AddSchemaDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addSubscription", GoMethod: "AddSubscription"},
			_jsii_.MemberMethod{JsiiMethod: "addToSchema", GoMethod: "AddToSchema"},
			_jsii_.MemberMethod{JsiiMethod: "addType", GoMethod: "AddType"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "apiKey", GoGetter: "ApiKey"},
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
			_jsii_.MemberProperty{JsiiProperty: "graphqlUrl", GoGetter: "GraphqlUrl"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addNoneDataSource", GoMethod: "AddNoneDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSchemaDependency", GoMethod: "AddSchemaDependency"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
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
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.GraphqlType",
		reflect.TypeOf((*GraphqlType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "argsToString", GoMethod: "ArgsToString"},
			_jsii_.MemberMethod{JsiiMethod: "directivesToString", GoMethod: "DirectivesToString"},
			_jsii_.MemberProperty{JsiiProperty: "intermediateType", GoGetter: "IntermediateType"},
			_jsii_.MemberProperty{JsiiProperty: "isList", GoGetter: "IsList"},
			_jsii_.MemberProperty{JsiiProperty: "isRequired", GoGetter: "IsRequired"},
			_jsii_.MemberProperty{JsiiProperty: "isRequiredList", GoGetter: "IsRequiredList"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_GraphqlType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IField)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.GraphqlTypeOptions",
		reflect.TypeOf((*GraphqlTypeOptions)(nil)).Elem(),
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
		"aws-cdk-lib.aws_appsync.IAppsyncFunction",
		reflect.TypeOf((*IAppsyncFunction)(nil)).Elem(),
		[]_jsii_.Member{
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
		"aws-cdk-lib.aws_appsync.IField",
		reflect.TypeOf((*IField)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "argsToString", GoMethod: "ArgsToString"},
			_jsii_.MemberMethod{JsiiMethod: "directivesToString", GoMethod: "DirectivesToString"},
			_jsii_.MemberProperty{JsiiProperty: "fieldOptions", GoGetter: "FieldOptions"},
			_jsii_.MemberProperty{JsiiProperty: "intermediateType", GoGetter: "IntermediateType"},
			_jsii_.MemberProperty{JsiiProperty: "isList", GoGetter: "IsList"},
			_jsii_.MemberProperty{JsiiProperty: "isRequired", GoGetter: "IsRequired"},
			_jsii_.MemberProperty{JsiiProperty: "isRequiredList", GoGetter: "IsRequiredList"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IField{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.IGraphqlApi",
		reflect.TypeOf((*IGraphqlApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addNoneDataSource", GoMethod: "AddNoneDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSchemaDependency", GoMethod: "AddSchemaDependency"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "createResolver", GoMethod: "CreateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IGraphqlApi{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appsync.IIntermediateType",
		reflect.TypeOf((*IIntermediateType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addField", GoMethod: "AddField"},
			_jsii_.MemberMethod{JsiiMethod: "attribute", GoMethod: "Attribute"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "directives", GoGetter: "Directives"},
			_jsii_.MemberProperty{JsiiProperty: "interfaceTypes", GoGetter: "InterfaceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "intermediateType", GoGetter: "IntermediateType"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "resolvers", GoGetter: "Resolvers"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_IIntermediateType{}
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
		"aws-cdk-lib.aws_appsync.InputType",
		reflect.TypeOf((*InputType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addField", GoMethod: "AddField"},
			_jsii_.MemberMethod{JsiiMethod: "attribute", GoMethod: "Attribute"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_InputType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIntermediateType)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.InterfaceType",
		reflect.TypeOf((*InterfaceType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addField", GoMethod: "AddField"},
			_jsii_.MemberMethod{JsiiMethod: "attribute", GoMethod: "Attribute"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "directives", GoGetter: "Directives"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_InterfaceType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIntermediateType)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.IntermediateTypeOptions",
		reflect.TypeOf((*IntermediateTypeOptions)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.ObjectType",
		reflect.TypeOf((*ObjectType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addField", GoMethod: "AddField"},
			_jsii_.MemberMethod{JsiiMethod: "attribute", GoMethod: "Attribute"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "directives", GoGetter: "Directives"},
			_jsii_.MemberMethod{JsiiMethod: "generateResolver", GoMethod: "GenerateResolver"},
			_jsii_.MemberProperty{JsiiProperty: "interfaceTypes", GoGetter: "InterfaceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "resolvers", GoGetter: "Resolvers"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ObjectType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InterfaceType)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIntermediateType)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ObjectTypeOptions",
		reflect.TypeOf((*ObjectTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.OpenIdConnectConfig",
		reflect.TypeOf((*OpenIdConnectConfig)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.ResolvableField",
		reflect.TypeOf((*ResolvableField)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "argsToString", GoMethod: "ArgsToString"},
			_jsii_.MemberMethod{JsiiMethod: "directivesToString", GoMethod: "DirectivesToString"},
			_jsii_.MemberProperty{JsiiProperty: "fieldOptions", GoGetter: "FieldOptions"},
			_jsii_.MemberProperty{JsiiProperty: "intermediateType", GoGetter: "IntermediateType"},
			_jsii_.MemberProperty{JsiiProperty: "isList", GoGetter: "IsList"},
			_jsii_.MemberProperty{JsiiProperty: "isRequired", GoGetter: "IsRequired"},
			_jsii_.MemberProperty{JsiiProperty: "isRequiredList", GoGetter: "IsRequiredList"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_ResolvableField{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Field)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IField)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.ResolvableFieldOptions",
		reflect.TypeOf((*ResolvableFieldOptions)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.Schema",
		reflect.TypeOf((*Schema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMutation", GoMethod: "AddMutation"},
			_jsii_.MemberMethod{JsiiMethod: "addQuery", GoMethod: "AddQuery"},
			_jsii_.MemberMethod{JsiiMethod: "addSubscription", GoMethod: "AddSubscription"},
			_jsii_.MemberMethod{JsiiMethod: "addToSchema", GoMethod: "AddToSchema"},
			_jsii_.MemberMethod{JsiiMethod: "addType", GoMethod: "AddType"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
		},
		func() interface{} {
			return &jsiiProxy_Schema{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.SchemaOptions",
		reflect.TypeOf((*SchemaOptions)(nil)).Elem(),
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
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appsync.Type",
		reflect.TypeOf((*Type)(nil)).Elem(),
		map[string]interface{}{
			"ID": Type_ID,
			"STRING": Type_STRING,
			"INT": Type_INT,
			"FLOAT": Type_FLOAT,
			"BOOLEAN": Type_BOOLEAN,
			"AWS_DATE": Type_AWS_DATE,
			"AWS_TIME": Type_AWS_TIME,
			"AWS_DATE_TIME": Type_AWS_DATE_TIME,
			"AWS_TIMESTAMP": Type_AWS_TIMESTAMP,
			"AWS_EMAIL": Type_AWS_EMAIL,
			"AWS_JSON": Type_AWS_JSON,
			"AWS_URL": Type_AWS_URL,
			"AWS_PHONE": Type_AWS_PHONE,
			"AWS_IP_ADDRESS": Type_AWS_IP_ADDRESS,
			"INTERMEDIATE": Type_INTERMEDIATE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appsync.UnionType",
		reflect.TypeOf((*UnionType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addField", GoMethod: "AddField"},
			_jsii_.MemberMethod{JsiiMethod: "attribute", GoMethod: "Attribute"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_UnionType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIntermediateType)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appsync.UnionTypeOptions",
		reflect.TypeOf((*UnionTypeOptions)(nil)).Elem(),
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
}
