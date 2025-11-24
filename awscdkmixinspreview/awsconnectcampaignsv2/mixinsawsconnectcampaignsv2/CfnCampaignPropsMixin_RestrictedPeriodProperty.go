package mixinsawsconnectcampaignsv2


// Contains information about a restricted period.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   restrictedPeriodProperty := &RestrictedPeriodProperty{
//   	EndDate: jsii.String("endDate"),
//   	Name: jsii.String("name"),
//   	StartDate: jsii.String("startDate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html
//
type CfnCampaignPropsMixin_RestrictedPeriodProperty struct {
	// The end date of the restricted period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html#cfn-connectcampaignsv2-campaign-restrictedperiod-enddate
	//
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// The name of the restricted period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html#cfn-connectcampaignsv2-campaign-restrictedperiod-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The start date of the restricted period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html#cfn-connectcampaignsv2-campaign-restrictedperiod-startdate
	//
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
}

