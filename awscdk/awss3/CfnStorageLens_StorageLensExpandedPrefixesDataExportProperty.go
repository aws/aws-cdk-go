package awss3


// Expanded Prefixes Data Export.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sses3 interface{}
//
//   storageLensExpandedPrefixesDataExportProperty := &StorageLensExpandedPrefixesDataExportProperty{
//   	S3BucketDestination: &S3BucketDestinationProperty{
//   		AccountId: jsii.String("accountId"),
//   		Arn: jsii.String("arn"),
//   		Format: jsii.String("format"),
//   		OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   		// the properties below are optional
//   		Encryption: &EncryptionProperty{
//   			Ssekms: &SSEKMSProperty{
//   				KeyId: jsii.String("keyId"),
//   			},
//   			Sses3: sses3,
//   		},
//   		Prefix: jsii.String("prefix"),
//   	},
//   	StorageLensTableDestination: &StorageLensTableDestinationProperty{
//   		IsEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Encryption: &EncryptionProperty{
//   			Ssekms: &SSEKMSProperty{
//   				KeyId: jsii.String("keyId"),
//   			},
//   			Sses3: sses3,
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensexpandedprefixesdataexport.html
//
type CfnStorageLens_StorageLensExpandedPrefixesDataExportProperty struct {
	// S3 bucket destination settings for the Amazon S3 Storage Lens metrics export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensexpandedprefixesdataexport.html#cfn-s3-storagelens-storagelensexpandedprefixesdataexport-s3bucketdestination
	//
	S3BucketDestination interface{} `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
	// S3 Tables destination settings for the Amazon S3 Storage Lens metrics export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensexpandedprefixesdataexport.html#cfn-s3-storagelens-storagelensexpandedprefixesdataexport-storagelenstabledestination
	//
	StorageLensTableDestination interface{} `field:"optional" json:"storageLensTableDestination" yaml:"storageLensTableDestination"`
}

