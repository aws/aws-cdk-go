package awss3


// What kind of server-side encryption to apply to this bucket.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyEncryptedBucket"), &bucketProps{
//   	encryption: s3.bucketEncryption_KMS,
//   })
//
//   // you can access the encryption key:
//   assert(bucket.encryptionKey instanceof kms.key)
//
type BucketEncryption string

const (
	// Objects in the bucket are not encrypted.
	BucketEncryption_UNENCRYPTED BucketEncryption = "UNENCRYPTED"
	// Server-side KMS encryption with a master key managed by KMS.
	BucketEncryption_KMS_MANAGED BucketEncryption = "KMS_MANAGED"
	// Server-side encryption with a master key managed by S3.
	BucketEncryption_S3_MANAGED BucketEncryption = "S3_MANAGED"
	// Server-side encryption with a KMS key managed by the user.
	//
	// If `encryptionKey` is specified, this key will be used, otherwise, one will be defined.
	BucketEncryption_KMS BucketEncryption = "KMS"
)

