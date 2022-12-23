package integtests


// The result from a SdkQuery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiCallResponse interface{}
//
//   awsApiCallResult := &awsApiCallResult{
//   	apiCallResponse: apiCallResponse,
//   }
//
// Experimental.
type AwsApiCallResult struct {
	// The full api response.
	// Experimental.
	ApiCallResponse interface{} `field:"required" json:"apiCallResponse" yaml:"apiCallResponse"`
}

