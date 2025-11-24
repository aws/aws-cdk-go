package mixinsawsarcregionswitch


// Defines a condition that can automatically trigger the execution of a Region switch plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   triggerProperty := &TriggerProperty{
//   	Action: jsii.String("action"),
//   	Conditions: []interface{}{
//   		&TriggerConditionProperty{
//   			AssociatedAlarmName: jsii.String("associatedAlarmName"),
//   			Condition: jsii.String("condition"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	MinDelayMinutesBetweenExecutions: jsii.Number(123),
//   	TargetRegion: jsii.String("targetRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html
//
type CfnPlanPropsMixin_TriggerProperty struct {
	// The action to perform when the trigger fires.
	//
	// Valid values include ACTIVATE and DEACTIVATE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html#cfn-arcregionswitch-plan-trigger-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The conditions that must be met for the trigger to fire.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html#cfn-arcregionswitch-plan-trigger-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The description for a trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html#cfn-arcregionswitch-plan-trigger-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The minimum time, in minutes, that must elapse between automatic executions of the plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html#cfn-arcregionswitch-plan-trigger-mindelayminutesbetweenexecutions
	//
	MinDelayMinutesBetweenExecutions *float64 `field:"optional" json:"minDelayMinutesBetweenExecutions" yaml:"minDelayMinutesBetweenExecutions"`
	// The AWS Region for a trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html#cfn-arcregionswitch-plan-trigger-targetregion
	//
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
}

