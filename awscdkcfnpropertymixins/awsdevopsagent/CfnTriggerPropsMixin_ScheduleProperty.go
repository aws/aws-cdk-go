package awsdevopsagent


// Schedule configuration for a time-based trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   scheduleProperty := &ScheduleProperty{
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-trigger-schedule.html
//
type CfnTriggerPropsMixin_ScheduleProperty struct {
	// A cron or rate expression that defines when the trigger fires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-trigger-schedule.html#cfn-devopsagent-trigger-schedule-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

