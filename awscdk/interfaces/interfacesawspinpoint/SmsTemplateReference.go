package interfacesawspinpoint


// A reference to a SmsTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smsTemplateReference := &SmsTemplateReference{
//   	SmsTemplateArn: jsii.String("smsTemplateArn"),
//   	SmsTemplateId: jsii.String("smsTemplateId"),
//   }
//
type SmsTemplateReference struct {
	// The ARN of the SmsTemplate resource.
	SmsTemplateArn *string `field:"required" json:"smsTemplateArn" yaml:"smsTemplateArn"`
	// The Id of the SmsTemplate resource.
	SmsTemplateId *string `field:"required" json:"smsTemplateId" yaml:"smsTemplateId"`
}

