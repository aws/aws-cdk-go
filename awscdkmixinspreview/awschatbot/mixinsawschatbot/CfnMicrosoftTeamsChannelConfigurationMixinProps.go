package mixinsawschatbot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMicrosoftTeamsChannelConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMicrosoftTeamsChannelConfigurationMixinProps := &CfnMicrosoftTeamsChannelConfigurationMixinProps{
//   	ConfigurationName: jsii.String("configurationName"),
//   	CustomizationResourceArns: []*string{
//   		jsii.String("customizationResourceArns"),
//   	},
//   	GuardrailPolicies: []*string{
//   		jsii.String("guardrailPolicies"),
//   	},
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	LoggingLevel: jsii.String("loggingLevel"),
//   	SnsTopicArns: []*string{
//   		jsii.String("snsTopicArns"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TeamId: jsii.String("teamId"),
//   	TeamsChannelId: jsii.String("teamsChannelId"),
//   	TeamsChannelName: jsii.String("teamsChannelName"),
//   	TeamsTenantId: jsii.String("teamsTenantId"),
//   	UserRoleRequired: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html
//
type CfnMicrosoftTeamsChannelConfigurationMixinProps struct {
	// The name of the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-configurationname
	//
	ConfigurationName *string `field:"optional" json:"configurationName" yaml:"configurationName"`
	// Links a list of resource ARNs (for example, custom action ARNs) to a Microsoft Teams channel configuration for  .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-customizationresourcearns
	//
	CustomizationResourceArns *[]*string `field:"optional" json:"customizationResourceArns" yaml:"customizationResourceArns"`
	// The list of IAM policy ARNs that are applied as channel guardrails.
	//
	// The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-guardrailpolicies
	//
	GuardrailPolicies *[]*string `field:"optional" json:"guardrailPolicies" yaml:"guardrailPolicies"`
	// The ARN of the IAM role that defines the permissions for  .
	//
	// This is a user-defined role that  will assume. This is not the service-linked role. For more information, see [IAM Policies for  in chat applications](https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-iamrolearn
	//
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Specifies the logging level for this configuration. This property affects the log entries pushed to Amazon CloudWatch Logs.
	//
	// Logging levels include `ERROR` , `INFO` , or `NONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-logginglevel
	//
	// Default: - "NONE".
	//
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// The ARNs of the SNS topics that deliver notifications to  .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-snstopicarns
	//
	SnsTopicArns *[]*string `field:"optional" json:"snsTopicArns" yaml:"snsTopicArns"`
	// The tags to add to the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the Microsoft Team authorized with  .
	//
	// To get the team ID, you must perform the initial authorization flow with Microsoft Teams in the  in chat applications console. Then you can copy and paste the team ID from the console. For more details, see steps 1-3 in [Tutorial: Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html) in the *in chat applications Administrator Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-teamid
	//
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// The ID of the Microsoft Teams channel.
	//
	// To get the channel ID, open Microsoft Teams, right click on the channel name in the left pane, then choose *Copy* . An example of the channel ID syntax is: `19%3ab6ef35dc342d56ba5654e6fc6d25a071%40thread.tacv2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-teamschannelid
	//
	TeamsChannelId *string `field:"optional" json:"teamsChannelId" yaml:"teamsChannelId"`
	// The name of the Microsoft Teams channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-teamschannelname
	//
	TeamsChannelName *string `field:"optional" json:"teamsChannelName" yaml:"teamsChannelName"`
	// The ID of the Microsoft Teams tenant.
	//
	// To get the tenant ID, you must perform the initial authorization flow with Microsoft Teams in the  in chat applications console. Then you can copy and paste the tenant ID from the console. For more details, see steps 1-3 in [Tutorial: Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html) in the *in chat applications Administrator Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-teamstenantid
	//
	TeamsTenantId *string `field:"optional" json:"teamsTenantId" yaml:"teamsTenantId"`
	// Enables use of a user role requirement in your chat configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-userrolerequired
	//
	// Default: - false.
	//
	UserRoleRequired interface{} `field:"optional" json:"userRoleRequired" yaml:"userRoleRequired"`
}

