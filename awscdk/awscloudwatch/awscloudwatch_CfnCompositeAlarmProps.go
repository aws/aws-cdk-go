package awscloudwatch


// Properties for defining a `CfnCompositeAlarm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCompositeAlarmProps := &cfnCompositeAlarmProps{
//   	alarmRule: jsii.String("alarmRule"),
//
//   	// the properties below are optional
//   	actionsEnabled: jsii.Boolean(false),
//   	actionsSuppressor: jsii.String("actionsSuppressor"),
//   	actionsSuppressorExtensionPeriod: jsii.Number(123),
//   	actionsSuppressorWaitPeriod: jsii.Number(123),
//   	alarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	alarmDescription: jsii.String("alarmDescription"),
//   	alarmName: jsii.String("alarmName"),
//   	insufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	okActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   }
//
type CfnCompositeAlarmProps struct {
	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state.
	//
	// For each alarm that you reference, you designate a function that specifies whether that alarm needs to be in ALARM state, OK state, or INSUFFICIENT_DATA state. You can use operators (AND, OR and NOT) to combine multiple functions in a single expression. You can use parenthesis to logically group the functions in your expression.
	//
	// You can use either alarm names or ARNs to reference the other alarms that are to be evaluated.
	//
	// Functions can include the following:
	//
	// - ALARM("alarm-name or alarm-ARN") is TRUE if the named alarm is in ALARM state.
	// - OK("alarm-name or alarm-ARN") is TRUE if the named alarm is in OK state.
	// - INSUFFICIENT_DATA("alarm-name or alarm-ARN") is TRUE if the named alarm is in INSUFFICIENT_DATA state.
	// - TRUE always evaluates to TRUE.
	// - FALSE always evaluates to FALSE.
	//
	// TRUE and FALSE are useful for testing a complex AlarmRule structure, and for testing your alarm actions.
	//
	// For more information about `AlarmRule` syntax, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	AlarmRule *string `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm.
	//
	// The default is TRUE.
	ActionsEnabled interface{} `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// `AWS::CloudWatch::CompositeAlarm.ActionsSuppressor`.
	ActionsSuppressor *string `field:"optional" json:"actionsSuppressor" yaml:"actionsSuppressor"`
	// `AWS::CloudWatch::CompositeAlarm.ActionsSuppressorExtensionPeriod`.
	ActionsSuppressorExtensionPeriod *float64 `field:"optional" json:"actionsSuppressorExtensionPeriod" yaml:"actionsSuppressorExtensionPeriod"`
	// `AWS::CloudWatch::CompositeAlarm.ActionsSuppressorWaitPeriod`.
	ActionsSuppressorWaitPeriod *float64 `field:"optional" json:"actionsSuppressorWaitPeriod" yaml:"actionsSuppressorWaitPeriod"`
	// The actions to execute when this alarm transitions to the ALARM state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// The description for the composite alarm.
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// The name for the composite alarm.
	//
	// This name must be unique within your AWS account.
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// The actions to execute when this alarm transitions to the OK state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
}

