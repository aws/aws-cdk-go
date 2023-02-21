// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha


// HTTP response codes for Endpoint invocations.
// Experimental.
type InvocationHttpResponseCode string

const (
	// 4xx response codes from Endpoint invocations.
	// Experimental.
	InvocationHttpResponseCode_INVOCATION_4XX_ERRORS InvocationHttpResponseCode = "INVOCATION_4XX_ERRORS"
	// 5xx response codes from Endpoint invocations.
	// Experimental.
	InvocationHttpResponseCode_INVOCATION_5XX_ERRORS InvocationHttpResponseCode = "INVOCATION_5XX_ERRORS"
)

