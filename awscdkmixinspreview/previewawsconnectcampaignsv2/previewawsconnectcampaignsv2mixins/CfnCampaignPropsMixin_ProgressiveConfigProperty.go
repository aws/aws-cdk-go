package previewawsconnectcampaignsv2mixins


// Contains the progressive outbound mode configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   progressiveConfigProperty := &ProgressiveConfigProperty{
//   	BandwidthAllocation: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-progressiveconfig.html
//
type CfnCampaignPropsMixin_ProgressiveConfigProperty struct {
	// Bandwidth allocation for the progressive outbound mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-progressiveconfig.html#cfn-connectcampaignsv2-campaign-progressiveconfig-bandwidthallocation
	//
	BandwidthAllocation *float64 `field:"optional" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
}

