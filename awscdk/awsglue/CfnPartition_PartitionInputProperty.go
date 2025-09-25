package awsglue


// The structure used to create and update a partition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var skewedColumnValueLocationMaps interface{}
//
//   partitionInputProperty := &PartitionInputProperty{
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//
//   	// the properties below are optional
//   	Parameters: parameters,
//   	StorageDescriptor: &StorageDescriptorProperty{
//   		BucketColumns: []*string{
//   			jsii.String("bucketColumns"),
//   		},
//   		Columns: []interface{}{
//   			&ColumnProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Comment: jsii.String("comment"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Compressed: jsii.Boolean(false),
//   		InputFormat: jsii.String("inputFormat"),
//   		Location: jsii.String("location"),
//   		NumberOfBuckets: jsii.Number(123),
//   		OutputFormat: jsii.String("outputFormat"),
//   		Parameters: parameters,
//   		SchemaReference: &SchemaReferenceProperty{
//   			SchemaId: &SchemaIdProperty{
//   				RegistryName: jsii.String("registryName"),
//   				SchemaArn: jsii.String("schemaArn"),
//   				SchemaName: jsii.String("schemaName"),
//   			},
//   			SchemaVersionId: jsii.String("schemaVersionId"),
//   			SchemaVersionNumber: jsii.Number(123),
//   		},
//   		SerdeInfo: &SerdeInfoProperty{
//   			Name: jsii.String("name"),
//   			Parameters: parameters,
//   			SerializationLibrary: jsii.String("serializationLibrary"),
//   		},
//   		SkewedInfo: &SkewedInfoProperty{
//   			SkewedColumnNames: []*string{
//   				jsii.String("skewedColumnNames"),
//   			},
//   			SkewedColumnValueLocationMaps: skewedColumnValueLocationMaps,
//   			SkewedColumnValues: []*string{
//   				jsii.String("skewedColumnValues"),
//   			},
//   		},
//   		SortColumns: []interface{}{
//   			&OrderProperty{
//   				Column: jsii.String("column"),
//
//   				// the properties below are optional
//   				SortOrder: jsii.Number(123),
//   			},
//   		},
//   		StoredAsSubDirectories: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-partitioninput.html
//
type CfnPartition_PartitionInputProperty struct {
	// The values of the partition.
	//
	// Although this parameter is not required by the SDK, you must specify this parameter for a valid input.
	//
	// The values for the keys for the new partition must be passed as an array of String objects that must be ordered in the same order as the partition keys appearing in the Amazon S3 prefix. Otherwise AWS Glue will add the values to the wrong keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-partitioninput.html#cfn-glue-partition-partitioninput-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// These key-value pairs define partition parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-partitioninput.html#cfn-glue-partition-partitioninput-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Provides information about the physical location where the partition is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-partitioninput.html#cfn-glue-partition-partitioninput-storagedescriptor
	//
	StorageDescriptor interface{} `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
}

