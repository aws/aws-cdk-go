package awspinpoint


// Specifies the content and settings for an SMS message that's sent to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignSmsMessageProperty := &campaignSmsMessageProperty{
//   	body: jsii.String("body"),
//   	entityId: jsii.String("entityId"),
//   	messageType: jsii.String("messageType"),
//   	originationNumber: jsii.String("originationNumber"),
//   	senderId: jsii.String("senderId"),
//   	templateId: jsii.String("templateId"),
//   }
//
type CfnCampaign_CampaignSmsMessageProperty struct {
	// The body of the SMS message.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The entity ID or Principal Entity (PE) id received from the regulatory body for sending SMS in your country.
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// The SMS message type.
	//
	// Valid values are `TRANSACTIONAL` (for messages that are critical or time-sensitive, such as a one-time passwords) and `PROMOTIONAL` (for messsages that aren't critical or time-sensitive, such as marketing messages).
	MessageType *string `field:"optional" json:"messageType" yaml:"messageType"`
	// The long code to send the SMS message from.
	//
	// This value should be one of the dedicated long codes that's assigned to your AWS account. Although it isn't required, we recommend that you specify the long code using an E.164 format to ensure prompt and accurate delivery of the message. For example, +12065550100.
	OriginationNumber *string `field:"optional" json:"originationNumber" yaml:"originationNumber"`
	// The alphabetic Sender ID to display as the sender of the message on a recipient's device.
	//
	// Support for sender IDs varies by country or region. To specify a phone number as the sender, omit this parameter and use `OriginationNumber` instead. For more information about support for Sender ID by country, see the [Amazon Pinpoint User Guide](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-countries.html) .
	SenderId *string `field:"optional" json:"senderId" yaml:"senderId"`
	// The template ID received from the regulatory body for sending SMS in your country.
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
}

