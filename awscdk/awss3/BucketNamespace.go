package awss3


// The namespace for the bucket name when using `bucketNamePrefix`.
//
// Determines how CloudFormation generates the unique portion of the bucket name.
//
// Example:
//   s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	BucketNamePrefix: jsii.String("my-app"),
//   	BucketNamespace: s3.BucketNamespace_ACCOUNT_REGIONAL,
//   })
//
type BucketNamespace string

const (
	// The bucket name is globally unique.
	BucketNamespace_GLOBAL BucketNamespace = "GLOBAL"
	// The bucket name is unique within the account and region.
	BucketNamespace_ACCOUNT_REGIONAL BucketNamespace = "ACCOUNT_REGIONAL"
)

