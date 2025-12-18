package awss3tables

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTableProps := &CfnTableProps{
//   	Namespace: jsii.String("namespace"),
//   	OpenTableFormat: jsii.String("openTableFormat"),
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   	TableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	Compaction: &CompactionProperty{
//   		Status: jsii.String("status"),
//   		TargetFileSizeMb: jsii.Number(123),
//   	},
//   	IcebergMetadata: &IcebergMetadataProperty{
//   		IcebergSchema: &IcebergSchemaProperty{
//   			SchemaFieldList: []interface{}{
//   				&SchemaFieldProperty{
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Required: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	SnapshotManagement: &SnapshotManagementProperty{
//   		MaxSnapshotAgeHours: jsii.Number(123),
//   		MinSnapshotsToKeep: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	StorageClassConfiguration: &StorageClassConfigurationProperty{
//   		StorageClass: jsii.String("storageClass"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WithoutMetadata: jsii.String("withoutMetadata"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html
//
type CfnTableProps struct {
	// The name of the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html#cfn-s3tables-table-namespace
	//
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The format of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html#cfn-s3tables-table-opentableformat
	//
	OpenTableFormat *string `field:"required" json:"openTableFormat" yaml:"openTableFormat"`
	// The Amazon Resource Name (ARN) of the table bucket to create the table in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html#cfn-s3tables-table-tablebucketarn
	//
	TableBucketArn *string `field:"required" json:"tableBucketArn" yaml:"tableBucketArn"`
	// The name for the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html#cfn-s3tables-table-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Contains details about the compaction settings for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html#cfn-s3tables-table-compaction
	//
	Compaction interface{} `field:"optional" json:"compaction" yaml:"compaction"`
	// Contains details about the metadata for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html#cfn-s3tables-table-icebergmetadata
	//
	IcebergMetadata interface{} `field:"optional" json:"icebergMetadata" yaml:"icebergMetadata"`
	// Contains details about the Iceberg snapshot management settings for the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html#cfn-s3tables-table-snapshotmanagement
	//
	SnapshotManagement interface{} `field:"optional" json:"snapshotManagement" yaml:"snapshotManagement"`
	// The configuration details for the storage class of tables or table buckets.
	//
	// This allows you to optimize storage costs by selecting the appropriate storage class based on your access patterns and performance requirements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html#cfn-s3tables-table-storageclassconfiguration
	//
	StorageClassConfiguration interface{} `field:"optional" json:"storageClassConfiguration" yaml:"storageClassConfiguration"`
	// User tags (key-value pairs) to associate with the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html#cfn-s3tables-table-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indicates that you don't want to specify a schema for the table.
	//
	// This property is mutually exclusive to `IcebergMetadata` , and its only possible value is `Yes` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html#cfn-s3tables-table-withoutmetadata
	//
	WithoutMetadata *string `field:"optional" json:"withoutMetadata" yaml:"withoutMetadata"`
}

