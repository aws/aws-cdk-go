package previewawscomputeoptimizermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scheduleProperty := &ScheduleProperty{
//   	ExecutionWindowInMinutes: jsii.Number(123),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-schedule.html
//
type CfnAutomationRulePropsMixin_ScheduleProperty struct {
	// Execution window duration in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-schedule.html#cfn-computeoptimizer-automationrule-schedule-executionwindowinminutes
	//
	ExecutionWindowInMinutes *float64 `field:"optional" json:"executionWindowInMinutes" yaml:"executionWindowInMinutes"`
	// Schedule expression (e.g., cron or rate expression).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-schedule.html#cfn-computeoptimizer-automationrule-schedule-scheduleexpression
	//
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// IANA timezone identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-schedule.html#cfn-computeoptimizer-automationrule-schedule-scheduleexpressiontimezone
	//
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
}

