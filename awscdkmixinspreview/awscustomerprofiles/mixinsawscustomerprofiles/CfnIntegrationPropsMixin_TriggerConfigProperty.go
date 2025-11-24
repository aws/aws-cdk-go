package mixinsawscustomerprofiles


// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   triggerConfigProperty := &TriggerConfigProperty{
//   	TriggerProperties: &TriggerPropertiesProperty{
//   		Scheduled: &ScheduledTriggerPropertiesProperty{
//   			DataPullMode: jsii.String("dataPullMode"),
//   			FirstExecutionFrom: jsii.Number(123),
//   			ScheduleEndTime: jsii.Number(123),
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//   			ScheduleOffset: jsii.Number(123),
//   			ScheduleStartTime: jsii.Number(123),
//   			Timezone: jsii.String("timezone"),
//   		},
//   	},
//   	TriggerType: jsii.String("triggerType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-triggerconfig.html
//
type CfnIntegrationPropsMixin_TriggerConfigProperty struct {
	// Specifies the configuration details of a schedule-triggered flow that you define.
	//
	// Currently, these settings only apply to the Scheduled trigger type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-triggerconfig.html#cfn-customerprofiles-integration-triggerconfig-triggerproperties
	//
	TriggerProperties interface{} `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
	// Specifies the type of flow trigger.
	//
	// It can be OnDemand, Scheduled, or Event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-triggerconfig.html#cfn-customerprofiles-integration-triggerconfig-triggertype
	//
	TriggerType *string `field:"optional" json:"triggerType" yaml:"triggerType"`
}

