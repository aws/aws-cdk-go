// CDK Integration Testing Constructs
package awscdkintegtestsalpha


// Construct that creates a custom resource that will perform a query using the AWS SDK.
//
// Example:
//   var myAppStack stack
//
//
//   awscdkintegtestsalpha.NewAwsApiCall(myAppStack, jsii.String("GetObject"), &AwsApiCallProps{
//   	Service: jsii.String("S3"),
//   	Api: jsii.String("getObject"),
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
	// Restrict the data returned by the API call to specific paths in the API response.
	//
	// Use this to limit the data returned by the custom
	// resource if working with API calls that could potentially result in custom
	// response objects exceeding the hard limit of 4096 bytes.
	// Experimental.
	OutputPaths *[]*string `field:"optional" json:"outputPaths" yaml:"outputPaths"`
	// Any parameters to pass to the api call.
	// Experimental.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

