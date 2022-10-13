package awselasticloadbalancingv2


// A fixed response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fixedResponse := &fixedResponse{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	contentType: awscdk.Aws_elasticloadbalancingv2.contentType_TEXT_PLAIN,
//   	messageBody: jsii.String("messageBody"),
//   }
//
// Deprecated: superceded by `ListenerAction.fixedResponse()`.
type FixedResponse struct {
	// The HTTP response code (2XX, 4XX or 5XX).
	// Deprecated: superceded by `ListenerAction.fixedResponse()`.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The content type.
	// Deprecated: superceded by `ListenerAction.fixedResponse()`.
	ContentType ContentType `field:"optional" json:"contentType" yaml:"contentType"`
	// The message.
	// Deprecated: superceded by `ListenerAction.fixedResponse()`.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
}

