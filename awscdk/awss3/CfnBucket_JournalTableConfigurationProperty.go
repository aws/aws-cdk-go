package awss3


// The journal table configuration for an S3 Metadata configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   journalTableConfigurationProperty := &JournalTableConfigurationProperty{
//   	RecordExpiration: &RecordExpirationProperty{
//   		Expiration: jsii.String("expiration"),
//
//   		// the properties below are optional
//   		Days: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	EncryptionConfiguration: &MetadataTableEncryptionConfigurationProperty{
//   		SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   		// the properties below are optional
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	TableArn: jsii.String("tableArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-journaltableconfiguration.html
//
type CfnBucket_JournalTableConfigurationProperty struct {
	// The journal table record expiration settings for the journal table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-journaltableconfiguration.html#cfn-s3-bucket-journaltableconfiguration-recordexpiration
	//
	RecordExpiration interface{} `field:"required" json:"recordExpiration" yaml:"recordExpiration"`
	// The encryption configuration for the journal table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-journaltableconfiguration.html#cfn-s3-bucket-journaltableconfiguration-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The Amazon Resource Name (ARN) for the journal table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-journaltableconfiguration.html#cfn-s3-bucket-journaltableconfiguration-tablearn
	//
	TableArn *string `field:"optional" json:"tableArn" yaml:"tableArn"`
	// The name of the journal table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-journaltableconfiguration.html#cfn-s3-bucket-journaltableconfiguration-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

