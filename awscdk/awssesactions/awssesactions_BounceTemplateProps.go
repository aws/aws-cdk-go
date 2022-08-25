package awssesactions


// Construction properties for a BounceTemplate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bounceTemplateProps := &bounceTemplateProps{
//   	message: jsii.String("message"),
//   	smtpReplyCode: jsii.String("smtpReplyCode"),
//
//   	// the properties below are optional
//   	statusCode: jsii.String("statusCode"),
//   }
//
type BounceTemplateProps struct {
	// Human-readable text to include in the bounce message.
	Message *string `field:"required" json:"message" yaml:"message"`
	// The SMTP reply code, as defined by RFC 5321.
	// See: https://tools.ietf.org/html/rfc5321
	//
	SmtpReplyCode *string `field:"required" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// The SMTP enhanced status code, as defined by RFC 3463.
	// See: https://tools.ietf.org/html/rfc3463
	//
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

