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
//   triggerPropertiesProperty := &TriggerPropertiesProperty{
//   	Scheduled: &ScheduledTriggerPropertiesProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//
//   		// the properties below are optional
//   		DataPullMode: jsii.String("dataPullMode"),
//   		FirstExecutionFrom: jsii.Number(123),
//   		ScheduleEndTime: jsii.Number(123),
//   		ScheduleOffset: jsii.Number(123),
//   		ScheduleStartTime: jsii.Number(123),
//   		Timezone: jsii.String("timezone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-triggerproperties.html
//
type CfnIntegration_TriggerPropertiesProperty struct {
	// Specifies the configuration details of a schedule-triggered flow that you define.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-triggerproperties.html#cfn-customerprofiles-integration-triggerproperties-scheduled
	//
	Scheduled interface{} `field:"optional" json:"scheduled" yaml:"scheduled"`
}

