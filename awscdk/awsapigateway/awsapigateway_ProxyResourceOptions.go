package awsapigateway


// Example:
//   var resource resource
//   var handler function
//
//   proxy := resource.addProxy(&proxyResourceOptions{
//   	defaultIntegration: apigateway.NewLambdaIntegration(handler),
//
//   	// "false" will require explicitly adding methods on the `proxy` resource
//   	anyMethod: jsii.Boolean(true),
//   })
//
type ProxyResourceOptions struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// Adds an "ANY" method to this resource.
	//
	// If set to `false`, you will have to explicitly
	// add methods to this resource after it's created.
	AnyMethod *bool `field:"optional" json:"anyMethod" yaml:"anyMethod"`
}

