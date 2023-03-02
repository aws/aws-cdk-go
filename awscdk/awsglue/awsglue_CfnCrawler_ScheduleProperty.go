package awsglue


// A scheduling object using a `cron` statement to schedule an event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleProperty := &scheduleProperty{
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   }
//
type CfnCrawler_ScheduleProperty struct {
	// A `cron` expression used to specify the schedule.
	//
	// For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html) . For example, to run something every day at 12:15 UTC, specify `cron(15 12 * * ? *)` .
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

