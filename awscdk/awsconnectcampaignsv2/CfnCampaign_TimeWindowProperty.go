package awsconnectcampaignsv2


// Contains information about a time window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//
//   	// the properties below are optional
//   	RestrictedPeriods: &RestrictedPeriodsProperty{
//   		RestrictedPeriodList: []interface{}{
//   			&RestrictedPeriodProperty{
//   				EndDate: jsii.String("endDate"),
//   				StartDate: jsii.String("startDate"),
//
//   				// the properties below are optional
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timewindow.html
//
type CfnCampaign_TimeWindowProperty struct {
	// The open hours configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timewindow.html#cfn-connectcampaignsv2-campaign-timewindow-openhours
	//
	OpenHours interface{} `field:"required" json:"openHours" yaml:"openHours"`
	// The restricted periods configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timewindow.html#cfn-connectcampaignsv2-campaign-timewindow-restrictedperiods
	//
	RestrictedPeriods interface{} `field:"optional" json:"restrictedPeriods" yaml:"restrictedPeriods"`
}

