package awsses


// The action for a rule to take. Only one of the contained actions can be set.
//
// > This data type is a UNION, so only one of the following members can be specified when used or returned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var drop interface{}
//
//   ruleActionProperty := &RuleActionProperty{
//   	AddHeader: &AddHeaderActionProperty{
//   		HeaderName: jsii.String("headerName"),
//   		HeaderValue: jsii.String("headerValue"),
//   	},
//   	Archive: &ArchiveActionProperty{
//   		TargetArchive: jsii.String("targetArchive"),
//
//   		// the properties below are optional
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	},
//   	DeliverToMailbox: &DeliverToMailboxActionProperty{
//   		MailboxArn: jsii.String("mailboxArn"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	},
//   	Drop: drop,
//   	Relay: &RelayActionProperty{
//   		Relay: jsii.String("relay"),
//
//   		// the properties below are optional
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   		MailFrom: jsii.String("mailFrom"),
//   	},
//   	ReplaceRecipient: &ReplaceRecipientActionProperty{
//   		ReplaceWith: []*string{
//   			jsii.String("replaceWith"),
//   		},
//   	},
//   	Send: &SendActionProperty{
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	},
//   	WriteToS3: &S3ActionProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		S3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   		S3Prefix: jsii.String("s3Prefix"),
//   		S3SseKmsKeyId: jsii.String("s3SseKmsKeyId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html
//
type CfnMailManagerRuleSet_RuleActionProperty struct {
	// This action adds a header.
	//
	// This can be used to add arbitrary email headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-addheader
	//
	AddHeader interface{} `field:"optional" json:"addHeader" yaml:"addHeader"`
	// This action archives the email.
	//
	// This can be used to deliver an email to an archive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-archive
	//
	Archive interface{} `field:"optional" json:"archive" yaml:"archive"`
	// This action delivers an email to a WorkMail mailbox.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-delivertomailbox
	//
	DeliverToMailbox interface{} `field:"optional" json:"deliverToMailbox" yaml:"deliverToMailbox"`
	// This action terminates the evaluation of rules in the rule set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-drop
	//
	Drop interface{} `field:"optional" json:"drop" yaml:"drop"`
	// This action relays the email to another SMTP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-relay
	//
	Relay interface{} `field:"optional" json:"relay" yaml:"relay"`
	// The action replaces certain or all recipients with a different set of recipients.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-replacerecipient
	//
	ReplaceRecipient interface{} `field:"optional" json:"replaceRecipient" yaml:"replaceRecipient"`
	// This action sends the email to the internet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-send
	//
	Send interface{} `field:"optional" json:"send" yaml:"send"`
	// This action writes the MIME content of the email to an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-writetos3
	//
	WriteToS3 interface{} `field:"optional" json:"writeToS3" yaml:"writeToS3"`
}

