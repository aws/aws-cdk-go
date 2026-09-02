package awsmsk


// Configuration of topic in a channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicConfigurationProperty := &TopicConfigurationProperty{
//   	RecordConverter: &RecordConverterProperty{
//   		ValueConverter: jsii.String("valueConverter"),
//   	},
//   	TopicArn: jsii.String("topicArn"),
//
//   	// the properties below are optional
//   	RecordSchema: &RecordSchemaProperty{
//   		GsrArn: jsii.String("gsrArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-topicconfiguration.html
//
type CfnChannel_TopicConfigurationProperty struct {
	// Record converter configuration for a topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-topicconfiguration.html#cfn-msk-channel-topicconfiguration-recordconverter
	//
	RecordConverter interface{} `field:"required" json:"recordConverter" yaml:"recordConverter"`
	// The Amazon Resource Name (ARN) that uniquely identifies the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-topicconfiguration.html#cfn-msk-channel-topicconfiguration-topicarn
	//
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// Record schema configuration for a topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-topicconfiguration.html#cfn-msk-channel-topicconfiguration-recordschema
	//
	RecordSchema interface{} `field:"optional" json:"recordSchema" yaml:"recordSchema"`
}

