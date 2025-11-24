package mixinsawstimestream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTablePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var magneticStoreWriteProperties interface{}
//   var retentionProperties interface{}
//
//   cfnTableMixinProps := &CfnTableMixinProps{
//   	DatabaseName: jsii.String("databaseName"),
//   	MagneticStoreWriteProperties: magneticStoreWriteProperties,
//   	RetentionProperties: retentionProperties,
//   	Schema: &SchemaProperty{
//   		CompositePartitionKey: []interface{}{
//   			&PartitionKeyProperty{
//   				EnforcementInRecord: jsii.String("enforcementInRecord"),
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	TableName: jsii.String("tableName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-table.html
//
type CfnTableMixinProps struct {
	// The name of the Timestream database that contains this table.
	//
	// *Length Constraints* : Minimum length of 3 bytes. Maximum length of 256 bytes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-table.html#cfn-timestream-table-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Contains properties to set on the table when enabling magnetic store writes.
	//
	// This object has the following attributes:
	//
	// - *EnableMagneticStoreWrites* : A `boolean` flag to enable magnetic store writes.
	// - *MagneticStoreRejectedDataLocation* : The location to write error reports for records rejected, asynchronously, during magnetic store writes. Only `S3Configuration` objects are allowed. The `S3Configuration` object has the following attributes:
	//
	// - *BucketName* : The name of the S3 bucket.
	// - *EncryptionOption* : The encryption option for the S3 location. Valid values are S3 server-side encryption with an S3 managed key ( `SSE_S3` ) or AWS managed key ( `SSE_KMS` ).
	// - *KmsKeyId* : The AWS KMS key ID to use when encrypting with an AWS managed key.
	// - *ObjectKeyPrefix* : The prefix to use option for the objects stored in S3.
	//
	// Both `BucketName` and `EncryptionOption` are *required* when `S3Configuration` is specified. If you specify `SSE_KMS` as your `EncryptionOption` then `KmsKeyId` is *required* .
	//
	// `EnableMagneticStoreWrites` attribute is *required* when `MagneticStoreWriteProperties` is specified. `MagneticStoreRejectedDataLocation` attribute is *required* when `EnableMagneticStoreWrites` is set to `true` .
	//
	// See the following examples:
	//
	// *JSON*
	//
	// ```json
	// { "Type" : AWS::Timestream::Table", "Properties":{ "DatabaseName":"TestDatabase", "TableName":"TestTable", "MagneticStoreWriteProperties":{ "EnableMagneticStoreWrites":true, "MagneticStoreRejectedDataLocation":{ "S3Configuration":{ "BucketName":" amzn-s3-demo-bucket ", "EncryptionOption":"SSE_KMS", "KmsKeyId":"1234abcd-12ab-34cd-56ef-1234567890ab", "ObjectKeyPrefix":"prefix" } } } }
	// }
	// ```
	//
	// *YAML*
	//
	// ```
	// Type: AWS::Timestream::Table
	// DependsOn: TestDatabase
	// Properties: TableName: "TestTable" DatabaseName: "TestDatabase" MagneticStoreWriteProperties: EnableMagneticStoreWrites: true MagneticStoreRejectedDataLocation: S3Configuration: BucketName: " amzn-s3-demo-bucket " EncryptionOption: "SSE_KMS" KmsKeyId: "1234abcd-12ab-34cd-56ef-1234567890ab" ObjectKeyPrefix: "prefix"
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-table.html#cfn-timestream-table-magneticstorewriteproperties
	//
	MagneticStoreWriteProperties interface{} `field:"optional" json:"magneticStoreWriteProperties" yaml:"magneticStoreWriteProperties"`
	// The retention duration for the memory store and magnetic store. This object has the following attributes:.
	//
	// - *MemoryStoreRetentionPeriodInHours* : Retention duration for memory store, in hours.
	// - *MagneticStoreRetentionPeriodInDays* : Retention duration for magnetic store, in days.
	//
	// Both attributes are of type `string` . Both attributes are *required* when `RetentionProperties` is specified.
	//
	// See the following examples:
	//
	// *JSON*
	//
	// `{ "Type" : AWS::Timestream::Table", "Properties" : { "DatabaseName" : "TestDatabase", "TableName" : "TestTable", "RetentionProperties" : { "MemoryStoreRetentionPeriodInHours": "24", "MagneticStoreRetentionPeriodInDays": "7" } } }`
	//
	// *YAML*
	//
	// ```
	// Type: AWS::Timestream::Table
	// DependsOn: TestDatabase
	// Properties: TableName: "TestTable" DatabaseName: "TestDatabase" RetentionProperties: MemoryStoreRetentionPeriodInHours: "24" MagneticStoreRetentionPeriodInDays: "7"
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-table.html#cfn-timestream-table-retentionproperties
	//
	RetentionProperties interface{} `field:"optional" json:"retentionProperties" yaml:"retentionProperties"`
	// The schema of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-table.html#cfn-timestream-table-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// The name of the Timestream table.
	//
	// *Length Constraints* : Minimum length of 3 bytes. Maximum length of 256 bytes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-table.html#cfn-timestream-table-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// The tags to add to the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-table.html#cfn-timestream-table-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

