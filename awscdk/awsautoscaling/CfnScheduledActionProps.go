package awsautoscaling


// Properties for defining a `CfnScheduledAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduledActionProps := &CfnScheduledActionProps{
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//
//   	// the properties below are optional
//   	DesiredCapacity: jsii.Number(123),
//   	EndTime: jsii.String("endTime"),
//   	MaxSize: jsii.Number(123),
//   	MinSize: jsii.Number(123),
//   	Recurrence: jsii.String("recurrence"),
//   	StartTime: jsii.String("startTime"),
//   	TimeZone: jsii.String("timeZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scheduledaction.html
//
type CfnScheduledActionProps struct {
	// The name of the Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scheduledaction.html#cfn-autoscaling-scheduledaction-autoscalinggroupname
	//
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.
	//
	// It can scale beyond this capacity if you add more scaling conditions.
	//
	// > You must specify at least one of the following properties: `MaxSize` , `MinSize` , or `DesiredCapacity` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scheduledaction.html#cfn-autoscaling-scheduledaction-desiredcapacity
	//
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// The date and time for the recurring schedule to end, in UTC.
	//
	// For example, `"2021-06-01T00:00:00Z"` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scheduledaction.html#cfn-autoscaling-scheduledaction-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The maximum size of the Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scheduledaction.html#cfn-autoscaling-scheduledaction-maxsize
	//
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum size of the Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scheduledaction.html#cfn-autoscaling-scheduledaction-minsize
	//
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// The recurring schedule for this action.
	//
	// This format consists of five fields separated by white spaces: [Minute] [Hour] [Day_of_Month] [Month_of_Year] [Day_of_Week]. The value must be in quotes (for example, `"30 0 1 1,6,12 *"` ). For more information about this format, see [Crontab](https://docs.aws.amazon.com/http://crontab.org) .
	//
	// When `StartTime` and `EndTime` are specified with `Recurrence` , they form the boundaries of when the recurring action starts and stops.
	//
	// Cron expressions use Universal Coordinated Time (UTC) by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scheduledaction.html#cfn-autoscaling-scheduledaction-recurrence
	//
	Recurrence *string `field:"optional" json:"recurrence" yaml:"recurrence"`
	// The date and time for this action to start, in YYYY-MM-DDThh:mm:ssZ format in UTC/GMT only and in quotes (for example, `"2021-06-01T00:00:00Z"` ).
	//
	// If you specify `Recurrence` and `StartTime` , Amazon EC2 Auto Scaling performs the action at this time, and then performs the action based on the specified recurrence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scheduledaction.html#cfn-autoscaling-scheduledaction-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Specifies the time zone for a cron expression.
	//
	// If a time zone is not provided, UTC is used by default.
	//
	// Valid values are the canonical names of the IANA time zones, derived from the IANA Time Zone Database (such as `Etc/GMT+9` or `Pacific/Tahiti` ). For more information, see [https://en.wikipedia.org/wiki/List_of_tz_database_time_zones](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scheduledaction.html#cfn-autoscaling-scheduledaction-timezone
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

