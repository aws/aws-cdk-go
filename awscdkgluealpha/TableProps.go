// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &TableProps{
//   	Database: myDatabase,
//   	Columns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	PartitionKeys: []*column{
//   		&column{
//   			Name: jsii.String("year"),
//   			Type: glue.Schema_SMALL_INT(),
//   		},
//   		&column{
//   			Name: jsii.String("month"),
//   			Type: glue.Schema_SMALL_INT(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   })
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
	// S3 bucket in which to store data.
	// Experimental.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// Indicates whether the table's data is compressed or not.
	// Experimental.
	Compressed *bool `field:"optional" json:"compressed" yaml:"compressed"`
	// Description of the table.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enables partition filtering.
	// See: https://docs.aws.amazon.com/athena/latest/ug/glue-best-practices.html#glue-best-practices-partition-index
	//
	// Experimental.
	EnablePartitionFiltering *bool `field:"optional" json:"enablePartitionFiltering" yaml:"enablePartitionFiltering"`
	// The kind of encryption to secure the data with.
	//
	// You can only provide this option if you are not explicitly passing in a bucket.
	//
	// If you choose `SSE-KMS`, you *can* provide an un-managed KMS key with `encryptionKey`.
	// If you choose `CSE-KMS`, you *must* provide an un-managed KMS key with `encryptionKey`.
	// Experimental.
	Encryption TableEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for bucket encryption.
	//
	// The `encryption` property must be `SSE-KMS` or `CSE-KMS`.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Partition indexes on the table.
	//
	// A maximum of 3 indexes
	// are allowed on a table. Keys in the index must be part
	// of the table's partition keys.
	// Experimental.
	PartitionIndexes *[]*PartitionIndex `field:"optional" json:"partitionIndexes" yaml:"partitionIndexes"`
	// Partition columns of the table.
	// Experimental.
	PartitionKeys *[]*Column `field:"optional" json:"partitionKeys" yaml:"partitionKeys"`
	// S3 prefix under which table objects are stored.
	// Experimental.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
	// Indicates whether the table data is stored in subdirectories.
	// Experimental.
	StoredAsSubDirectories *bool `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
	// Name of the table.
	// Experimental.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

