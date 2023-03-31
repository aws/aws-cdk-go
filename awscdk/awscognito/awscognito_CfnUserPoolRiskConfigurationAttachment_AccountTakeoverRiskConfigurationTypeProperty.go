package awscognito


// Configuration for mitigation actions and notification for different levels of risk detected for a potential account takeover.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountTakeoverRiskConfigurationTypeProperty := &accountTakeoverRiskConfigurationTypeProperty{
//   	actions: &accountTakeoverActionsTypeProperty{
//   		highAction: &accountTakeoverActionTypeProperty{
//   			eventAction: jsii.String("eventAction"),
//   			notify: jsii.Boolean(false),
//   		},
//   		lowAction: &accountTakeoverActionTypeProperty{
//   			eventAction: jsii.String("eventAction"),
//   			notify: jsii.Boolean(false),
//   		},
//   		mediumAction: &accountTakeoverActionTypeProperty{
//   			eventAction: jsii.String("eventAction"),
//   			notify: jsii.Boolean(false),
//   		},
//   	},
//
//   	// the properties below are optional
//   	notifyConfiguration: &notifyConfigurationTypeProperty{
//   		sourceArn: jsii.String("sourceArn"),
//
//   		// the properties below are optional
//   		blockEmail: &notifyEmailTypeProperty{
//   			subject: jsii.String("subject"),
//
//   			// the properties below are optional
//   			htmlBody: jsii.String("htmlBody"),
//   			textBody: jsii.String("textBody"),
//   		},
//   		from: jsii.String("from"),
//   		mfaEmail: &notifyEmailTypeProperty{
//   			subject: jsii.String("subject"),
//
//   			// the properties below are optional
//   			htmlBody: jsii.String("htmlBody"),
//   			textBody: jsii.String("textBody"),
//   		},
//   		noActionEmail: &notifyEmailTypeProperty{
//   			subject: jsii.String("subject"),
//
//   			// the properties below are optional
//   			htmlBody: jsii.String("htmlBody"),
//   			textBody: jsii.String("textBody"),
//   		},
//   		replyTo: jsii.String("replyTo"),
//   	},
//   }
//
type CfnUserPoolRiskConfigurationAttachment_AccountTakeoverRiskConfigurationTypeProperty struct {
	// Account takeover risk configuration actions.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The notify configuration used to construct email notifications.
	NotifyConfiguration interface{} `field:"optional" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

