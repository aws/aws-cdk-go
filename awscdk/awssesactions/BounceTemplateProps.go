package awssesactions


// Construction properties for a BounceTemplate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bounceTemplateProps := &BounceTemplateProps{
//   	Message: jsii.String("message"),
//   	SmtpReplyCode: jsii.String("smtpReplyCode"),
//
//   	// the properties below are optional
//   	StatusCode: jsii.String("statusCode"),
//   }
//
// Experimental.
type BounceTemplateProps struct {
	// Human-readable text to include in the bounce message.
	// Experimental.
	Message *string `field:"required" json:"message" yaml:"message"`
	// The SMTP reply code, as defined by RFC 5321.
	// See: https://tools.ietf.org/html/rfc5321
	//
	// Experimental.
	SmtpReplyCode *string `field:"required" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// The SMTP enhanced status code, as defined by RFC 3463.
	// See: https://tools.ietf.org/html/rfc3463
	//
	// Experimental.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

