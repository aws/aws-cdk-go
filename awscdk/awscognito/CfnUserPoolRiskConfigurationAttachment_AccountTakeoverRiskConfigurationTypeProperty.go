package awscognito


// The settings for automated responses and notification templates for adaptive authentication with threat protection features.
//
// This data type is a request parameter of `API_SetRiskConfiguration` and a response parameter of `API_DescribeRiskConfiguration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//
//   	// the properties below are optional
//   	NotifyConfiguration: &NotifyConfigurationTypeProperty{
//   		SourceArn: jsii.String("sourceArn"),
//
//   		// the properties below are optional
//   		BlockEmail: &NotifyEmailTypeProperty{
//   			Subject: jsii.String("subject"),
//
//   			// the properties below are optional
//   			HtmlBody: jsii.String("htmlBody"),
//   			TextBody: jsii.String("textBody"),
//   		},
//   		From: jsii.String("from"),
//   		MfaEmail: &NotifyEmailTypeProperty{
//   			Subject: jsii.String("subject"),
//
//   			// the properties below are optional
//   			HtmlBody: jsii.String("htmlBody"),
//   			TextBody: jsii.String("textBody"),
//   		},
//   		NoActionEmail: &NotifyEmailTypeProperty{
//   			Subject: jsii.String("subject"),
//
//   			// the properties below are optional
//   			HtmlBody: jsii.String("htmlBody"),
//   			TextBody: jsii.String("textBody"),
//   		},
//   		ReplyTo: jsii.String("replyTo"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype.html
//
type CfnUserPoolRiskConfigurationAttachment_AccountTakeoverRiskConfigurationTypeProperty struct {
	// A list of account-takeover actions for each level of risk that Amazon Cognito might assess with threat protection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The settings for composing and sending an email message when threat protection assesses a risk level with adaptive authentication.
	//
	// When you choose to notify users in `AccountTakeoverRiskConfiguration` , Amazon Cognito sends an email message using the method and template that you set with this data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype-notifyconfiguration
	//
	NotifyConfiguration interface{} `field:"optional" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

