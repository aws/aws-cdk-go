package integtests


// Options for creating an SDKQuery provider.
//
// Example:
//   var myAppStack stack
//
//
//   awscdk.NewAwsApiCall(myAppStack, jsii.String("GetObject"), &awsApiCallProps{
//   	service: jsii.String("S3"),
//   	api: jsii.String("getObject"),
//   })
//
// Experimental.
type AwsApiCallProps struct {
	// The api call to make, i.e. getBucketLifecycle.
	// Experimental.
	Api *string `field:"required" json:"api" yaml:"api"`
	// The AWS service, i.e. S3.
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Any parameters to pass to the api call.
	// Experimental.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

