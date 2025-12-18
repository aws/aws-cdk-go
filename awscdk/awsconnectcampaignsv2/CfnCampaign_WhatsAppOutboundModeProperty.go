package awsconnectcampaignsv2


// Contains information about the WhatsApp outbound mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentlessConfig interface{}
//
//   whatsAppOutboundModeProperty := &WhatsAppOutboundModeProperty{
//   	AgentlessConfig: agentlessConfig,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-whatsappoutboundmode.html
//
type CfnCampaign_WhatsAppOutboundModeProperty struct {
	// Agentless config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-whatsappoutboundmode.html#cfn-connectcampaignsv2-campaign-whatsappoutboundmode-agentlessconfig
	//
	AgentlessConfig interface{} `field:"optional" json:"agentlessConfig" yaml:"agentlessConfig"`
}

