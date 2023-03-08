package awscustomerprofiles


// Specifies the configuration details that control the trigger for a flow.
//
// Currently, these settings only apply to the Scheduled trigger type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerPropertiesProperty := &triggerPropertiesProperty{
//   	scheduled: &scheduledTriggerPropertiesProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//
//   		// the properties below are optional
//   		dataPullMode: jsii.String("dataPullMode"),
//   		firstExecutionFrom: jsii.Number(123),
//   		scheduleEndTime: jsii.Number(123),
//   		scheduleOffset: jsii.Number(123),
//   		scheduleStartTime: jsii.Number(123),
//   		timezone: jsii.String("timezone"),
//   	},
//   }
//
type CfnIntegration_TriggerPropertiesProperty struct {
	// Specifies the configuration details of a schedule-triggered flow that you define.
	Scheduled interface{} `field:"optional" json:"scheduled" yaml:"scheduled"`
}

