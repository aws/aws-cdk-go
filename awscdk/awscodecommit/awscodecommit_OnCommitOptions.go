package awscodecommit

import (
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// Options for the onCommit() method.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var repo repository
//
//   myTopic := sns.NewTopic(this, jsii.String("Topic"))
//
//   repo.onCommit(jsii.String("OnCommit"), &OnCommitOptions{
//   	Target: targets.NewSnsTopic(myTopic),
//   })
//
// Experimental.
type OnCommitOptions struct {
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
	// The branch to monitor.
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
}

