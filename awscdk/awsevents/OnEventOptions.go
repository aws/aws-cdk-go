package awsevents

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Standard set of options for `onXxx` event handlers on construct.
//
// Example:
//   // Lambda function containing logic that evaluates compliance with the rule.
//   evalComplianceFn := lambda.NewFunction(this, jsii.String("CustomFunction"), &FunctionProps{
//   	Code: lambda.AssetCode_FromInline(jsii.String("exports.handler = (event) => console.log(event);")),
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   })
//
//   // A custom rule that runs on configuration changes of EC2 instances
//   customRule := config.NewCustomRule(this, jsii.String("Custom"), &CustomRuleProps{
//   	ConfigurationChanges: jsii.Boolean(true),
//   	LambdaFunction: evalComplianceFn,
//   	RuleScope: config.RuleScope_FromResource(config.ResourceType_EC2_INSTANCE()),
//   })
//
//   // A rule to detect stack drifts
//   driftRule := config.NewCloudFormationStackDriftDetectionCheck(this, jsii.String("Drift"))
//
//   // Topic to which compliance notification events will be published
//   complianceTopic := sns.NewTopic(this, jsii.String("ComplianceTopic"))
//
//   // Send notification on compliance change events
//   driftRule.onComplianceChange(jsii.String("ComplianceChange"), &OnEventOptions{
//   	Target: targets.NewSnsTopic(complianceTopic),
//   })
//
type OnEventOptions struct {
	// The scope to use if the source of the rule and its target are in different Stacks (but in the same account & region).
	//
	// This helps dealing with cycles that often arise in these situations.
	// Default: - none (the main scope will be used, even for cross-stack Events).
	//
	CrossStackScope constructs.Construct `field:"optional" json:"crossStackScope" yaml:"crossStackScope"`
	// A description of the rule's purpose.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	// Default: - No additional filtering based on an event pattern.
	//
	EventPattern *EventPattern `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	// Default: AWS CloudFormation generates a unique physical ID.
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	// Default: - No target is added to the rule. Use `addTarget()` to add a target.
	//
	Target IRuleTarget `field:"optional" json:"target" yaml:"target"`
}

