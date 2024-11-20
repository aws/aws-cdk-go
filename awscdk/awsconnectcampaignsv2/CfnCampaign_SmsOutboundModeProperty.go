package awsconnectcampaignsv2


// SMS Outbound Mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentlessConfig interface{}
//
//   smsOutboundModeProperty := &SmsOutboundModeProperty{
//   	AgentlessConfig: agentlessConfig,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smsoutboundmode.html
//
type CfnCampaign_SmsOutboundModeProperty struct {
	// Agentless config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smsoutboundmode.html#cfn-connectcampaignsv2-campaign-smsoutboundmode-agentlessconfig
	//
	AgentlessConfig interface{} `field:"optional" json:"agentlessConfig" yaml:"agentlessConfig"`
}

