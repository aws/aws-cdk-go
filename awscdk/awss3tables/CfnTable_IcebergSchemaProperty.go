package awss3tables


// Contains details about the schema for an Iceberg table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergSchemaProperty := &IcebergSchemaProperty{
//   	SchemaFieldList: []interface{}{
//   		&SchemaFieldProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Required: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergschema.html
//
type CfnTable_IcebergSchemaProperty struct {
	// The schema fields for the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergschema.html#cfn-s3tables-table-icebergschema-schemafieldlist
	//
	SchemaFieldList interface{} `field:"required" json:"schemaFieldList" yaml:"schemaFieldList"`
}

