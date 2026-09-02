package awsdevopsagent


// The condition that causes the trigger to fire.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionProperty := &ConditionProperty{
//   	Schedule: &ScheduleProperty{
//   		Expression: jsii.String("expression"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-trigger-condition.html
//
type CfnTrigger_ConditionProperty struct {
	// Schedule configuration for a time-based trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-trigger-condition.html#cfn-devopsagent-trigger-condition-schedule
	//
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
}

