package awsevents


// Properties for defining a `CfnApiDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiDestinationProps := &cfnApiDestinationProps{
//   	connectionArn: jsii.String("connectionArn"),
//   	httpMethod: jsii.String("httpMethod"),
//   	invocationEndpoint: jsii.String("invocationEndpoint"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	invocationRateLimitPerSecond: jsii.Number(123),
//   	name: jsii.String("name"),
//   }
//
type CfnApiDestinationProps struct {
	// The ARN of the connection to use for the API destination.
	//
	// The destination endpoint must support the authorization type specified for the connection.
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The method to use for the request to the HTTP invocation endpoint.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The URL to the HTTP invocation endpoint for the API destination.
	InvocationEndpoint *string `field:"required" json:"invocationEndpoint" yaml:"invocationEndpoint"`
	// A description for the API destination to create.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum number of requests per second to send to the HTTP invocation endpoint.
	InvocationRateLimitPerSecond *float64 `field:"optional" json:"invocationRateLimitPerSecond" yaml:"invocationRateLimitPerSecond"`
	// The name for the API destination to create.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

