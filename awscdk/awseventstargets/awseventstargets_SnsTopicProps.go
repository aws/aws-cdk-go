package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// Customize the SNS Topic Event Target.
//
// Example:
//   var onCommitRule rule
//   var topic topic
//
//
//   onCommitRule.addTarget(targets.NewSnsTopic(topic, &snsTopicProps{
//   	message: events.ruleTargetInput.fromText(
//   	fmt.Sprintf("A commit was pushed to the repository %v on branch %v", codecommit.referenceEvent.repositoryName, codecommit.*referenceEvent.referenceName)),
//   }))
//
// Experimental.
type SnsTopicProps struct {
	// The message to send to the topic.
	// Experimental.
	Message awsevents.RuleTargetInput `field:"optional" json:"message" yaml:"message"`
}

