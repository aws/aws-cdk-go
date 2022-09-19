package awsses


// Properties for a receipt rule action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   receiptRuleActionConfig := &receiptRuleActionConfig{
//   	addHeaderAction: &addHeaderActionConfig{
//   		headerName: jsii.String("headerName"),
//   		headerValue: jsii.String("headerValue"),
//   	},
//   	bounceAction: &bounceActionConfig{
//   		message: jsii.String("message"),
//   		sender: jsii.String("sender"),
//   		smtpReplyCode: jsii.String("smtpReplyCode"),
//
//   		// the properties below are optional
//   		statusCode: jsii.String("statusCode"),
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	lambdaAction: &lambdaActionConfig{
//   		functionArn: jsii.String("functionArn"),
//
//   		// the properties below are optional
//   		invocationType: jsii.String("invocationType"),
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	s3Action: &s3ActionConfig{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   		objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	snsAction: &sNSActionConfig{
//   		encoding: jsii.String("encoding"),
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	stopAction: &stopActionConfig{
//   		scope: jsii.String("scope"),
//
//   		// the properties below are optional
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	workmailAction: &workmailActionConfig{
//   		organizationArn: jsii.String("organizationArn"),
//
//   		// the properties below are optional
//   		topicArn: jsii.String("topicArn"),
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

