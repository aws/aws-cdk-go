package mixinsawsmsk


// Properties for CfnConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationMixinProps := &CfnConfigurationMixinProps{
//   	Description: jsii.String("description"),
//   	KafkaVersionsList: []*string{
//   		jsii.String("kafkaVersionsList"),
//   	},
//   	LatestRevision: &LatestRevisionProperty{
//   		CreationTime: jsii.String("creationTime"),
//   		Description: jsii.String("description"),
//   		Revision: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	ServerProperties: jsii.String("serverProperties"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html
//
type CfnConfigurationMixinProps struct {
	// The description of the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html#cfn-msk-configuration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The [versions of Apache Kafka](https://docs.aws.amazon.com/msk/latest/developerguide/supported-kafka-versions.html) with which you can use this MSK configuration.
	//
	// When you update the `KafkaVersionsList` property, CloudFormation recreates a new configuration with the updated property before deleting the old configuration. Such an update requires a [resource replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) . To successfully update `KafkaVersionsList` , you must also update the `Name` property in the same operation.
	//
	// If your configuration is attached with any clusters created using the AWS Management Console or AWS CLI , you'll need to manually delete the old configuration from the console after the update completes.
	//
	// For more information, see [Can’t update KafkaVersionsList in MSK configuration](https://docs.aws.amazon.com/msk/latest/developerguide/troubleshooting.html#troubleshoot-kafkaversionslist-cfn-update-failure) in the *Amazon MSK Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html#cfn-msk-configuration-kafkaversionslist
	//
	KafkaVersionsList *[]*string `field:"optional" json:"kafkaVersionsList" yaml:"kafkaVersionsList"`
	// Latest revision of the MSK configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html#cfn-msk-configuration-latestrevision
	//
	LatestRevision interface{} `field:"optional" json:"latestRevision" yaml:"latestRevision"`
	// The name of the configuration.
	//
	// Configuration names are strings that match the regex "^[0-9A-Za-z][0-9A-Za-z-]{0,}$".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html#cfn-msk-configuration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Contents of the `server.properties` file. When using the console, the SDK, or the AWS CLI , the contents of `server.properties` can be in plaintext.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html#cfn-msk-configuration-serverproperties
	//
	ServerProperties *string `field:"optional" json:"serverProperties" yaml:"serverProperties"`
}

