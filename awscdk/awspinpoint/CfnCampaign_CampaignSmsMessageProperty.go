package awspinpoint


// Specifies the content and settings for an SMS message that's sent to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignSmsMessageProperty := &CampaignSmsMessageProperty{
//   	Body: jsii.String("body"),
//   	EntityId: jsii.String("entityId"),
//   	MessageType: jsii.String("messageType"),
//   	OriginationNumber: jsii.String("originationNumber"),
//   	SenderId: jsii.String("senderId"),
//   	TemplateId: jsii.String("templateId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html
//
type CfnCampaign_CampaignSmsMessageProperty struct {
	// The body of the SMS message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The entity ID or Principal Entity (PE) id received from the regulatory body for sending SMS in your country.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-entityid
	//
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// The SMS message type.
	//
	// Valid values are `TRANSACTIONAL` (for messages that are critical or time-sensitive, such as a one-time passwords) and `PROMOTIONAL` (for messsages that aren't critical or time-sensitive, such as marketing messages).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-messagetype
	//
	MessageType *string `field:"optional" json:"messageType" yaml:"messageType"`
	// The long code to send the SMS message from.
	//
	// This value should be one of the dedicated long codes that's assigned to your AWS account. Although it isn't required, we recommend that you specify the long code using an E.164 format to ensure prompt and accurate delivery of the message. For example, +12065550100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-originationnumber
	//
	OriginationNumber *string `field:"optional" json:"originationNumber" yaml:"originationNumber"`
	// The alphabetic Sender ID to display as the sender of the message on a recipient's device.
	//
	// Support for sender IDs varies by country or region. To specify a phone number as the sender, omit this parameter and use `OriginationNumber` instead. For more information about support for Sender ID by country, see the [Amazon Pinpoint User Guide](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-countries.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-senderid
	//
	SenderId *string `field:"optional" json:"senderId" yaml:"senderId"`
	// The template ID received from the regulatory body for sending SMS in your country.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-templateid
	//
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
}

