package mixinsawsiot


// Specifies a Kafka header using key-value pairs when you create a Rule’s Kafka Action.
//
// You can use these headers to route data from IoT clients to downstream Kafka clusters without modifying your message payload.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kafkaActionHeaderProperty := &KafkaActionHeaderProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaactionheader.html
//
type CfnTopicRulePropsMixin_KafkaActionHeaderProperty struct {
	// The key of the Kafka header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaactionheader.html#cfn-iot-topicrule-kafkaactionheader-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the Kafka header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kafkaactionheader.html#cfn-iot-topicrule-kafkaactionheader-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

