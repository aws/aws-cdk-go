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
//   storageDescriptorProperty := &StorageDescriptorProperty{
//   	BucketColumns: []*string{
//   		jsii.String("bucketColumns"),
//   	},
//   	Columns: []interface{}{
//   		&ColumnProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Comment: jsii.String("comment"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Compressed: jsii.Boolean(false),
//   	InputFormat: jsii.String("inputFormat"),
//   	Location: jsii.String("location"),
//   	NumberOfBuckets: jsii.Number(123),
//   	OutputFormat: jsii.String("outputFormat"),
//   	Parameters: parameters,
//   	SchemaReference: &SchemaReferenceProperty{
//   		SchemaId: &SchemaIdProperty{
//   			RegistryName: jsii.String("registryName"),
//   			SchemaArn: jsii.String("schemaArn"),
//   			SchemaName: jsii.String("schemaName"),
//   		},
//   		SchemaVersionId: jsii.String("schemaVersionId"),
//   		SchemaVersionNumber: jsii.Number(123),
//   	},
//   	SerdeInfo: &SerdeInfoProperty{
//   		Name: jsii.String("name"),
//   		Parameters: parameters,
//   		SerializationLibrary: jsii.String("serializationLibrary"),
//   	},
//   	SkewedInfo: &SkewedInfoProperty{
//   		SkewedColumnNames: []*string{
//   			jsii.String("skewedColumnNames"),
//   		},
//   		SkewedColumnValueLocationMaps: skewedColumnValueLocationMaps,
//   		SkewedColumnValues: []*string{
//   			jsii.String("skewedColumnValues"),
//   		},
//   	},
//   	SortColumns: []interface{}{
//   		&OrderProperty{
//   			Column: jsii.String("column"),
//
//   			// the properties below are optional
//   			SortOrder: jsii.Number(123),
//   		},
//   	},
//   	StoredAsSubDirectories: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html
//
type CfnPartition_StorageDescriptorProperty struct {
	// A list of reducer grouping columns, clustering columns, and bucketing columns in the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-bucketcolumns
	//
	BucketColumns *[]*string `field:"optional" json:"bucketColumns" yaml:"bucketColumns"`
	// A list of the `Columns` in the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// `True` if the data in the table is compressed, or `False` if not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-compressed
	//
	Compressed interface{} `field:"optional" json:"compressed" yaml:"compressed"`
	// The input format: `SequenceFileInputFormat` (binary), or `TextInputFormat` , or a custom format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-inputformat
	//
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// The physical location of the table.
	//
	// By default, this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The number of buckets.
	//
	// You must specify this property if the partition contains any dimension columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-numberofbuckets
	//
	NumberOfBuckets *float64 `field:"optional" json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// The output format: `SequenceFileOutputFormat` (binary), or `IgnoreKeyTextOutputFormat` , or a custom format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-outputformat
	//
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// The user-supplied properties in key-value form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An object that references a schema stored in the AWS Glue Schema Registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-schemareference
	//
	SchemaReference interface{} `field:"optional" json:"schemaReference" yaml:"schemaReference"`
	// The serialization/deserialization (SerDe) information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-serdeinfo
	//
	SerdeInfo interface{} `field:"optional" json:"serdeInfo" yaml:"serdeInfo"`
	// The information about values that appear frequently in a column (skewed values).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-skewedinfo
	//
	SkewedInfo interface{} `field:"optional" json:"skewedInfo" yaml:"skewedInfo"`
	// A list specifying the sort order of each bucket in the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-sortcolumns
	//
	SortColumns interface{} `field:"optional" json:"sortColumns" yaml:"sortColumns"`
	// `True` if the table data is stored in subdirectories, or `False` if not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-storedassubdirectories
	//
	StoredAsSubDirectories interface{} `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

