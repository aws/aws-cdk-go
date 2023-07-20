package awsevents


// Properties for defining a `CfnApiDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiDestinationProps := &CfnApiDestinationProps{
//   	ConnectionArn: jsii.String("connectionArn"),
//   	HttpMethod: jsii.String("httpMethod"),
//   	InvocationEndpoint: jsii.String("invocationEndpoint"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	InvocationRateLimitPerSecond: jsii.Number(123),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html
//
type CfnApiDestinationProps struct {
	// The ARN of the connection to use for the API destination.
	//
	// The destination endpoint must support the authorization type specified for the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-connectionarn
	//
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The method to use for the request to the HTTP invocation endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-httpmethod
	//
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The URL to the HTTP invocation endpoint for the API destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-invocationendpoint
	//
	InvocationEndpoint *string `field:"required" json:"invocationEndpoint" yaml:"invocationEndpoint"`
	// A description for the API destination to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum number of requests per second to send to the HTTP invocation endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-invocationratelimitpersecond
	//
	InvocationRateLimitPerSecond *float64 `field:"optional" json:"invocationRateLimitPerSecond" yaml:"invocationRateLimitPerSecond"`
	// The name for the API destination to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

