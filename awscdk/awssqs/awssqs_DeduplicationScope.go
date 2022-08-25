package awssqs


// What kind of deduplication scope to apply.
// Experimental.
type DeduplicationScope string

const (
	// Deduplication occurs at the message group level.
	// Experimental.
	DeduplicationScope_MESSAGE_GROUP DeduplicationScope = "MESSAGE_GROUP"
	// Deduplication occurs at the message queue level.
	// Experimental.
	DeduplicationScope_QUEUE DeduplicationScope = "QUEUE"
)

