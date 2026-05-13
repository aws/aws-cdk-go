package awschime

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAppInstanceBotPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAppInstanceBotMixinProps := &CfnAppInstanceBotMixinProps{
//   	AppInstanceArn: jsii.String("appInstanceArn"),
//   	Configuration: &ConfigurationProperty{
//   		Lex: &LexConfigurationProperty{
//   			InvokedBy: &InvokedByProperty{
//   				StandardMessages: jsii.String("standardMessages"),
//   				TargetedMessages: jsii.String("targetedMessages"),
//   			},
//   			LexBotAliasArn: jsii.String("lexBotAliasArn"),
//   			LocaleId: jsii.String("localeId"),
//   			RespondsTo: jsii.String("respondsTo"),
//   			WelcomeIntent: jsii.String("welcomeIntent"),
//   		},
//   	},
//   	Metadata: jsii.String("metadata"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstancebot.html
//
type CfnAppInstanceBotMixinProps struct {
	// The ARN of the AppInstance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstancebot.html#cfn-chime-appinstancebot-appinstancearn
	//
	AppInstanceArn *string `field:"optional" json:"appInstanceArn" yaml:"appInstanceArn"`
	// A structure that contains configuration data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstancebot.html#cfn-chime-appinstancebot-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The metadata of the AppInstanceBot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstancebot.html#cfn-chime-appinstancebot-metadata
	//
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// The name of the AppInstanceBot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstancebot.html#cfn-chime-appinstancebot-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags assigned to the AppInstanceBot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstancebot.html#cfn-chime-appinstancebot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

