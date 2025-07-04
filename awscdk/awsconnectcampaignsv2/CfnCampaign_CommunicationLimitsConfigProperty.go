package awsconnectcampaignsv2


// Contains the communication limits configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   communicationLimitsConfigProperty := &CommunicationLimitsConfigProperty{
//   	AllChannelsSubtypes: &CommunicationLimitsProperty{
//   		CommunicationLimitList: []interface{}{
//   			&CommunicationLimitProperty{
//   				Frequency: jsii.Number(123),
//   				MaxCountPerRecipient: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   		},
//   	},
//   	InstanceLimitsHandling: jsii.String("instanceLimitsHandling"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimitsconfig.html
//
type CfnCampaign_CommunicationLimitsConfigProperty struct {
	// The CommunicationLimits that apply to all channel subtypes defined in an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimitsconfig.html#cfn-connectcampaignsv2-campaign-communicationlimitsconfig-allchannelssubtypes
	//
	AllChannelsSubtypes interface{} `field:"optional" json:"allChannelsSubtypes" yaml:"allChannelsSubtypes"`
	// Opt-in or Opt-out from instance-level limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimitsconfig.html#cfn-connectcampaignsv2-campaign-communicationlimitsconfig-instancelimitshandling
	//
	InstanceLimitsHandling *string `field:"optional" json:"instanceLimitsHandling" yaml:"instanceLimitsHandling"`
}

