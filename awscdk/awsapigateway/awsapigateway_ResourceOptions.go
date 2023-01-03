package awsapigateway


// Example:
//   var resource resource
//
//
//   subtree := resource.addResource(jsii.String("subtree"), &resourceOptions{
//   	defaultCorsPreflightOptions: &corsOptions{
//   		allowOrigins: []*string{
//   			jsii.String("https://amazon.com"),
//   		},
//   	},
//   })
//
// Experimental.
type ResourceOptions struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	// Experimental.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	// Experimental.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	// Experimental.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
}

