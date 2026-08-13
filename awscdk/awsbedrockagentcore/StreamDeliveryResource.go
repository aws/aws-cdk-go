package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
)

// A delivery target for real-time streaming of memory record lifecycle events.
//
// Instances are created through the static factory methods, one per delivery
// target type, for example `StreamDeliveryResource.kinesis()`.
//
// Example:
//   stream := kinesis.NewStream(this, jsii.String("MemoryEventStream"))
//
//   memory := agentcore.NewMemory(this, jsii.String("MemoryWithStreamDelivery"), &MemoryProps{
//   	MemoryName: jsii.String("memory_with_stream"),
//   	StreamDeliveryResources: []StreamDeliveryResource{
//   		agentcore.StreamDeliveryResource_Kinesis(stream, &KinesisStreamDeliveryOptions{
//   			ContentConfigurations: []StreamDeliveryContentConfiguration{
//   				&StreamDeliveryContentConfiguration{
//   					Type: agentcore.StreamDeliveryContentType_MEMORY_RECORDS,
//   					// Streams complete memory record bodies, which may include sensitive data
//   					Level: agentcore.StreamDeliveryContentLevel_FULL_CONTENT,
//   				},
//   			},
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/memory-record-streaming.html
//
type StreamDeliveryResource interface {
	// Content configurations defining what is delivered to the stream.
	ContentConfigurations() *[]*StreamDeliveryContentConfiguration
	// The Kinesis Data Stream that memory record events are delivered to.
	Stream() awskinesis.IStream
}

// The jsii proxy struct for StreamDeliveryResource
type jsiiProxy_StreamDeliveryResource struct {
	_ byte // padding
}

func (j *jsiiProxy_StreamDeliveryResource) ContentConfigurations() *[]*StreamDeliveryContentConfiguration {
	var returns *[]*StreamDeliveryContentConfiguration
	_jsii_.Get(
		j,
		"contentConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamDeliveryResource) Stream() awskinesis.IStream {
	var returns awskinesis.IStream
	_jsii_.Get(
		j,
		"stream",
		&returns,
	)
	return returns
}


// Deliver memory record lifecycle events to an Amazon Kinesis Data Stream.
//
// The memory execution role is automatically granted write permissions to
// the stream.
func StreamDeliveryResource_Kinesis(stream awskinesis.IStream, options *KinesisStreamDeliveryOptions) StreamDeliveryResource {
	_init_.Initialize()

	if err := validateStreamDeliveryResource_KinesisParameters(stream, options); err != nil {
		panic(err)
	}
	var returns StreamDeliveryResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.StreamDeliveryResource",
		"kinesis",
		[]interface{}{stream, options},
		&returns,
	)

	return returns
}

