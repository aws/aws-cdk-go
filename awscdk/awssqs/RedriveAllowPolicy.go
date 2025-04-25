package awssqs


// Permission settings for the dead letter source queue.
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
type RedriveAllowPolicy struct {
	// Permission settings for source queues that can designate this queue as their dead-letter queue.
	// Default: - `RedrivePermission.BY_QUEUE` if `sourceQueues` is specified,`RedrivePermission.ALLOW_ALL` otherwise.
	//
	RedrivePermission RedrivePermission `field:"optional" json:"redrivePermission" yaml:"redrivePermission"`
	// Source queues that can designate this queue as their dead-letter queue.
	//
	// When `redrivePermission` is set to `RedrivePermission.BY_QUEUE`, this parameter is required.
	//
	// You can specify up to 10 source queues.
	// To allow more than 10 source queues to specify dead-letter queues, set the `redrivePermission` to
	// `RedrivePermission.ALLOW_ALL`.
	//
	// When `redrivePermission` is either `RedrivePermission.ALLOW_ALL` or `RedrivePermission.DENY_ALL`,
	// this parameter cannot be set.
	// Default: - Required when `redrivePermission` is `RedrivePermission.BY_QUEUE`, cannot be defined otherwise.
	//
	SourceQueues *[]IQueue `field:"optional" json:"sourceQueues" yaml:"sourceQueues"`
}

