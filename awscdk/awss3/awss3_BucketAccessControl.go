package awss3


// Default bucket access control types.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewS3PutObjectAction(bucket, &S3PutObjectActionProps{
//   			AccessControl: s3.BucketAccessControl_PUBLIC_READ,
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html
//
// Experimental.
type BucketAccessControl string

const (
	// Owner gets FULL_CONTROL.
	//
	// No one else has access rights.
	// Experimental.
	BucketAccessControl_PRIVATE BucketAccessControl = "PRIVATE"
	// Owner gets FULL_CONTROL.
	//
	// The AllUsers group gets READ access.
	// Experimental.
	BucketAccessControl_PUBLIC_READ BucketAccessControl = "PUBLIC_READ"
	// Owner gets FULL_CONTROL.
	//
	// The AllUsers group gets READ and WRITE access.
	// Granting this on a bucket is generally not recommended.
	// Experimental.
	BucketAccessControl_PUBLIC_READ_WRITE BucketAccessControl = "PUBLIC_READ_WRITE"
	// Owner gets FULL_CONTROL.
	//
	// The AuthenticatedUsers group gets READ access.
	// Experimental.
	BucketAccessControl_AUTHENTICATED_READ BucketAccessControl = "AUTHENTICATED_READ"
	// The LogDelivery group gets WRITE and READ_ACP permissions on the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html
	//
	// Experimental.
	BucketAccessControl_LOG_DELIVERY_WRITE BucketAccessControl = "LOG_DELIVERY_WRITE"
	// Object owner gets FULL_CONTROL.
	//
	// Bucket owner gets READ access.
	// If you specify this canned ACL when creating a bucket, Amazon S3 ignores it.
	// Experimental.
	BucketAccessControl_BUCKET_OWNER_READ BucketAccessControl = "BUCKET_OWNER_READ"
	// Both the object owner and the bucket owner get FULL_CONTROL over the object.
	//
	// If you specify this canned ACL when creating a bucket, Amazon S3 ignores it.
	// Experimental.
	BucketAccessControl_BUCKET_OWNER_FULL_CONTROL BucketAccessControl = "BUCKET_OWNER_FULL_CONTROL"
	// Owner gets FULL_CONTROL.
	//
	// Amazon EC2 gets READ access to GET an Amazon Machine Image (AMI) bundle from Amazon S3.
	// Experimental.
	BucketAccessControl_AWS_EXEC_READ BucketAccessControl = "AWS_EXEC_READ"
)

