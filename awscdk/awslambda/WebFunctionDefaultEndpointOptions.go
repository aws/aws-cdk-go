package awslambda


// Settings of a WebFunction endpoint other than its auth type.
//
// Used for the endpoint a WebFunction creates for itself, whose auth type is set by
// `WebFunctionProps.defaultEndpointAuth`.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var codeBucket IBucket
//
//
//   lambda.NewWebFunction(this, jsii.String("GlobalApi"), &WebFunctionProps{
//   	Code: lambda.WebFunctionCode_FromBucket(codeBucket, jsii.String("app.zip"), jsii.String("version-id")),
//   	Runtime: lambda.WebFunctionRuntime_NODEJS_24_X(),
//   	DefaultEndpointAuth: lambda.WebFunctionAuthType_IAM_AUTH,
//   	DefaultEndpoint: &WebFunctionDefaultEndpointOptions{
//   		EndpointLocation: lambda.WebFunctionEndpointLocation_MultiRegion([]*string{
//   			jsii.String("us-east-1"),
//   			jsii.String("eu-west-1"),
//   			jsii.String("ap-northeast-1"),
//   		}),
//   	},
//   })
//
type WebFunctionDefaultEndpointOptions struct {
	// A description of the endpoint.
	//
	// Up to 256 characters.
	// Default: - No description.
	//
	EndpointDescription *string `field:"optional" json:"endpointDescription" yaml:"endpointDescription"`
	// Where the endpoint serves traffic: the home Region only, or a set of replica Regions.
	//
	// Changing it replaces the endpoint, which yields a new `domainName`. If you set
	// `endpointName`, change it at the same time.
	// Default: WebFunctionEndpointLocation.homeRegion()
	//
	EndpointLocation WebFunctionEndpointLocation `field:"optional" json:"endpointLocation" yaml:"endpointLocation"`
	// The name of the WebFunction endpoint.
	//
	// Must be 2 to 64 characters of letters, numbers, hyphens or underscores, starting and
	// ending with a letter or number.
	//
	// Changing this property replaces the endpoint, which yields a new `domainName`.
	//
	// `endpointLocation` also replaces the endpoint when it changes, and CloudFormation cannot
	// replace a resource that keeps its name. Without this property the name is derived from the
	// location, so that works. With it, change the name together with the location, or the
	// deployment fails.
	// Default: - for the endpoint a WebFunction creates, `prod` in the home Region and, for a
	// regional endpoint, `prod-` followed by a hash of its location; for any other endpoint, a
	// name derived from its construct path and location.
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The maximum number of concurrent execution environments for the endpoint.
	//
	// This is an optional cap on top of your account's vCPU quota, not a reservation:
	// it further constrains how far the endpoint scales. Must be 2 or greater.
	// Default: - Scaling is limited only by your account's vCPU quota.
	//
	MaxEnvironments *float64 `field:"optional" json:"maxEnvironments" yaml:"maxEnvironments"`
	// The maximum request rate per second for the endpoint, up to 10000.
	//
	// This is an optional cap on top of your account's rate limit quota. Specify 0 to
	// reject all new requests.
	//
	// The service accepts only certain values: 0, 100 through 1000 in increments of 100,
	// and 2000 through 10000 in increments of 1000. Which of those are available varies by
	// Region, and an unsupported value is rejected at deploy time with the values that
	// Region does accept. Only the lower bound is checked at synth, since the schema
	// declares nothing more.
	// Default: - The request rate is limited only by your account's rate limit quota.
	//
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

