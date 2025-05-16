package awsamazonmq


// Properties for defining a `CfnConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationProps := &CfnConfigurationProps{
//   	EngineType: jsii.String("engineType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AuthenticationStrategy: jsii.String("authenticationStrategy"),
//   	Data: jsii.String("data"),
//   	Description: jsii.String("description"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	Tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html
//
type CfnConfigurationProps struct {
	// The type of broker engine.
	//
	// Note: Currently, Amazon MQ only supports ACTIVEMQ for creating and editing broker configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-enginetype
	//
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// The name of the configuration.
	//
	// This value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~). This value must be 1-150 characters long.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional.
	//
	// The authentication strategy associated with the configuration. The default is `SIMPLE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-authenticationstrategy
	//
	AuthenticationStrategy *string `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// The base64-encoded XML configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The description of the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version of the broker engine.
	//
	// For a list of supported engine versions, see [](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Create tags when creating the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-tags
	//
	Tags *[]*CfnConfiguration_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

