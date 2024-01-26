package awscloudfront


// Properties for creating a CloudFront Function.
//
// Example:
//   var s3Bucket bucket
//   // Add a cloudfront Function to a Distribution
//   cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: cloudfront.FunctionCode_FromInline(jsii.String("function handler(event) { return event.request }")),
//   	Runtime: cloudfront.FunctionRuntime_JS_2_0(),
//   })
//   cloudfront.NewDistribution(this, jsii.String("distro"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3Origin(s3Bucket),
//   		FunctionAssociations: []functionAssociation{
//   			&functionAssociation{
//   				Function: cfFunction,
//   				EventType: cloudfront.FunctionEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
type FunctionProps struct {
	// The source code of the function.
	Code FunctionCode `field:"required" json:"code" yaml:"code"`
	// A comment to describe the function.
	// Default: - same as `functionName`.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A name to identify the function.
	// Default: - generated from the `id`.
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The Key Value Store to associate with this function.
	//
	// In order to associate a Key Value Store, the `runtime` must be
	// `cloudfront-js-2.0` or newer.
	// Default: - no key value store is associated.
	//
	KeyValueStore IKeyValueStore `field:"optional" json:"keyValueStore" yaml:"keyValueStore"`
	// The runtime environment for the function.
	// Default: FunctionRuntime.JS_1_0 (unless `keyValueStore` is specified, then `FunctionRuntime.JS_2_0`)
	//
	Runtime FunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
}

