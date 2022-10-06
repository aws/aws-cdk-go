package awsconnectcampaigns


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
	// `CfnCampaign.DialerConfigProperty.PredictiveDialerConfig`.
	PredictiveDialerConfig interface{} `field:"optional" json:"predictiveDialerConfig" yaml:"predictiveDialerConfig"`
	// `CfnCampaign.DialerConfigProperty.ProgressiveDialerConfig`.
	ProgressiveDialerConfig interface{} `field:"optional" json:"progressiveDialerConfig" yaml:"progressiveDialerConfig"`
}

