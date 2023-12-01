package awsapigatewayv2


// Properties to initialize an instance of `HttpApi`.
//
// Example:
//   import "github.com/aws-samples/dummy/awscdklib/awsapigatewayv2integrations"
//
//   var handler function
//   var dn domainName
//
//
//   apiDemo := apigwv2.NewHttpApi(this, jsii.String("DemoApi"), &HttpApiProps{
//   	DefaultIntegration: awscdklibawsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("DefaultIntegration"), handler),
//   	// https://${dn.domainName}/demo goes to apiDemo $default stage
//   	DefaultDomainMapping: &DomainMappingOptions{
//   		DomainName: dn,
//   		MappingKey: jsii.String("demo"),
//   	},
//   })
//
type HttpApiProps struct {
	// Name for the HTTP API resource.
	// Default: - id of the HttpApi construct.
	//
	ApiName *string `field:"optional" json:"apiName" yaml:"apiName"`
	// Specifies a CORS configuration for an API.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html
	//
	// Default: - CORS disabled.
	//
	CorsPreflight *CorsPreflightOptions `field:"optional" json:"corsPreflight" yaml:"corsPreflight"`
	// Whether a default stage and deployment should be automatically created.
	// Default: true.
	//
	CreateDefaultStage *bool `field:"optional" json:"createDefaultStage" yaml:"createDefaultStage"`
	// Default OIDC scopes attached to all routes in the gateway, unless explicitly configured on the route.
	//
	// The scopes are used with a COGNITO_USER_POOLS authorizer to authorize the method invocation.
	// Default: - no default authorization scopes.
	//
	DefaultAuthorizationScopes *[]*string `field:"optional" json:"defaultAuthorizationScopes" yaml:"defaultAuthorizationScopes"`
	// Default Authorizer applied to all routes in the gateway.
	// Default: - no default authorizer.
	//
	DefaultAuthorizer IHttpRouteAuthorizer `field:"optional" json:"defaultAuthorizer" yaml:"defaultAuthorizer"`
	// Configure a custom domain with the API mapping resource to the HTTP API.
	// Default: - no default domain mapping configured. meaningless if `createDefaultStage` is `false`.
	//
	DefaultDomainMapping *DomainMappingOptions `field:"optional" json:"defaultDomainMapping" yaml:"defaultDomainMapping"`
	// An integration that will be configured on the catch-all route ($default).
	// Default: - none.
	//
	DefaultIntegration HttpRouteIntegration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// The description of the API.
	// Default: - none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke your API using the default endpoint.
	//
	// By default, clients can invoke your API with the default
	// `https://{api_id}.execute-api.{region}.amazonaws.com` endpoint. Enable
	// this if you would like clients to use your custom domain name.
	// Default: false execute-api endpoint enabled.
	//
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
}

