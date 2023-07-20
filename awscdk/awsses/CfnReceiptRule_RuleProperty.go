package awsses


// Receipt rules enable you to specify which actions Amazon SES should take when it receives mail on behalf of one or more email addresses or domains that you own.
//
// Each receipt rule defines a set of email addresses or domains that it applies to. If the email addresses or domains match at least one recipient address of the message, Amazon SES executes all of the receipt rule's actions on the message.
//
// For information about setting up receipt rules, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &RuleProperty{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			AddHeaderAction: &AddHeaderActionProperty{
//   				HeaderName: jsii.String("headerName"),
//   				HeaderValue: jsii.String("headerValue"),
//   			},
//   			BounceAction: &BounceActionProperty{
//   				Message: jsii.String("message"),
//   				Sender: jsii.String("sender"),
//   				SmtpReplyCode: jsii.String("smtpReplyCode"),
//
//   				// the properties below are optional
//   				StatusCode: jsii.String("statusCode"),
//   				TopicArn: jsii.String("topicArn"),
//   			},
//   			LambdaAction: &LambdaActionProperty{
//   				FunctionArn: jsii.String("functionArn"),
//
//   				// the properties below are optional
//   				InvocationType: jsii.String("invocationType"),
//   				TopicArn: jsii.String("topicArn"),
//   			},
//   			S3Action: &S3ActionProperty{
//   				BucketName: jsii.String("bucketName"),
//
//   				// the properties below are optional
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   				ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   				TopicArn: jsii.String("topicArn"),
//   			},
//   			SnsAction: &SNSActionProperty{
//   				Encoding: jsii.String("encoding"),
//   				TopicArn: jsii.String("topicArn"),
//   			},
//   			StopAction: &StopActionProperty{
//   				Scope: jsii.String("scope"),
//
//   				// the properties below are optional
//   				TopicArn: jsii.String("topicArn"),
//   			},
//   			WorkmailAction: &WorkmailActionProperty{
//   				OrganizationArn: jsii.String("organizationArn"),
//
//   				// the properties below are optional
//   				TopicArn: jsii.String("topicArn"),
//   			},
//   		},
//   	},
//   	Enabled: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	Recipients: []*string{
//   		jsii.String("recipients"),
//   	},
//   	ScanEnabled: jsii.Boolean(false),
//   	TlsPolicy: jsii.String("tlsPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-rule.html
//
type CfnReceiptRule_RuleProperty struct {
	// An ordered list of actions to perform on messages that match at least one of the recipient email addresses or domains specified in the receipt rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-rule.html#cfn-ses-receiptrule-rule-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// If `true` , the receipt rule is active.
	//
	// The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-rule.html#cfn-ses-receiptrule-rule-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The name of the receipt rule. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), dashes (-), or periods (.).
	// - Start and end with a letter or number.
	// - Contain 64 characters or fewer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-rule.html#cfn-ses-receiptrule-rule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The recipient domains and email addresses that the receipt rule applies to.
	//
	// If this field is not specified, this rule matches all recipients on all verified domains.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-rule.html#cfn-ses-receiptrule-rule-recipients
	//
	Recipients *[]*string `field:"optional" json:"recipients" yaml:"recipients"`
	// If `true` , then messages that this receipt rule applies to are scanned for spam and viruses.
	//
	// The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-rule.html#cfn-ses-receiptrule-rule-scanenabled
	//
	ScanEnabled interface{} `field:"optional" json:"scanEnabled" yaml:"scanEnabled"`
	// Specifies whether Amazon SES should require that incoming email is delivered over a connection encrypted with Transport Layer Security (TLS).
	//
	// If this parameter is set to `Require` , Amazon SES bounces emails that are not received over TLS. The default is `Optional` .
	//
	// Valid Values: `Require | Optional`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-rule.html#cfn-ses-receiptrule-rule-tlspolicy
	//
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

