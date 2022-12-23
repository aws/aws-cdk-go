package awsapigateway


// Example:
//   var domain domainName
//   var api1 restApi
//   var api2 restApi
//
//
//   domain.addBasePathMapping(api1, &basePathMappingOptions{
//   	basePath: jsii.String("go-to-api1"),
//   })
//   domain.addBasePathMapping(api2, &basePathMappingOptions{
//   	basePath: jsii.String("boom"),
//   })
//
type BasePathMappingOptions struct {
	// Whether to attach the base path mapping to a stage.
	//
	// Use this property to create a base path mapping without attaching it to the Rest API default stage.
	// This property is ignored if `stage` is provided.
	AttachToStage *bool `field:"optional" json:"attachToStage" yaml:"attachToStage"`
	// The base path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The Deployment stage of API [disable-awslint:ref-via-interface].
	Stage Stage `field:"optional" json:"stage" yaml:"stage"`
}

