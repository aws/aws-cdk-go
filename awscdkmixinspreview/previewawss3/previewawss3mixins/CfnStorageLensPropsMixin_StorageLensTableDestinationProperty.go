package previewawss3mixins


// S3 Tables destination settings for the Amazon S3 Storage Lens metrics export.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var sses3 interface{}
//
//   storageLensTableDestinationProperty := &StorageLensTableDestinationProperty{
//   	Encryption: &EncryptionProperty{
//   		Ssekms: &SSEKMSProperty{
//   			KeyId: jsii.String("keyId"),
//   		},
//   		Sses3: sses3,
//   	},
//   	IsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelenstabledestination.html
//
type CfnStorageLensPropsMixin_StorageLensTableDestinationProperty struct {
	// Configures the server-side encryption for Amazon S3 Storage Lens report files with either S3-managed keys (SSE-S3) or KMS-managed keys (SSE-KMS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelenstabledestination.html#cfn-s3-storagelens-storagelenstabledestination-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Specifies whether the export to S3 Tables is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelenstabledestination.html#cfn-s3-storagelens-storagelenstabledestination-isenabled
	//
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

