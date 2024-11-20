package awsconnectcampaignsv2


// Restricted period.
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
	// Date in ISO 8601 format, e.g. 2024-01-01.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html#cfn-connectcampaignsv2-campaign-restrictedperiod-enddate
	//
	EndDate *string `field:"required" json:"endDate" yaml:"endDate"`
	// Date in ISO 8601 format, e.g. 2024-01-01.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html#cfn-connectcampaignsv2-campaign-restrictedperiod-startdate
	//
	StartDate *string `field:"required" json:"startDate" yaml:"startDate"`
	// The name of a restricted period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiod.html#cfn-connectcampaignsv2-campaign-restrictedperiod-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

