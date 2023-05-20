package awscdkintegtestsalpha


// Options to perform an AWS JavaScript V2 API call.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var parameters interface{}
//
//   awsApiCallOptions := &AwsApiCallOptions{
//   	Api: jsii.String("api"),
//   	Service: jsii.String("service"),
//
//   	// the properties below are optional
//   	OutputPaths: []*string{
//   		jsii.String("outputPaths"),
//   	},
//   	Parameters: parameters,
//   }
//
// Experimental.
type AwsApiCallOptions struct {
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

