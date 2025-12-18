package awsconnectcampaignsv2


// The outbound configuration for WhatsApp.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   whatsAppOutboundConfigProperty := &WhatsAppOutboundConfigProperty{
//   	ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   	WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-whatsappoutboundconfig.html
//
type CfnCampaign_WhatsAppOutboundConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Connect source WhatsApp phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-whatsappoutboundconfig.html#cfn-connectcampaignsv2-campaign-whatsappoutboundconfig-connectsourcephonenumberarn
	//
	ConnectSourcePhoneNumberArn *string `field:"required" json:"connectSourcePhoneNumberArn" yaml:"connectSourcePhoneNumberArn"`
	// The Amazon Resource Name (ARN) of the Amazon Q in Connect template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-whatsappoutboundconfig.html#cfn-connectcampaignsv2-campaign-whatsappoutboundconfig-wisdomtemplatearn
	//
	WisdomTemplateArn *string `field:"required" json:"wisdomTemplateArn" yaml:"wisdomTemplateArn"`
}

