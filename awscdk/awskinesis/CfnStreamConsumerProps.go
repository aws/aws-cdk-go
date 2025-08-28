package awskinesis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnStreamConsumer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamConsumerProps := &CfnStreamConsumerProps{
//   	ConsumerName: jsii.String("consumerName"),
//   	StreamArn: jsii.String("streamArn"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html
//
type CfnStreamConsumerProps struct {
	// The name of the consumer is something you choose when you register the consumer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html#cfn-kinesis-streamconsumer-consumername
	//
	ConsumerName *string `field:"required" json:"consumerName" yaml:"consumerName"`
	// The ARN of the stream with which you registered the consumer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html#cfn-kinesis-streamconsumer-streamarn
	//
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
	// An array of tags to be added to a specified Kinesis resource.
	//
	// A tag consists of a required key and an optional value. You can specify up to 50 tag key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html#cfn-kinesis-streamconsumer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

