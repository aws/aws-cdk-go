package awsconnectcampaigns


// Contains progressive dialer configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   progressiveDialerConfigProperty := &ProgressiveDialerConfigProperty{
//   	BandwidthAllocation: jsii.Number(123),
//
//   	// the properties below are optional
//   	DialingCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-progressivedialerconfig.html
//
type CfnCampaign_ProgressiveDialerConfigProperty struct {
	// Bandwidth allocation for the progressive dialer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-progressivedialerconfig.html#cfn-connectcampaigns-campaign-progressivedialerconfig-bandwidthallocation
	//
	BandwidthAllocation *float64 `field:"required" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-progressivedialerconfig.html#cfn-connectcampaigns-campaign-progressivedialerconfig-dialingcapacity
	//
	DialingCapacity *float64 `field:"optional" json:"dialingCapacity" yaml:"dialingCapacity"`
}

