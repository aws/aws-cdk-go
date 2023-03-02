package awscustomerprofiles


// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerConfigProperty := &triggerConfigProperty{
//   	triggerType: jsii.String("triggerType"),
//
//   	// the properties below are optional
//   	triggerProperties: &triggerPropertiesProperty{
//   		scheduled: &scheduledTriggerPropertiesProperty{
//   			scheduleExpression: jsii.String("scheduleExpression"),
//
//   			// the properties below are optional
//   			dataPullMode: jsii.String("dataPullMode"),
//   			firstExecutionFrom: jsii.Number(123),
//   			scheduleEndTime: jsii.Number(123),
//   			scheduleOffset: jsii.Number(123),
//   			scheduleStartTime: jsii.Number(123),
//   			timezone: jsii.String("timezone"),
//   		},
//   	},
//   }
//
type CfnIntegration_TriggerConfigProperty struct {
	// Specifies the type of flow trigger.
	//
	// It can be OnDemand, Scheduled, or Event.
	TriggerType *string `field:"required" json:"triggerType" yaml:"triggerType"`
	// Specifies the configuration details of a schedule-triggered flow that you define.
	//
	// Currently, these settings only apply to the Scheduled trigger type.
	TriggerProperties interface{} `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
}

