package previewawss3tablesmixins


// Contains details about the metadata for an Iceberg table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   icebergMetadataProperty := &IcebergMetadataProperty{
//   	IcebergSchema: &IcebergSchemaProperty{
//   		SchemaFieldList: []interface{}{
//   			&SchemaFieldProperty{
//   				Name: jsii.String("name"),
//   				Required: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html
//
type CfnTablePropsMixin_IcebergMetadataProperty struct {
	// The schema for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergmetadata.html#cfn-s3tables-table-icebergmetadata-icebergschema
	//
	IcebergSchema interface{} `field:"optional" json:"icebergSchema" yaml:"icebergSchema"`
}

