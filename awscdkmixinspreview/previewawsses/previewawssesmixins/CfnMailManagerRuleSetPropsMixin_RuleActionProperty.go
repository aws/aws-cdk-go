package previewawssesmixins


// The action for a rule to take. Only one of the contained actions can be set.
//
// > This data type is a UNION, so only one of the following members can be specified when used or returned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var drop interface{}
//
//   ruleActionProperty := &RuleActionProperty{
//   	AddHeader: &AddHeaderActionProperty{
//   		HeaderName: jsii.String("headerName"),
//   		HeaderValue: jsii.String("headerValue"),
//   	},
//   	Archive: &ArchiveActionProperty{
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   		TargetArchive: jsii.String("targetArchive"),
//   	},
//   	DeliverToMailbox: &DeliverToMailboxActionProperty{
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   		MailboxArn: jsii.String("mailboxArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	DeliverToQBusiness: &DeliverToQBusinessActionProperty{
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   		ApplicationId: jsii.String("applicationId"),
//   		IndexId: jsii.String("indexId"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	Drop: drop,
//   	PublishToSns: &SnsActionProperty{
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   		Encoding: jsii.String("encoding"),
//   		PayloadType: jsii.String("payloadType"),
//   		RoleArn: jsii.String("roleArn"),
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	Relay: &RelayActionProperty{
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   		MailFrom: jsii.String("mailFrom"),
//   		Relay: jsii.String("relay"),
//   	},
//   	ReplaceRecipient: &ReplaceRecipientActionProperty{
//   		ReplaceWith: []*string{
//   			jsii.String("replaceWith"),
//   		},
//   	},
//   	Send: &SendActionProperty{
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	WriteToS3: &S3ActionProperty{
//   		ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Prefix: jsii.String("s3Prefix"),
//   		S3SseKmsKeyId: jsii.String("s3SseKmsKeyId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html
//
type CfnMailManagerRuleSetPropsMixin_RuleActionProperty struct {
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
	// This action delivers an email to an Amazon Q Business application for ingestion into its knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-delivertoqbusiness
	//
	DeliverToQBusiness interface{} `field:"optional" json:"deliverToQBusiness" yaml:"deliverToQBusiness"`
	// This action terminates the evaluation of rules in the rule set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-drop
	//
	Drop interface{} `field:"optional" json:"drop" yaml:"drop"`
	// This action publishes the email content to an Amazon SNS topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleaction.html#cfn-ses-mailmanagerruleset-ruleaction-publishtosns
	//
	PublishToSns interface{} `field:"optional" json:"publishToSns" yaml:"publishToSns"`
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

