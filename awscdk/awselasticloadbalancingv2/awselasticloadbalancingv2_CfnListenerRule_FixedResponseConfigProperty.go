package awselasticloadbalancingv2


// Specifies information required when returning a custom HTTP response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fixedResponseConfigProperty := &fixedResponseConfigProperty{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	contentType: jsii.String("contentType"),
//   	messageBody: jsii.String("messageBody"),
//   }
//
type CfnListenerRule_FixedResponseConfigProperty struct {
	// The HTTP response code (2XX, 4XX, or 5XX).
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The content type.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The message.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
}

