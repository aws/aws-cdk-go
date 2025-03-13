package awscdkintegtestsalpha


// Response from the HttpCall resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var body interface{}
//   var headers interface{}
//
//   httpResponseWrapper := &HttpResponseWrapper{
//   	ApiCallResponse: &HttpResponse{
//   		Body: body,
//   		Headers: map[string]interface{}{
//   			"headersKey": headers,
//   		},
//   		Ok: jsii.Boolean(false),
//   		Status: jsii.Number(123),
//   		StatusText: jsii.String("statusText"),
//   	},
//   }
//
// Experimental.
type HttpResponseWrapper struct {
	// The Response from the fetch request.
	// Experimental.
	ApiCallResponse *HttpResponse `field:"required" json:"apiCallResponse" yaml:"apiCallResponse"`
}

