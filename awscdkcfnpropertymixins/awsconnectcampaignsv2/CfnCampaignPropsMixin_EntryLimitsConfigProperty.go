package awsconnectcampaignsv2


// Entry limits config for a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   entryLimitsConfigProperty := &EntryLimitsConfigProperty{
//   	MaxEntryCount: jsii.Number(123),
//   	MinEntryInterval: jsii.String("minEntryInterval"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-entrylimitsconfig.html
//
type CfnCampaignPropsMixin_EntryLimitsConfigProperty struct {
	// Maximum number of entries per participant.
	//
	// 0 indicates unlimited entries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-entrylimitsconfig.html#cfn-connectcampaignsv2-campaign-entrylimitsconfig-maxentrycount
	//
	MaxEntryCount *float64 `field:"optional" json:"maxEntryCount" yaml:"maxEntryCount"`
	// Time duration in ISO 8601 format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-entrylimitsconfig.html#cfn-connectcampaignsv2-campaign-entrylimitsconfig-minentryinterval
	//
	MinEntryInterval *string `field:"optional" json:"minEntryInterval" yaml:"minEntryInterval"`
}

