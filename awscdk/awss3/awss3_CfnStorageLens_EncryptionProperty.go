package awss3


// This resource contains the type of server-side encryption used for Amazon S3 Storage Lens.
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
//   encryptionProperty := &encryptionProperty{
//   	ssekms: &sSEKMSProperty{
//   		keyId: jsii.String("keyId"),
//   	},
//   	sses3: sses3,
//   }
//
type CfnStorageLens_EncryptionProperty struct {
	// `CfnStorageLens.EncryptionProperty.SSEKMS`.
	Ssekms interface{} `field:"optional" json:"ssekms" yaml:"ssekms"`
	// `CfnStorageLens.EncryptionProperty.SSES3`.
	Sses3 interface{} `field:"optional" json:"sses3" yaml:"sses3"`
}

