package awsconnectcampaignsv2


// Open Hours config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openHoursProperty := &OpenHoursProperty{
//   	DailyHours: []interface{}{
//   		&DailyHourProperty{
//   			Key: jsii.String("key"),
//   			Value: []interface{}{
//   				&TimeRangeProperty{
//   					EndTime: jsii.String("endTime"),
//   					StartTime: jsii.String("startTime"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-openhours.html
//
type CfnCampaign_OpenHoursProperty struct {
	// Daily Hours map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-openhours.html#cfn-connectcampaignsv2-campaign-openhours-dailyhours
	//
	DailyHours interface{} `field:"required" json:"dailyHours" yaml:"dailyHours"`
}

