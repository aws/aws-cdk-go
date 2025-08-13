package awsarcregionswitch


// Defines a condition that can automatically trigger the execution of a Region switch plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerProperty := &TriggerProperty{
//   	Action: jsii.String("action"),
//   	Conditions: []interface{}{
//   		&TriggerConditionProperty{
//   			AssociatedAlarmName: jsii.String("associatedAlarmName"),
//   			Condition: jsii.String("condition"),
//   		},
//   	},
//   	MinDelayMinutesBetweenExecutions: jsii.Number(123),
//   	TargetRegion: jsii.String("targetRegion"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html
//
type CfnPlan_TriggerProperty struct {
	// The action to perform when the trigger fires.
	//
	// Valid values include ACTIVATE and DEACTIVATE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html#cfn-arcregionswitch-plan-trigger-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The conditions that must be met for the trigger to fire.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html#cfn-arcregionswitch-plan-trigger-conditions
	//
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The minimum time, in minutes, that must elapse between automatic executions of the plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html#cfn-arcregionswitch-plan-trigger-mindelayminutesbetweenexecutions
	//
	MinDelayMinutesBetweenExecutions *float64 `field:"required" json:"minDelayMinutesBetweenExecutions" yaml:"minDelayMinutesBetweenExecutions"`
	// The AWS Region for a trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html#cfn-arcregionswitch-plan-trigger-targetregion
	//
	TargetRegion *string `field:"required" json:"targetRegion" yaml:"targetRegion"`
	// The description for a trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-trigger.html#cfn-arcregionswitch-plan-trigger-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

