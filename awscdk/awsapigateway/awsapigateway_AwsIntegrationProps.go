package awsapigateway


// Example:
//   getMessageIntegration := apigateway.NewAwsIntegration(&awsIntegrationProps{
//   	service: jsii.String("sqs"),
//   	path: jsii.String("queueName"),
//   	region: jsii.String("eu-west-1"),
//   })
//
type AwsIntegrationProps struct {
	// The name of the integrated AWS service (e.g. `s3`).
	Service *string `field:"required" json:"service" yaml:"service"`
	// The AWS action to perform in the integration.
	//
	// Use `actionParams` to specify key-value params for the action.
	//
	// Mutually exclusive with `path`.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Parameters for the action.
	//
	// `action` must be set, and `path` must be undefined.
	// The action params will be URL encoded.
	ActionParameters *map[string]*string `field:"optional" json:"actionParameters" yaml:"actionParameters"`
	// The integration's HTTP method type.
	IntegrationHttpMethod *string `field:"optional" json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// Integration options, such as content handling, request/response mapping, etc.
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// The path to use for path-base APIs.
	//
	// For example, for S3 GET, you can set path to `bucket/key`.
	// For lambda, you can set path to `2015-03-31/functions/${function-arn}/invocations`
	//
	// Mutually exclusive with the `action` options.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Use AWS_PROXY integration.
	Proxy *bool `field:"optional" json:"proxy" yaml:"proxy"`
	// The region of the integrated AWS service.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A designated subdomain supported by certain AWS service for fast host-name lookup.
	Subdomain *string `field:"optional" json:"subdomain" yaml:"subdomain"`
}

