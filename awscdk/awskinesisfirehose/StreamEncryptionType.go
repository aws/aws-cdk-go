package awskinesisfirehose


// Options for server-side encryption of a delivery stream.
type StreamEncryptionType string

const (
	// Data in the stream is stored unencrypted.
	StreamEncryptionType_UNENCRYPTED StreamEncryptionType = "UNENCRYPTED"
	// Data in the stream is stored encrypted by a KMS key managed by the customer.
	StreamEncryptionType_CUSTOMER_MANAGED StreamEncryptionType = "CUSTOMER_MANAGED"
	// Data in the stream is stored encrypted by a KMS key owned by AWS and managed for use in multiple AWS accounts.
	StreamEncryptionType_AWS_OWNED StreamEncryptionType = "AWS_OWNED"
)

