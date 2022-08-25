package awss3


// Default bucket access control types.
//
// Example:
//   websiteBucket := s3.NewBucket(this, jsii.String("WebsiteBucket"), &bucketProps{
//   	websiteIndexDocument: jsii.String("index.html"),
//   	publicReadAccess: jsii.Boolean(true),
//   })
//
//   s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &bucketDeploymentProps{
//   	sources: []iSource{
//   		s3deploy.source.asset(jsii.String("./website-dist")),
//   	},
//   	destinationBucket: websiteBucket,
//   	destinationKeyPrefix: jsii.String("web/static"),
//   	 // optional prefix in destination bucket
//   	metadata: &userDefinedObjectMetadata{
//   		a: jsii.String("1"),
//   		b: jsii.String("2"),
//   	},
//   	 // user-defined metadata
//
//   	// system-defined metadata
//   	contentType: jsii.String("text/html"),
//   	contentLanguage: jsii.String("en"),
//   	storageClass: s3deploy.storageClass_INTELLIGENT_TIERING,
//   	serverSideEncryption: s3deploy.serverSideEncryption_AES_256,
//   	cacheControl: []cacheControl{
//   		s3deploy.*cacheControl.setPublic(),
//   		s3deploy.*cacheControl.maxAge(awscdk.Duration.hours(jsii.Number(1))),
//   	},
//   	accessControl: s3.bucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html
//
type BucketAccessControl string

const (
	// Owner gets FULL_CONTROL.
	//
	// No one else has access rights.
	BucketAccessControl_PRIVATE BucketAccessControl = "PRIVATE"
	// Owner gets FULL_CONTROL.
	//
	// The AllUsers group gets READ access.
	BucketAccessControl_PUBLIC_READ BucketAccessControl = "PUBLIC_READ"
	// Owner gets FULL_CONTROL.
	//
	// The AllUsers group gets READ and WRITE access.
	// Granting this on a bucket is generally not recommended.
	BucketAccessControl_PUBLIC_READ_WRITE BucketAccessControl = "PUBLIC_READ_WRITE"
	// Owner gets FULL_CONTROL.
	//
	// The AuthenticatedUsers group gets READ access.
	BucketAccessControl_AUTHENTICATED_READ BucketAccessControl = "AUTHENTICATED_READ"
	// The LogDelivery group gets WRITE and READ_ACP permissions on the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html
	//
	BucketAccessControl_LOG_DELIVERY_WRITE BucketAccessControl = "LOG_DELIVERY_WRITE"
	// Object owner gets FULL_CONTROL.
	//
	// Bucket owner gets READ access.
	// If you specify this canned ACL when creating a bucket, Amazon S3 ignores it.
	BucketAccessControl_BUCKET_OWNER_READ BucketAccessControl = "BUCKET_OWNER_READ"
	// Both the object owner and the bucket owner get FULL_CONTROL over the object.
	//
	// If you specify this canned ACL when creating a bucket, Amazon S3 ignores it.
	BucketAccessControl_BUCKET_OWNER_FULL_CONTROL BucketAccessControl = "BUCKET_OWNER_FULL_CONTROL"
	// Owner gets FULL_CONTROL.
	//
	// Amazon EC2 gets READ access to GET an Amazon Machine Image (AMI) bundle from Amazon S3.
	BucketAccessControl_AWS_EXEC_READ BucketAccessControl = "AWS_EXEC_READ"
)

