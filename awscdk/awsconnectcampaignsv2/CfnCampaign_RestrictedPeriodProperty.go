package awsconnectcampaignsv2


// Contains information about a restricted period.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restrictedPeriodProperty := &RestrictedPeriodProperty{
//   	EndDate: jsii.String("endDate"),
//   	StartDate: jsii.String("startDate"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html
//
type CfnCampaign_RestrictedPeriodProperty struct {
	// The end date of the restricted period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html#cfn-connectcampaignsv2-campaign-restrictedperiod-enddate
	//
	EndDate *string `field:"required" json:"endDate" yaml:"endDate"`
	// The start date of the restricted period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html#cfn-connectcampaignsv2-campaign-restrictedperiod-startdate
	//
	StartDate *string `field:"required" json:"startDate" yaml:"startDate"`
	// The name of the restricted period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html#cfn-connectcampaignsv2-campaign-restrictedperiod-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

