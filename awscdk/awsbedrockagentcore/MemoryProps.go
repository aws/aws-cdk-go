package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Memory resource.
//
// Example:
//   // Create a Kinesis Data Stream
//   stream := kinesis.NewStream(this, jsii.String("MemoryEventStream"), &StreamProps{
//   	StreamName: jsii.String("memory-events"),
//   })
//
//   memory := agentcore.NewMemory(this, jsii.String("MemoryWithStreamDelivery"), &MemoryProps{
//   	MemoryName: jsii.String("memory_with_stream"),
//   	Description: jsii.String("Memory with Kinesis stream delivery"),
//   	ExpirationDuration: cdk.Duration_Days(jsii.Number(90)),
//   	StreamDeliveryResources: []StreamDeliveryResource{
//   		agentcore.StreamDeliveryResource_Kinesis(stream, &KinesisStreamDeliveryOptions{
//   			ContentConfigurations: []StreamDeliveryContentConfiguration{
//   				&StreamDeliveryContentConfiguration{
//   					Type: agentcore.StreamDeliveryContentType_MEMORY_RECORDS,
//   					Level: agentcore.StreamDeliveryContentLevel_METADATA_ONLY,
//   				},
//   			},
//   		}),
//   	},
//   })
//
type MemoryProps struct {
	// Optional description for the memory Valid characters are a-z, A-Z, 0-9, _ (underscore), - (hyphen) and spaces The description can have up to 200 characters.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM role that provides permissions for the memory to access AWS services when using custom strategies.
	// Default: - A new role will be created.
	//
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Short-term memory expiration in days (between 7 and 365).
	//
	// Sets the short-term (raw event) memory retention.
	// Events older than the specified duration will expire and no longer be stored.
	// Default: - 90 days.
	//
	ExpirationDuration awscdk.Duration `field:"optional" json:"expirationDuration" yaml:"expirationDuration"`
	// Custom KMS key to use for encryption.
	// Default: - Your data is encrypted with a key that AWS owns and manages for you.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The name of the memory Valid characters are a-z, A-Z, 0-9, _ (underscore) The name must start with a letter and can be up to 48 characters long Pattern: [a-zA-Z][a-zA-Z0-9_]{0,47}.
	// Default: - auto generate.
	//
	MemoryName *string `field:"optional" json:"memoryName" yaml:"memoryName"`
	// If you need long-term memory for context recall across sessions, you can setup memory extraction strategies to extract the relevant memory from the raw events.
	// Default: - No extraction strategies (short term memory only).
	//
	MemoryStrategies *[]IMemoryStrategy `field:"optional" json:"memoryStrategies" yaml:"memoryStrategies"`
	// Stream delivery resources for real-time push-based streaming of memory record lifecycle events (created, updated, deleted) to Amazon Kinesis Data Streams.
	//
	// The memory execution role will automatically be granted write permissions to each stream.
	//
	// Only one stream delivery resource is currently supported (CloudFormation maximum);
	// providing more than one fails at synth with `TooManyStreamDeliveryResources`:
	//
	// ```ts
	// declare const stream: kinesis.IStream;
	// new agentcore.Memory(this, 'Memory', {
	//   streamDeliveryResources: [
	//     agentcore.StreamDeliveryResource.kinesis(stream, {
	//       contentConfigurations: [{
	//         type: agentcore.StreamDeliveryContentType.MEMORY_RECORDS,
	//         level: agentcore.StreamDeliveryContentLevel.METADATA_ONLY,
	//       }],
	//     }),
	//   ],
	// });
	// ```.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/memory-record-streaming.html
	//
	// Default: - No stream delivery (events are not pushed to Kinesis).
	//
	StreamDeliveryResources *[]StreamDeliveryResource `field:"optional" json:"streamDeliveryResources" yaml:"streamDeliveryResources"`
	// Tags (optional) A list of key:value pairs of tags to apply to this memory resource.
	// Default: - no tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

