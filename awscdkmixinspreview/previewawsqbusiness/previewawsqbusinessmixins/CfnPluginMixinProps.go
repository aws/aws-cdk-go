package previewawsqbusinessmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPluginPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var noAuthConfiguration interface{}
//
//   cfnPluginMixinProps := &CfnPluginMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	AuthConfiguration: &PluginAuthConfigurationProperty{
//   		BasicAuthConfiguration: &BasicAuthConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		NoAuthConfiguration: noAuthConfiguration,
//   		OAuth2ClientCredentialConfiguration: &OAuth2ClientCredentialConfigurationProperty{
//   			AuthorizationUrl: jsii.String("authorizationUrl"),
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   			TokenUrl: jsii.String("tokenUrl"),
//   		},
//   	},
//   	CustomPluginConfiguration: &CustomPluginConfigurationProperty{
//   		ApiSchema: &APISchemaProperty{
//   			Payload: jsii.String("payload"),
//   			S3: &S3Property{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		ApiSchemaType: jsii.String("apiSchemaType"),
//   		Description: jsii.String("description"),
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	ServerUrl: jsii.String("serverUrl"),
//   	State: jsii.String("state"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html
//
type CfnPluginMixinProps struct {
	// The identifier of the application that will contain the plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Authentication configuration information for an Amazon Q Business plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-authconfiguration
	//
	AuthConfiguration interface{} `field:"optional" json:"authConfiguration" yaml:"authConfiguration"`
	// Configuration information required to create a custom plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-custompluginconfiguration
	//
	CustomPluginConfiguration interface{} `field:"optional" json:"customPluginConfiguration" yaml:"customPluginConfiguration"`
	// The name of the plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The plugin server URL used for configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-serverurl
	//
	ServerUrl *string `field:"optional" json:"serverUrl" yaml:"serverUrl"`
	// The current status of the plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// A list of key-value pairs that identify or categorize the data source connector.
	//
	// You can also use tags to help control access to the data source connector. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + -
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

