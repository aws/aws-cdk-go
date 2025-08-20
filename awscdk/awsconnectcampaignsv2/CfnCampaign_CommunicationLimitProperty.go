package awsconnectcampaignsv2


// Contains information about a communication limit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   communicationLimitProperty := &CommunicationLimitProperty{
//   	Frequency: jsii.Number(123),
//   	MaxCountPerRecipient: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimit.html
//
type CfnCampaign_CommunicationLimitProperty struct {
	// The frequency of communication limit evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimit.html#cfn-connectcampaignsv2-campaign-communicationlimit-frequency
	//
	Frequency *float64 `field:"required" json:"frequency" yaml:"frequency"`
	// The maximum outreaching count for each recipient.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimit.html#cfn-connectcampaignsv2-campaign-communicationlimit-maxcountperrecipient
	//
	MaxCountPerRecipient *float64 `field:"required" json:"maxCountPerRecipient" yaml:"maxCountPerRecipient"`
	// The unit of communication limit evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-communicationlimit.html#cfn-connectcampaignsv2-campaign-communicationlimit-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}

