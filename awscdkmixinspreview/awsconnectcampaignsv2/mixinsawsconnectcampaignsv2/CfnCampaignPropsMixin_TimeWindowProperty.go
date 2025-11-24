package mixinsawsconnectcampaignsv2


// Contains information about a time window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeWindowProperty := &TimeWindowProperty{
//   	OpenHours: &OpenHoursProperty{
//   		DailyHours: []interface{}{
//   			&DailyHourProperty{
//   				Key: jsii.String("key"),
//   				Value: []interface{}{
//   					&TimeRangeProperty{
//   						EndTime: jsii.String("endTime"),
//   						StartTime: jsii.String("startTime"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RestrictedPeriods: &RestrictedPeriodsProperty{
//   		RestrictedPeriodList: []interface{}{
//   			&RestrictedPeriodProperty{
//   				EndDate: jsii.String("endDate"),
//   				Name: jsii.String("name"),
//   				StartDate: jsii.String("startDate"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timewindow.html
//
type CfnCampaignPropsMixin_TimeWindowProperty struct {
	// The open hours configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timewindow.html#cfn-connectcampaignsv2-campaign-timewindow-openhours
	//
	OpenHours interface{} `field:"optional" json:"openHours" yaml:"openHours"`
	// The restricted periods configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timewindow.html#cfn-connectcampaignsv2-campaign-timewindow-restrictedperiods
	//
	RestrictedPeriods interface{} `field:"optional" json:"restrictedPeriods" yaml:"restrictedPeriods"`
}

