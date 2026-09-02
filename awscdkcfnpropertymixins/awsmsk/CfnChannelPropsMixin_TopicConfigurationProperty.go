package awsmsk


// Configuration of topic in a channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   topicConfigurationProperty := &TopicConfigurationProperty{
//   	RecordConverter: &RecordConverterProperty{
//   		ValueConverter: jsii.String("valueConverter"),
//   	},
//   	RecordSchema: &RecordSchemaProperty{
//   		GsrArn: jsii.String("gsrArn"),
//   	},
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-topicconfiguration.html
//
type CfnChannelPropsMixin_TopicConfigurationProperty struct {
	// Record converter configuration for a topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-topicconfiguration.html#cfn-msk-channel-topicconfiguration-recordconverter
	//
	RecordConverter interface{} `field:"optional" json:"recordConverter" yaml:"recordConverter"`
	// Record schema configuration for a topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-topicconfiguration.html#cfn-msk-channel-topicconfiguration-recordschema
	//
	RecordSchema interface{} `field:"optional" json:"recordSchema" yaml:"recordSchema"`
	// The Amazon Resource Name (ARN) that uniquely identifies the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-topicconfiguration.html#cfn-msk-channel-topicconfiguration-topicarn
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

