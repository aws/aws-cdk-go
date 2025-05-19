package awskinesis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// A reference to a stream.
//
// The easiest way to instantiate is to call
// `stream.export()`. Then, the consumer can use `Stream.import(this, ref)` and
// get a `Stream`.
//
// Example:
//   importedStream := kinesis.Stream_FromStreamAttributes(this, jsii.String("ImportedEncryptedStream"), &StreamAttributes{
//   	StreamArn: jsii.String("arn:aws:kinesis:us-east-2:123456789012:stream/f3j09j2230j"),
//   	EncryptionKey: kms.Key_FromKeyArn(this, jsii.String("key"), jsii.String("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012")),
//   })
//
type StreamAttributes struct {
	// The ARN of the stream.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
	// The KMS key securing the contents of the stream if encryption is enabled.
	// Default: - No encryption.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

