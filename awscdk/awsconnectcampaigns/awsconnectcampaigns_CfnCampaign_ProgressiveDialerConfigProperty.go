package awsconnectcampaigns


// Contains progressive dialer configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   progressiveDialerConfigProperty := &progressiveDialerConfigProperty{
//   	bandwidthAllocation: jsii.Number(123),
//   }
//
type CfnCampaign_ProgressiveDialerConfigProperty struct {
	// Bandwidth allocation for the progressive dialer.
	BandwidthAllocation *float64 `field:"required" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
}

