package awscdkappsyncalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.AddFieldOptions",
		reflect.TypeOf((*AddFieldOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.ApiKeyConfig",
		reflect.TypeOf((*ApiKeyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.AppsyncFunction",
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
		"@aws-cdk/aws-appsync-alpha.AppsyncFunctionAttributes",
		reflect.TypeOf((*AppsyncFunctionAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.AppsyncFunctionProps",
		reflect.TypeOf((*AppsyncFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.Assign",
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
		"@aws-cdk/aws-appsync-alpha.AttributeValues",
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
		"@aws-cdk/aws-appsync-alpha.AttributeValuesStep",
		reflect.TypeOf((*AttributeValuesStep)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "is", GoMethod: "Is"},
		},
		func() interface{} {
			return &jsiiProxy_AttributeValuesStep{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.AuthorizationConfig",
		reflect.TypeOf((*AuthorizationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.AuthorizationMode",
		reflect.TypeOf((*AuthorizationMode)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-appsync-alpha.AuthorizationType",
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
		"@aws-cdk/aws-appsync-alpha.AwsIamConfig",
		reflect.TypeOf((*AwsIamConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.BackedDataSource",
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
		"@aws-cdk/aws-appsync-alpha.BackedDataSourceProps",
		reflect.TypeOf((*BackedDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.BaseAppsyncFunctionProps",
		reflect.TypeOf((*BaseAppsyncFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.BaseDataSource",
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
		"@aws-cdk/aws-appsync-alpha.BaseDataSourceProps",
		reflect.TypeOf((*BaseDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.BaseResolverProps",
		reflect.TypeOf((*BaseResolverProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.BaseTypeOptions",
		reflect.TypeOf((*BaseTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.CachingConfig",
		reflect.TypeOf((*CachingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.DataSourceOptions",
		reflect.TypeOf((*DataSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.Directive",
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
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.DomainOptions",
		reflect.TypeOf((*DomainOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.DynamoDbDataSource",
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
		"@aws-cdk/aws-appsync-alpha.DynamoDbDataSourceProps",
		reflect.TypeOf((*DynamoDbDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.ElasticsearchDataSource",
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
		"@aws-cdk/aws-appsync-alpha.ElasticsearchDataSourceProps",
		reflect.TypeOf((*ElasticsearchDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.EnumType",
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
		"@aws-cdk/aws-appsync-alpha.EnumTypeOptions",
		reflect.TypeOf((*EnumTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.ExtendedDataSourceProps",
		reflect.TypeOf((*ExtendedDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.ExtendedResolverProps",
		reflect.TypeOf((*ExtendedResolverProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.Field",
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
		"@aws-cdk/aws-appsync-alpha.FieldLogLevel",
		reflect.TypeOf((*FieldLogLevel)(nil)).Elem(),
		map[string]interface{}{
			"NONE": FieldLogLevel_NONE,
			"ERROR": FieldLogLevel_ERROR,
			"ALL": FieldLogLevel_ALL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.FieldOptions",
		reflect.TypeOf((*FieldOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.GraphqlApi",
		reflect.TypeOf((*GraphqlApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addElasticsearchDataSource", GoMethod: "AddElasticsearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addMutation", GoMethod: "AddMutation"},
			_jsii_.MemberMethod{JsiiMethod: "addNoneDataSource", GoMethod: "AddNoneDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addOpenSearchDataSource", GoMethod: "AddOpenSearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addQuery", GoMethod: "AddQuery"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSchemaDependency", GoMethod: "AddSchemaDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addSubscription", GoMethod: "AddSubscription"},
			_jsii_.MemberMethod{JsiiMethod: "addToSchema", GoMethod: "AddToSchema"},
			_jsii_.MemberMethod{JsiiMethod: "addType", GoMethod: "AddType"},
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
			_jsii_.MemberProperty{JsiiProperty: "graphqlUrl", GoGetter: "GraphqlUrl"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
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
		"@aws-cdk/aws-appsync-alpha.GraphqlApiAttributes",
		reflect.TypeOf((*GraphqlApiAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.GraphqlApiBase",
		reflect.TypeOf((*GraphqlApiBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addElasticsearchDataSource", GoMethod: "AddElasticsearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addNoneDataSource", GoMethod: "AddNoneDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addOpenSearchDataSource", GoMethod: "AddOpenSearchDataSource"},
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
		"@aws-cdk/aws-appsync-alpha.GraphqlApiProps",
		reflect.TypeOf((*GraphqlApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.GraphqlType",
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
		"@aws-cdk/aws-appsync-alpha.GraphqlTypeOptions",
		reflect.TypeOf((*GraphqlTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.HttpDataSource",
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
		"@aws-cdk/aws-appsync-alpha.HttpDataSourceOptions",
		reflect.TypeOf((*HttpDataSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.HttpDataSourceProps",
		reflect.TypeOf((*HttpDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-appsync-alpha.IAppsyncFunction",
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
		"@aws-cdk/aws-appsync-alpha.IField",
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
		"@aws-cdk/aws-appsync-alpha.IGraphqlApi",
		reflect.TypeOf((*IGraphqlApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDynamoDbDataSource", GoMethod: "AddDynamoDbDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addElasticsearchDataSource", GoMethod: "AddElasticsearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addHttpDataSource", GoMethod: "AddHttpDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaDataSource", GoMethod: "AddLambdaDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addNoneDataSource", GoMethod: "AddNoneDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addOpenSearchDataSource", GoMethod: "AddOpenSearchDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addRdsDataSource", GoMethod: "AddRdsDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSchemaDependency", GoMethod: "AddSchemaDependency"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
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
		"@aws-cdk/aws-appsync-alpha.IIntermediateType",
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
		"@aws-cdk/aws-appsync-alpha.IamResource",
		reflect.TypeOf((*IamResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resourceArns", GoMethod: "ResourceArns"},
		},
		func() interface{} {
			return &jsiiProxy_IamResource{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.InputType",
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
		"@aws-cdk/aws-appsync-alpha.InterfaceType",
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
		"@aws-cdk/aws-appsync-alpha.IntermediateTypeOptions",
		reflect.TypeOf((*IntermediateTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.KeyCondition",
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
		"@aws-cdk/aws-appsync-alpha.LambdaAuthorizerConfig",
		reflect.TypeOf((*LambdaAuthorizerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.LambdaDataSource",
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
		"@aws-cdk/aws-appsync-alpha.LambdaDataSourceProps",
		reflect.TypeOf((*LambdaDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.LogConfig",
		reflect.TypeOf((*LogConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		reflect.TypeOf((*MappingTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "renderTemplate", GoMethod: "RenderTemplate"},
		},
		func() interface{} {
			return &jsiiProxy_MappingTemplate{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.NoneDataSource",
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
		"@aws-cdk/aws-appsync-alpha.NoneDataSourceProps",
		reflect.TypeOf((*NoneDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.ObjectType",
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
		"@aws-cdk/aws-appsync-alpha.ObjectTypeOptions",
		reflect.TypeOf((*ObjectTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.OpenIdConnectConfig",
		reflect.TypeOf((*OpenIdConnectConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.OpenSearchDataSource",
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
		"@aws-cdk/aws-appsync-alpha.OpenSearchDataSourceProps",
		reflect.TypeOf((*OpenSearchDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.PartitionKey",
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
		"@aws-cdk/aws-appsync-alpha.PartitionKeyStep",
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
		"@aws-cdk/aws-appsync-alpha.PrimaryKey",
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
		"@aws-cdk/aws-appsync-alpha.RdsDataSource",
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
		"@aws-cdk/aws-appsync-alpha.RdsDataSourceProps",
		reflect.TypeOf((*RdsDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.ResolvableField",
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
		"@aws-cdk/aws-appsync-alpha.ResolvableFieldOptions",
		reflect.TypeOf((*ResolvableFieldOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.Resolver",
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
		"@aws-cdk/aws-appsync-alpha.ResolverProps",
		reflect.TypeOf((*ResolverProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.Schema",
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
		"@aws-cdk/aws-appsync-alpha.SchemaOptions",
		reflect.TypeOf((*SchemaOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.SortKeyStep",
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
		"@aws-cdk/aws-appsync-alpha.Type",
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
		"@aws-cdk/aws-appsync-alpha.UnionType",
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
		"@aws-cdk/aws-appsync-alpha.UnionTypeOptions",
		reflect.TypeOf((*UnionTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appsync-alpha.UserPoolConfig",
		reflect.TypeOf((*UserPoolConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-appsync-alpha.UserPoolDefaultAction",
		reflect.TypeOf((*UserPoolDefaultAction)(nil)).Elem(),
		map[string]interface{}{
			"ALLOW": UserPoolDefaultAction_ALLOW,
			"DENY": UserPoolDefaultAction_DENY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appsync-alpha.Values",
		reflect.TypeOf((*Values)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Values{}
		},
	)
}
