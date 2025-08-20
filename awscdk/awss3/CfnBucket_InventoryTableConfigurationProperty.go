package awss3


// The inventory table configuration for an S3 Metadata configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inventoryTableConfigurationProperty := &InventoryTableConfigurationProperty{
//   	ConfigurationState: jsii.String("configurationState"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventorytableconfiguration.html
//
type CfnBucket_InventoryTableConfigurationProperty struct {
	// The configuration state of the inventory table, indicating whether the inventory table is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventorytableconfiguration.html#cfn-s3-bucket-inventorytableconfiguration-configurationstate
	//
	ConfigurationState *string `field:"required" json:"configurationState" yaml:"configurationState"`
	// The encryption configuration for the inventory table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventorytableconfiguration.html#cfn-s3-bucket-inventorytableconfiguration-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The Amazon Resource Name (ARN) for the inventory table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventorytableconfiguration.html#cfn-s3-bucket-inventorytableconfiguration-tablearn
	//
	TableArn *string `field:"optional" json:"tableArn" yaml:"tableArn"`
	// The name of the inventory table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventorytableconfiguration.html#cfn-s3-bucket-inventorytableconfiguration-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

