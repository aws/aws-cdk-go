package awssecretsmanager


// The rotation schedule and window.
//
// We recommend you use `ScheduleExpression` to set a cron or rate expression for the schedule and `Duration` to set the length of the rotation window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rotationRulesProperty := &RotationRulesProperty{
//   	AutomaticallyAfterDays: jsii.Number(123),
//   	Duration: jsii.String("duration"),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-rotationrules.html
//
type CfnRotationSchedule_RotationRulesProperty struct {
	// The number of days between automatic scheduled rotations of the secret.
	//
	// You can use this value to check that your secret meets your compliance guidelines for how often secrets must be rotated.
	//
	// In `DescribeSecret` and `ListSecrets` , this value is calculated from the rotation schedule after every successful rotation. In `RotateSecret` , you can set the rotation schedule in `RotationRules` with `AutomaticallyAfterDays` or `ScheduleExpression` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-rotationrules.html#cfn-secretsmanager-rotationschedule-rotationrules-automaticallyafterdays
	//
	AutomaticallyAfterDays *float64 `field:"optional" json:"automaticallyAfterDays" yaml:"automaticallyAfterDays"`
	// The length of the rotation window in hours, for example `3h` for a three hour window.
	//
	// Secrets Manager rotates your secret at any time during this window. The window must not extend into the next rotation window or the next UTC day. The window starts according to the `ScheduleExpression` . If you don't specify a `Duration` , for a `ScheduleExpression` in hours, the window automatically closes after one hour. For a `ScheduleExpression` in days, the window automatically closes at the end of the UTC day. For more information, including examples, see [Schedule expressions in Secrets Manager rotation](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_schedule.html) in the *Secrets Manager Users Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-rotationrules.html#cfn-secretsmanager-rotationschedule-rotationrules-duration
	//
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// A `cron()` or `rate()` expression that defines the schedule for rotating your secret.
	//
	// Secrets Manager rotation schedules use UTC time zone. Secrets Manager rotates your secret any time during a rotation window.
	//
	// Secrets Manager `rate()` expressions represent the interval in hours or days that you want to rotate your secret, for example `rate(12 hours)` or `rate(10 days)` . You can rotate a secret as often as every four hours. If you use a `rate()` expression, the rotation window starts at midnight. For a rate in hours, the default rotation window closes after one hour. For a rate in days, the default rotation window closes at the end of the day. You can set the `Duration` to change the rotation window. The rotation window must not extend into the next UTC day or into the next rotation window.
	//
	// You can use a `cron()` expression to create a rotation schedule that is more detailed than a rotation interval. For more information, including examples, see [Schedule expressions in Secrets Manager rotation](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_schedule.html) in the *Secrets Manager Users Guide* . For a cron expression that represents a schedule in hours, the default rotation window closes after one hour. For a cron expression that represents a schedule in days, the default rotation window closes at the end of the day. You can set the `Duration` to change the rotation window. The rotation window must not extend into the next UTC day or into the next rotation window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-rotationrules.html#cfn-secretsmanager-rotationschedule-rotationrules-scheduleexpression
	//
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

