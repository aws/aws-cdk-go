package awsconnectcampaignsv2


// The outbound configuration for email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailOutboundConfigProperty := &EmailOutboundConfigProperty{
//   	ConnectSourceEmailAddress: jsii.String("connectSourceEmailAddress"),
//   	WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//
//   	// the properties below are optional
//   	SourceEmailAddressDisplayName: jsii.String("sourceEmailAddressDisplayName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailoutboundconfig.html
//
type CfnCampaign_EmailOutboundConfigProperty struct {
	// The Amazon Connect source email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailoutboundconfig.html#cfn-connectcampaignsv2-campaign-emailoutboundconfig-connectsourceemailaddress
	//
	ConnectSourceEmailAddress *string `field:"required" json:"connectSourceEmailAddress" yaml:"connectSourceEmailAddress"`
	// The Amazon Resource Name (ARN) of the Amazon Q in Connect template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailoutboundconfig.html#cfn-connectcampaignsv2-campaign-emailoutboundconfig-wisdomtemplatearn
	//
	WisdomTemplateArn *string `field:"required" json:"wisdomTemplateArn" yaml:"wisdomTemplateArn"`
	// The display name for the Amazon Connect source email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailoutboundconfig.html#cfn-connectcampaignsv2-campaign-emailoutboundconfig-sourceemailaddressdisplayname
	//
	SourceEmailAddressDisplayName *string `field:"optional" json:"sourceEmailAddressDisplayName" yaml:"sourceEmailAddressDisplayName"`
}

