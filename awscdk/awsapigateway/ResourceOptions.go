package awsapigateway


// Example:
//   var resource resource
//
//
//   subtree := resource.AddResource(jsii.String("subtree"), &ResourceOptions{
//   	DefaultCorsPreflightOptions: &CorsOptions{
//   		AllowOrigins: []*string{
//   			jsii.String("https://amazon.com"),
//   		},
//   	},
//   })
//
type ResourceOptions struct {
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
}

