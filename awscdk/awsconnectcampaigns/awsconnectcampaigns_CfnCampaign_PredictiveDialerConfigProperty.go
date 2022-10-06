package awsconnectcampaigns


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveDialerConfigProperty := &predictiveDialerConfigProperty{
//   	bandwidthAllocation: jsii.Number(123),
//   }
//
type CfnCampaign_PredictiveDialerConfigProperty struct {
	// `CfnCampaign.PredictiveDialerConfigProperty.BandwidthAllocation`.
	BandwidthAllocation *float64 `field:"required" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
}

