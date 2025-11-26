package previewawscognitomixins


// Properties for CfnUserPoolRiskConfigurationAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserPoolRiskConfigurationAttachmentMixinProps := &CfnUserPoolRiskConfigurationAttachmentMixinProps{
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
//   		NotifyConfiguration: &NotifyConfigurationTypeProperty{
//   			BlockEmail: &NotifyEmailTypeProperty{
//   				HtmlBody: jsii.String("htmlBody"),
//   				Subject: jsii.String("subject"),
//   				TextBody: jsii.String("textBody"),
//   			},
//   			From: jsii.String("from"),
//   			MfaEmail: &NotifyEmailTypeProperty{
//   				HtmlBody: jsii.String("htmlBody"),
//   				Subject: jsii.String("subject"),
//   				TextBody: jsii.String("textBody"),
//   			},
//   			NoActionEmail: &NotifyEmailTypeProperty{
//   				HtmlBody: jsii.String("htmlBody"),
//   				Subject: jsii.String("subject"),
//   				TextBody: jsii.String("textBody"),
//   			},
//   			ReplyTo: jsii.String("replyTo"),
//   			SourceArn: jsii.String("sourceArn"),
//   		},
//   	},
//   	ClientId: jsii.String("clientId"),
//   	CompromisedCredentialsRiskConfiguration: &CompromisedCredentialsRiskConfigurationTypeProperty{
//   		Actions: &CompromisedCredentialsActionsTypeProperty{
//   			EventAction: jsii.String("eventAction"),
//   		},
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
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html
//
type CfnUserPoolRiskConfigurationAttachmentMixinProps struct {
	// The settings for automated responses and notification templates for adaptive authentication with threat protection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfiguration
	//
	AccountTakeoverRiskConfiguration interface{} `field:"optional" json:"accountTakeoverRiskConfiguration" yaml:"accountTakeoverRiskConfiguration"`
	// The app client where this configuration is applied.
	//
	// When this parameter isn't present, the risk configuration applies to all user pool app clients that don't have client-level settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Settings for compromised-credentials actions and authentication types with threat protection in full-function `ENFORCED` mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfiguration
	//
	CompromisedCredentialsRiskConfiguration interface{} `field:"optional" json:"compromisedCredentialsRiskConfiguration" yaml:"compromisedCredentialsRiskConfiguration"`
	// Exceptions to the risk evaluation configuration, including always-allow and always-block IP address ranges.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-riskexceptionconfiguration
	//
	RiskExceptionConfiguration interface{} `field:"optional" json:"riskExceptionConfiguration" yaml:"riskExceptionConfiguration"`
	// The ID of the user pool that has the risk configuration applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-userpoolid
	//
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

