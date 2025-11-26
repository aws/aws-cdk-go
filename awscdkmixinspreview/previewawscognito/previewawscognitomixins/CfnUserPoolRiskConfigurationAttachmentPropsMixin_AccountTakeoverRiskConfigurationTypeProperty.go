package previewawscognitomixins


// The settings for automated responses and notification templates for adaptive authentication with advanced security features.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accountTakeoverRiskConfigurationTypeProperty := &AccountTakeoverRiskConfigurationTypeProperty{
//   	Actions: &AccountTakeoverActionsTypeProperty{
//   		HighAction: &AccountTakeoverActionTypeProperty{
//   			EventAction: jsii.String("eventAction"),
//   			Notify: jsii.Boolean(false),
//   		},
//   		LowAction: &AccountTakeoverActionTypeProperty{
//   			EventAction: jsii.String("eventAction"),
//   			Notify: jsii.Boolean(false),
//   		},
//   		MediumAction: &AccountTakeoverActionTypeProperty{
//   			EventAction: jsii.String("eventAction"),
//   			Notify: jsii.Boolean(false),
//   		},
//   	},
//   	NotifyConfiguration: &NotifyConfigurationTypeProperty{
//   		BlockEmail: &NotifyEmailTypeProperty{
//   			HtmlBody: jsii.String("htmlBody"),
//   			Subject: jsii.String("subject"),
//   			TextBody: jsii.String("textBody"),
//   		},
//   		From: jsii.String("from"),
//   		MfaEmail: &NotifyEmailTypeProperty{
//   			HtmlBody: jsii.String("htmlBody"),
//   			Subject: jsii.String("subject"),
//   			TextBody: jsii.String("textBody"),
//   		},
//   		NoActionEmail: &NotifyEmailTypeProperty{
//   			HtmlBody: jsii.String("htmlBody"),
//   			Subject: jsii.String("subject"),
//   			TextBody: jsii.String("textBody"),
//   		},
//   		ReplyTo: jsii.String("replyTo"),
//   		SourceArn: jsii.String("sourceArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype.html
//
type CfnUserPoolRiskConfigurationAttachmentPropsMixin_AccountTakeoverRiskConfigurationTypeProperty struct {
	// A list of account-takeover actions for each level of risk that Amazon Cognito might assess with threat protection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The settings for composing and sending an email message when threat protection assesses a risk level with adaptive authentication.
	//
	// When you choose to notify users in `AccountTakeoverRiskConfiguration` , Amazon Cognito sends an email message using the method and template that you set with this data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype-notifyconfiguration
	//
	NotifyConfiguration interface{} `field:"optional" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

