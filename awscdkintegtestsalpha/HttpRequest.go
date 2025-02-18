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
//
//   	// the properties below are optional
//   	FlattenResponse: jsii.String("flattenResponse"),
//   }
//
// Experimental.
type HttpRequest struct {
	// Parameters from the custom resource.
	// Experimental.
	Parameters *HttpRequestParameters `field:"required" json:"parameters" yaml:"parameters"`
	// Whether or not to flatten the response from the HTTP request.
	//
	// Valid values are 'true' or 'false' as strings
	//
	// Typically when using an HttpRequest you will be passing it as the
	// `actual` value to an assertion provider so this would be set
	// to 'false' (you want the actual response).
	//
	// If you are using the HttpRequest to perform more of a query to return
	// a single value to use, then this should be set to 'true'.
	// Default: 'false'.
	//
	// Experimental.
	FlattenResponse *string `field:"optional" json:"flattenResponse" yaml:"flattenResponse"`
}

