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
//   ruleProperty := &ruleProperty{
//   	actions: []interface{}{
//   		&actionProperty{
//   			addHeaderAction: &addHeaderActionProperty{
//   				headerName: jsii.String("headerName"),
//   				headerValue: jsii.String("headerValue"),
//   			},
//   			bounceAction: &bounceActionProperty{
//   				message: jsii.String("message"),
//   				sender: jsii.String("sender"),
//   				smtpReplyCode: jsii.String("smtpReplyCode"),
//
//   				// the properties below are optional
//   				statusCode: jsii.String("statusCode"),
//   				topicArn: jsii.String("topicArn"),
//   			},
//   			lambdaAction: &lambdaActionProperty{
//   				functionArn: jsii.String("functionArn"),
//
//   				// the properties below are optional
//   				invocationType: jsii.String("invocationType"),
//   				topicArn: jsii.String("topicArn"),
//   			},
//   			s3Action: &s3ActionProperty{
//   				bucketName: jsii.String("bucketName"),
//
//   				// the properties below are optional
//   				kmsKeyArn: jsii.String("kmsKeyArn"),
//   				objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   				topicArn: jsii.String("topicArn"),
//   			},
//   			snsAction: &sNSActionProperty{
//   				encoding: jsii.String("encoding"),
//   				topicArn: jsii.String("topicArn"),
//   			},
//   			stopAction: &stopActionProperty{
//   				scope: jsii.String("scope"),
//
//   				// the properties below are optional
//   				topicArn: jsii.String("topicArn"),
//   			},
//   			workmailAction: &workmailActionProperty{
//   				organizationArn: jsii.String("organizationArn"),
//
//   				// the properties below are optional
//   				topicArn: jsii.String("topicArn"),
//   			},
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	recipients: []*string{
//   		jsii.String("recipients"),
//   	},
//   	scanEnabled: jsii.Boolean(false),
//   	tlsPolicy: jsii.String("tlsPolicy"),
//   }
//
type CfnReceiptRule_RuleProperty struct {
	// An ordered list of actions to perform on messages that match at least one of the recipient email addresses or domains specified in the receipt rule.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// If `true` , the receipt rule is active.
	//
	// The default value is `false` .
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The name of the receipt rule. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), dashes (-), or periods (.).
	// - Start and end with a letter or number.
	// - Contain 64 characters or fewer.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The recipient domains and email addresses that the receipt rule applies to.
	//
	// If this field is not specified, this rule matches all recipients on all verified domains.
	Recipients *[]*string `field:"optional" json:"recipients" yaml:"recipients"`
	// If `true` , then messages that this receipt rule applies to are scanned for spam and viruses.
	//
	// The default value is `false` .
	ScanEnabled interface{} `field:"optional" json:"scanEnabled" yaml:"scanEnabled"`
	// Specifies whether Amazon SES should require that incoming email is delivered over a connection encrypted with Transport Layer Security (TLS).
	//
	// If this parameter is set to `Require` , Amazon SES bounces emails that are not received over TLS. The default is `Optional` .
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

