package awscognito


// Properties for defining a `CfnUserPoolRiskConfigurationAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolRiskConfigurationAttachmentProps := &CfnUserPoolRiskConfigurationAttachmentProps{
//   	ClientId: jsii.String("clientId"),
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	AccountTakeoverRiskConfiguration: &AccountTakeoverRiskConfigurationTypeProperty{
//   		Actions: &AccountTakeoverActionsTypeProperty{
//   			HighAction: &AccountTakeoverActionTypeProperty{
//   				EventAction: jsii.String("eventAction"),
//   				Notify: jsii.Boolean(false),
//   			},
//   			LowAction: &AccountTakeoverActionTypeProperty{
//   				EventAction: jsii.String("eventAction"),
//   				Notify: jsii.Boolean(false),
//   			},
//   			MediumAction: &AccountTakeoverActionTypeProperty{
//   				EventAction: jsii.String("eventAction"),
//   				Notify: jsii.Boolean(false),
//   			},
//   		},
//
//   		// the properties below are optional
//   		NotifyConfiguration: &NotifyConfigurationTypeProperty{
//   			SourceArn: jsii.String("sourceArn"),
//
//   			// the properties below are optional
//   			BlockEmail: &NotifyEmailTypeProperty{
//   				Subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				HtmlBody: jsii.String("htmlBody"),
//   				TextBody: jsii.String("textBody"),
//   			},
//   			From: jsii.String("from"),
//   			MfaEmail: &NotifyEmailTypeProperty{
//   				Subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				HtmlBody: jsii.String("htmlBody"),
//   				TextBody: jsii.String("textBody"),
//   			},
//   			NoActionEmail: &NotifyEmailTypeProperty{
//   				Subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				HtmlBody: jsii.String("htmlBody"),
//   				TextBody: jsii.String("textBody"),
//   			},
//   			ReplyTo: jsii.String("replyTo"),
//   		},
//   	},
//   	CompromisedCredentialsRiskConfiguration: &CompromisedCredentialsRiskConfigurationTypeProperty{
//   		Actions: &CompromisedCredentialsActionsTypeProperty{
//   			EventAction: jsii.String("eventAction"),
//   		},
//
//   		// the properties below are optional
//   		EventFilter: []*string{
//   			jsii.String("eventFilter"),
//   		},
//   	},
//   	RiskExceptionConfiguration: &RiskExceptionConfigurationTypeProperty{
//   		BlockedIpRangeList: []*string{
//   			jsii.String("blockedIpRangeList"),
//   		},
//   		SkippedIpRangeList: []*string{
//   			jsii.String("skippedIpRangeList"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html
//
type CfnUserPoolRiskConfigurationAttachmentProps struct {
	// The app client where this configuration is applied.
	//
	// When this parameter isn't present, the risk configuration applies to all user pool app clients that don't have client-level settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The ID of the user pool that has the risk configuration applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-userpoolid
	//
	UserPoolId interface{} `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// The settings for automated responses and notification templates for adaptive authentication with threat protection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfiguration
	//
	AccountTakeoverRiskConfiguration interface{} `field:"optional" json:"accountTakeoverRiskConfiguration" yaml:"accountTakeoverRiskConfiguration"`
	// Settings for compromised-credentials actions and authentication types with threat protection in full-function `ENFORCED` mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfiguration
	//
	CompromisedCredentialsRiskConfiguration interface{} `field:"optional" json:"compromisedCredentialsRiskConfiguration" yaml:"compromisedCredentialsRiskConfiguration"`
	// Exceptions to the risk evaluation configuration, including always-allow and always-block IP address ranges.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-riskexceptionconfiguration
	//
	RiskExceptionConfiguration interface{} `field:"optional" json:"riskExceptionConfiguration" yaml:"riskExceptionConfiguration"`
}

