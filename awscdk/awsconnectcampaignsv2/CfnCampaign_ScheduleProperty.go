package awsconnectcampaignsv2


// Campaign schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleProperty := &ScheduleProperty{
//   	EndTime: jsii.String("endTime"),
//   	StartTime: jsii.String("startTime"),
//
//   	// the properties below are optional
//   	RefreshFrequency: jsii.String("refreshFrequency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-schedule.html
//
type CfnCampaign_ScheduleProperty struct {
	// Timestamp with no UTC offset or timezone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-schedule.html#cfn-connectcampaignsv2-campaign-schedule-endtime
	//
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Timestamp with no UTC offset or timezone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-schedule.html#cfn-connectcampaignsv2-campaign-schedule-starttime
	//
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Time duration in ISO 8601 format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-schedule.html#cfn-connectcampaignsv2-campaign-schedule-refreshfrequency
	//
	RefreshFrequency *string `field:"optional" json:"refreshFrequency" yaml:"refreshFrequency"`
}

