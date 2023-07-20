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
	// The app client ID.
	//
	// You can specify the risk configuration for a single client (with a specific ClientId) or for all clients (by setting the ClientId to `ALL` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The user pool ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// The account takeover risk configuration object, including the `NotifyConfiguration` object and `Actions` to take if there is an account takeover.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfiguration
	//
	AccountTakeoverRiskConfiguration interface{} `field:"optional" json:"accountTakeoverRiskConfiguration" yaml:"accountTakeoverRiskConfiguration"`
	// The compromised credentials risk configuration object, including the `EventFilter` and the `EventAction` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfiguration
	//
	CompromisedCredentialsRiskConfiguration interface{} `field:"optional" json:"compromisedCredentialsRiskConfiguration" yaml:"compromisedCredentialsRiskConfiguration"`
	// The configuration to override the risk decision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-riskexceptionconfiguration
	//
	RiskExceptionConfiguration interface{} `field:"optional" json:"riskExceptionConfiguration" yaml:"riskExceptionConfiguration"`
}

