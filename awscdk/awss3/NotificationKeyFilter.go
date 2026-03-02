package awss3


// Example:
//   var myLambda Function
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   filter := &NotificationKeyFilter{
//   	Prefix: jsii.String("home/myusername/*"),
//   }
//   bucket.AddEventNotification(s3.EventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), filter)
//
type NotificationKeyFilter struct {
	// S3 keys must have the specified prefix.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// S3 keys must have the specified suffix.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

