package awssns


// The throughput quota and deduplication behavior to apply for the FIFO topic.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("MyTopic"), &TopicProps{
//   	Fifo: jsii.Boolean(true),
//   	FifoThroughputScope: sns.FifoThroughputScope_TOPIC,
//   })
//
type FifoThroughputScope string

const (
	// Topic scope - Throughput: 3000 messages per second and a bandwidth of 20MB per second.
	//
	// - Deduplication: Message deduplication is verified on the entire FIFO topic.
	FifoThroughputScope_TOPIC FifoThroughputScope = "TOPIC"
	// Message group scope - Throughput: Maximum regional limits.
	//
	// - Deduplication: Message deduplication is only verified within a message group.
	FifoThroughputScope_MESSAGE_GROUP FifoThroughputScope = "MESSAGE_GROUP"
)

