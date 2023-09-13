package awscdkgluealpha


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   var database database
//   var dataFormat dataFormat
//   var storageParameter storageParameter
//
//   tableBaseProps := &TableBaseProps{
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
//   	Compressed: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	EnablePartitionFiltering: jsii.Boolean(false),
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
//   	StorageParameters: []*storageParameter{
//   		storageParameter,
//   	},
//   	StoredAsSubDirectories: jsii.Boolean(false),
//   	TableName: jsii.String("tableName"),
//   }
//
// Experimental.
type TableBaseProps struct {
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
	//   var glueDatabase iDatabase
	//
	//   table := glue.NewTable(this, jsii.String("Table"), &S3TableProps{
	//   	StorageParameters: []storageParameter{
	//   		glue.*storageParameter_SkipHeaderLineCount(jsii.Number(1)),
	//   		glue.*storageParameter_CompressionType(glue.CompressionType_GZIP),
	//   		glue.*storageParameter_Custom(jsii.String("foo"), jsii.String("bar")),
	//   		glue.*storageParameter_*Custom(jsii.String("separatorChar"), jsii.String(",")),
	//   		glue.*storageParameter_*Custom(glue.StorageParameters_WRITE_PARALLEL, jsii.String("off")),
	//   	},
	//   	// ...
	//   	Database: glueDatabase,
	//   	Columns: []column{
	//   		&column{
	//   			Name: jsii.String("col1"),
	//   			Type: glue.Schema_STRING(),
	//   		},
	//   	},
	//   	DataFormat: glue.DataFormat_CSV(),
	//   })
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
}

