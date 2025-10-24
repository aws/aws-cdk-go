package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var fn Function
//
//
//   bucket := s3.NewBucket(this, jsii.String("mybucket"))
//
//   fn.AddEventSource(awscdk.NewS3EventSource(bucket, &S3EventSourceProps{
//   	Events: []EventType{
//   		s3.EventType_OBJECT_CREATED,
//   		s3.EventType_OBJECT_REMOVED,
//   	},
//   	Filters: []NotificationKeyFilter{
//   		&NotificationKeyFilter{
//   			Prefix: jsii.String("subdir/"),
//   		},
//   	},
//   }))
//
type S3EventSourceProps struct {
	// The s3 event types that will trigger the notification.
	Events *[]awss3.EventType `field:"required" json:"events" yaml:"events"`
	// S3 object key filter rules to determine which objects trigger this event.
	//
	// Each filter must include a `prefix` and/or `suffix` that will be matched
	// against the s3 object key. Refer to the S3 Developer Guide for details
	// about allowed filter rules.
	Filters *[]*awss3.NotificationKeyFilter `field:"optional" json:"filters" yaml:"filters"`
}

