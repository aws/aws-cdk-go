package awss3


// Supported types of notification destinations.
// Experimental.
type BucketNotificationDestinationType string

const (
	// Experimental.
	BucketNotificationDestinationType_LAMBDA BucketNotificationDestinationType = "LAMBDA"
	// Experimental.
	BucketNotificationDestinationType_QUEUE BucketNotificationDestinationType = "QUEUE"
	// Experimental.
	BucketNotificationDestinationType_TOPIC BucketNotificationDestinationType = "TOPIC"
)

