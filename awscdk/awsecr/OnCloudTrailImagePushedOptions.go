package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options for the onCloudTrailImagePushed method.
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
//   onCloudTrailImagePushedOptions := &OnCloudTrailImagePushedOptions{
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
//   	ImageTag: jsii.String("imageTag"),
//   	RuleName: jsii.String("ruleName"),
//   	Target: ruleTarget,
//   }
//
type OnCloudTrailImagePushedOptions struct {
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
	EventPattern *awsevents.EventPattern `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	// Default: AWS CloudFormation generates a unique physical ID.
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	// Default: - No target is added to the rule. Use `addTarget()` to add a target.
	//
	Target awsevents.IRuleTarget `field:"optional" json:"target" yaml:"target"`
	// Only watch changes to this image tag.
	// Default: - Watch changes to all tags.
	//
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

