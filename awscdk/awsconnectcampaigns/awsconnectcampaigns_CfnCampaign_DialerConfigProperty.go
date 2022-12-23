package awsconnectcampaigns


// Contains dialer configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dialerConfigProperty := &dialerConfigProperty{
//   	predictiveDialerConfig: &predictiveDialerConfigProperty{
//   		bandwidthAllocation: jsii.Number(123),
//   	},
//   	progressiveDialerConfig: &progressiveDialerConfigProperty{
//   		bandwidthAllocation: jsii.Number(123),
//   	},
//   }
//
type CfnCampaign_DialerConfigProperty struct {
	// The configuration of the predictive dialer.
	PredictiveDialerConfig interface{} `field:"optional" json:"predictiveDialerConfig" yaml:"predictiveDialerConfig"`
	// The configuration of the progressive dialer.
	ProgressiveDialerConfig interface{} `field:"optional" json:"progressiveDialerConfig" yaml:"progressiveDialerConfig"`
}

