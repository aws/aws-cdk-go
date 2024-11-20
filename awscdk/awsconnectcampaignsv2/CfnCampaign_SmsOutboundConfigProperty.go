package awsconnectcampaignsv2


// Default SMS outbound config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smsOutboundConfigProperty := &SmsOutboundConfigProperty{
//   	ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   	WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smsoutboundconfig.html
//
type CfnCampaign_SmsOutboundConfigProperty struct {
	// Arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smsoutboundconfig.html#cfn-connectcampaignsv2-campaign-smsoutboundconfig-connectsourcephonenumberarn
	//
	ConnectSourcePhoneNumberArn *string `field:"required" json:"connectSourcePhoneNumberArn" yaml:"connectSourcePhoneNumberArn"`
	// Arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smsoutboundconfig.html#cfn-connectcampaignsv2-campaign-smsoutboundconfig-wisdomtemplatearn
	//
	WisdomTemplateArn *string `field:"required" json:"wisdomTemplateArn" yaml:"wisdomTemplateArn"`
}

