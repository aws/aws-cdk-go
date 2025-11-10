package awsconnectcampaignsv2


// Contains preview outbound mode configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   previewConfigProperty := &PreviewConfigProperty{
//   	BandwidthAllocation: jsii.Number(123),
//   	TimeoutConfig: &TimeoutConfigProperty{
//   		DurationInSeconds: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	AgentActions: []*string{
//   		jsii.String("agentActions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-previewconfig.html
//
type CfnCampaign_PreviewConfigProperty struct {
	// Bandwidth allocation for the preview outbound mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-previewconfig.html#cfn-connectcampaignsv2-campaign-previewconfig-bandwidthallocation
	//
	BandwidthAllocation *float64 `field:"required" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
	// Countdown timer configuration for preview outbound mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-previewconfig.html#cfn-connectcampaignsv2-campaign-previewconfig-timeoutconfig
	//
	TimeoutConfig interface{} `field:"required" json:"timeoutConfig" yaml:"timeoutConfig"`
	// Agent actions for the preview outbound mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-previewconfig.html#cfn-connectcampaignsv2-campaign-previewconfig-agentactions
	//
	AgentActions *[]*string `field:"optional" json:"agentActions" yaml:"agentActions"`
}

