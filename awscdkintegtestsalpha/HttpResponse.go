package awscdkintegtestsalpha


// Response from fetch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var headers interface{}
//   var responseJson interface{}
//
//   httpResponse := &HttpResponse{
//   	Headers: map[string]interface{}{
//   		"headersKey": headers,
//   	},
//   	Ok: jsii.Boolean(false),
//   	ResponseJson: map[string]interface{}{
//   		"responseJsonKey": responseJson,
//   	},
//   	Status: jsii.Number(123),
//   	StatusText: jsii.String("statusText"),
//   }
//
// Experimental.
type HttpResponse struct {
	// Headers associated with the response.
	// Experimental.
	Headers *map[string]interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Indicates whether the response was successful.
	//
	// status range 200-299.
	// Experimental.
	Ok *bool `field:"optional" json:"ok" yaml:"ok"`
	// The response as JSON.
	// Experimental.
	ResponseJson *map[string]interface{} `field:"optional" json:"responseJson" yaml:"responseJson"`
	// Status code of the response.
	// Experimental.
	Status *float64 `field:"optional" json:"status" yaml:"status"`
	// The status message corresponding to the status code.
	// Experimental.
	StatusText *string `field:"optional" json:"statusText" yaml:"statusText"`
}

