package awsbedrockagentcore


// Content detail level for stream delivery.
//
// Controls how much detail is included in each delivered record.
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
type StreamDeliveryContentLevel string

const (
	// Deliver only metadata (record ID, timestamps, event type).
	StreamDeliveryContentLevel_METADATA_ONLY StreamDeliveryContentLevel = "METADATA_ONLY"
	// Deliver full content including the memory record body.
	StreamDeliveryContentLevel_FULL_CONTENT StreamDeliveryContentLevel = "FULL_CONTENT"
)

