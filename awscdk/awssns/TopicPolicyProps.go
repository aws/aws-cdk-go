package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties to associate SNS topics with a policy.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//   policyDocument := iam.NewPolicyDocument(&PolicyDocumentProps{
//   	AssignSids: jsii.Boolean(true),
//   	Statements: []policyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("sns:Subscribe"),
//   			},
//   			Principals: []iPrincipal{
//   				iam.NewAnyPrincipal(),
//   			},
//   			Resources: []*string{
//   				topic.TopicArn,
//   			},
//   		}),
//   	},
//   })
//
//   topicPolicy := sns.NewTopicPolicy(this, jsii.String("Policy"), &TopicPolicyProps{
//   	Topics: []iTopic{
//   		topic,
//   	},
//   	PolicyDocument: PolicyDocument,
//   })
//
type TopicPolicyProps struct {
	// The set of topics this policy applies to.
	Topics *[]ITopic `field:"required" json:"topics" yaml:"topics"`
	// Adds a statement to enforce encryption of data in transit when publishing to the topic.
	//
	// For more information, see https://docs.aws.amazon.com/sns/latest/dg/sns-security-best-practices.html#enforce-encryption-data-in-transit.
	// Default: false.
	//
	EnforceSSL *bool `field:"optional" json:"enforceSSL" yaml:"enforceSSL"`
	// IAM policy document to apply to topic(s).
	// Default: empty policy document.
	//
	PolicyDocument awsiam.PolicyDocument `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

