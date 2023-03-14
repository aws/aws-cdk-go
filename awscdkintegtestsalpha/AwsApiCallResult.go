// CDK Integration Testing Constructs
package awscdkintegtestsalpha


// The result from a SdkQuery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var apiCallResponse interface{}
//
//   awsApiCallResult := &AwsApiCallResult{
//   	ApiCallResponse: apiCallResponse,
//   }
//
// Experimental.
type AwsApiCallResult struct {
	// The full api response.
	// Experimental.
	ApiCallResponse interface{} `field:"required" json:"apiCallResponse" yaml:"apiCallResponse"`
}

