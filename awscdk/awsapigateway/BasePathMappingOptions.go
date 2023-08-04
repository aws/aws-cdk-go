package awsapigateway


// Example:
//   var domain domainName
//   var api1 restApi
//   var api2 restApi
//
//
//   domain.AddBasePathMapping(api1, &BasePathMappingOptions{
//   	BasePath: jsii.String("go-to-api1"),
//   })
//   domain.AddBasePathMapping(api2, &BasePathMappingOptions{
//   	BasePath: jsii.String("boom"),
//   })
//
type BasePathMappingOptions struct {
	// Whether to attach the base path mapping to a stage.
	//
	// Use this property to create a base path mapping without attaching it to the Rest API default stage.
	// This property is ignored if `stage` is provided.
	// Default: - true.
	//
	AttachToStage *bool `field:"optional" json:"attachToStage" yaml:"attachToStage"`
	// The base path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	// Default: - map requests from the domain root (e.g. `example.com`). If this
	// is undefined, no additional mappings will be allowed on this domain name.
	//
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The Deployment stage of API [disable-awslint:ref-via-interface].
	// Default: - map to deploymentStage of restApi otherwise stage needs to pass in URL.
	//
	Stage Stage `field:"optional" json:"stage" yaml:"stage"`
}

