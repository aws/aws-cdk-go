package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Example:
//   var myDatabase Database
//
//   glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
//   	Database: myDatabase,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("data"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	PartitionKeys: []Column{
//   		&Column{
//   			Name: jsii.String("date"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   	PartitionProjection: map[string]PartitionProjectionConfiguration{
//   		"date": glue.PartitionProjectionConfiguration_date(&DatePartitionProjectionConfigurationProps{
//   			"min": jsii.String("2020-01-01"),
//   			"max": jsii.String("2023-12-31"),
//   			"format": jsii.String("yyyy-MM-dd"),
//   			"interval": jsii.Number(1),
//   			 // optional, defaults to 1
//   			"intervalUnit": glue.DateIntervalUnit_DAYS,
//   		}),
//   	},
//   })
//
// Experimental.
type S3TableProps struct {
	// Columns of the table.
	// Experimental.
	Columns *[]*Column `field:"required" json:"columns" yaml:"columns"`
	// Database in which to store the table.
	// Experimental.
	Database IDatabase `field:"required" json:"database" yaml:"database"`
	// Storage type of the table's data.
	// Experimental.
	DataFormat DataFormat `field:"required" json:"dataFormat" yaml:"dataFormat"`
	// Indicates whether the table's data is compressed or not.
	// Default: false.
	//
	// Experimental.
	Compressed *bool `field:"optional" json:"compressed" yaml:"compressed"`
	// Description of the table.
	// Default: generated.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enables partition filtering.
	// See: https://docs.aws.amazon.com/athena/latest/ug/glue-best-practices.html#glue-best-practices-partition-index
	//
	// Default: - The parameter is not defined.
	//
	// Experimental.
	EnablePartitionFiltering *bool `field:"optional" json:"enablePartitionFiltering" yaml:"enablePartitionFiltering"`
	// The key/value pairs define properties associated with the table.
	//
	// The key/value pairs that are allowed to be submitted are not limited, however their functionality is not guaranteed.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-parameters
	//
	// Default: - The parameter is not defined.
	//
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Partition indexes on the table.
	//
	// A maximum of 3 indexes
	// are allowed on a table. Keys in the index must be part
	// of the table's partition keys.
	// Default: table has no partition indexes.
	//
	// Experimental.
	PartitionIndexes *[]*PartitionIndex `field:"optional" json:"partitionIndexes" yaml:"partitionIndexes"`
	// Partition columns of the table.
	// Default: table is not partitioned.
	//
	// Experimental.
	PartitionKeys *[]*Column `field:"optional" json:"partitionKeys" yaml:"partitionKeys"`
	// Partition projection configuration for this table.
	//
	// Partition projection allows Athena to automatically add new partitions
	// without requiring `ALTER TABLE ADD PARTITION` statements.
	// See: https://docs.aws.amazon.com/athena/latest/ug/partition-projection.html
	//
	// Default: - No partition projection.
	//
	// Experimental.
	PartitionProjection *map[string]PartitionProjectionConfiguration `field:"optional" json:"partitionProjection" yaml:"partitionProjection"`
	// The user-supplied properties for the description of the physical storage of this table.
	//
	// These properties help describe the format of the data that is stored within the crawled data sources.
	//
	// The key/value pairs that are allowed to be submitted are not limited, however their functionality is not guaranteed.
	//
	// Some keys will be auto-populated by glue crawlers, however, you can override them by specifying the key and value in this property.
	//
	// Example:
	//      declare const glueDatabase: glue.IDatabase;
	//      const table = new glue.Table(this, 'Table', {
	//        storageParameters: [
	//            glue.StorageParameter.skipHeaderLineCount(1),
	//            glue.StorageParameter.compressionType(glue.CompressionType.GZIP),
	//            glue.StorageParameter.custom('foo', 'bar'), // Will have no effect
	//            glue.StorageParameter.custom('separatorChar', ','), // Will describe the separator char used in the data
	//            glue.StorageParameter.custom(glue.StorageParameters.WRITE_PARALLEL, 'off'),
	//        ],
	//        // ...
	//        database: glueDatabase,
	//        columns: [{
	//            name: 'col1',
	//            type: glue.Schema.STRING,
	//        }],
	//        dataFormat: glue.DataFormat.CSV,
	//      });
	//
	// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_
	//
	// Default: - The parameter is not defined.
	//
	// Experimental.
	StorageParameters *[]StorageParameter `field:"optional" json:"storageParameters" yaml:"storageParameters"`
	// Indicates whether the table data is stored in subdirectories.
	// Default: false.
	//
	// Experimental.
	StoredAsSubDirectories *bool `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
	// Name of the table.
	// Default: - generated by CDK.
	//
	// Experimental.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// S3 bucket in which to store data.
	// Default: one is created for you.
	//
	// Experimental.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// The kind of encryption to secure the data with.
	//
	// You can only provide this option if you are not explicitly passing in a bucket.
	//
	// If you choose `SSE-KMS`, you *can* provide an un-managed KMS key with `encryptionKey`.
	// If you choose `CSE-KMS`, you *must* provide an un-managed KMS key with `encryptionKey`.
	// Default: BucketEncryption.S3_MANAGED
	//
	// Experimental.
	Encryption TableEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for bucket encryption.
	//
	// The `encryption` property must be `SSE-KMS` or `CSE-KMS`.
	// Default: key is managed by KMS.
	//
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// S3 prefix under which table objects are stored.
	// Default: - No prefix. The data will be stored under the root of the bucket.
	//
	// Experimental.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
}

