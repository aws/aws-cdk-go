package previewawss3tablesmixins


// Contains details about the metadata for an Iceberg table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   icebergMetadataProperty := &IcebergMetadataProperty{
//   	IcebergPartitionSpec: &IcebergPartitionSpecProperty{
//   		Fields: []interface{}{
//   			&IcebergPartitionFieldProperty{
//   				FieldId: jsii.Number(123),
//   				Name: jsii.String("name"),
//   				SourceId: jsii.Number(123),
//   				Transform: jsii.String("transform"),
//   			},
//   		},
//   		SpecId: jsii.Number(123),
//   	},
//   	IcebergSchema: &IcebergSchemaProperty{
//   		SchemaFieldList: []interface{}{
//   			&SchemaFieldProperty{
//   				Id: jsii.Number(123),
//   				Name: jsii.String("name"),
//   				Required: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//   			},
//   		},
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
//   		OrderId: jsii.Number(123),
//   	},
//   	TableProperties: map[string]*string{
//   		"tablePropertiesKey": jsii.String("tableProperties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html
//
type CfnTablePropsMixin_IcebergMetadataProperty struct {
	// Partition specification for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html#cfn-s3tables-table-icebergmetadata-icebergpartitionspec
	//
	IcebergPartitionSpec interface{} `field:"optional" json:"icebergPartitionSpec" yaml:"icebergPartitionSpec"`
	// The schema for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html#cfn-s3tables-table-icebergmetadata-icebergschema
	//
	IcebergSchema interface{} `field:"optional" json:"icebergSchema" yaml:"icebergSchema"`
	// Sort order specification for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html#cfn-s3tables-table-icebergmetadata-icebergsortorder
	//
	IcebergSortOrder interface{} `field:"optional" json:"icebergSortOrder" yaml:"icebergSortOrder"`
	// Iceberg table properties (e.g., format-version, write.parquet.compression-codec).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html#cfn-s3tables-table-icebergmetadata-tableproperties
	//
	TableProperties interface{} `field:"optional" json:"tableProperties" yaml:"tableProperties"`
}

