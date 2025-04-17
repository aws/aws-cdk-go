package awss3


// Supported types of notification destinations.
type BucketNotificationDestinationType string

const (
	BucketNotificationDestinationType_LAMBDA BucketNotificationDestinationType = "LAMBDA"
	BucketNotificationDestinationType_QUEUE BucketNotificationDestinationType = "QUEUE"
	BucketNotificationDestinationType_TOPIC BucketNotificationDestinationType = "TOPIC"
)

