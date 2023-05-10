package awss3


// What kind of server-side encryption to apply to this bucket.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyEncryptedBucket"), &BucketProps{
//   	Encryption: s3.BucketEncryption_KMS,
//   })
//
//   // you can access the encryption key:
//   assert(bucket.EncryptionKey instanceof kms.Key)
//
// Experimental.
type BucketEncryption string

const (
	// Objects in the bucket are not encrypted.
	// Experimental.
	BucketEncryption_UNENCRYPTED BucketEncryption = "UNENCRYPTED"
	// Server-side KMS encryption with a master key managed by KMS.
	// Experimental.
	BucketEncryption_KMS_MANAGED BucketEncryption = "KMS_MANAGED"
	// Server-side encryption with a master key managed by S3.
	// Experimental.
	BucketEncryption_S3_MANAGED BucketEncryption = "S3_MANAGED"
	// Server-side encryption with a KMS key managed by the user.
	//
	// If `encryptionKey` is specified, this key will be used, otherwise, one will be defined.
	// Experimental.
	BucketEncryption_KMS BucketEncryption = "KMS"
)

