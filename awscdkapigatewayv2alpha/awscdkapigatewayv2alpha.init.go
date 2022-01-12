package awscdkapigatewayv2alpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.AddRoutesOptions",
		reflect.TypeOf((*AddRoutesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.ApiMapping",
		reflect.TypeOf((*ApiMapping)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiMappingId", GoGetter: "ApiMappingId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mappingKey", GoGetter: "MappingKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApiMapping{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApiMapping)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.ApiMappingAttributes",
		reflect.TypeOf((*ApiMappingAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.ApiMappingProps",
		reflect.TypeOf((*ApiMappingProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-alpha.AuthorizerPayloadVersion",
		reflect.TypeOf((*AuthorizerPayloadVersion)(nil)).Elem(),
		map[string]interface{}{
			"VERSION_1_0": AuthorizerPayloadVersion_VERSION_1_0,
			"VERSION_2_0": AuthorizerPayloadVersion_VERSION_2_0,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.BatchHttpRouteOptions",
		reflect.TypeOf((*BatchHttpRouteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-alpha.CorsHttpMethod",
		reflect.TypeOf((*CorsHttpMethod)(nil)).Elem(),
		map[string]interface{}{
			"ANY": CorsHttpMethod_ANY,
			"DELETE": CorsHttpMethod_DELETE,
			"GET": CorsHttpMethod_GET,
			"HEAD": CorsHttpMethod_HEAD,
			"OPTIONS": CorsHttpMethod_OPTIONS,
			"PATCH": CorsHttpMethod_PATCH,
			"POST": CorsHttpMethod_POST,
			"PUT": CorsHttpMethod_PUT,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.CorsPreflightOptions",
		reflect.TypeOf((*CorsPreflightOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.DomainMappingOptions",
		reflect.TypeOf((*DomainMappingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.DomainName",
		reflect.TypeOf((*DomainName)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEndpoint", GoMethod: "AddEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "regionalDomainName", GoGetter: "RegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "regionalHostedZoneId", GoGetter: "RegionalHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DomainName{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDomainName)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.DomainNameAttributes",
		reflect.TypeOf((*DomainNameAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.DomainNameProps",
		reflect.TypeOf((*DomainNameProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.EndpointOptions",
		reflect.TypeOf((*EndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-alpha.EndpointType",
		reflect.TypeOf((*EndpointType)(nil)).Elem(),
		map[string]interface{}{
			"EDGE": EndpointType_EDGE,
			"REGIONAL": EndpointType_REGIONAL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.GrantInvokeOptions",
		reflect.TypeOf((*GrantInvokeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		reflect.TypeOf((*HttpApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoutes", GoMethod: "AddRoutes"},
			_jsii_.MemberMethod{JsiiMethod: "addStage", GoMethod: "AddStage"},
			_jsii_.MemberMethod{JsiiMethod: "addVpcLink", GoMethod: "AddVpcLink"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "defaultStage", GoGetter: "DefaultStage"},
			_jsii_.MemberProperty{JsiiProperty: "disableExecuteApiEndpoint", GoGetter: "DisableExecuteApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpApiId", GoGetter: "HttpApiId"},
			_jsii_.MemberProperty{JsiiProperty: "httpApiName", GoGetter: "HttpApiName"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricClientError", GoMethod: "MetricClientError"},
			_jsii_.MemberMethod{JsiiMethod: "metricCount", GoMethod: "MetricCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDataProcessed", GoMethod: "MetricDataProcessed"},
			_jsii_.MemberMethod{JsiiMethod: "metricIntegrationLatency", GoMethod: "MetricIntegrationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricServerError", GoMethod: "MetricServerError"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_HttpApi{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApi)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IHttpApi)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApiAttributes",
		reflect.TypeOf((*HttpApiAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApiProps",
		reflect.TypeOf((*HttpApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpAuthorizer",
		reflect.TypeOf((*HttpAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
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
			j := jsiiProxy_HttpAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IHttpAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpAuthorizerAttributes",
		reflect.TypeOf((*HttpAuthorizerAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpAuthorizerProps",
		reflect.TypeOf((*HttpAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpAuthorizerType",
		reflect.TypeOf((*HttpAuthorizerType)(nil)).Elem(),
		map[string]interface{}{
			"IAM": HttpAuthorizerType_IAM,
			"JWT": HttpAuthorizerType_JWT,
			"LAMBDA": HttpAuthorizerType_LAMBDA,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpConnectionType",
		reflect.TypeOf((*HttpConnectionType)(nil)).Elem(),
		map[string]interface{}{
			"VPC_LINK": HttpConnectionType_VPC_LINK,
			"INTERNET": HttpConnectionType_INTERNET,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpIntegration",
		reflect.TypeOf((*HttpIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpApi", GoGetter: "HttpApi"},
			_jsii_.MemberProperty{JsiiProperty: "integrationId", GoGetter: "IntegrationId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HttpIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IHttpIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpIntegrationProps",
		reflect.TypeOf((*HttpIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpIntegrationType",
		reflect.TypeOf((*HttpIntegrationType)(nil)).Elem(),
		map[string]interface{}{
			"LAMBDA_PROXY": HttpIntegrationType_LAMBDA_PROXY,
			"HTTP_PROXY": HttpIntegrationType_HTTP_PROXY,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpMethod",
		reflect.TypeOf((*HttpMethod)(nil)).Elem(),
		map[string]interface{}{
			"ANY": HttpMethod_ANY,
			"DELETE": HttpMethod_DELETE,
			"GET": HttpMethod_GET,
			"HEAD": HttpMethod_HEAD,
			"OPTIONS": HttpMethod_OPTIONS,
			"PATCH": HttpMethod_PATCH,
			"POST": HttpMethod_POST,
			"PUT": HttpMethod_PUT,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpNoneAuthorizer",
		reflect.TypeOf((*HttpNoneAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpNoneAuthorizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRoute",
		reflect.TypeOf((*HttpRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "httpApi", GoGetter: "HttpApi"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routeArn", GoGetter: "RouteArn"},
			_jsii_.MemberProperty{JsiiProperty: "routeId", GoGetter: "RouteId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HttpRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IHttpRoute)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteAuthorizerBindOptions",
		reflect.TypeOf((*HttpRouteAuthorizerBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteAuthorizerConfig",
		reflect.TypeOf((*HttpRouteAuthorizerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteIntegration",
		reflect.TypeOf((*HttpRouteIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_HttpRouteIntegration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteIntegrationBindOptions",
		reflect.TypeOf((*HttpRouteIntegrationBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteIntegrationConfig",
		reflect.TypeOf((*HttpRouteIntegrationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteKey",
		reflect.TypeOf((*HttpRouteKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
		},
		func() interface{} {
			return &jsiiProxy_HttpRouteKey{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteProps",
		reflect.TypeOf((*HttpRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpStage",
		reflect.TypeOf((*HttpStage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "baseApi", GoGetter: "BaseApi"},
			_jsii_.MemberProperty{JsiiProperty: "domainUrl", GoGetter: "DomainUrl"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricClientError", GoMethod: "MetricClientError"},
			_jsii_.MemberMethod{JsiiMethod: "metricCount", GoMethod: "MetricCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDataProcessed", GoMethod: "MetricDataProcessed"},
			_jsii_.MemberMethod{JsiiMethod: "metricIntegrationLatency", GoMethod: "MetricIntegrationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricServerError", GoMethod: "MetricServerError"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_HttpStage{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IHttpStage)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpStageAttributes",
		reflect.TypeOf((*HttpStageAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpStageOptions",
		reflect.TypeOf((*HttpStageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpStageProps",
		reflect.TypeOf((*HttpStageProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IApi",
		reflect.TypeOf((*IApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
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
		"@aws-cdk/aws-apigatewayv2-alpha.IApiMapping",
		reflect.TypeOf((*IApiMapping)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiMappingId", GoGetter: "ApiMappingId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IApiMapping{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IAuthorizer",
		reflect.TypeOf((*IAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IDomainName",
		reflect.TypeOf((*IDomainName)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "regionalDomainName", GoGetter: "RegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "regionalHostedZoneId", GoGetter: "RegionalHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDomainName{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IHttpApi",
		reflect.TypeOf((*IHttpApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVpcLink", GoMethod: "AddVpcLink"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "httpApiId", GoGetter: "HttpApiId"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricClientError", GoMethod: "MetricClientError"},
			_jsii_.MemberMethod{JsiiMethod: "metricCount", GoMethod: "MetricCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDataProcessed", GoMethod: "MetricDataProcessed"},
			_jsii_.MemberMethod{JsiiMethod: "metricIntegrationLatency", GoMethod: "MetricIntegrationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricServerError", GoMethod: "MetricServerError"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IHttpApi{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApi)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IHttpAuthorizer",
		reflect.TypeOf((*IHttpAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IHttpAuthorizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IHttpIntegration",
		reflect.TypeOf((*IHttpIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "httpApi", GoGetter: "HttpApi"},
			_jsii_.MemberProperty{JsiiProperty: "integrationId", GoGetter: "IntegrationId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IHttpIntegration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIntegration)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IHttpRoute",
		reflect.TypeOf((*IHttpRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "httpApi", GoGetter: "HttpApi"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "routeArn", GoGetter: "RouteArn"},
			_jsii_.MemberProperty{JsiiProperty: "routeId", GoGetter: "RouteId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IHttpRoute{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRoute)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IHttpRouteAuthorizer",
		reflect.TypeOf((*IHttpRouteAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IHttpRouteAuthorizer{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IHttpStage",
		reflect.TypeOf((*IHttpStage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "domainUrl", GoGetter: "DomainUrl"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricClientError", GoMethod: "MetricClientError"},
			_jsii_.MemberMethod{JsiiMethod: "metricCount", GoMethod: "MetricCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDataProcessed", GoMethod: "MetricDataProcessed"},
			_jsii_.MemberMethod{JsiiMethod: "metricIntegrationLatency", GoMethod: "MetricIntegrationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricServerError", GoMethod: "MetricServerError"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_IHttpStage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStage)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IIntegration",
		reflect.TypeOf((*IIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "integrationId", GoGetter: "IntegrationId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IMappingValue",
		reflect.TypeOf((*IMappingValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IMappingValue{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IRoute",
		reflect.TypeOf((*IRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routeId", GoGetter: "RouteId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IStage",
		reflect.TypeOf((*IStage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_IStage{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IVpcLink",
		reflect.TypeOf((*IVpcLink)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "vpcLinkId", GoGetter: "VpcLinkId"},
		},
		func() interface{} {
			j := jsiiProxy_IVpcLink{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IWebSocketApi",
		reflect.TypeOf((*IWebSocketApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IWebSocketApi{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApi)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IWebSocketAuthorizer",
		reflect.TypeOf((*IWebSocketAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IWebSocketAuthorizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IWebSocketIntegration",
		reflect.TypeOf((*IWebSocketIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "integrationId", GoGetter: "IntegrationId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "webSocketApi", GoGetter: "WebSocketApi"},
		},
		func() interface{} {
			j := jsiiProxy_IWebSocketIntegration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIntegration)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IWebSocketRoute",
		reflect.TypeOf((*IWebSocketRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routeId", GoGetter: "RouteId"},
			_jsii_.MemberProperty{JsiiProperty: "routeKey", GoGetter: "RouteKey"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "webSocketApi", GoGetter: "WebSocketApi"},
		},
		func() interface{} {
			j := jsiiProxy_IWebSocketRoute{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRoute)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IWebSocketRouteAuthorizer",
		reflect.TypeOf((*IWebSocketRouteAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IWebSocketRouteAuthorizer{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apigatewayv2-alpha.IWebSocketStage",
		reflect.TypeOf((*IWebSocketStage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "callbackUrl", GoGetter: "CallbackUrl"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_IWebSocketStage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.MTLSConfig",
		reflect.TypeOf((*MTLSConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		reflect.TypeOf((*MappingValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_MappingValue{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMappingValue)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.ParameterMapping",
		reflect.TypeOf((*ParameterMapping)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "appendHeader", GoMethod: "AppendHeader"},
			_jsii_.MemberMethod{JsiiMethod: "appendQueryString", GoMethod: "AppendQueryString"},
			_jsii_.MemberMethod{JsiiMethod: "custom", GoMethod: "Custom"},
			_jsii_.MemberProperty{JsiiProperty: "mappings", GoGetter: "Mappings"},
			_jsii_.MemberMethod{JsiiMethod: "overwriteHeader", GoMethod: "OverwriteHeader"},
			_jsii_.MemberMethod{JsiiMethod: "overwritePath", GoMethod: "OverwritePath"},
			_jsii_.MemberMethod{JsiiMethod: "overwriteQueryString", GoMethod: "OverwriteQueryString"},
			_jsii_.MemberMethod{JsiiMethod: "removeHeader", GoMethod: "RemoveHeader"},
			_jsii_.MemberMethod{JsiiMethod: "removeQueryString", GoMethod: "RemoveQueryString"},
		},
		func() interface{} {
			return &jsiiProxy_ParameterMapping{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.PayloadFormatVersion",
		reflect.TypeOf((*PayloadFormatVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_PayloadFormatVersion{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-alpha.SecurityPolicy",
		reflect.TypeOf((*SecurityPolicy)(nil)).Elem(),
		map[string]interface{}{
			"TLS_1_0": SecurityPolicy_TLS_1_0,
			"TLS_1_2": SecurityPolicy_TLS_1_2,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.StageAttributes",
		reflect.TypeOf((*StageAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.StageOptions",
		reflect.TypeOf((*StageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.VpcLink",
		reflect.TypeOf((*VpcLink)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSecurityGroups", GoMethod: "AddSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addSubnets", GoMethod: "AddSubnets"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "vpcLinkId", GoGetter: "VpcLinkId"},
		},
		func() interface{} {
			j := jsiiProxy_VpcLink{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpcLink)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.VpcLinkAttributes",
		reflect.TypeOf((*VpcLinkAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.VpcLinkProps",
		reflect.TypeOf((*VpcLinkProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketApi",
		reflect.TypeOf((*WebSocketApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantManageConnections", GoMethod: "GrantManageConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webSocketApiName", GoGetter: "WebSocketApiName"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketApi{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApi)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWebSocketApi)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketApiKeySelectionExpression",
		reflect.TypeOf((*WebSocketApiKeySelectionExpression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "customApiKeySelector", GoGetter: "CustomApiKeySelector"},
		},
		func() interface{} {
			return &jsiiProxy_WebSocketApiKeySelectionExpression{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketApiProps",
		reflect.TypeOf((*WebSocketApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketAuthorizer",
		reflect.TypeOf((*WebSocketAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
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
			j := jsiiProxy_WebSocketAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWebSocketAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketAuthorizerAttributes",
		reflect.TypeOf((*WebSocketAuthorizerAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketAuthorizerProps",
		reflect.TypeOf((*WebSocketAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketAuthorizerType",
		reflect.TypeOf((*WebSocketAuthorizerType)(nil)).Elem(),
		map[string]interface{}{
			"LAMBDA": WebSocketAuthorizerType_LAMBDA,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketIntegration",
		reflect.TypeOf((*WebSocketIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "integrationId", GoGetter: "IntegrationId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webSocketApi", GoGetter: "WebSocketApi"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWebSocketIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketIntegrationProps",
		reflect.TypeOf((*WebSocketIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketIntegrationType",
		reflect.TypeOf((*WebSocketIntegrationType)(nil)).Elem(),
		map[string]interface{}{
			"AWS_PROXY": WebSocketIntegrationType_AWS_PROXY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketNoneAuthorizer",
		reflect.TypeOf((*WebSocketNoneAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketNoneAuthorizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWebSocketRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRoute",
		reflect.TypeOf((*WebSocketRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "integrationResponseId", GoGetter: "IntegrationResponseId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routeId", GoGetter: "RouteId"},
			_jsii_.MemberProperty{JsiiProperty: "routeKey", GoGetter: "RouteKey"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webSocketApi", GoGetter: "WebSocketApi"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWebSocketRoute)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRouteAuthorizerBindOptions",
		reflect.TypeOf((*WebSocketRouteAuthorizerBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRouteAuthorizerConfig",
		reflect.TypeOf((*WebSocketRouteAuthorizerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRouteIntegration",
		reflect.TypeOf((*WebSocketRouteIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_WebSocketRouteIntegration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRouteIntegrationBindOptions",
		reflect.TypeOf((*WebSocketRouteIntegrationBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRouteIntegrationConfig",
		reflect.TypeOf((*WebSocketRouteIntegrationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRouteOptions",
		reflect.TypeOf((*WebSocketRouteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRouteProps",
		reflect.TypeOf((*WebSocketRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		reflect.TypeOf((*WebSocketStage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "baseApi", GoGetter: "BaseApi"},
			_jsii_.MemberProperty{JsiiProperty: "callbackUrl", GoGetter: "CallbackUrl"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantManagementApiAccess", GoMethod: "GrantManagementApiAccess"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketStage{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStage)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWebSocketStage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStageAttributes",
		reflect.TypeOf((*WebSocketStageAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStageProps",
		reflect.TypeOf((*WebSocketStageProps)(nil)).Elem(),
	)
}
