package awstimestream


// Configuration of the schedule of the query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleConfigurationProperty := &scheduleConfigurationProperty{
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   }
//
type CfnScheduledQuery_ScheduleConfigurationProperty struct {
	// An expression that denotes when to trigger the scheduled query run.
	//
	// This can be a cron expression or a rate expression.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}

