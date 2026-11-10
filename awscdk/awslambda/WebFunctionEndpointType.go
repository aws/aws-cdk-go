package awslambda


// The endpoint type for a WebFunction endpoint.
//
// Choose it with `WebFunctionEndpointLocation` when defining an endpoint.
//
// Example:
//   endpoint := lambda.WebFunctionEndpoint_FromWebFunctionEndpointAttributes(this, jsii.String("Endpoint"), &WebFunctionEndpointAttributes{
//   	EndpointArn: jsii.String("arn:aws:lambda:us-west-2:123456789012:web-function/my-function/endpoint/prod"),
//   	DomainName: jsii.String("abcd1234.lambda-web.us-west-2.on.aws"),
//   	EndpointType: lambda.WebFunctionEndpointType_HOME_REGION,
//   	AuthType: lambda.WebFunctionAuthType_IAM_AUTH,
//   })
//
type WebFunctionEndpointType string

const (
	// Single-region endpoint deployed in the stack's region only.
	//
	// The function is accessible from a URL in this region.
	WebFunctionEndpointType_HOME_REGION WebFunctionEndpointType = "HOME_REGION"
	// Multi-region endpoint with managed latency-based routing.
	//
	// The function is replicated to the specified replica regions
	// and users are routed to the nearest region.
	WebFunctionEndpointType_MULTI_REGION WebFunctionEndpointType = "MULTI_REGION"
	// Per-region endpoint: a separate regional endpoint is exposed for each replica region (no managed cross-region routing).
	WebFunctionEndpointType_PER_REGION WebFunctionEndpointType = "PER_REGION"
)

