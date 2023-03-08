package awscognito


// Properties for defining a `CfnUserPoolRiskConfigurationAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolRiskConfigurationAttachmentProps := &cfnUserPoolRiskConfigurationAttachmentProps{
//   	clientId: jsii.String("clientId"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	accountTakeoverRiskConfiguration: &accountTakeoverRiskConfigurationTypeProperty{
//   		actions: &accountTakeoverActionsTypeProperty{
//   			highAction: &accountTakeoverActionTypeProperty{
//   				eventAction: jsii.String("eventAction"),
//   				notify: jsii.Boolean(false),
//   			},
//   			lowAction: &accountTakeoverActionTypeProperty{
//   				eventAction: jsii.String("eventAction"),
//   				notify: jsii.Boolean(false),
//   			},
//   			mediumAction: &accountTakeoverActionTypeProperty{
//   				eventAction: jsii.String("eventAction"),
//   				notify: jsii.Boolean(false),
//   			},
//   		},
//
//   		// the properties below are optional
//   		notifyConfiguration: &notifyConfigurationTypeProperty{
//   			sourceArn: jsii.String("sourceArn"),
//
//   			// the properties below are optional
//   			blockEmail: &notifyEmailTypeProperty{
//   				subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				htmlBody: jsii.String("htmlBody"),
//   				textBody: jsii.String("textBody"),
//   			},
//   			from: jsii.String("from"),
//   			mfaEmail: &notifyEmailTypeProperty{
//   				subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				htmlBody: jsii.String("htmlBody"),
//   				textBody: jsii.String("textBody"),
//   			},
//   			noActionEmail: &notifyEmailTypeProperty{
//   				subject: jsii.String("subject"),
//
//   				// the properties below are optional
//   				htmlBody: jsii.String("htmlBody"),
//   				textBody: jsii.String("textBody"),
//   			},
//   			replyTo: jsii.String("replyTo"),
//   		},
//   	},
//   	compromisedCredentialsRiskConfiguration: &compromisedCredentialsRiskConfigurationTypeProperty{
//   		actions: &compromisedCredentialsActionsTypeProperty{
//   			eventAction: jsii.String("eventAction"),
//   		},
//
//   		// the properties below are optional
//   		eventFilter: []*string{
//   			jsii.String("eventFilter"),
//   		},
//   	},
//   	riskExceptionConfiguration: &riskExceptionConfigurationTypeProperty{
//   		blockedIpRangeList: []*string{
//   			jsii.String("blockedIpRangeList"),
//   		},
//   		skippedIpRangeList: []*string{
//   			jsii.String("skippedIpRangeList"),
//   		},
//   	},
//   }
//
type CfnUserPoolRiskConfigurationAttachmentProps struct {
	// The app client ID.
	//
	// You can specify the risk configuration for a single client (with a specific ClientId) or for all clients (by setting the ClientId to `ALL` ).
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The user pool ID.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// The account takeover risk configuration object, including the `NotifyConfiguration` object and `Actions` to take if there is an account takeover.
	AccountTakeoverRiskConfiguration interface{} `field:"optional" json:"accountTakeoverRiskConfiguration" yaml:"accountTakeoverRiskConfiguration"`
	// The compromised credentials risk configuration object, including the `EventFilter` and the `EventAction` .
	CompromisedCredentialsRiskConfiguration interface{} `field:"optional" json:"compromisedCredentialsRiskConfiguration" yaml:"compromisedCredentialsRiskConfiguration"`
	// The configuration to override the risk decision.
	RiskExceptionConfiguration interface{} `field:"optional" json:"riskExceptionConfiguration" yaml:"riskExceptionConfiguration"`
}

