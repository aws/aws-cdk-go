package previewawspinpointmixins


// Specifies the contents of a message that's sent through a custom channel to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   campaignCustomMessageProperty := &CampaignCustomMessageProperty{
//   	Data: jsii.String("data"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigncustommessage.html
//
type CfnCampaignPropsMixin_CampaignCustomMessageProperty struct {
	// The raw, JSON-formatted string to use as the payload for the message.
	//
	// The maximum size is 5 KB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigncustommessage.html#cfn-pinpoint-campaign-campaigncustommessage-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
}

