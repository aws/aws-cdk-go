package previewawsconnectcampaignsv2mixins


// Contains information about the SMS outbound mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var agentlessConfig interface{}
//
//   smsOutboundModeProperty := &SmsOutboundModeProperty{
//   	AgentlessConfig: agentlessConfig,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smsoutboundmode.html
//
type CfnCampaignPropsMixin_SmsOutboundModeProperty struct {
	// Contains agentless outbound mode configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smsoutboundmode.html#cfn-connectcampaignsv2-campaign-smsoutboundmode-agentlessconfig
	//
	AgentlessConfig interface{} `field:"optional" json:"agentlessConfig" yaml:"agentlessConfig"`
}

