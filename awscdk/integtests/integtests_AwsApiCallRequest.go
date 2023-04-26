package integtests


// A AWS JavaScript SDK V2 request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   awsApiCallRequest := &AwsApiCallRequest{
//   	Api: jsii.String("api"),
//   	Service: jsii.String("service"),
//
//   	// the properties below are optional
//   	FlattenResponse: jsii.String("flattenResponse"),
//   	Parameters: parameters,
//   }
//
// Experimental.
type AwsApiCallRequest struct {
	// The AWS api call to make i.e. getBucketLifecycle.
	// Experimental.
	Api *string `field:"required" json:"api" yaml:"api"`
	// The AWS service i.e. S3.
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Whether or not to flatten the response from the api call.
	//
	// Valid values are 'true' or 'false' as strings
	//
	// Typically when using an SdkRequest you will be passing it as the
	// `actual` value to an assertion provider so this would be set
	// to 'false' (you want the actual response).
	//
	// If you are using the SdkRequest to perform more of a query to return
	// a single value to use, then this should be set to 'true'. For example,
	// you could make a StepFunctions.startExecution api call and retreive the
	// `executionArn` from the response.
	// Experimental.
	FlattenResponse *string `field:"optional" json:"flattenResponse" yaml:"flattenResponse"`
	// Any parameters to pass to the api call.
	// Experimental.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

