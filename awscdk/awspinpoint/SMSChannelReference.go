package awspinpoint


// A reference to a SMSChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sMSChannelReference := &SMSChannelReference{
//   	SmsChannelId: jsii.String("smsChannelId"),
//   }
//
type SMSChannelReference struct {
	// The Id of the SMSChannel resource.
	SmsChannelId *string `field:"required" json:"smsChannelId" yaml:"smsChannelId"`
}

