package awscloudfront


// Represents a CloudFront function and event type when using CF Functions.
//
// The type of the `AddBehaviorOptions.functionAssociations` property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//
//   functionAssociation := &FunctionAssociation{
//   	EventType: awscdk.Aws_cloudfront.FunctionEventType_VIEWER_REQUEST,
//   	Function: function_,
//   }
//
type FunctionAssociation struct {
	// The type of event which should invoke the function.
	EventType FunctionEventType `field:"required" json:"eventType" yaml:"eventType"`
	// The CloudFront function that will be invoked.
	Function IFunction `field:"required" json:"function" yaml:"function"`
}

