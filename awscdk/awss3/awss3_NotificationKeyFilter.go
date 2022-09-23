package awss3


// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myQueue queue
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   bucket.addEventNotification(s3.eventType_OBJECT_REMOVED,
//   s3n.NewSqsDestination(myQueue), &notificationKeyFilter{
//   	prefix: jsii.String("foo/"),
//   	suffix: jsii.String(".jpg"),
//   })
//
type NotificationKeyFilter struct {
	// S3 keys must have the specified prefix.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// S3 keys must have the specified suffix.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

