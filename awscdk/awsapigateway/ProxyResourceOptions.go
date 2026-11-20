package awsapigateway


// Example:
//   var resource Resource
//   var handler Function
//
//   proxy := resource.AddProxy(&ProxyResourceOptions{
//   	DefaultIntegration: apigateway.NewLambdaIntegration(handler),
//
//   	// "false" will require explicitly adding methods on the `proxy` resource
//   	AnyMethod: jsii.Boolean(true),
//   })
//
type ProxyResourceOptions struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	// Default: - CORS is disabled.
	//
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	// Default: - Inherited from parent.
	//
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	// Default: - Inherited from parent.
	//
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// Adds an "ANY" method to this resource.
	//
	// If set to `false`, you will have to explicitly
	// add methods to this resource after it's created.
	// Default: true.
	//
	AnyMethod *bool `field:"optional" json:"anyMethod" yaml:"anyMethod"`
	// Automatically configure path parameter mapping for the proxy resource.
	//
	// When enabled, this property automatically configures both the method request
	// and integration request to properly handle the proxy path parameter:
	// - Adds `method.request.path.proxy: true` to method request parameters
	// - Adds `integration.request.path.proxy: 'method.request.path.proxy'` to integration request parameters
	//
	// This matches the behavior of the AWS Console when creating proxy resources
	// and eliminates the need to manually configure path parameters.
	//
	// This is meant for integrations whose URI uses `{proxy}`, such as HTTP
	// integrations. Lambda proxy (`AWS_PROXY`) integrations do not need it, as
	// Lambda already receives the full request path. The mapping is added even
	// when the integration URI has no `{proxy}`.
	//
	// Mock integrations are left unchanged. When the proxy is added to the root
	// resource, the method that is also added to the root has no `{proxy}` path
	// parameter, so it is not configured.
	// Default: false.
	//
	AutoConfigurePathParameter *bool `field:"optional" json:"autoConfigurePathParameter" yaml:"autoConfigurePathParameter"`
}

