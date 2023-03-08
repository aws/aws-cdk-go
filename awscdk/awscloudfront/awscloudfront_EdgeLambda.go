package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Represents a Lambda function version and event type when using Lambda@Edge.
//
// The type of the {@link AddBehaviorOptions.edgeLambdas} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var version version
//
//   edgeLambda := &edgeLambda{
//   	eventType: awscdk.Aws_cloudfront.lambdaEdgeEventType_ORIGIN_REQUEST,
//   	functionVersion: version,
//
//   	// the properties below are optional
//   	includeBody: jsii.Boolean(false),
//   }
//
// Experimental.
type EdgeLambda struct {
	// The type of event in response to which should the function be invoked.
	// Experimental.
	EventType LambdaEdgeEventType `field:"required" json:"eventType" yaml:"eventType"`
	// The version of the Lambda function that will be invoked.
	//
	// **Note**: it's not possible to use the '$LATEST' function version for Lambda@Edge!
	// Experimental.
	FunctionVersion awslambda.IVersion `field:"required" json:"functionVersion" yaml:"functionVersion"`
	// Allows a Lambda function to have read access to the body content.
	//
	// Only valid for "request" event types (`ORIGIN_REQUEST` or `VIEWER_REQUEST`).
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html
	// Experimental.
	IncludeBody *bool `field:"optional" json:"includeBody" yaml:"includeBody"`
}

