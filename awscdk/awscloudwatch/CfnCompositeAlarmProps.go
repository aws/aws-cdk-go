package awscloudwatch


// Properties for defining a `CfnCompositeAlarm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCompositeAlarmProps := &CfnCompositeAlarmProps{
//   	AlarmRule: jsii.String("alarmRule"),
//
//   	// the properties below are optional
//   	ActionsEnabled: jsii.Boolean(false),
//   	ActionsSuppressor: jsii.String("actionsSuppressor"),
//   	ActionsSuppressorExtensionPeriod: jsii.Number(123),
//   	ActionsSuppressorWaitPeriod: jsii.Number(123),
//   	AlarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	AlarmDescription: jsii.String("alarmDescription"),
//   	AlarmName: jsii.String("alarmName"),
//   	InsufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	OkActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmrule
	//
	AlarmRule *string `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm.
	//
	// The default is TRUE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-actionsenabled
	//
	ActionsEnabled interface{} `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// Actions will be suppressed if the suppressor alarm is in the `ALARM` state.
	//
	// `ActionsSuppressor` can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-actionssuppressor
	//
	ActionsSuppressor *string `field:"optional" json:"actionsSuppressor" yaml:"actionsSuppressor"`
	// The maximum time in seconds that the composite alarm waits after suppressor alarm goes out of the `ALARM` state.
	//
	// After this time, the composite alarm performs its actions.
	//
	// > `ExtensionPeriod` is required only when `ActionsSuppressor` is specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-actionssuppressorextensionperiod
	//
	ActionsSuppressorExtensionPeriod *float64 `field:"optional" json:"actionsSuppressorExtensionPeriod" yaml:"actionsSuppressorExtensionPeriod"`
	// The maximum time in seconds that the composite alarm waits for the suppressor alarm to go into the `ALARM` state.
	//
	// After this time, the composite alarm performs its actions.
	//
	// > `WaitPeriod` is required only when `ActionsSuppressor` is specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-actionssuppressorwaitperiod
	//
	ActionsSuppressorWaitPeriod *float64 `field:"optional" json:"actionsSuppressorWaitPeriod" yaml:"actionsSuppressorWaitPeriod"`
	// The actions to execute when this alarm transitions to the ALARM state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmactions
	//
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// The description for the composite alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmdescription
	//
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// The name for the composite alarm.
	//
	// This name must be unique within your AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmname
	//
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-insufficientdataactions
	//
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// The actions to execute when this alarm transitions to the OK state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-okactions
	//
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
}

