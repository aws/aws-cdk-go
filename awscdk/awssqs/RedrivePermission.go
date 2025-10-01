package awssqs


// The permission type that defines which source queues can specify the current queue as the dead-letter queue.
//
// Example:
//   var sourceQueue iQueue
//
//
//   // Only the sourceQueue can specify this queue as the dead-letter queue.
//   queue1 := sqs.NewQueue(this, jsii.String("Queue2"), &QueueProps{
//   	RedriveAllowPolicy: &RedriveAllowPolicy{
//   		SourceQueues: []*iQueue{
//   			sourceQueue,
//   		},
//   	},
//   })
//
//   // No source queues can specify this queue as the dead-letter queue.
//   queue2 := sqs.NewQueue(this, jsii.String("Queue"), &QueueProps{
//   	RedriveAllowPolicy: &RedriveAllowPolicy{
//   		RedrivePermission: sqs.RedrivePermission_DENY_ALL,
//   	},
//   })
//
type RedrivePermission string

const (
	// Any source queues in this AWS account in the same Region can specify this queue as the dead-letter queue.
	RedrivePermission_ALLOW_ALL RedrivePermission = "ALLOW_ALL"
	// No source queues can specify this queue as the dead-letter queue.
	RedrivePermission_DENY_ALL RedrivePermission = "DENY_ALL"
	// Only queues specified by the `sourceQueueArns` parameter can specify this queue as the dead-letter queue.
	RedrivePermission_BY_QUEUE RedrivePermission = "BY_QUEUE"
)

