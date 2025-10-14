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
//   				Required: jsii.Boolean(false),
//   			},
//   		},
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
}

