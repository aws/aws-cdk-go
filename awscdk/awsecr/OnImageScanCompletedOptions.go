package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// Options for the OnImageScanCompleted method.
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
//   onImageScanCompletedOptions := &OnImageScanCompletedOptions{
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
//   	ImageTags: []*string{
//   		jsii.String("imageTags"),
//   	},
//   	RuleName: jsii.String("ruleName"),
//   	Target: ruleTarget,
//   }
//
// Experimental.
type OnImageScanCompletedOptions struct {
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
	// Only watch changes to the image tags spedified.
	//
	// Leave it undefined to watch the full repository.
	// Experimental.
	ImageTags *[]*string `field:"optional" json:"imageTags" yaml:"imageTags"`
}

