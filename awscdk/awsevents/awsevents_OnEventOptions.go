package awsevents

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Standard set of options for `onXxx` event handlers on construct.
//
// Example:
//   // Lambda function containing logic that evaluates compliance with the rule.
//   evalComplianceFn := lambda.NewFunction(this, jsii.String("CustomFunction"), &functionProps{
//   	code: lambda.assetCode.fromInline(jsii.String("exports.handler = (event) => console.log(event);")),
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   })
//
//   // A custom rule that runs on configuration changes of EC2 instances
//   customRule := config.NewCustomRule(this, jsii.String("Custom"), &customRuleProps{
//   	configurationChanges: jsii.Boolean(true),
//   	lambdaFunction: evalComplianceFn,
//   	ruleScope: config.ruleScope.fromResource(config.resourceType_EC2_INSTANCE()),
//   })
//
//   // A rule to detect stack drifts
//   driftRule := config.NewCloudFormationStackDriftDetectionCheck(this, jsii.String("Drift"))
//
//   // Topic to which compliance notification events will be published
//   complianceTopic := sns.NewTopic(this, jsii.String("ComplianceTopic"))
//
//   // Send notification on compliance change events
//   driftRule.onComplianceChange(jsii.String("ComplianceChange"), &onEventOptions{
//   	target: targets.NewSnsTopic(complianceTopic),
//   })
//
type OnEventOptions struct {
	// The scope to use if the source of the rule and its target are in different Stacks (but in the same account & region).
	//
	// This helps dealing with cycles that often arise in these situations.
	CrossStackScope constructs.Construct `field:"optional" json:"crossStackScope" yaml:"crossStackScope"`
	// A description of the rule's purpose.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	EventPattern *EventPattern `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	Target IRuleTarget `field:"optional" json:"target" yaml:"target"`
}

