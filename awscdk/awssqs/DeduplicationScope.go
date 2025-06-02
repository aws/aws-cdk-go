package awssqs


// What kind of deduplication scope to apply.
type DeduplicationScope string

const (
	// Deduplication occurs at the message group level.
	DeduplicationScope_MESSAGE_GROUP DeduplicationScope = "MESSAGE_GROUP"
	// Deduplication occurs at the message queue level.
	DeduplicationScope_QUEUE DeduplicationScope = "QUEUE"
)

