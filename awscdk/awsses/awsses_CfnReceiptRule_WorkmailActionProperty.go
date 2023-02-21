package awsses


// When included in a receipt rule, this action calls Amazon WorkMail and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
//
// It usually isn't necessary to set this up manually, because Amazon WorkMail adds the rule automatically during its setup procedure.
//
// For information using a receipt rule to call Amazon WorkMail, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-workmail.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workmailActionProperty := &workmailActionProperty{
//   	organizationArn: jsii.String("organizationArn"),
//
//   	// the properties below are optional
//   	topicArn: jsii.String("topicArn"),
//   }
//
type CfnReceiptRule_WorkmailActionProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon WorkMail organization. Amazon WorkMail ARNs use the following format:.
	//
	// `arn:aws:workmail:<region>:<awsAccountId>:organization/<workmailOrganizationId>`
	//
	// You can find the ID of your organization by using the [ListOrganizations](https://docs.aws.amazon.com/workmail/latest/APIReference/API_ListOrganizations.html) operation in Amazon WorkMail. Amazon WorkMail organization IDs begin with " `m-` ", followed by a string of alphanumeric characters.
	//
	// For information about Amazon WorkMail organizations, see the [Amazon WorkMail Administrator Guide](https://docs.aws.amazon.com/workmail/latest/adminguide/organizations_overview.html) .
	OrganizationArn *string `field:"required" json:"organizationArn" yaml:"organizationArn"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the WorkMail action is called.
	//
	// You can find the ARN of a topic by using the [ListTopics](https://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html) operation in Amazon SNS.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

