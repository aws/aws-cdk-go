package mixinsawsconnectcampaignsv2


// The outbound configuration for email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   emailOutboundConfigProperty := &EmailOutboundConfigProperty{
//   	ConnectSourceEmailAddress: jsii.String("connectSourceEmailAddress"),
//   	SourceEmailAddressDisplayName: jsii.String("sourceEmailAddressDisplayName"),
//   	WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailoutboundconfig.html
//
type CfnCampaignPropsMixin_EmailOutboundConfigProperty struct {
	// The Amazon Connect source email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailoutboundconfig.html#cfn-connectcampaignsv2-campaign-emailoutboundconfig-connectsourceemailaddress
	//
	ConnectSourceEmailAddress *string `field:"optional" json:"connectSourceEmailAddress" yaml:"connectSourceEmailAddress"`
	// The display name for the Amazon Connect source email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailoutboundconfig.html#cfn-connectcampaignsv2-campaign-emailoutboundconfig-sourceemailaddressdisplayname
	//
	SourceEmailAddressDisplayName *string `field:"optional" json:"sourceEmailAddressDisplayName" yaml:"sourceEmailAddressDisplayName"`
	// The Amazon Resource Name (ARN) of the Amazon Q in Connect template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailoutboundconfig.html#cfn-connectcampaignsv2-campaign-emailoutboundconfig-wisdomtemplatearn
	//
	WisdomTemplateArn *string `field:"optional" json:"wisdomTemplateArn" yaml:"wisdomTemplateArn"`
}

