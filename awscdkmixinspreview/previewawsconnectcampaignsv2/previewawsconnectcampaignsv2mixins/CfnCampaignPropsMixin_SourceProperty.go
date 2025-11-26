package previewawsconnectcampaignsv2mixins


// Contains source configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceProperty := &SourceProperty{
//   	CustomerProfilesSegmentArn: jsii.String("customerProfilesSegmentArn"),
//   	EventTrigger: &EventTriggerProperty{
//   		CustomerProfilesDomainArn: jsii.String("customerProfilesDomainArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-source.html
//
type CfnCampaignPropsMixin_SourceProperty struct {
	// The Amazon Resource Name (ARN) of the Customer Profiles segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-source.html#cfn-connectcampaignsv2-campaign-source-customerprofilessegmentarn
	//
	CustomerProfilesSegmentArn *string `field:"optional" json:"customerProfilesSegmentArn" yaml:"customerProfilesSegmentArn"`
	// The event trigger of the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-source.html#cfn-connectcampaignsv2-campaign-source-eventtrigger
	//
	EventTrigger interface{} `field:"optional" json:"eventTrigger" yaml:"eventTrigger"`
}

