package awskinesis


// Enables or updates server-side encryption using an AWS KMS key for a specified stream.
//
// > When invoking this API, you must use either the `StreamARN` or the `StreamName` parameter, or both. It is recommended that you use the `StreamARN` input parameter when you invoke this API.
//
// Starting encryption is an asynchronous operation. Upon receiving the request, Kinesis Data Streams returns immediately and sets the status of the stream to `UPDATING` . After the update is complete, Kinesis Data Streams sets the status of the stream back to `ACTIVE` . Updating or applying encryption normally takes a few seconds to complete, but it can take minutes. You can continue to read and write data to your stream while its status is `UPDATING` . Once the status of the stream is `ACTIVE` , encryption begins for records written to the stream.
//
// API Limits: You can successfully apply a new AWS KMS key for server-side encryption 25 times in a rolling 24-hour period.
//
// Note: It can take up to 5 seconds after the stream is in an `ACTIVE` status before all records written to the stream are encrypted. After you enable encryption, you can verify that encryption is applied by inspecting the API response from `PutRecord` or `PutRecords` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamEncryptionProperty := &StreamEncryptionProperty{
//   	EncryptionType: jsii.String("encryptionType"),
//   	KeyId: jsii.String("keyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html
//
type CfnStream_StreamEncryptionProperty struct {
	// The encryption type to use.
	//
	// The only valid value is `KMS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html#cfn-kinesis-stream-streamencryption-encryptiontype
	//
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// The GUID for the customer-managed AWS KMS key to use for encryption.
	//
	// This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You can also use a master key owned by Kinesis Data Streams by specifying the alias `aws/kinesis` .
	//
	// - Key ARN example: `arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012`
	// - Alias ARN example: `arn:aws:kms:us-east-1:123456789012:alias/MyAliasName`
	// - Globally unique key ID example: `12345678-1234-1234-1234-123456789012`
	// - Alias name example: `alias/MyAliasName`
	// - Master key owned by Kinesis Data Streams: `alias/aws/kinesis`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html#cfn-kinesis-stream-streamencryption-keyid
	//
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

