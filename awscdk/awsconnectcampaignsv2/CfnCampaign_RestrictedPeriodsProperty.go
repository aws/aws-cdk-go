package awsconnectcampaignsv2


// Contains information about restricted periods.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restrictedPeriodsProperty := &RestrictedPeriodsProperty{
//   	RestrictedPeriodList: []interface{}{
//   		&RestrictedPeriodProperty{
//   			EndDate: jsii.String("endDate"),
//   			StartDate: jsii.String("startDate"),
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiods.html
//
type CfnCampaign_RestrictedPeriodsProperty struct {
	// The restricted period list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-restrictedperiods.html#cfn-connectcampaignsv2-campaign-restrictedperiods-restrictedperiodlist
	//
	RestrictedPeriodList interface{} `field:"required" json:"restrictedPeriodList" yaml:"restrictedPeriodList"`
}

