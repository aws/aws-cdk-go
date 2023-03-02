package awsappflow


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
//   	triggerProperties: &scheduledTriggerPropertiesProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//
//   		// the properties below are optional
//   		dataPullMode: jsii.String("dataPullMode"),
//   		firstExecutionFrom: jsii.Number(123),
//   		flowErrorDeactivationThreshold: jsii.Number(123),
//   		scheduleEndTime: jsii.Number(123),
//   		scheduleOffset: jsii.Number(123),
//   		scheduleStartTime: jsii.Number(123),
//   		timeZone: jsii.String("timeZone"),
//   	},
//   }
//
type CfnFlow_TriggerConfigProperty struct {
	// Specifies the type of flow trigger.
	//
	// This can be `OnDemand` , `Scheduled` , or `Event` .
	TriggerType *string `field:"required" json:"triggerType" yaml:"triggerType"`
	// Specifies the configuration details of a schedule-triggered flow as defined by the user.
	//
	// Currently, these settings only apply to the `Scheduled` trigger type.
	TriggerProperties interface{} `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
}

