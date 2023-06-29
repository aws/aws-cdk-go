package awscdkintegtestsalpha


// Request to the HttpCall resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   httpRequest := &HttpRequest{
//   	Parameters: &HttpRequestParameters{
//   		Url: jsii.String("url"),
//
//   		// the properties below are optional
//   		FetchOptions: &FetchOptions{
//   			Body: jsii.String("body"),
//   			Headers: map[string]*string{
//   				"headersKey": jsii.String("headers"),
//   			},
//   			Method: jsii.String("method"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   }
//
// Experimental.
type HttpRequest struct {
	// Parameters from the custom resource.
	// Experimental.
	Parameters *HttpRequestParameters `field:"required" json:"parameters" yaml:"parameters"`
}

