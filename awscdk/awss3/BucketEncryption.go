package awss3


// What kind of server-side encryption to apply to this bucket.
//
// Example:
//   var application application
//
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	Versioned: jsii.Boolean(true),
//   	Encryption: s3.BucketEncryption_KMS,
//   })
//
//   appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
//   	Application: Application,
//   	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
//   })
//
type BucketEncryption string

const (
	// Previous option.
	//
	// Buckets can not be unencrypted now.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/serv-side-encryption.html
	//
	// Deprecated: S3 applies server-side encryption with SSE-S3 for every bucket
	// that default encryption is not configured.
	BucketEncryption_UNENCRYPTED BucketEncryption = "UNENCRYPTED"
	// Server-side KMS encryption with a master key managed by KMS.
	BucketEncryption_KMS_MANAGED BucketEncryption = "KMS_MANAGED"
	// Server-side encryption with a master key managed by S3.
	BucketEncryption_S3_MANAGED BucketEncryption = "S3_MANAGED"
	// Server-side encryption with a KMS key managed by the user.
	//
	// If `encryptionKey` is specified, this key will be used, otherwise, one will be defined.
	BucketEncryption_KMS BucketEncryption = "KMS"
	// Double server-side KMS encryption with a master key managed by KMS.
	BucketEncryption_DSSE_MANAGED BucketEncryption = "DSSE_MANAGED"
	// Double server-side encryption with a KMS key managed by the user.
	//
	// If `encryptionKey` is specified, this key will be used, otherwise, one will be defined.
	BucketEncryption_DSSE BucketEncryption = "DSSE"
)

