package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   emailConfigurationProperty := &EmailConfigurationProperty{
//   	ConfigurationSet: jsii.String("configurationSet"),
//   	EmailSendingAccount: jsii.String("emailSendingAccount"),
//   	From: jsii.String("from"),
//   	ReplyToEmailAddress: jsii.String("replyToEmailAddress"),
//   	SourceArn: jsii.String("sourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-emailconfiguration.html
//
type CfnUserPoolRegionalConfigurationAttachmentPropsMixin_EmailConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-emailconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-emailconfiguration-configurationset
	//
	ConfigurationSet *string `field:"optional" json:"configurationSet" yaml:"configurationSet"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-emailconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-emailconfiguration-emailsendingaccount
	//
	EmailSendingAccount *string `field:"optional" json:"emailSendingAccount" yaml:"emailSendingAccount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-emailconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-emailconfiguration-from
	//
	From *string `field:"optional" json:"from" yaml:"from"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-emailconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-emailconfiguration-replytoemailaddress
	//
	ReplyToEmailAddress *string `field:"optional" json:"replyToEmailAddress" yaml:"replyToEmailAddress"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-emailconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-emailconfiguration-sourcearn
	//
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

