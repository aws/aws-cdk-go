package awsdevopsagent


// Properties for CfnTriggerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var action interface{}
//
//   cfnTriggerMixinProps := &CfnTriggerMixinProps{
//   	Action: action,
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	Condition: &ConditionProperty{
//   		Schedule: &ScheduleProperty{
//   			Expression: jsii.String("expression"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html
//
type CfnTriggerMixinProps struct {
	// The action to perform when the trigger fires.
	//
	// A JSON object containing actionType and task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html#cfn-devopsagent-trigger-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// The unique identifier of the parent Agent Space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html#cfn-devopsagent-trigger-agentspaceid
	//
	AgentSpaceId *string `field:"optional" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The condition that causes the trigger to fire.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html#cfn-devopsagent-trigger-condition
	//
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// The status of the trigger.
	//
	// Active triggers fire on schedule; Inactive triggers are paused.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html#cfn-devopsagent-trigger-status
	//
	// Default: - "Active".
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The type of trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html#cfn-devopsagent-trigger-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

