package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties for defining an EventBridge Rule.
//
// Example:
//   import redshiftserverless "github.com/aws/aws-cdk-go/awscdk"
//
//   var workgroup cfnWorkgroup
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   rule.AddTarget(targets.NewRedshiftQuery(workgroup.AttrWorkgroupWorkgroupArn, &RedshiftQueryProps{
//   	Database: jsii.String("dev"),
//   	DeadLetterQueue: dlq,
//   	Sql: []*string{
//   		jsii.String("SELECT * FROM foo"),
//   		jsii.String("SELECT * FROM baz"),
//   	},
//   }))
//
type RuleProps struct {
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
	// Indicates whether the rule is enabled.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The event bus to associate with this rule.
	// Default: - The default event bus.
	//
	EventBus IEventBus `field:"optional" json:"eventBus" yaml:"eventBus"`
	// The role that is used for target invocation.
	//
	// Must be assumable by principal `events.amazonaws.com`.
	// Default: - No role associated.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The schedule or rate (frequency) that determines when EventBridge runs the rule.
	//
	// You must specify this property, the `eventPattern` property, or both.
	//
	// For more information, see Schedule Expression Syntax for
	// Rules in the Amazon EventBridge User Guide.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/scheduled-events.html
	//
	// Default: - None.
	//
	Schedule Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Targets to invoke when this rule matches an event.
	//
	// Input will be the full matched event. If you wish to specify custom
	// target input, use `addTarget(target[, inputOptions])`.
	// Default: - No targets.
	//
	Targets *[]IRuleTarget `field:"optional" json:"targets" yaml:"targets"`
}

