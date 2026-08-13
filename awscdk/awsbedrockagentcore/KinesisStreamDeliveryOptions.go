package awsbedrockagentcore


// Options for delivering memory record events to a Kinesis Data Stream.
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
type KinesisStreamDeliveryOptions struct {
	// Content configurations defining what to deliver to the stream.
	//
	// Currently exactly one configuration is supported.
	ContentConfigurations *[]*StreamDeliveryContentConfiguration `field:"required" json:"contentConfigurations" yaml:"contentConfigurations"`
}

