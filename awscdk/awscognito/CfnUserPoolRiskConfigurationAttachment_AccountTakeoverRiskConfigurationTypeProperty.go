package awscognito


// Configuration for mitigation actions and notification for different levels of risk detected for a potential account takeover.
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
	// Account takeover risk configuration actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The notify configuration used to construct email notifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype-notifyconfiguration
	//
	NotifyConfiguration interface{} `field:"optional" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

