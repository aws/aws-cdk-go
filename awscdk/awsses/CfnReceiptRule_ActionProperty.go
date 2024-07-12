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
//   actionProperty := &ActionProperty{
//   	AddHeaderAction: &AddHeaderActionProperty{
//   		HeaderName: jsii.String("headerName"),
//   		HeaderValue: jsii.String("headerValue"),
//   	},
//   	BounceAction: &BounceActionProperty{
//   		Message: jsii.String("message"),
//   		Sender: jsii.String("sender"),
//   		SmtpReplyCode: jsii.String("smtpReplyCode"),
//
//   		// the properties below are optional
//   		StatusCode: jsii.String("statusCode"),
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	LambdaAction: &LambdaActionProperty{
//   		FunctionArn: jsii.String("functionArn"),
//
//   		// the properties below are optional
//   		InvocationType: jsii.String("invocationType"),
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	S3Action: &S3ActionProperty{
//   		BucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	SnsAction: &SNSActionProperty{
//   		Encoding: jsii.String("encoding"),
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	StopAction: &StopActionProperty{
//   		Scope: jsii.String("scope"),
//
//   		// the properties below are optional
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	WorkmailAction: &WorkmailActionProperty{
//   		OrganizationArn: jsii.String("organizationArn"),
//
//   		// the properties below are optional
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html
//
type CfnReceiptRule_ActionProperty struct {
	// Adds a header to the received email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-addheaderaction
	//
	AddHeaderAction interface{} `field:"optional" json:"addHeaderAction" yaml:"addHeaderAction"`
	// Rejects the received email by returning a bounce response to the sender and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-bounceaction
	//
	BounceAction interface{} `field:"optional" json:"bounceAction" yaml:"bounceAction"`
	// Calls an AWS Lambda function, and optionally, publishes a notification to Amazon SNS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-lambdaaction
	//
	LambdaAction interface{} `field:"optional" json:"lambdaAction" yaml:"lambdaAction"`
	// Saves the received message to an Amazon Simple Storage Service (Amazon S3) bucket and, optionally, publishes a notification to Amazon SNS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-s3action
	//
	S3Action interface{} `field:"optional" json:"s3Action" yaml:"s3Action"`
	// Publishes the email content within a notification to Amazon SNS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-snsaction
	//
	SnsAction interface{} `field:"optional" json:"snsAction" yaml:"snsAction"`
	// Terminates the evaluation of the receipt rule set and optionally publishes a notification to Amazon SNS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-stopaction
	//
	StopAction interface{} `field:"optional" json:"stopAction" yaml:"stopAction"`
	// Calls Amazon WorkMail and, optionally, publishes a notification to Amazon SNS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-workmailaction
	//
	WorkmailAction interface{} `field:"optional" json:"workmailAction" yaml:"workmailAction"`
}

