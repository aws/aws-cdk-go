package previewawsconnectcampaignsv2mixins


// Contains information about email outbound mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var agentlessConfig interface{}
//
//   emailOutboundModeProperty := &EmailOutboundModeProperty{
//   	AgentlessConfig: agentlessConfig,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailoutboundmode.html
//
type CfnCampaignPropsMixin_EmailOutboundModeProperty struct {
	// The agentless outbound mode configuration for email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailoutboundmode.html#cfn-connectcampaignsv2-campaign-emailoutboundmode-agentlessconfig
	//
	AgentlessConfig interface{} `field:"optional" json:"agentlessConfig" yaml:"agentlessConfig"`
}

