package awscognito


// The email configuration of your user pool.
//
// The email configuration type sets your preferred sending method, AWS Region, and sender for messages from your user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailConfigurationProperty := &emailConfigurationProperty{
//   	configurationSet: jsii.String("configurationSet"),
//   	emailSendingAccount: jsii.String("emailSendingAccount"),
//   	from: jsii.String("from"),
//   	replyToEmailAddress: jsii.String("replyToEmailAddress"),
//   	sourceArn: jsii.String("sourceArn"),
//   }
//
type CfnUserPool_EmailConfigurationProperty struct {
	// The set of configuration rules that can be applied to emails sent using Amazon SES.
	//
	// A configuration set is applied to an email by including a reference to the configuration set in the headers of the email. Once applied, all of the rules in that configuration set are applied to the email. Configuration sets can be used to apply the following types of rules to emails:
	//
	// - Event publishing – Amazon SES can track the number of send, delivery, open, click, bounce, and complaint events for each email sent. Use event publishing to send information about these events to other AWS services such as SNS and CloudWatch.
	// - IP pool management – When leasing dedicated IP addresses with Amazon SES, you can create groups of IP addresses, called dedicated IP pools. You can then associate the dedicated IP pools with configuration sets.
	ConfigurationSet *string `field:"optional" json:"configurationSet" yaml:"configurationSet"`
	// Specifies whether Amazon Cognito uses its built-in functionality to send your users email messages, or uses your Amazon Simple Email Service email configuration.
	//
	// Specify one of the following values:
	//
	// - **COGNITO_DEFAULT** - When Amazon Cognito emails your users, it uses its built-in email functionality. When you use the default option, Amazon Cognito allows only a limited number of emails each day for your user pool. For typical production environments, the default email limit is less than the required delivery volume. To achieve a higher delivery volume, specify DEVELOPER to use your Amazon SES email configuration.
	//
	// To look up the email delivery limit for the default option, see [Limits in](https://docs.aws.amazon.com/cognito/latest/developerguide/limits.html) in the *Developer Guide* .
	//
	// The default FROM address is `no-reply@verificationemail.com` . To customize the FROM address, provide the Amazon Resource Name (ARN) of an Amazon SES verified email address for the `SourceArn` parameter.
	// - **DEVELOPER** - When Amazon Cognito emails your users, it uses your Amazon SES configuration. Amazon Cognito calls Amazon SES on your behalf to send email from your verified email address. When you use this option, the email delivery limits are the same limits that apply to your Amazon SES verified email address in your AWS account .
	//
	// If you use this option, provide the ARN of an Amazon SES verified email address for the `SourceArn` parameter.
	//
	// Before Amazon Cognito can email your users, it requires additional permissions to call Amazon SES on your behalf. When you update your user pool with this option, Amazon Cognito creates a *service-linked role* , which is a type of role, in your AWS account . This role contains the permissions that allow to access Amazon SES and send email messages with your address. For more information about the service-linked role that Amazon Cognito creates, see [Using Service-Linked Roles for Amazon Cognito](https://docs.aws.amazon.com/cognito/latest/developerguide/using-service-linked-roles.html) in the *Amazon Cognito Developer Guide* .
	EmailSendingAccount *string `field:"optional" json:"emailSendingAccount" yaml:"emailSendingAccount"`
	// Identifies either the sender's email address or the sender's name with their email address.
	//
	// For example, `testuser@example.com` or `Test User <testuser@example.com>` . This address appears before the body of the email.
	From *string `field:"optional" json:"from" yaml:"from"`
	// The destination to which the receiver of the email should reply.
	ReplyToEmailAddress *string `field:"optional" json:"replyToEmailAddress" yaml:"replyToEmailAddress"`
	// The ARN of a verified email address in Amazon SES.
	//
	// Amazon Cognito uses this email address in one of the following ways, depending on the value that you specify for the `EmailSendingAccount` parameter:
	//
	// - If you specify `COGNITO_DEFAULT` , Amazon Cognito uses this address as the custom FROM address when it emails your users using its built-in email account.
	// - If you specify `DEVELOPER` , Amazon Cognito emails your users with this address by calling Amazon SES on your behalf.
	//
	// The Region value of the `SourceArn` parameter must indicate a supported AWS Region of your user pool. Typically, the Region in the `SourceArn` and the user pool Region are the same. For more information, see [Amazon SES email configuration regions](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-email.html#user-pool-email-developer-region-mapping) in the [Amazon Cognito Developer Guide](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools.html) .
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

