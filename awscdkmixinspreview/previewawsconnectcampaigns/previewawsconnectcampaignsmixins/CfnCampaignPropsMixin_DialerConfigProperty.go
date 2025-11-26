package previewawsconnectcampaignsmixins


// Contains dialer configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dialerConfigProperty := &DialerConfigProperty{
//   	AgentlessDialerConfig: &AgentlessDialerConfigProperty{
//   		DialingCapacity: jsii.Number(123),
//   	},
//   	PredictiveDialerConfig: &PredictiveDialerConfigProperty{
//   		BandwidthAllocation: jsii.Number(123),
//   		DialingCapacity: jsii.Number(123),
//   	},
//   	ProgressiveDialerConfig: &ProgressiveDialerConfigProperty{
//   		BandwidthAllocation: jsii.Number(123),
//   		DialingCapacity: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-dialerconfig.html
//
type CfnCampaignPropsMixin_DialerConfigProperty struct {
	// The configuration of the agentless dialer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-dialerconfig.html#cfn-connectcampaigns-campaign-dialerconfig-agentlessdialerconfig
	//
	AgentlessDialerConfig interface{} `field:"optional" json:"agentlessDialerConfig" yaml:"agentlessDialerConfig"`
	// The configuration of the predictive dialer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-dialerconfig.html#cfn-connectcampaigns-campaign-dialerconfig-predictivedialerconfig
	//
	PredictiveDialerConfig interface{} `field:"optional" json:"predictiveDialerConfig" yaml:"predictiveDialerConfig"`
	// The configuration of the progressive dialer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-dialerconfig.html#cfn-connectcampaigns-campaign-dialerconfig-progressivedialerconfig
	//
	ProgressiveDialerConfig interface{} `field:"optional" json:"progressiveDialerConfig" yaml:"progressiveDialerConfig"`
}

