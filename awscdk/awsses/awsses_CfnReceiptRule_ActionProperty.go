package awsses


// An action that Amazon SES can take when it receives an email on behalf of one or more email addresses or domains that you own.
//
// An instance of this data type can represent only one action.
//
// For information about setting up receipt rules, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	addHeaderAction: &addHeaderActionProperty{
//   		headerName: jsii.String("headerName"),
//   		headerValue: jsii.String("headerValue"),
//   	},
//   	bounceAction: &bounceActionProperty{
//   		message: jsii.String("message"),
//   		sender: jsii.String("sender"),
//   		smtpReplyCode: jsii.String("smtpReplyCode"),
//
//   		// the properties below are optional
//   		statusCode: jsii.String("statusCode"),
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	lambdaAction: &lambdaActionProperty{
//   		functionArn: jsii.String("functionArn"),
//
//   		// the properties below are optional
//   		invocationType: jsii.String("invocationType"),
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	s3Action: &s3ActionProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   		objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	snsAction: &sNSActionProperty{
//   		encoding: jsii.String("encoding"),
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	stopAction: &stopActionProperty{
//   		scope: jsii.String("scope"),
//
//   		// the properties below are optional
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	workmailAction: &workmailActionProperty{
//   		organizationArn: jsii.String("organizationArn"),
//
//   		// the properties below are optional
//   		topicArn: jsii.String("topicArn"),
//   	},
//   }
//
type CfnReceiptRule_ActionProperty struct {
	// Adds a header to the received email.
	AddHeaderAction interface{} `field:"optional" json:"addHeaderAction" yaml:"addHeaderAction"`
	// Rejects the received email by returning a bounce response to the sender and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
	BounceAction interface{} `field:"optional" json:"bounceAction" yaml:"bounceAction"`
	// Calls an AWS Lambda function, and optionally, publishes a notification to Amazon SNS.
	LambdaAction interface{} `field:"optional" json:"lambdaAction" yaml:"lambdaAction"`
	// Saves the received message to an Amazon Simple Storage Service (Amazon S3) bucket and, optionally, publishes a notification to Amazon SNS.
	S3Action interface{} `field:"optional" json:"s3Action" yaml:"s3Action"`
	// Publishes the email content within a notification to Amazon SNS.
	SnsAction interface{} `field:"optional" json:"snsAction" yaml:"snsAction"`
	// Terminates the evaluation of the receipt rule set and optionally publishes a notification to Amazon SNS.
	StopAction interface{} `field:"optional" json:"stopAction" yaml:"stopAction"`
	// Calls Amazon WorkMail and, optionally, publishes a notification to Amazon Amazon SNS.
	WorkmailAction interface{} `field:"optional" json:"workmailAction" yaml:"workmailAction"`
}

