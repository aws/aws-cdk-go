package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var database database
//   var dataFormat dataFormat
//   var key key
//   var storageParameter storageParameter
//
//   tableProps := &TableProps{
//   	Columns: []column{
//   		&column{
//   			Name: jsii.String("name"),
//   			Type: &Type{
//   				InputString: jsii.String("inputString"),
//   				IsPrimitive: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			Comment: jsii.String("comment"),
//   		},
//   	},
//   	Database: database,
//   	DataFormat: dataFormat,
//
//   	// the properties below are optional
//   	Bucket: bucket,
//   	Compressed: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	EnablePartitionFiltering: jsii.Boolean(false),
//   	Encryption: glue_alpha.TableEncryption_S3_MANAGED,
//   	EncryptionKey: key,
//   	PartitionIndexes: []partitionIndex{
//   		&partitionIndex{
//   			KeyNames: []*string{
//   				jsii.String("keyNames"),
//   			},
//
//   			// the properties below are optional
//   			IndexName: jsii.String("indexName"),
//   		},
//   	},
//   	PartitionKeys: []*column{
//   		&column{
//   			Name: jsii.String("name"),
//   			Type: &Type{
//   				InputString: jsii.String("inputString"),
//   				IsPrimitive: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			Comment: jsii.String("comment"),
//   		},
//   	},
//   	S3Prefix: jsii.String("s3Prefix"),
//   	StorageParameters: []*storageParameter{
//   		storageParameter,
//   	},
//   	StoredAsSubDirectories: jsii.Boolean(false),
//   	TableName: jsii.String("tableName"),
//   }
//
// Experimental.
type TableProps struct {
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

