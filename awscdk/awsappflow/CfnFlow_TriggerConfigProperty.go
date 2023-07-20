package awsappflow


// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerConfigProperty := &TriggerConfigProperty{
//   	TriggerType: jsii.String("triggerType"),
//
//   	// the properties below are optional
//   	TriggerProperties: &ScheduledTriggerPropertiesProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//
//   		// the properties below are optional
//   		DataPullMode: jsii.String("dataPullMode"),
//   		FirstExecutionFrom: jsii.Number(123),
//   		FlowErrorDeactivationThreshold: jsii.Number(123),
//   		ScheduleEndTime: jsii.Number(123),
//   		ScheduleOffset: jsii.Number(123),
//   		ScheduleStartTime: jsii.Number(123),
//   		TimeZone: jsii.String("timeZone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-triggerconfig.html
//
type CfnFlow_TriggerConfigProperty struct {
	// Specifies the type of flow trigger.
	//
	// This can be `OnDemand` , `Scheduled` , or `Event` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-triggerconfig.html#cfn-appflow-flow-triggerconfig-triggertype
	//
	TriggerType *string `field:"required" json:"triggerType" yaml:"triggerType"`
	// Specifies the configuration details of a schedule-triggered flow as defined by the user.
	//
	// Currently, these settings only apply to the `Scheduled` trigger type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-triggerconfig.html#cfn-appflow-flow-triggerconfig-triggerproperties
	//
	TriggerProperties interface{} `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
}

