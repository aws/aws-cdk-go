package awss3


// The ObjectOwnership of the bucket.
//
// Example:
//   accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"), &BucketProps{
//   	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
//   })
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	ServerAccessLogsBucket: accessLogsBucket,
//   	ServerAccessLogsPrefix: jsii.String("logs"),
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html
//
type ObjectOwnership string

const (
	// ACLs are disabled, and the bucket owner automatically owns and has full control over every object in the bucket.
	//
	// ACLs no longer affect permissions to data in the S3 bucket.
	// The bucket uses policies to define access control.
	ObjectOwnership_BUCKET_OWNER_ENFORCED ObjectOwnership = "BUCKET_OWNER_ENFORCED"
	// The bucket owner will own the object if the object is uploaded with the bucket-owner-full-control canned ACL.
	//
	// Without this setting and
	// canned ACL, the object is uploaded and remains owned by the uploading account.
	ObjectOwnership_BUCKET_OWNER_PREFERRED ObjectOwnership = "BUCKET_OWNER_PREFERRED"
	// The uploading account will own the object.
	ObjectOwnership_OBJECT_WRITER ObjectOwnership = "OBJECT_WRITER"
)

