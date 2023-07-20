package awss3


// This resource contains the type of server-side encryption used to encrypt an Amazon S3 Storage Lens metrics export.
//
// For valid values, see the [StorageLensDataExportEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_StorageLensDataExportEncryption.html) in the *Amazon S3 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sses3 interface{}
//
//   encryptionProperty := &EncryptionProperty{
//   	Ssekms: &SSEKMSProperty{
//   		KeyId: jsii.String("keyId"),
//   	},
//   	Sses3: sses3,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-encryption.html
//
type CfnStorageLens_EncryptionProperty struct {
	// Specifies the use of AWS Key Management Service keys (SSE-KMS) to encrypt the S3 Storage Lens metrics export file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-encryption.html#cfn-s3-storagelens-encryption-ssekms
	//
	Ssekms interface{} `field:"optional" json:"ssekms" yaml:"ssekms"`
	// Specifies the use of an Amazon S3-managed key (SSE-S3) to encrypt the S3 Storage Lens metrics export file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-encryption.html#cfn-s3-storagelens-encryption-sses3
	//
	Sses3 interface{} `field:"optional" json:"sses3" yaml:"sses3"`
}

