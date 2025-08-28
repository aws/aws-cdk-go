package awss3


// Creates a V2 Amazon S3 Metadata configuration of a general purpose bucket.
//
// For more information, see [Accelerating data discovery with S3 Metadata](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataConfigurationProperty := &MetadataConfigurationProperty{
//   	JournalTableConfiguration: &JournalTableConfigurationProperty{
//   		RecordExpiration: &RecordExpirationProperty{
//   			Expiration: jsii.String("expiration"),
//
//   			// the properties below are optional
//   			Days: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		EncryptionConfiguration: &MetadataTableEncryptionConfigurationProperty{
//   			SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   			// the properties below are optional
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   		TableArn: jsii.String("tableArn"),
//   		TableName: jsii.String("tableName"),
//   	},
//
//   	// the properties below are optional
//   	Destination: &MetadataDestinationProperty{
//   		TableBucketType: jsii.String("tableBucketType"),
//
//   		// the properties below are optional
//   		TableBucketArn: jsii.String("tableBucketArn"),
//   		TableNamespace: jsii.String("tableNamespace"),
//   	},
//   	InventoryTableConfiguration: &InventoryTableConfigurationProperty{
//   		ConfigurationState: jsii.String("configurationState"),
//
//   		// the properties below are optional
//   		EncryptionConfiguration: &MetadataTableEncryptionConfigurationProperty{
//   			SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   			// the properties below are optional
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   		TableArn: jsii.String("tableArn"),
//   		TableName: jsii.String("tableName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadataconfiguration.html
//
type CfnBucket_MetadataConfigurationProperty struct {
	// The journal table configuration for a metadata configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadataconfiguration.html#cfn-s3-bucket-metadataconfiguration-journaltableconfiguration
	//
	JournalTableConfiguration interface{} `field:"required" json:"journalTableConfiguration" yaml:"journalTableConfiguration"`
	// The destination information for the S3 Metadata configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadataconfiguration.html#cfn-s3-bucket-metadataconfiguration-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The inventory table configuration for a metadata configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadataconfiguration.html#cfn-s3-bucket-metadataconfiguration-inventorytableconfiguration
	//
	InventoryTableConfiguration interface{} `field:"optional" json:"inventoryTableConfiguration" yaml:"inventoryTableConfiguration"`
}

