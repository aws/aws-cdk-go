package awsapigatewayv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.AddRoutesOptions",
		reflect.TypeOf((*AddRoutesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.ApiMapping",
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
		"aws-cdk-lib.aws_apigatewayv2.ApiMappingAttributes",
		reflect.TypeOf((*ApiMappingAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.ApiMappingProps",
		reflect.TypeOf((*ApiMappingProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.AuthorizerPayloadVersion",
		reflect.TypeOf((*AuthorizerPayloadVersion)(nil)).Elem(),
		map[string]interface{}{
			"VERSION_1_0": AuthorizerPayloadVersion_VERSION_1_0,
			"VERSION_2_0": AuthorizerPayloadVersion_VERSION_2_0,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.BatchHttpRouteOptions",
		reflect.TypeOf((*BatchHttpRouteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnApi",
		reflect.TypeOf((*CfnApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeySelectionExpression", GoGetter: "ApiKeySelectionExpression"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiEndpoint", GoGetter: "AttrApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiId", GoGetter: "AttrApiId"},
			_jsii_.MemberProperty{JsiiProperty: "basePath", GoGetter: "BasePath"},
			_jsii_.MemberProperty{JsiiProperty: "body", GoGetter: "Body"},
			_jsii_.MemberProperty{JsiiProperty: "bodyS3Location", GoGetter: "BodyS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "corsConfiguration", GoGetter: "CorsConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsArn", GoGetter: "CredentialsArn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "disableExecuteApiEndpoint", GoGetter: "DisableExecuteApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "disableSchemaValidation", GoGetter: "DisableSchemaValidation"},
			_jsii_.MemberProperty{JsiiProperty: "failOnWarnings", GoGetter: "FailOnWarnings"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "protocolType", GoGetter: "ProtocolType"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "routeKey", GoGetter: "RouteKey"},
			_jsii_.MemberProperty{JsiiProperty: "routeSelectionExpression", GoGetter: "RouteSelectionExpression"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApi{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnApi.BodyS3LocationProperty",
		reflect.TypeOf((*CfnApi_BodyS3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnApi.CorsProperty",
		reflect.TypeOf((*CfnApi_CorsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides",
		reflect.TypeOf((*CfnApiGatewayManagedOverrides)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "integration", GoGetter: "Integration"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "route", GoGetter: "Route"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApiGatewayManagedOverrides{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides.AccessLogSettingsProperty",
		reflect.TypeOf((*CfnApiGatewayManagedOverrides_AccessLogSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides.IntegrationOverridesProperty",
		reflect.TypeOf((*CfnApiGatewayManagedOverrides_IntegrationOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides.RouteOverridesProperty",
		reflect.TypeOf((*CfnApiGatewayManagedOverrides_RouteOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides.RouteSettingsProperty",
		reflect.TypeOf((*CfnApiGatewayManagedOverrides_RouteSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides.StageOverridesProperty",
		reflect.TypeOf((*CfnApiGatewayManagedOverrides_StageOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverridesProps",
		reflect.TypeOf((*CfnApiGatewayManagedOverridesProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiMapping",
		reflect.TypeOf((*CfnApiMapping)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "apiMappingKey", GoGetter: "ApiMappingKey"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrApiMappingId", GoGetter: "AttrApiMappingId"},
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
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApiMapping{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiMappingProps",
		reflect.TypeOf((*CfnApiMappingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiProps",
		reflect.TypeOf((*CfnApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnAuthorizer",
		reflect.TypeOf((*CfnAuthorizer)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrAuthorizerId", GoGetter: "AttrAuthorizerId"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerCredentialsArn", GoGetter: "AuthorizerCredentialsArn"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerPayloadFormatVersion", GoGetter: "AuthorizerPayloadFormatVersion"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerResultTtlInSeconds", GoGetter: "AuthorizerResultTtlInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerType", GoGetter: "AuthorizerType"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerUri", GoGetter: "AuthorizerUri"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableSimpleResponses", GoGetter: "EnableSimpleResponses"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "identitySource", GoGetter: "IdentitySource"},
			_jsii_.MemberProperty{JsiiProperty: "identityValidationExpression", GoGetter: "IdentityValidationExpression"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "jwtConfiguration", GoGetter: "JwtConfiguration"},
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
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnAuthorizer.JWTConfigurationProperty",
		reflect.TypeOf((*CfnAuthorizer_JWTConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnAuthorizerProps",
		reflect.TypeOf((*CfnAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnDeployment",
		reflect.TypeOf((*CfnDeployment)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrDeploymentId", GoGetter: "AttrDeploymentId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
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
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeployment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnDeploymentProps",
		reflect.TypeOf((*CfnDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnDomainName",
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
			_jsii_.MemberProperty{JsiiProperty: "attrDomainNameArn", GoGetter: "AttrDomainNameArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrRegionalDomainName", GoGetter: "AttrRegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "attrRegionalHostedZoneId", GoGetter: "AttrRegionalHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameConfigurations", GoGetter: "DomainNameConfigurations"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "mutualTlsAuthentication", GoGetter: "MutualTlsAuthentication"},
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
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainName{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnDomainName.DomainNameConfigurationProperty",
		reflect.TypeOf((*CfnDomainName_DomainNameConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnDomainName.MutualTlsAuthenticationProperty",
		reflect.TypeOf((*CfnDomainName_MutualTlsAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnDomainNameProps",
		reflect.TypeOf((*CfnDomainNameProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
		reflect.TypeOf((*CfnIntegration)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrIntegrationId", GoGetter: "AttrIntegrationId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionId", GoGetter: "ConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "contentHandlingStrategy", GoGetter: "ContentHandlingStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsArn", GoGetter: "CredentialsArn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "integrationMethod", GoGetter: "IntegrationMethod"},
			_jsii_.MemberProperty{JsiiProperty: "integrationSubtype", GoGetter: "IntegrationSubtype"},
			_jsii_.MemberProperty{JsiiProperty: "integrationType", GoGetter: "IntegrationType"},
			_jsii_.MemberProperty{JsiiProperty: "integrationUri", GoGetter: "IntegrationUri"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughBehavior", GoGetter: "PassthroughBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersion", GoGetter: "PayloadFormatVersion"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "requestParameters", GoGetter: "RequestParameters"},
			_jsii_.MemberProperty{JsiiProperty: "requestTemplates", GoGetter: "RequestTemplates"},
			_jsii_.MemberProperty{JsiiProperty: "responseParameters", GoGetter: "ResponseParameters"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "templateSelectionExpression", GoGetter: "TemplateSelectionExpression"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInMillis", GoGetter: "TimeoutInMillis"},
			_jsii_.MemberProperty{JsiiProperty: "tlsConfig", GoGetter: "TlsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration.ResponseParameterListProperty",
		reflect.TypeOf((*CfnIntegration_ResponseParameterListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration.ResponseParameterMapProperty",
		reflect.TypeOf((*CfnIntegration_ResponseParameterMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration.ResponseParameterProperty",
		reflect.TypeOf((*CfnIntegration_ResponseParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration.TlsConfigProperty",
		reflect.TypeOf((*CfnIntegration_TlsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegrationProps",
		reflect.TypeOf((*CfnIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegrationResponse",
		reflect.TypeOf((*CfnIntegrationResponse)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrIntegrationResponseId", GoGetter: "AttrIntegrationResponseId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "contentHandlingStrategy", GoGetter: "ContentHandlingStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "integrationId", GoGetter: "IntegrationId"},
			_jsii_.MemberProperty{JsiiProperty: "integrationResponseKey", GoGetter: "IntegrationResponseKey"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "responseParameters", GoGetter: "ResponseParameters"},
			_jsii_.MemberProperty{JsiiProperty: "responseTemplates", GoGetter: "ResponseTemplates"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "templateSelectionExpression", GoGetter: "TemplateSelectionExpression"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIntegrationResponse{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegrationResponseProps",
		reflect.TypeOf((*CfnIntegrationResponseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnModel",
		reflect.TypeOf((*CfnModel)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrModelId", GoGetter: "AttrModelId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "contentType", GoGetter: "ContentType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
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
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModel{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnModelProps",
		reflect.TypeOf((*CfnModelProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnRoute",
		reflect.TypeOf((*CfnRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyRequired", GoGetter: "ApiKeyRequired"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrRouteId", GoGetter: "AttrRouteId"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationScopes", GoGetter: "AuthorizationScopes"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationType", GoGetter: "AuthorizationType"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "modelSelectionExpression", GoGetter: "ModelSelectionExpression"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "requestModels", GoGetter: "RequestModels"},
			_jsii_.MemberProperty{JsiiProperty: "requestParameters", GoGetter: "RequestParameters"},
			_jsii_.MemberProperty{JsiiProperty: "routeKey", GoGetter: "RouteKey"},
			_jsii_.MemberProperty{JsiiProperty: "routeResponseSelectionExpression", GoGetter: "RouteResponseSelectionExpression"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnRoute.ParameterConstraintsProperty",
		reflect.TypeOf((*CfnRoute_ParameterConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnRouteProps",
		reflect.TypeOf((*CfnRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnRouteResponse",
		reflect.TypeOf((*CfnRouteResponse)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrRouteResponseId", GoGetter: "AttrRouteResponseId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "modelSelectionExpression", GoGetter: "ModelSelectionExpression"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "responseModels", GoGetter: "ResponseModels"},
			_jsii_.MemberProperty{JsiiProperty: "responseParameters", GoGetter: "ResponseParameters"},
			_jsii_.MemberProperty{JsiiProperty: "routeId", GoGetter: "RouteId"},
			_jsii_.MemberProperty{JsiiProperty: "routeResponseKey", GoGetter: "RouteResponseKey"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRouteResponse{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnRouteResponse.ParameterConstraintsProperty",
		reflect.TypeOf((*CfnRouteResponse_ParameterConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnRouteResponseProps",
		reflect.TypeOf((*CfnRouteResponseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnStage",
		reflect.TypeOf((*CfnStage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessLogSettings", GoGetter: "AccessLogSettings"},
			_jsii_.MemberProperty{JsiiProperty: "accessPolicyId", GoGetter: "AccessPolicyId"},
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
			_jsii_.MemberProperty{JsiiProperty: "autoDeploy", GoGetter: "AutoDeploy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificateId", GoGetter: "ClientCertificateId"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteSettings", GoGetter: "DefaultRouteSettings"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentId", GoGetter: "DeploymentId"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
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
			_jsii_.MemberProperty{JsiiProperty: "routeSettings", GoGetter: "RouteSettings"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberProperty{JsiiProperty: "stageVariables", GoGetter: "StageVariables"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStage{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnStage.AccessLogSettingsProperty",
		reflect.TypeOf((*CfnStage_AccessLogSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnStage.RouteSettingsProperty",
		reflect.TypeOf((*CfnStage_RouteSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnStageProps",
		reflect.TypeOf((*CfnStageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.CfnVpcLink",
		reflect.TypeOf((*CfnVpcLink)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrVpcLinkId", GoGetter: "AttrVpcLinkId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcLink{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.CfnVpcLinkProps",
		reflect.TypeOf((*CfnVpcLinkProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.ContentHandling",
		reflect.TypeOf((*ContentHandling)(nil)).Elem(),
		map[string]interface{}{
			"CONVERT_TO_BINARY": ContentHandling_CONVERT_TO_BINARY,
			"CONVERT_TO_TEXT": ContentHandling_CONVERT_TO_TEXT,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.CorsHttpMethod",
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
		"aws-cdk-lib.aws_apigatewayv2.CorsPreflightOptions",
		reflect.TypeOf((*CorsPreflightOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.DomainMappingOptions",
		reflect.TypeOf((*DomainMappingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.DomainName",
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
		"aws-cdk-lib.aws_apigatewayv2.DomainNameAttributes",
		reflect.TypeOf((*DomainNameAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.DomainNameProps",
		reflect.TypeOf((*DomainNameProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.EndpointOptions",
		reflect.TypeOf((*EndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.EndpointType",
		reflect.TypeOf((*EndpointType)(nil)).Elem(),
		map[string]interface{}{
			"EDGE": EndpointType_EDGE,
			"REGIONAL": EndpointType_REGIONAL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.GrantInvokeOptions",
		reflect.TypeOf((*GrantInvokeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.HttpApi",
		reflect.TypeOf((*HttpApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoutes", GoMethod: "AddRoutes"},
			_jsii_.MemberMethod{JsiiMethod: "addStage", GoMethod: "AddStage"},
			_jsii_.MemberMethod{JsiiMethod: "addVpcLink", GoMethod: "AddVpcLink"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForExecuteApi", GoMethod: "ArnForExecuteApi"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAuthorizationScopes", GoGetter: "DefaultAuthorizationScopes"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAuthorizer", GoGetter: "DefaultAuthorizer"},
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
		"aws-cdk-lib.aws_apigatewayv2.HttpApiAttributes",
		reflect.TypeOf((*HttpApiAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.HttpApiProps",
		reflect.TypeOf((*HttpApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.HttpAuthorizer",
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
		"aws-cdk-lib.aws_apigatewayv2.HttpAuthorizerAttributes",
		reflect.TypeOf((*HttpAuthorizerAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.HttpAuthorizerProps",
		reflect.TypeOf((*HttpAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.HttpAuthorizerType",
		reflect.TypeOf((*HttpAuthorizerType)(nil)).Elem(),
		map[string]interface{}{
			"IAM": HttpAuthorizerType_IAM,
			"JWT": HttpAuthorizerType_JWT,
			"LAMBDA": HttpAuthorizerType_LAMBDA,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.HttpConnectionType",
		reflect.TypeOf((*HttpConnectionType)(nil)).Elem(),
		map[string]interface{}{
			"VPC_LINK": HttpConnectionType_VPC_LINK,
			"INTERNET": HttpConnectionType_INTERNET,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.HttpIntegration",
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
		"aws-cdk-lib.aws_apigatewayv2.HttpIntegrationProps",
		reflect.TypeOf((*HttpIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.HttpIntegrationSubtype",
		reflect.TypeOf((*HttpIntegrationSubtype)(nil)).Elem(),
		map[string]interface{}{
			"EVENTBRIDGE_PUT_EVENTS": HttpIntegrationSubtype_EVENTBRIDGE_PUT_EVENTS,
			"SQS_SEND_MESSAGE": HttpIntegrationSubtype_SQS_SEND_MESSAGE,
			"SQS_RECEIVE_MESSAGE": HttpIntegrationSubtype_SQS_RECEIVE_MESSAGE,
			"SQS_DELETE_MESSAGE": HttpIntegrationSubtype_SQS_DELETE_MESSAGE,
			"SQS_PURGE_QUEUE": HttpIntegrationSubtype_SQS_PURGE_QUEUE,
			"APPCONFIG_GET_CONFIGURATION": HttpIntegrationSubtype_APPCONFIG_GET_CONFIGURATION,
			"KINESIS_PUT_RECORD": HttpIntegrationSubtype_KINESIS_PUT_RECORD,
			"STEPFUNCTIONS_START_EXECUTION": HttpIntegrationSubtype_STEPFUNCTIONS_START_EXECUTION,
			"STEPFUNCTIONS_START_SYNC_EXECUTION": HttpIntegrationSubtype_STEPFUNCTIONS_START_SYNC_EXECUTION,
			"STEPFUNCTIONS_STOP_EXECUTION": HttpIntegrationSubtype_STEPFUNCTIONS_STOP_EXECUTION,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.HttpIntegrationType",
		reflect.TypeOf((*HttpIntegrationType)(nil)).Elem(),
		map[string]interface{}{
			"HTTP_PROXY": HttpIntegrationType_HTTP_PROXY,
			"AWS_PROXY": HttpIntegrationType_AWS_PROXY,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.HttpMethod",
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
		"aws-cdk-lib.aws_apigatewayv2.HttpNoneAuthorizer",
		reflect.TypeOf((*HttpNoneAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizationType", GoGetter: "AuthorizationType"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpNoneAuthorizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.HttpRoute",
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
		"aws-cdk-lib.aws_apigatewayv2.HttpRouteAuthorizerBindOptions",
		reflect.TypeOf((*HttpRouteAuthorizerBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.HttpRouteAuthorizerConfig",
		reflect.TypeOf((*HttpRouteAuthorizerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.HttpRouteIntegration",
		reflect.TypeOf((*HttpRouteIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "completeBind", GoMethod: "CompleteBind"},
		},
		func() interface{} {
			return &jsiiProxy_HttpRouteIntegration{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.HttpRouteIntegrationBindOptions",
		reflect.TypeOf((*HttpRouteIntegrationBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.HttpRouteIntegrationConfig",
		reflect.TypeOf((*HttpRouteIntegrationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.HttpRouteKey",
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
		"aws-cdk-lib.aws_apigatewayv2.HttpRouteProps",
		reflect.TypeOf((*HttpRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.HttpStage",
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
		"aws-cdk-lib.aws_apigatewayv2.HttpStageAttributes",
		reflect.TypeOf((*HttpStageAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.HttpStageOptions",
		reflect.TypeOf((*HttpStageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.HttpStageProps",
		reflect.TypeOf((*HttpStageProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_apigatewayv2.IApi",
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
		"aws-cdk-lib.aws_apigatewayv2.IApiMapping",
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
		"aws-cdk-lib.aws_apigatewayv2.IAuthorizer",
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
		"aws-cdk-lib.aws_apigatewayv2.IDomainName",
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
		"aws-cdk-lib.aws_apigatewayv2.IHttpApi",
		reflect.TypeOf((*IHttpApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVpcLink", GoMethod: "AddVpcLink"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForExecuteApi", GoMethod: "ArnForExecuteApi"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAuthorizationScopes", GoGetter: "DefaultAuthorizationScopes"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAuthorizer", GoGetter: "DefaultAuthorizer"},
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
		},
		func() interface{} {
			j := jsiiProxy_IHttpApi{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApi)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_apigatewayv2.IHttpAuthorizer",
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
		"aws-cdk-lib.aws_apigatewayv2.IHttpIntegration",
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
		"aws-cdk-lib.aws_apigatewayv2.IHttpRoute",
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
		"aws-cdk-lib.aws_apigatewayv2.IHttpRouteAuthorizer",
		reflect.TypeOf((*IHttpRouteAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IHttpRouteAuthorizer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_apigatewayv2.IHttpStage",
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
		"aws-cdk-lib.aws_apigatewayv2.IIntegration",
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
		"aws-cdk-lib.aws_apigatewayv2.IMappingValue",
		reflect.TypeOf((*IMappingValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IMappingValue{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_apigatewayv2.IRoute",
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
		"aws-cdk-lib.aws_apigatewayv2.IStage",
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
		"aws-cdk-lib.aws_apigatewayv2.IVpcLink",
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
		"aws-cdk-lib.aws_apigatewayv2.IWebSocketApi",
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
		"aws-cdk-lib.aws_apigatewayv2.IWebSocketAuthorizer",
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
		"aws-cdk-lib.aws_apigatewayv2.IWebSocketIntegration",
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
		"aws-cdk-lib.aws_apigatewayv2.IWebSocketRoute",
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
		"aws-cdk-lib.aws_apigatewayv2.IWebSocketRouteAuthorizer",
		reflect.TypeOf((*IWebSocketRouteAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IWebSocketRouteAuthorizer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_apigatewayv2.IWebSocketStage",
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
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.IntegrationCredentials",
		reflect.TypeOf((*IntegrationCredentials)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "credentialsArn", GoGetter: "CredentialsArn"},
		},
		func() interface{} {
			return &jsiiProxy_IntegrationCredentials{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.MTLSConfig",
		reflect.TypeOf((*MTLSConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
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
		"aws-cdk-lib.aws_apigatewayv2.ParameterMapping",
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
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.PassthroughBehavior",
		reflect.TypeOf((*PassthroughBehavior)(nil)).Elem(),
		map[string]interface{}{
			"WHEN_NO_MATCH": PassthroughBehavior_WHEN_NO_MATCH,
			"NEVER": PassthroughBehavior_NEVER,
			"WHEN_NO_TEMPLATES": PassthroughBehavior_WHEN_NO_TEMPLATES,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.PayloadFormatVersion",
		reflect.TypeOf((*PayloadFormatVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_PayloadFormatVersion{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.SecurityPolicy",
		reflect.TypeOf((*SecurityPolicy)(nil)).Elem(),
		map[string]interface{}{
			"TLS_1_0": SecurityPolicy_TLS_1_0,
			"TLS_1_2": SecurityPolicy_TLS_1_2,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.StageAttributes",
		reflect.TypeOf((*StageAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.StageOptions",
		reflect.TypeOf((*StageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.ThrottleSettings",
		reflect.TypeOf((*ThrottleSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.VpcLink",
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
		"aws-cdk-lib.aws_apigatewayv2.VpcLinkAttributes",
		reflect.TypeOf((*VpcLinkAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.VpcLinkProps",
		reflect.TypeOf((*VpcLinkProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApi",
		reflect.TypeOf((*WebSocketApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForExecuteApi", GoMethod: "ArnForExecuteApi"},
			_jsii_.MemberMethod{JsiiMethod: "arnForExecuteApiV2", GoMethod: "ArnForExecuteApiV2"},
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
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApiAttributes",
		reflect.TypeOf((*WebSocketApiAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApiKeySelectionExpression",
		reflect.TypeOf((*WebSocketApiKeySelectionExpression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "customApiKeySelector", GoGetter: "CustomApiKeySelector"},
		},
		func() interface{} {
			return &jsiiProxy_WebSocketApiKeySelectionExpression{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApiProps",
		reflect.TypeOf((*WebSocketApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketAuthorizer",
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
		"aws-cdk-lib.aws_apigatewayv2.WebSocketAuthorizerAttributes",
		reflect.TypeOf((*WebSocketAuthorizerAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketAuthorizerProps",
		reflect.TypeOf((*WebSocketAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketAuthorizerType",
		reflect.TypeOf((*WebSocketAuthorizerType)(nil)).Elem(),
		map[string]interface{}{
			"LAMBDA": WebSocketAuthorizerType_LAMBDA,
			"IAM": WebSocketAuthorizerType_IAM,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketIntegration",
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
		"aws-cdk-lib.aws_apigatewayv2.WebSocketIntegrationProps",
		reflect.TypeOf((*WebSocketIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketIntegrationType",
		reflect.TypeOf((*WebSocketIntegrationType)(nil)).Elem(),
		map[string]interface{}{
			"AWS_PROXY": WebSocketIntegrationType_AWS_PROXY,
			"MOCK": WebSocketIntegrationType_MOCK,
			"AWS": WebSocketIntegrationType_AWS,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketNoneAuthorizer",
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
		"aws-cdk-lib.aws_apigatewayv2.WebSocketRoute",
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
		"aws-cdk-lib.aws_apigatewayv2.WebSocketRouteAuthorizerBindOptions",
		reflect.TypeOf((*WebSocketRouteAuthorizerBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketRouteAuthorizerConfig",
		reflect.TypeOf((*WebSocketRouteAuthorizerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketRouteIntegration",
		reflect.TypeOf((*WebSocketRouteIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_WebSocketRouteIntegration{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketRouteIntegrationBindOptions",
		reflect.TypeOf((*WebSocketRouteIntegrationBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketRouteIntegrationConfig",
		reflect.TypeOf((*WebSocketRouteIntegrationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketRouteOptions",
		reflect.TypeOf((*WebSocketRouteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketRouteProps",
		reflect.TypeOf((*WebSocketRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketStage",
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
		"aws-cdk-lib.aws_apigatewayv2.WebSocketStageAttributes",
		reflect.TypeOf((*WebSocketStageAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketStageProps",
		reflect.TypeOf((*WebSocketStageProps)(nil)).Elem(),
	)
}
