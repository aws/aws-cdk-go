package awsqbusiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPlugin`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var noAuthConfiguration interface{}
//
//   cfnPluginProps := &CfnPluginProps{
//   	AuthConfiguration: &PluginAuthConfigurationProperty{
//   		BasicAuthConfiguration: &BasicAuthConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		NoAuthConfiguration: noAuthConfiguration,
//   		OAuth2ClientCredentialConfiguration: &OAuth2ClientCredentialConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//
//   			// the properties below are optional
//   			AuthorizationUrl: jsii.String("authorizationUrl"),
//   			TokenUrl: jsii.String("tokenUrl"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ApplicationId: jsii.String("applicationId"),
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
//   	ServerUrl: jsii.String("serverUrl"),
//   	State: jsii.String("state"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html
//
type CfnPluginProps struct {
	// Authentication configuration information for an Amazon Q Business plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-authconfiguration
	//
	AuthConfiguration interface{} `field:"required" json:"authConfiguration" yaml:"authConfiguration"`
	// The name of the plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The type of the plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The identifier of the application that will contain the plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Configuration information required to create a custom plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-plugin.html#cfn-qbusiness-plugin-custompluginconfiguration
	//
	CustomPluginConfiguration interface{} `field:"optional" json:"customPluginConfiguration" yaml:"customPluginConfiguration"`
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
}

