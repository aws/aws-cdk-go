package awsglue


// Describes the physical storage of table data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var skewedColumnValueLocationMaps interface{}
//
//   storageDescriptorProperty := &storageDescriptorProperty{
//   	bucketColumns: []*string{
//   		jsii.String("bucketColumns"),
//   	},
//   	columns: []interface{}{
//   		&columnProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			comment: jsii.String("comment"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	compressed: jsii.Boolean(false),
//   	inputFormat: jsii.String("inputFormat"),
//   	location: jsii.String("location"),
//   	numberOfBuckets: jsii.Number(123),
//   	outputFormat: jsii.String("outputFormat"),
//   	parameters: parameters,
//   	schemaReference: &schemaReferenceProperty{
//   		schemaId: &schemaIdProperty{
//   			registryName: jsii.String("registryName"),
//   			schemaArn: jsii.String("schemaArn"),
//   			schemaName: jsii.String("schemaName"),
//   		},
//   		schemaVersionId: jsii.String("schemaVersionId"),
//   		schemaVersionNumber: jsii.Number(123),
//   	},
//   	serdeInfo: &serdeInfoProperty{
//   		name: jsii.String("name"),
//   		parameters: parameters,
//   		serializationLibrary: jsii.String("serializationLibrary"),
//   	},
//   	skewedInfo: &skewedInfoProperty{
//   		skewedColumnNames: []*string{
//   			jsii.String("skewedColumnNames"),
//   		},
//   		skewedColumnValueLocationMaps: skewedColumnValueLocationMaps,
//   		skewedColumnValues: []*string{
//   			jsii.String("skewedColumnValues"),
//   		},
//   	},
//   	sortColumns: []interface{}{
//   		&orderProperty{
//   			column: jsii.String("column"),
//
//   			// the properties below are optional
//   			sortOrder: jsii.Number(123),
//   		},
//   	},
//   	storedAsSubDirectories: jsii.Boolean(false),
//   }
//
type CfnPartition_StorageDescriptorProperty struct {
	// A list of reducer grouping columns, clustering columns, and bucketing columns in the table.
	BucketColumns *[]*string `field:"optional" json:"bucketColumns" yaml:"bucketColumns"`
	// A list of the `Columns` in the table.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// `True` if the data in the table is compressed, or `False` if not.
	Compressed interface{} `field:"optional" json:"compressed" yaml:"compressed"`
	// The input format: `SequenceFileInputFormat` (binary), or `TextInputFormat` , or a custom format.
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// The physical location of the table.
	//
	// By default, this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The number of buckets.
	//
	// You must specify this property if the partition contains any dimension columns.
	NumberOfBuckets *float64 `field:"optional" json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// The output format: `SequenceFileOutputFormat` (binary), or `IgnoreKeyTextOutputFormat` , or a custom format.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// The user-supplied properties in key-value form.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An object that references a schema stored in the AWS Glue Schema Registry.
	SchemaReference interface{} `field:"optional" json:"schemaReference" yaml:"schemaReference"`
	// The serialization/deserialization (SerDe) information.
	SerdeInfo interface{} `field:"optional" json:"serdeInfo" yaml:"serdeInfo"`
	// The information about values that appear frequently in a column (skewed values).
	SkewedInfo interface{} `field:"optional" json:"skewedInfo" yaml:"skewedInfo"`
	// A list specifying the sort order of each bucket in the table.
	SortColumns interface{} `field:"optional" json:"sortColumns" yaml:"sortColumns"`
	// `True` if the table data is stored in subdirectories, or `False` if not.
	StoredAsSubDirectories interface{} `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

