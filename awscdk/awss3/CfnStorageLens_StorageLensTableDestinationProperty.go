package awss3


// S3 Tables destination settings for the Amazon S3 Storage Lens metrics export.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sses3 interface{}
//
//   storageLensTableDestinationProperty := &StorageLensTableDestinationProperty{
//   	IsEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	Encryption: &EncryptionProperty{
//   		Ssekms: &SSEKMSProperty{
//   			KeyId: jsii.String("keyId"),
//   		},
//   		Sses3: sses3,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelenstabledestination.html
//
type CfnStorageLens_StorageLensTableDestinationProperty struct {
	// Specifies whether the export to S3 Tables is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelenstabledestination.html#cfn-s3-storagelens-storagelenstabledestination-isenabled
	//
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Configures the server-side encryption for Amazon S3 Storage Lens report files with either S3-managed keys (SSE-S3) or KMS-managed keys (SSE-KMS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelenstabledestination.html#cfn-s3-storagelens-storagelenstabledestination-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
}

