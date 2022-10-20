package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options for the onCloudTrailPutObject method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//   var detail interface{}
//   var ruleTarget iRuleTarget
//
//   onCloudTrailBucketEventOptions := &onCloudTrailBucketEventOptions{
//   	crossStackScope: construct,
//   	description: jsii.String("description"),
//   	eventPattern: &eventPattern{
//   		account: []*string{
//   			jsii.String("account"),
//   		},
//   		detail: map[string]interface{}{
//   			"detailKey": detail,
//   		},
//   		detailType: []*string{
//   			jsii.String("detailType"),
//   		},
//   		id: []*string{
//   			jsii.String("id"),
//   		},
//   		region: []*string{
//   			jsii.String("region"),
//   		},
//   		resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		source: []*string{
//   			jsii.String("source"),
//   		},
//   		time: []*string{
//   			jsii.String("time"),
//   		},
//   		version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	paths: []*string{
//   		jsii.String("paths"),
//   	},
//   	ruleName: jsii.String("ruleName"),
//   	target: ruleTarget,
//   }
//
type OnCloudTrailBucketEventOptions struct {
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
	EventPattern *awsevents.EventPattern `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	Target awsevents.IRuleTarget `field:"optional" json:"target" yaml:"target"`
	// Only watch changes to these object paths.
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

