package awsconnectcampaignsv2


// The daily hours configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dailyHourProperty := &DailyHourProperty{
//   	Key: jsii.String("key"),
//   	Value: []interface{}{
//   		&TimeRangeProperty{
//   			EndTime: jsii.String("endTime"),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-dailyhour.html
//
type CfnCampaign_DailyHourProperty struct {
	// The key for DailyHour.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-dailyhour.html#cfn-connectcampaignsv2-campaign-dailyhour-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for DailyHour.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-dailyhour.html#cfn-connectcampaignsv2-campaign-dailyhour-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

