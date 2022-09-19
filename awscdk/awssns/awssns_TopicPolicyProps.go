package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties to associate SNS topics with a policy.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//   topicPolicy := sns.NewTopicPolicy(this, jsii.String("TopicPolicy"), &topicPolicyProps{
//   	topics: []iTopic{
//   		topic,
//   	},
//   })
//
//   topicPolicy.document.addStatements(iam.NewPolicyStatement(&policyStatementProps{
//   	actions: []*string{
//   		jsii.String("sns:Subscribe"),
//   	},
//   	principals: []iPrincipal{
//   		iam.NewAnyPrincipal(),
//   	},
//   	resources: []*string{
//   		topic.topicArn,
//   	},
//   }))
//
type TopicPolicyProps struct {
	// The set of topics this policy applies to.
	Topics *[]ITopic `field:"required" json:"topics" yaml:"topics"`
	// IAM policy document to apply to topic(s).
	PolicyDocument awsiam.PolicyDocument `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

