package previewawsconnectcampaignsv2mixins


// The event trigger of the campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventTriggerProperty := &EventTriggerProperty{
//   	CustomerProfilesDomainArn: jsii.String("customerProfilesDomainArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-eventtrigger.html
//
type CfnCampaignPropsMixin_EventTriggerProperty struct {
	// The Amazon Resource Name (ARN) of the Customer Profiles domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-eventtrigger.html#cfn-connectcampaignsv2-campaign-eventtrigger-customerprofilesdomainarn
	//
	CustomerProfilesDomainArn *string `field:"optional" json:"customerProfilesDomainArn" yaml:"customerProfilesDomainArn"`
}

