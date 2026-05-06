package awsglue


// Specifies an input structure that defines an Apache Iceberg metadata table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var properties interface{}
//
//   icebergInputProperty := &IcebergInputProperty{
//   	IcebergTableInput: &IcebergTableInputProperty{
//   		Location: jsii.String("location"),
//   		PartitionSpec: &IcebergPartitionSpecProperty{
//   			Fields: []interface{}{
//   				&IcebergPartitionFieldProperty{
//   					FieldId: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					SourceId: jsii.Number(123),
//   					Transform: jsii.String("transform"),
//   				},
//   			},
//   			SpecId: jsii.Number(123),
//   		},
//   		Properties: properties,
//   		Schema: &IcebergSchemaProperty{
//   			Fields: []interface{}{
//   				&IcebergStructFieldProperty{
//   					Doc: jsii.String("doc"),
//   					Id: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					Required: jsii.Boolean(false),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			IdentifierFieldIds: []interface{}{
//   				jsii.Number(123),
//   			},
//   			SchemaId: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   		WriteOrder: &IcebergSortOrderProperty{
//   			Fields: []interface{}{
//   				&IcebergSortFieldProperty{
//   					Direction: jsii.String("direction"),
//   					NullOrder: jsii.String("nullOrder"),
//   					SourceId: jsii.Number(123),
//   					Transform: jsii.String("transform"),
//   				},
//   			},
//   			OrderId: jsii.Number(123),
//   		},
//   	},
//   	MetadataOperation: jsii.String("metadataOperation"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-iceberginput.html
//
type CfnTablePropsMixin_IcebergInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-iceberginput.html#cfn-glue-table-iceberginput-icebergtableinput
	//
	IcebergTableInput interface{} `field:"optional" json:"icebergTableInput" yaml:"icebergTableInput"`
	// A required metadata operation.
	//
	// Can only be set to CREATE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-iceberginput.html#cfn-glue-table-iceberginput-metadataoperation
	//
	MetadataOperation *string `field:"optional" json:"metadataOperation" yaml:"metadataOperation"`
	// The table version for the Iceberg table.
	//
	// Defaults to 2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-iceberginput.html#cfn-glue-table-iceberginput-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

