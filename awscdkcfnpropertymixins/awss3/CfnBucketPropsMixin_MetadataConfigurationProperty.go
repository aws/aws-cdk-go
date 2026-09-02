package awss3


// Creates a V2 Amazon S3 Metadata configuration of a general purpose bucket.
//
// For more information, see [Accelerating data discovery with S3 Metadata](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   metadataConfigurationProperty := &MetadataConfigurationProperty{
//   	AnnotationTableConfiguration: &AnnotationTableConfigurationProperty{
//   		ConfigurationState: jsii.String("configurationState"),
//   		EncryptionConfiguration: &MetadataTableEncryptionConfigurationProperty{
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			SseAlgorithm: jsii.String("sseAlgorithm"),
//   		},
//   		Role: jsii.String("role"),
//   		TableArn: jsii.String("tableArn"),
//   		TableName: jsii.String("tableName"),
//   	},
//   	Destination: &MetadataDestinationProperty{
//   		TableBucketArn: jsii.String("tableBucketArn"),
//   		TableBucketType: jsii.String("tableBucketType"),
//   		TableNamespace: jsii.String("tableNamespace"),
//   	},
//   	InventoryTableConfiguration: &InventoryTableConfigurationProperty{
//   		ConfigurationState: jsii.String("configurationState"),
//   		EncryptionConfiguration: &MetadataTableEncryptionConfigurationProperty{
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			SseAlgorithm: jsii.String("sseAlgorithm"),
//   		},
//   		TableArn: jsii.String("tableArn"),
//   		TableName: jsii.String("tableName"),
//   	},
//   	JournalTableConfiguration: &JournalTableConfigurationProperty{
//   		EncryptionConfiguration: &MetadataTableEncryptionConfigurationProperty{
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			SseAlgorithm: jsii.String("sseAlgorithm"),
//   		},
//   		RecordExpiration: &RecordExpirationProperty{
//   			Days: jsii.Number(123),
//   			Expiration: jsii.String("expiration"),
//   		},
//   		TableArn: jsii.String("tableArn"),
//   		TableName: jsii.String("tableName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadataconfiguration.html
//
type CfnBucketPropsMixin_MetadataConfigurationProperty struct {
	// The annotation table configuration for an S3 Metadata configuration.
	//
	// The annotation table tracks all annotations on objects in your bucket so that you can query annotation data at scale.
	//   If you've disabled your annotation table configuration and now want to re-enable it, you must first manually delete the old annotation table from your AWS managed table bucket. Otherwise, the newly re-enabled annotation table configuration will enter a failed state because the annotation table already exists in the table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadataconfiguration.html#cfn-s3-bucket-metadataconfiguration-annotationtableconfiguration
	//
	AnnotationTableConfiguration interface{} `field:"optional" json:"annotationTableConfiguration" yaml:"annotationTableConfiguration"`
	// The destination information for the S3 Metadata configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadataconfiguration.html#cfn-s3-bucket-metadataconfiguration-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The inventory table configuration for a metadata configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadataconfiguration.html#cfn-s3-bucket-metadataconfiguration-inventorytableconfiguration
	//
	InventoryTableConfiguration interface{} `field:"optional" json:"inventoryTableConfiguration" yaml:"inventoryTableConfiguration"`
	// The journal table configuration for a metadata configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadataconfiguration.html#cfn-s3-bucket-metadataconfiguration-journaltableconfiguration
	//
	JournalTableConfiguration interface{} `field:"optional" json:"journalTableConfiguration" yaml:"journalTableConfiguration"`
}

