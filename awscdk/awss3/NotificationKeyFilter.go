package awss3


// Example:
//   var myQueue queue
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   bucket.AddEventNotification(s3.EventType_OBJECT_REMOVED,
//   s3n.NewSqsDestination(myQueue), &NotificationKeyFilter{
//   	Prefix: jsii.String("foo/"),
//   	Suffix: jsii.String(".jpg"),
//   })
//
// Experimental.
type NotificationKeyFilter struct {
	// S3 keys must have the specified prefix.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// S3 keys must have the specified suffix.
	// Experimental.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

