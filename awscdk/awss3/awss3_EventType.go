package awss3


// Notification event types.
//
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
type EventType string

const (
	// Amazon S3 APIs such as PUT, POST, and COPY can create an object.
	//
	// Using
	// these event types, you can enable notification when an object is created
	// using a specific API, or you can use the s3:ObjectCreated:* event type to
	// request notification regardless of the API that was used to create an
	// object.
	EventType_OBJECT_CREATED EventType = "OBJECT_CREATED"
	// Amazon S3 APIs such as PUT, POST, and COPY can create an object.
	//
	// Using
	// these event types, you can enable notification when an object is created
	// using a specific API, or you can use the s3:ObjectCreated:* event type to
	// request notification regardless of the API that was used to create an
	// object.
	EventType_OBJECT_CREATED_PUT EventType = "OBJECT_CREATED_PUT"
	// Amazon S3 APIs such as PUT, POST, and COPY can create an object.
	//
	// Using
	// these event types, you can enable notification when an object is created
	// using a specific API, or you can use the s3:ObjectCreated:* event type to
	// request notification regardless of the API that was used to create an
	// object.
	EventType_OBJECT_CREATED_POST EventType = "OBJECT_CREATED_POST"
	// Amazon S3 APIs such as PUT, POST, and COPY can create an object.
	//
	// Using
	// these event types, you can enable notification when an object is created
	// using a specific API, or you can use the s3:ObjectCreated:* event type to
	// request notification regardless of the API that was used to create an
	// object.
	EventType_OBJECT_CREATED_COPY EventType = "OBJECT_CREATED_COPY"
	// Amazon S3 APIs such as PUT, POST, and COPY can create an object.
	//
	// Using
	// these event types, you can enable notification when an object is created
	// using a specific API, or you can use the s3:ObjectCreated:* event type to
	// request notification regardless of the API that was used to create an
	// object.
	EventType_OBJECT_CREATED_COMPLETE_MULTIPART_UPLOAD EventType = "OBJECT_CREATED_COMPLETE_MULTIPART_UPLOAD"
	// By using the ObjectRemoved event types, you can enable notification when an object or a batch of objects is removed from a bucket.
	//
	// You can request notification when an object is deleted or a versioned
	// object is permanently deleted by using the s3:ObjectRemoved:Delete event
	// type. Or you can request notification when a delete marker is created for
	// a versioned object by using s3:ObjectRemoved:DeleteMarkerCreated. For
	// information about deleting versioned objects, see Deleting Object
	// Versions. You can also use a wildcard s3:ObjectRemoved:* to request
	// notification anytime an object is deleted.
	//
	// You will not receive event notifications from automatic deletes from
	// lifecycle policies or from failed operations.
	EventType_OBJECT_REMOVED EventType = "OBJECT_REMOVED"
	// By using the ObjectRemoved event types, you can enable notification when an object or a batch of objects is removed from a bucket.
	//
	// You can request notification when an object is deleted or a versioned
	// object is permanently deleted by using the s3:ObjectRemoved:Delete event
	// type. Or you can request notification when a delete marker is created for
	// a versioned object by using s3:ObjectRemoved:DeleteMarkerCreated. For
	// information about deleting versioned objects, see Deleting Object
	// Versions. You can also use a wildcard s3:ObjectRemoved:* to request
	// notification anytime an object is deleted.
	//
	// You will not receive event notifications from automatic deletes from
	// lifecycle policies or from failed operations.
	EventType_OBJECT_REMOVED_DELETE EventType = "OBJECT_REMOVED_DELETE"
	// By using the ObjectRemoved event types, you can enable notification when an object or a batch of objects is removed from a bucket.
	//
	// You can request notification when an object is deleted or a versioned
	// object is permanently deleted by using the s3:ObjectRemoved:Delete event
	// type. Or you can request notification when a delete marker is created for
	// a versioned object by using s3:ObjectRemoved:DeleteMarkerCreated. For
	// information about deleting versioned objects, see Deleting Object
	// Versions. You can also use a wildcard s3:ObjectRemoved:* to request
	// notification anytime an object is deleted.
	//
	// You will not receive event notifications from automatic deletes from
	// lifecycle policies or from failed operations.
	EventType_OBJECT_REMOVED_DELETE_MARKER_CREATED EventType = "OBJECT_REMOVED_DELETE_MARKER_CREATED"
	// Using restore object event types you can receive notifications for initiation and completion when restoring objects from the S3 Glacier storage class.
	//
	// You use s3:ObjectRestore:Post to request notification of object restoration
	// initiation.
	EventType_OBJECT_RESTORE_POST EventType = "OBJECT_RESTORE_POST"
	// Using restore object event types you can receive notifications for initiation and completion when restoring objects from the S3 Glacier storage class.
	//
	// You use s3:ObjectRestore:Completed to request notification of
	// restoration completion.
	EventType_OBJECT_RESTORE_COMPLETED EventType = "OBJECT_RESTORE_COMPLETED"
	// Using restore object event types you can receive notifications for initiation and completion when restoring objects from the S3 Glacier storage class.
	//
	// You use s3:ObjectRestore:Delete to request notification of
	// restoration completion.
	EventType_OBJECT_RESTORE_DELETE EventType = "OBJECT_RESTORE_DELETE"
	// You can use this event type to request Amazon S3 to send a notification message when Amazon S3 detects that an object of the RRS storage class is lost.
	EventType_REDUCED_REDUNDANCY_LOST_OBJECT EventType = "REDUCED_REDUNDANCY_LOST_OBJECT"
	// You receive this notification event when an object that was eligible for replication using Amazon S3 Replication Time Control failed to replicate.
	EventType_REPLICATION_OPERATION_FAILED_REPLICATION EventType = "REPLICATION_OPERATION_FAILED_REPLICATION"
	// You receive this notification event when an object that was eligible for replication using Amazon S3 Replication Time Control exceeded the 15-minute threshold for replication.
	EventType_REPLICATION_OPERATION_MISSED_THRESHOLD EventType = "REPLICATION_OPERATION_MISSED_THRESHOLD"
	// You receive this notification event for an object that was eligible for replication using the Amazon S3 Replication Time Control feature replicated after the 15-minute threshold.
	EventType_REPLICATION_OPERATION_REPLICATED_AFTER_THRESHOLD EventType = "REPLICATION_OPERATION_REPLICATED_AFTER_THRESHOLD"
	// You receive this notification event for an object that was eligible for replication using Amazon S3 Replication Time Control but is no longer tracked by replication metrics.
	EventType_REPLICATION_OPERATION_NOT_TRACKED EventType = "REPLICATION_OPERATION_NOT_TRACKED"
	// By using the LifecycleExpiration event types, you can receive a notification when Amazon S3 deletes an object based on your S3 Lifecycle configuration.
	EventType_LIFECYCLE_EXPIRATION EventType = "LIFECYCLE_EXPIRATION"
	// The s3:LifecycleExpiration:Delete event type notifies you when an object in an unversioned bucket is deleted.
	//
	// It also notifies you when an object version is permanently deleted by an
	// S3 Lifecycle configuration.
	EventType_LIFECYCLE_EXPIRATION_DELETE EventType = "LIFECYCLE_EXPIRATION_DELETE"
	// The s3:LifecycleExpiration:DeleteMarkerCreated event type notifies you when S3 Lifecycle creates a delete marker when a current version of an object in versioned bucket is deleted.
	EventType_LIFECYCLE_EXPIRATION_DELETE_MARKER_CREATED EventType = "LIFECYCLE_EXPIRATION_DELETE_MARKER_CREATED"
	// You receive this notification event when an object is transitioned to another Amazon S3 storage class by an S3 Lifecycle configuration.
	EventType_LIFECYCLE_TRANSITION EventType = "LIFECYCLE_TRANSITION"
	// You receive this notification event when an object within the S3 Intelligent-Tiering storage class moved to the Archive Access tier or Deep Archive Access tier.
	EventType_INTELLIGENT_TIERING EventType = "INTELLIGENT_TIERING"
	// By using the ObjectTagging event types, you can enable notification when an object tag is added or deleted from an object.
	EventType_OBJECT_TAGGING EventType = "OBJECT_TAGGING"
	// The s3:ObjectTagging:Put event type notifies you when a tag is PUT on an object or an existing tag is updated.
	EventType_OBJECT_TAGGING_PUT EventType = "OBJECT_TAGGING_PUT"
	// The s3:ObjectTagging:Delete event type notifies you when a tag is removed from an object.
	EventType_OBJECT_TAGGING_DELETE EventType = "OBJECT_TAGGING_DELETE"
	// You receive this notification event when an ACL is PUT on an object or when an existing ACL is changed.
	//
	// An event is not generated when a request results in no change to an
	// object’s ACL.
	EventType_OBJECT_ACL_PUT EventType = "OBJECT_ACL_PUT"
)

