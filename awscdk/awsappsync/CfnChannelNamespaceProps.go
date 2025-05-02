package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnChannelNamespace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelNamespaceProps := &CfnChannelNamespaceProps{
//   	ApiId: jsii.String("apiId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CodeHandlers: jsii.String("codeHandlers"),
//   	CodeS3Location: jsii.String("codeS3Location"),
//   	HandlerConfigs: &HandlerConfigsProperty{
//   		OnPublish: &HandlerConfigProperty{
//   			Behavior: jsii.String("behavior"),
//   			Integration: &IntegrationProperty{
//   				DataSourceName: jsii.String("dataSourceName"),
//
//   				// the properties below are optional
//   				LambdaConfig: &LambdaConfigProperty{
//   					InvokeType: jsii.String("invokeType"),
//   				},
//   			},
//   		},
//   		OnSubscribe: &HandlerConfigProperty{
//   			Behavior: jsii.String("behavior"),
//   			Integration: &IntegrationProperty{
//   				DataSourceName: jsii.String("dataSourceName"),
//
//   				// the properties below are optional
//   				LambdaConfig: &LambdaConfigProperty{
//   					InvokeType: jsii.String("invokeType"),
//   				},
//   			},
//   		},
//   	},
//   	PublishAuthModes: []interface{}{
//   		&AuthModeProperty{
//   			AuthType: jsii.String("authType"),
//   		},
//   	},
//   	SubscribeAuthModes: []interface{}{
//   		&AuthModeProperty{
//   			AuthType: jsii.String("authType"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-channelnamespace.html
//
type CfnChannelNamespaceProps struct {
	// AppSync Api Id that this Channel Namespace belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-channelnamespace.html#cfn-appsync-channelnamespace-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Namespace indentifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-channelnamespace.html#cfn-appsync-channelnamespace-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// String of APPSYNC_JS code to be used by the handlers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-channelnamespace.html#cfn-appsync-channelnamespace-codehandlers
	//
	CodeHandlers *string `field:"optional" json:"codeHandlers" yaml:"codeHandlers"`
	// The Amazon S3 endpoint where the code is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-channelnamespace.html#cfn-appsync-channelnamespace-codes3location
	//
	CodeS3Location *string `field:"optional" json:"codeS3Location" yaml:"codeS3Location"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-channelnamespace.html#cfn-appsync-channelnamespace-handlerconfigs
	//
	HandlerConfigs interface{} `field:"optional" json:"handlerConfigs" yaml:"handlerConfigs"`
	// The authorization mode to use for publishing messages on the channel namespace.
	//
	// This configuration overrides the default `Api` authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-channelnamespace.html#cfn-appsync-channelnamespace-publishauthmodes
	//
	PublishAuthModes interface{} `field:"optional" json:"publishAuthModes" yaml:"publishAuthModes"`
	// The authorization mode to use for subscribing to messages on the channel namespace.
	//
	// This configuration overrides the default `Api` authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-channelnamespace.html#cfn-appsync-channelnamespace-subscribeauthmodes
	//
	SubscribeAuthModes interface{} `field:"optional" json:"subscribeAuthModes" yaml:"subscribeAuthModes"`
	// An arbitrary set of tags (key-value pairs) for this AppSync API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-channelnamespace.html#cfn-appsync-channelnamespace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

