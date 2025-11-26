package previewawsconnectcampaignsv2mixins


// Contains preview outbound mode configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   previewConfigProperty := &PreviewConfigProperty{
//   	AgentActions: []*string{
//   		jsii.String("agentActions"),
//   	},
//   	BandwidthAllocation: jsii.Number(123),
//   	TimeoutConfig: &TimeoutConfigProperty{
//   		DurationInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-previewconfig.html
//
type CfnCampaignPropsMixin_PreviewConfigProperty struct {
	// Agent actions for the preview outbound mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-previewconfig.html#cfn-connectcampaignsv2-campaign-previewconfig-agentactions
	//
	AgentActions *[]*string `field:"optional" json:"agentActions" yaml:"agentActions"`
	// Bandwidth allocation for the preview outbound mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-previewconfig.html#cfn-connectcampaignsv2-campaign-previewconfig-bandwidthallocation
	//
	BandwidthAllocation *float64 `field:"optional" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
	// Countdown timer configuration for preview outbound mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-previewconfig.html#cfn-connectcampaignsv2-campaign-previewconfig-timeoutconfig
	//
	TimeoutConfig interface{} `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

