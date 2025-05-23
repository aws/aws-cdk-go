package awscdkintegtestsalpha


// Options to pass to the JavaScript fetch api.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   fetchOptions := &FetchOptions{
//   	Body: jsii.String("body"),
//   	Headers: map[string]*string{
//   		"headersKey": jsii.String("headers"),
//   	},
//   	Method: jsii.String("method"),
//   	Port: jsii.Number(123),
//   }
//
// Experimental.
type FetchOptions struct {
	// Request body.
	// Default: - no body.
	//
	// Experimental.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Optional request headers.
	// Default: no headers.
	//
	// Experimental.
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// HTTP method.
	// Default: GET.
	//
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Optional port.
	// Default: default port for protocol.
	//
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

