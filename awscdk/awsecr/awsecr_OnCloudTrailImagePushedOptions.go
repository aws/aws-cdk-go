package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// Options for the onCloudTrailImagePushed method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var detail interface{}
//   var ruleTarget iRuleTarget
//
//   onCloudTrailImagePushedOptions := &onCloudTrailImagePushedOptions{
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
//   	imageTag: jsii.String("imageTag"),
//   	ruleName: jsii.String("ruleName"),
//   	target: ruleTarget,
//   }
//
// Experimental.
type OnCloudTrailImagePushedOptions struct {
	// A description of the rule's purpose.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	// Experimental.
	EventPattern *awsevents.EventPattern `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	// Experimental.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	// Experimental.
	Target awsevents.IRuleTarget `field:"optional" json:"target" yaml:"target"`
	// Only watch changes to this image tag.
	// Experimental.
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

