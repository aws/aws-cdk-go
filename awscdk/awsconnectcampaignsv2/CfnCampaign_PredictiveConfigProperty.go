package awsconnectcampaignsv2


// Predictive config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveConfigProperty := &PredictiveConfigProperty{
//   	BandwidthAllocation: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-predictiveconfig.html
//
type CfnCampaign_PredictiveConfigProperty struct {
	// The bandwidth allocation of a queue resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-predictiveconfig.html#cfn-connectcampaignsv2-campaign-predictiveconfig-bandwidthallocation
	//
	BandwidthAllocation *float64 `field:"required" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
}

