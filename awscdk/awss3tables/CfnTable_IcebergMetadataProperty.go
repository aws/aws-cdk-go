package awss3tables


// Contains details about the metadata for an Iceberg table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergMetadataProperty := &IcebergMetadataProperty{
//   	IcebergSchema: &IcebergSchemaProperty{
//   		SchemaFieldList: []interface{}{
//   			&SchemaFieldProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Id: jsii.Number(123),
//   				Required: jsii.Boolean(false),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	IcebergPartitionSpec: &IcebergPartitionSpecProperty{
//   		Fields: []interface{}{
//   			&IcebergPartitionFieldProperty{
//   				Name: jsii.String("name"),
//   				SourceId: jsii.Number(123),
//   				Transform: jsii.String("transform"),
//
//   				// the properties below are optional
//   				FieldId: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		SpecId: jsii.Number(123),
//   	},
//   	IcebergSortOrder: &IcebergSortOrderProperty{
//   		Fields: []interface{}{
//   			&IcebergSortFieldProperty{
//   				Direction: jsii.String("direction"),
//   				NullOrder: jsii.String("nullOrder"),
//   				SourceId: jsii.Number(123),
//   				Transform: jsii.String("transform"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		OrderId: jsii.Number(123),
//   	},
//   	TableProperties: map[string]*string{
//   		"tablePropertiesKey": jsii.String("tableProperties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html
//
type CfnTable_IcebergMetadataProperty struct {
	// The schema for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html#cfn-s3tables-table-icebergmetadata-icebergschema
	//
	IcebergSchema interface{} `field:"required" json:"icebergSchema" yaml:"icebergSchema"`
	// Partition specification for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html#cfn-s3tables-table-icebergmetadata-icebergpartitionspec
	//
	IcebergPartitionSpec interface{} `field:"optional" json:"icebergPartitionSpec" yaml:"icebergPartitionSpec"`
	// Sort order specification for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html#cfn-s3tables-table-icebergmetadata-icebergsortorder
	//
	IcebergSortOrder interface{} `field:"optional" json:"icebergSortOrder" yaml:"icebergSortOrder"`
	// Iceberg table properties (e.g., format-version, write.parquet.compression-codec).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html#cfn-s3tables-table-icebergmetadata-tableproperties
	//
	TableProperties interface{} `field:"optional" json:"tableProperties" yaml:"tableProperties"`
}

