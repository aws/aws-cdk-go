package previewawsamazonmqmixins


// Properties for CfnConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationMixinProps := &CfnConfigurationMixinProps{
//   	AuthenticationStrategy: jsii.String("authenticationStrategy"),
//   	Data: jsii.String("data"),
//   	Description: jsii.String("description"),
//   	EngineType: jsii.String("engineType"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	Name: jsii.String("name"),
//   	Tags: []TagsEntryProperty{
//   		&TagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html
//
type CfnConfigurationMixinProps struct {
	// Optional.
	//
	// The authentication strategy associated with the configuration. The default is `SIMPLE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-authenticationstrategy
	//
	AuthenticationStrategy *string `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// Amazon MQ for Active MQ: The base64-encoded XML configuration.
	//
	// Amazon MQ for RabbitMQ: the base64-encoded Cuttlefish configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The description of the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Required.
	//
	// The type of broker engine. Currently, Amazon MQ supports `ACTIVEMQ` and `RABBITMQ` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-enginetype
	//
	EngineType *string `field:"optional" json:"engineType" yaml:"engineType"`
	// The broker engine version.
	//
	// Defaults to the latest available version for the specified broker engine type. For more information, see the [ActiveMQ version management](https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/activemq-version-management.html) and the [RabbitMQ version management](https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/rabbitmq-version-management.html) sections in the Amazon MQ Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Required.
	//
	// The name of the configuration. This value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~). This value must be 1-150 characters long.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Create tags when creating the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-tags
	//
	Tags *[]*CfnConfigurationPropsMixin_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

