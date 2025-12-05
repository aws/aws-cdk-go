package previewawss3mixins


// Expanded Prefixes Data Export.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var sses3 interface{}
//
//   storageLensExpandedPrefixesDataExportProperty := &StorageLensExpandedPrefixesDataExportProperty{
//   	S3BucketDestination: &S3BucketDestinationProperty{
//   		AccountId: jsii.String("accountId"),
//   		Arn: jsii.String("arn"),
//   		Encryption: &EncryptionProperty{
//   			Ssekms: &SSEKMSProperty{
//   				KeyId: jsii.String("keyId"),
//   			},
//   			Sses3: sses3,
//   		},
//   		Format: jsii.String("format"),
//   		OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	StorageLensTableDestination: &StorageLensTableDestinationProperty{
//   		Encryption: &EncryptionProperty{
//   			Ssekms: &SSEKMSProperty{
//   				KeyId: jsii.String("keyId"),
//   			},
//   			Sses3: sses3,
//   		},
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensexpandedprefixesdataexport.html
//
type CfnStorageLensPropsMixin_StorageLensExpandedPrefixesDataExportProperty struct {
	// S3 bucket destination settings for the Amazon S3 Storage Lens metrics export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensexpandedprefixesdataexport.html#cfn-s3-storagelens-storagelensexpandedprefixesdataexport-s3bucketdestination
	//
	S3BucketDestination interface{} `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
	// S3 Tables destination settings for the Amazon S3 Storage Lens metrics export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensexpandedprefixesdataexport.html#cfn-s3-storagelens-storagelensexpandedprefixesdataexport-storagelenstabledestination
	//
	StorageLensTableDestination interface{} `field:"optional" json:"storageLensTableDestination" yaml:"storageLensTableDestination"`
}

