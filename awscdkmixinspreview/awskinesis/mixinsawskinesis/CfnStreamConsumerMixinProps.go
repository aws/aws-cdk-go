package mixinsawskinesis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStreamConsumerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStreamConsumerMixinProps := &CfnStreamConsumerMixinProps{
//   	ConsumerName: jsii.String("consumerName"),
//   	StreamArn: jsii.String("streamArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html
//
type CfnStreamConsumerMixinProps struct {
	// The name of the consumer is something you choose when you register the consumer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html#cfn-kinesis-streamconsumer-consumername
	//
	ConsumerName *string `field:"optional" json:"consumerName" yaml:"consumerName"`
	// The ARN of the stream with which you registered the consumer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html#cfn-kinesis-streamconsumer-streamarn
	//
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
	// An array of tags to be added to a specified Kinesis resource.
	//
	// A tag consists of a required key and an optional value. You can specify up to 50 tag key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html#cfn-kinesis-streamconsumer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

