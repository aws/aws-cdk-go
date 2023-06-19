package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   fn.AddEventSource(eventsources.NewS3EventSource(bucket, &S3EventSourceProps{
//   	Events: []eventType{
//   		s3.*eventType_OBJECT_CREATED,
//   		s3.*eventType_OBJECT_REMOVED,
//   	},
//   	Filters: []notificationKeyFilter{
//   		&notificationKeyFilter{
//   			Prefix: jsii.String("subdir/"),
//   		},
//   	},
//   }))
//
// Experimental.
type S3EventSourceProps struct {
	// The s3 event types that will trigger the notification.
	// Experimental.
	Events *[]awss3.EventType `field:"required" json:"events" yaml:"events"`
	// S3 object key filter rules to determine which objects trigger this event.
	//
	// Each filter must include a `prefix` and/or `suffix` that will be matched
	// against the s3 object key. Refer to the S3 Developer Guide for details
	// about allowed filter rules.
	// Experimental.
	Filters *[]*awss3.NotificationKeyFilter `field:"optional" json:"filters" yaml:"filters"`
}

