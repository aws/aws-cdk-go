package awsdevopsagent


// Properties for defining a `CfnTrigger`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var action interface{}
//
//   cfnTriggerProps := &CfnTriggerProps{
//   	Action: action,
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	Condition: &ConditionProperty{
//   		Schedule: &ScheduleProperty{
//   			Expression: jsii.String("expression"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html
//
type CfnTriggerProps struct {
	// The action to perform when the trigger fires.
	//
	// A JSON object containing actionType and task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html#cfn-devopsagent-trigger-action
	//
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// The unique identifier of the parent Agent Space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html#cfn-devopsagent-trigger-agentspaceid
	//
	AgentSpaceId *string `field:"required" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The condition that causes the trigger to fire.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html#cfn-devopsagent-trigger-condition
	//
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
	// The type of trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html#cfn-devopsagent-trigger-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The status of the trigger.
	//
	// Active triggers fire on schedule; Inactive triggers are paused.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html#cfn-devopsagent-trigger-status
	//
	// Default: - "Active".
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

