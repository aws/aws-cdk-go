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
//   partitionInputProperty := &partitionInputProperty{
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//
//   	// the properties below are optional
//   	parameters: parameters,
//   	storageDescriptor: &storageDescriptorProperty{
//   		bucketColumns: []*string{
//   			jsii.String("bucketColumns"),
//   		},
//   		columns: []interface{}{
//   			&columnProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				comment: jsii.String("comment"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		compressed: jsii.Boolean(false),
//   		inputFormat: jsii.String("inputFormat"),
//   		location: jsii.String("location"),
//   		numberOfBuckets: jsii.Number(123),
//   		outputFormat: jsii.String("outputFormat"),
//   		parameters: parameters,
//   		schemaReference: &schemaReferenceProperty{
//   			schemaId: &schemaIdProperty{
//   				registryName: jsii.String("registryName"),
//   				schemaArn: jsii.String("schemaArn"),
//   				schemaName: jsii.String("schemaName"),
//   			},
//   			schemaVersionId: jsii.String("schemaVersionId"),
//   			schemaVersionNumber: jsii.Number(123),
//   		},
//   		serdeInfo: &serdeInfoProperty{
//   			name: jsii.String("name"),
//   			parameters: parameters,
//   			serializationLibrary: jsii.String("serializationLibrary"),
//   		},
//   		skewedInfo: &skewedInfoProperty{
//   			skewedColumnNames: []*string{
//   				jsii.String("skewedColumnNames"),
//   			},
//   			skewedColumnValueLocationMaps: skewedColumnValueLocationMaps,
//   			skewedColumnValues: []*string{
//   				jsii.String("skewedColumnValues"),
//   			},
//   		},
//   		sortColumns: []interface{}{
//   			&orderProperty{
//   				column: jsii.String("column"),
//
//   				// the properties below are optional
//   				sortOrder: jsii.Number(123),
//   			},
//   		},
//   		storedAsSubDirectories: jsii.Boolean(false),
//   	},
//   }
//
type CfnPartition_PartitionInputProperty struct {
	// The values of the partition.
	//
	// Although this parameter is not required by the SDK, you must specify this parameter for a valid input.
	//
	// The values for the keys for the new partition must be passed as an array of String objects that must be ordered in the same order as the partition keys appearing in the Amazon S3 prefix. Otherwise AWS Glue will add the values to the wrong keys.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// These key-value pairs define partition parameters.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Provides information about the physical location where the partition is stored.
	StorageDescriptor interface{} `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
}

