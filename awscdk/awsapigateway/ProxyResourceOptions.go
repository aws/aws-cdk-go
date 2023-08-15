package awsapigateway


// Example:
//   var resource resource
//   var handler function
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
}

