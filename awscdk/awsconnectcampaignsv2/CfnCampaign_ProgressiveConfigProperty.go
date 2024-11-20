package awsconnectcampaignsv2


// Progressive config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   progressiveConfigProperty := &ProgressiveConfigProperty{
//   	BandwidthAllocation: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-progressiveconfig.html
//
type CfnCampaign_ProgressiveConfigProperty struct {
	// The bandwidth allocation of a queue resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-progressiveconfig.html#cfn-connectcampaignsv2-campaign-progressiveconfig-bandwidthallocation
	//
	BandwidthAllocation *float64 `field:"required" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
}

