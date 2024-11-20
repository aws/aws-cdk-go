package awsconnectcampaignsv2


// Communication limits config.
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimitsconfig.html
//
type CfnCampaign_CommunicationLimitsConfigProperty struct {
	// Communication limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimitsconfig.html#cfn-connectcampaignsv2-campaign-communicationlimitsconfig-allchannelssubtypes
	//
	AllChannelsSubtypes interface{} `field:"optional" json:"allChannelsSubtypes" yaml:"allChannelsSubtypes"`
}

