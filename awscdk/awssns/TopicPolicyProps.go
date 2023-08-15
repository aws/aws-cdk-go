package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties to associate SNS topics with a policy.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//   topicPolicy := sns.NewTopicPolicy(this, jsii.String("TopicPolicy"), &TopicPolicyProps{
//   	Topics: []iTopic{
//   		topic,
//   	},
//   })
//
//   topicPolicy.Document.AddStatements(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("sns:Subscribe"),
//   	},
//   	Principals: []iPrincipal{
//   		iam.NewAnyPrincipal(),
//   	},
//   	Resources: []*string{
//   		topic.TopicArn,
//   	},
//   }))
//
type TopicPolicyProps struct {
	// The set of topics this policy applies to.
	Topics *[]ITopic `field:"required" json:"topics" yaml:"topics"`
	// IAM policy document to apply to topic(s).
	// Default: empty policy document.
	//
	PolicyDocument awsiam.PolicyDocument `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

