package awskinesis


// What kind of server-side encryption to apply to this stream.
//
// Example:
//   key := kms.NewKey(this, jsii.String("MyKey"))
//
//   kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &streamProps{
//   	encryption: kinesis.streamEncryption_KMS,
//   	encryptionKey: key,
//   })
//
type StreamEncryption string

const (
	// Records in the stream are not encrypted.
	StreamEncryption_UNENCRYPTED StreamEncryption = "UNENCRYPTED"
	// Server-side encryption with a KMS key managed by the user.
	//
	// If `encryptionKey` is specified, this key will be used, otherwise, one will be defined.
	StreamEncryption_KMS StreamEncryption = "KMS"
	// Server-side encryption with a master key managed by Amazon Kinesis.
	StreamEncryption_MANAGED StreamEncryption = "MANAGED"
)

