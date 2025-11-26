package previewawsconnectcampaignsmixins


// Contains agentless dialer configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   agentlessDialerConfigProperty := &AgentlessDialerConfigProperty{
//   	DialingCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-agentlessdialerconfig.html
//
type CfnCampaignPropsMixin_AgentlessDialerConfigProperty struct {
	// The allocation of dialing capacity between multiple active campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-agentlessdialerconfig.html#cfn-connectcampaigns-campaign-agentlessdialerconfig-dialingcapacity
	//
	DialingCapacity *float64 `field:"optional" json:"dialingCapacity" yaml:"dialingCapacity"`
}

