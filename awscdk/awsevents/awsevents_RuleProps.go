package awsevents


// Properties for defining an EventBridge Rule.
//
// Example:
//   connection := events.NewConnection(this, jsii.String("Connection"), &connectionProps{
//   	authorization: events.authorization.apiKey(jsii.String("x-api-key"), awscdk.SecretValue.secretsManager(jsii.String("ApiSecretName"))),
//   	description: jsii.String("Connection with API Key x-api-key"),
//   })
//
//   destination := events.NewApiDestination(this, jsii.String("Destination"), &apiDestinationProps{
//   	connection: connection,
//   	endpoint: jsii.String("https://example.com"),
//   	description: jsii.String("Calling example.com with API key x-api-key"),
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
//   	targets: []iRuleTarget{
//   		targets.NewApiDestination(destination),
//   	},
//   })
//
// Experimental.
type RuleProps struct {
	// A description of the rule's purpose.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the rule is enabled.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The event bus to associate with this rule.
	// Experimental.
	EventBus IEventBus `field:"optional" json:"eventBus" yaml:"eventBus"`
	// Describes which events EventBridge routes to the specified target.
	//
	// These routed events are matched events. For more information, see Events
	// and Event Patterns in the Amazon EventBridge User Guide.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	// You must specify this property (either via props or via
	// `addEventPattern`), the `scheduleExpression` property, or both. The
	// method `addEventPattern` can be used to add filter values to the event
	// pattern.
	//
	// Experimental.
	EventPattern *EventPattern `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	// Experimental.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The schedule or rate (frequency) that determines when EventBridge runs the rule.
	//
	// For more information, see Schedule Expression Syntax for
	// Rules in the Amazon EventBridge User Guide.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/scheduled-events.html
	//
	// You must specify this property, the `eventPattern` property, or both.
	//
	// Experimental.
	Schedule Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Targets to invoke when this rule matches an event.
	//
	// Input will be the full matched event. If you wish to specify custom
	// target input, use `addTarget(target[, inputOptions])`.
	// Experimental.
	Targets *[]IRuleTarget `field:"optional" json:"targets" yaml:"targets"`
}

