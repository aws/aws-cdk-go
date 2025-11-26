package previewawsconnectcampaignsv2mixins


// Contains information about communication limits.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   communicationLimitsProperty := &CommunicationLimitsProperty{
//   	CommunicationLimitList: []interface{}{
//   		&CommunicationLimitProperty{
//   			Frequency: jsii.Number(123),
//   			MaxCountPerRecipient: jsii.Number(123),
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimits.html
//
type CfnCampaignPropsMixin_CommunicationLimitsProperty struct {
	// The list of CommunicationLimits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimits.html#cfn-connectcampaignsv2-campaign-communicationlimits-communicationlimitlist
	//
	CommunicationLimitList interface{} `field:"optional" json:"communicationLimitList" yaml:"communicationLimitList"`
}

