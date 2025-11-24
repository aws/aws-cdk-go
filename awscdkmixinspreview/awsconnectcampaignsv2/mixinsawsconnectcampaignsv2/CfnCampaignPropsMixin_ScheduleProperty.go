package mixinsawsconnectcampaignsv2


// Contains the schedule configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scheduleProperty := &ScheduleProperty{
//   	EndTime: jsii.String("endTime"),
//   	RefreshFrequency: jsii.String("refreshFrequency"),
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-schedule.html
//
type CfnCampaignPropsMixin_ScheduleProperty struct {
	// The end time of the schedule in UTC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-schedule.html#cfn-connectcampaignsv2-campaign-schedule-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The refresh frequency of the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-schedule.html#cfn-connectcampaignsv2-campaign-schedule-refreshfrequency
	//
	RefreshFrequency *string `field:"optional" json:"refreshFrequency" yaml:"refreshFrequency"`
	// The start time of the schedule in UTC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-schedule.html#cfn-connectcampaignsv2-campaign-schedule-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

