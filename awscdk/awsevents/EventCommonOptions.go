package awsevents

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Common options for Events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//   var detail interface{}
//
//   eventCommonOptions := &EventCommonOptions{
//   	CrossStackScope: construct,
//   	Description: jsii.String("description"),
//   	EventPattern: &EventPattern{
//   		Account: []*string{
//   			jsii.String("account"),
//   		},
//   		Detail: map[string]interface{}{
//   			"detailKey": detail,
//   		},
//   		DetailType: []*string{
//   			jsii.String("detailType"),
//   		},
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Source: []*string{
//   			jsii.String("source"),
//   		},
//   		Time: []*string{
//   			jsii.String("time"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	RuleName: jsii.String("ruleName"),
//   }
//
type EventCommonOptions struct {
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
}

