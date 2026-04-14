package awss3tables


// Contains details about the schema version 2 (V2) for an Iceberg table that supports Apache Iceberg Nested Types (struct, list, map).
//
// Primitive types are also supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var type interface{}
//
//   icebergSchemaV2Property := &IcebergSchemaV2Property{
//   	SchemaV2FieldList: []interface{}{
//   		&SchemaV2FieldProperty{
//   			Id: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			Required: jsii.Boolean(false),
//   			Type: type,
//
//   			// the properties below are optional
//   			Doc: jsii.String("doc"),
//   		},
//   	},
//   	SchemaV2FieldType: jsii.String("schemaV2FieldType"),
//
//   	// the properties below are optional
//   	IdentifierFieldIds: []interface{}{
//   		jsii.Number(123),
//   	},
//   	SchemaId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergschemav2.html
//
type CfnTable_IcebergSchemaV2Property struct {
	// List of schema fields that support nested types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergschemav2.html#cfn-s3tables-table-icebergschemav2-schemav2fieldlist
	//
	SchemaV2FieldList interface{} `field:"required" json:"schemaV2FieldList" yaml:"schemaV2FieldList"`
	// The type of the top-level schema, which is always 'struct'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergschemav2.html#cfn-s3tables-table-icebergschemav2-schemav2fieldtype
	//
	SchemaV2FieldType *string `field:"required" json:"schemaV2FieldType" yaml:"schemaV2FieldType"`
	// A list of field IDs that are used as the identifier fields for the table.
	//
	// Identifier fields uniquely identify a row in the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergschemav2.html#cfn-s3tables-table-icebergschemav2-identifierfieldids
	//
	IdentifierFieldIds interface{} `field:"optional" json:"identifierFieldIds" yaml:"identifierFieldIds"`
	// An optional unique identifier for the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergschemav2.html#cfn-s3tables-table-icebergschemav2-schemaid
	//
	SchemaId *float64 `field:"optional" json:"schemaId" yaml:"schemaId"`
}

