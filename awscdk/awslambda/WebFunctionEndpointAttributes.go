package awslambda


// Attributes for importing an existing WebFunction endpoint.
//
// Example:
//   endpoint := lambda.WebFunctionEndpoint_FromWebFunctionEndpointAttributes(this, jsii.String("Endpoint"), &WebFunctionEndpointAttributes{
//   	EndpointArn: jsii.String("arn:aws:lambda:us-west-2:123456789012:web-function/my-function/endpoint/prod"),
//   	DomainName: jsii.String("abcd1234.lambda-web.us-west-2.on.aws"),
//   	EndpointType: lambda.WebFunctionEndpointType_HOME_REGION,
//   	AuthType: lambda.WebFunctionAuthType_IAM_AUTH,
//   })
//
type WebFunctionEndpointAttributes struct {
	// How the endpoint authenticates requests.
	//
	// For an imported endpoint this is taken as given: it is not verified against the
	// deployed endpoint and must match it, because downstream constructs rely on it for
	// a synth-time security decision. Like `domainName`, it cannot be derived from the ARN.
	AuthType WebFunctionAuthType `field:"required" json:"authType" yaml:"authType"`
	// The domain name of the endpoint.
	//
	// The service assigns it, so it cannot be derived from the ARN.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The ARN of the endpoint, `arn:<partition>:lambda:<region>:<account>:web-function/<function>/endpoint/<endpoint>`.
	EndpointArn *string `field:"required" json:"endpointArn" yaml:"endpointArn"`
	// The endpoint type.
	EndpointType WebFunctionEndpointType `field:"required" json:"endpointType" yaml:"endpointType"`
}

