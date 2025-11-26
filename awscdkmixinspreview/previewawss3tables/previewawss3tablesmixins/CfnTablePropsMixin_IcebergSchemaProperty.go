package previewawss3tablesmixins


// Contains details about the schema for an Iceberg table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   icebergSchemaProperty := &IcebergSchemaProperty{
//   	SchemaFieldList: []interface{}{
//   		&SchemaFieldProperty{
//   			Name: jsii.String("name"),
//   			Required: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergschema.html
//
type CfnTablePropsMixin_IcebergSchemaProperty struct {
	// The schema fields for the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergschema.html#cfn-s3tables-table-icebergschema-schemafieldlist
	//
	SchemaFieldList interface{} `field:"optional" json:"schemaFieldList" yaml:"schemaFieldList"`
}

