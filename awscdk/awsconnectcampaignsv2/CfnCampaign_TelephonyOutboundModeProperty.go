package awsconnectcampaignsv2


// Contains information about telephony outbound mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentlessConfig interface{}
//
//   telephonyOutboundModeProperty := &TelephonyOutboundModeProperty{
//   	AgentlessConfig: agentlessConfig,
//   	PredictiveConfig: &PredictiveConfigProperty{
//   		BandwidthAllocation: jsii.Number(123),
//   	},
//   	PreviewConfig: &PreviewConfigProperty{
//   		BandwidthAllocation: jsii.Number(123),
//   		TimeoutConfig: &TimeoutConfigProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		AgentActions: []*string{
//   			jsii.String("agentActions"),
//   		},
//   	},
//   	ProgressiveConfig: &ProgressiveConfigProperty{
//   		BandwidthAllocation: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundmode.html
//
type CfnCampaign_TelephonyOutboundModeProperty struct {
	// The agentless outbound mode configuration for telephony.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundmode.html#cfn-connectcampaignsv2-campaign-telephonyoutboundmode-agentlessconfig
	//
	AgentlessConfig interface{} `field:"optional" json:"agentlessConfig" yaml:"agentlessConfig"`
	// Contains predictive outbound mode configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundmode.html#cfn-connectcampaignsv2-campaign-telephonyoutboundmode-predictiveconfig
	//
	PredictiveConfig interface{} `field:"optional" json:"predictiveConfig" yaml:"predictiveConfig"`
	// Preview config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundmode.html#cfn-connectcampaignsv2-campaign-telephonyoutboundmode-previewconfig
	//
	PreviewConfig interface{} `field:"optional" json:"previewConfig" yaml:"previewConfig"`
	// Contains progressive telephony outbound mode configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundmode.html#cfn-connectcampaignsv2-campaign-telephonyoutboundmode-progressiveconfig
	//
	ProgressiveConfig interface{} `field:"optional" json:"progressiveConfig" yaml:"progressiveConfig"`
}

