package awsses


// WorkmailAction configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workmailActionConfig := &WorkmailActionConfig{
//   	OrganizationArn: jsii.String("organizationArn"),
//
//   	// the properties below are optional
//   	TopicArn: jsii.String("topicArn"),
//   }
//
type WorkmailActionConfig struct {
	// The Amazon Resource Name (ARN) of the Amazon WorkMail organization.
	OrganizationArn *string `field:"required" json:"organizationArn" yaml:"organizationArn"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the WorkMail action is called.
	// Default: - No notification is sent to SNS.
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

