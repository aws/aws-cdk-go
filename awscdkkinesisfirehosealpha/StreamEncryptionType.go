package awscdkkinesisfirehosealpha


// Options for server-side encryption of a delivery stream.
// Experimental.
type StreamEncryptionType string

const (
	// Data in the stream is stored unencrypted.
	// Experimental.
	StreamEncryptionType_UNENCRYPTED StreamEncryptionType = "UNENCRYPTED"
	// Data in the stream is stored encrypted by a KMS key managed by the customer.
	// Experimental.
	StreamEncryptionType_CUSTOMER_MANAGED StreamEncryptionType = "CUSTOMER_MANAGED"
	// Data in the stream is stored encrypted by a KMS key owned by AWS and managed for use in multiple AWS accounts.
	// Experimental.
	StreamEncryptionType_AWS_OWNED StreamEncryptionType = "AWS_OWNED"
)

