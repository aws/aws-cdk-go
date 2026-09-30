package awsbcm


// Defines when and how often a scheduled report runs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   scheduleConfigProperty := &ScheduleConfigProperty{
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	ScheduleExpressionTimeZone: jsii.String("scheduleExpressionTimeZone"),
//   	SchedulePeriod: &SchedulePeriodProperty{
//   		EndTime: jsii.String("endTime"),
//   		StartTime: jsii.String("startTime"),
//   	},
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-scheduleconfig.html
//
type CfnScheduledReportPropsMixin_ScheduleConfigProperty struct {
	// The schedule expression that specifies when to trigger the scheduled report run.
	//
	// This value must be a cron expression consisting of six fields separated by white spaces: cron(minutes hours day_of_month month day_of_week year).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-scheduleconfig.html#cfn-bcm-scheduledreport-scheduleconfig-scheduleexpression
	//
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The time zone for the schedule expression, for example, UTC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-scheduleconfig.html#cfn-bcm-scheduledreport-scheduleconfig-scheduleexpressiontimezone
	//
	ScheduleExpressionTimeZone *string `field:"optional" json:"scheduleExpressionTimeZone" yaml:"scheduleExpressionTimeZone"`
	// The window during which the schedule is active.
	//
	// When omitted the service defaults it. EndTime must be within three years of the time of the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-scheduleconfig.html#cfn-bcm-scheduledreport-scheduleconfig-scheduleperiod
	//
	SchedulePeriod interface{} `field:"optional" json:"schedulePeriod" yaml:"schedulePeriod"`
	// Whether the schedule is active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-scheduleconfig.html#cfn-bcm-scheduledreport-scheduleconfig-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

