package awsbedrockagentcore


// Content type for stream delivery.
//
// Defines what kind of memory content is delivered to the Kinesis stream.
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
type StreamDeliveryContentType string

const (
	// Deliver memory record lifecycle events (created, updated, deleted).
	StreamDeliveryContentType_MEMORY_RECORDS StreamDeliveryContentType = "MEMORY_RECORDS"
)

