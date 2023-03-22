package awsamazonmq


// Properties for defining a `CfnConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationProps := &cfnConfigurationProps{
//   	data: jsii.String("data"),
//   	engineType: jsii.String("engineType"),
//   	engineVersion: jsii.String("engineVersion"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	authenticationStrategy: jsii.String("authenticationStrategy"),
//   	description: jsii.String("description"),
//   	tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConfigurationProps struct {
	// The base64-encoded XML configuration.
	Data *string `field:"required" json:"data" yaml:"data"`
	// The type of broker engine.
	//
	// Note: Currently, Amazon MQ only supports ACTIVEMQ for creating and editing broker configurations.
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// The version of the broker engine.
	//
	// For a list of supported engine versions, see [](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html)
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// The name of the configuration.
	//
	// This value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~). This value must be 1-150 characters long.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional.
	//
	// The authentication strategy associated with the configuration. The default is `SIMPLE` .
	AuthenticationStrategy *string `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// The description of the configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Create tags when creating the configuration.
	Tags *[]*CfnConfiguration_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

