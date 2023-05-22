package integtests


// Options to perform an AWS JavaScript V2 API call.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   awsApiCallOptions := &AwsApiCallOptions{
//   	Api: jsii.String("api"),
//   	Service: jsii.String("service"),
//
//   	// the properties below are optional
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
	// Any parameters to pass to the api call.
	// Experimental.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

