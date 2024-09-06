package awscloudfront


// The level of permissions granted to the CloudFront Distribution when configuring OAC.
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   s3Origin := origins.S3BucketOrigin_WithOriginAccessControl(myBucket, &S3BucketOriginWithOACProps{
//   	OriginAccessLevels: []accessLevel{
//   		cloudfront.*accessLevel_READ,
//   		cloudfront.*accessLevel_WRITE,
//   		cloudfront.*accessLevel_DELETE,
//   	},
//   })
//
type AccessLevel string

const (
	// Grants read permissions to CloudFront Distribution.
	AccessLevel_READ AccessLevel = "READ"
	// Grants write permission to CloudFront Distribution.
	AccessLevel_WRITE AccessLevel = "WRITE"
	// Grants delete permission to CloudFront Distribution.
	AccessLevel_DELETE AccessLevel = "DELETE"
)

