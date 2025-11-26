package previewawsconnectcampaignsv2mixins


// Contains predictive outbound mode configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   predictiveConfigProperty := &PredictiveConfigProperty{
//   	BandwidthAllocation: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-predictiveconfig.html
//
type CfnCampaignPropsMixin_PredictiveConfigProperty struct {
	// Bandwidth allocation for the predictive outbound mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-predictiveconfig.html#cfn-connectcampaignsv2-campaign-predictiveconfig-bandwidthallocation
	//
	BandwidthAllocation *float64 `field:"optional" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
}

