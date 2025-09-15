package awsses


// Properties for a receipt rule action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   receiptRuleActionConfig := &ReceiptRuleActionConfig{
//   	AddHeaderAction: &AddHeaderActionConfig{
//   		HeaderName: jsii.String("headerName"),
//   		HeaderValue: jsii.String("headerValue"),
//   	},
//   	BounceAction: &BounceActionConfig{
//   		Message: jsii.String("message"),
//   		Sender: jsii.String("sender"),
//   		SmtpReplyCode: jsii.String("smtpReplyCode"),
//
//   		// the properties below are optional
//   		StatusCode: jsii.String("statusCode"),
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	LambdaAction: &LambdaActionConfig{
//   		FunctionArn: jsii.String("functionArn"),
//
//   		// the properties below are optional
//   		InvocationType: jsii.String("invocationType"),
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	S3Action: &S3ActionConfig{
//   		BucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	SnsAction: &SNSActionConfig{
//   		Encoding: jsii.String("encoding"),
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	StopAction: &StopActionConfig{
//   		Scope: jsii.String("scope"),
//
//   		// the properties below are optional
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	WorkmailAction: &WorkmailActionConfig{
//   		OrganizationArn: jsii.String("organizationArn"),
//
//   		// the properties below are optional
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   }
//
type ReceiptRuleActionConfig struct {
	// Adds a header to the received email.
	AddHeaderAction *AddHeaderActionConfig `field:"optional" json:"addHeaderAction" yaml:"addHeaderAction"`
	// Rejects the received email by returning a bounce response to the sender and, optionally, publishes a notification to Amazon SNS.
	BounceAction *BounceActionConfig `field:"optional" json:"bounceAction" yaml:"bounceAction"`
	// Calls an AWS Lambda function, and optionally, publishes a notification to Amazon SNS.
	LambdaAction *LambdaActionConfig `field:"optional" json:"lambdaAction" yaml:"lambdaAction"`
	// Saves the received message to an Amazon S3 bucket and, optionally, publishes a notification to Amazon SNS.
	S3Action *S3ActionConfig `field:"optional" json:"s3Action" yaml:"s3Action"`
	// Publishes the email content within a notification to Amazon SNS.
	SnsAction *SNSActionConfig `field:"optional" json:"snsAction" yaml:"snsAction"`
	// Terminates the evaluation of the receipt rule set and optionally publishes a notification to Amazon SNS.
	StopAction *StopActionConfig `field:"optional" json:"stopAction" yaml:"stopAction"`
	// Calls Amazon WorkMail and, optionally, publishes a notification to Amazon SNS.
	WorkmailAction *WorkmailActionConfig `field:"optional" json:"workmailAction" yaml:"workmailAction"`
}

