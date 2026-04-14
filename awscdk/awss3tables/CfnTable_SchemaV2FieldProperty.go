package awss3tables


// Contains details about a schema field for an Iceberg table that supports nested types (struct, list, map).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var type interface{}
//
//   schemaV2FieldProperty := &SchemaV2FieldProperty{
//   	Id: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Required: jsii.Boolean(false),
//   	Type: type,
//
//   	// the properties below are optional
//   	Doc: jsii.String("doc"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-schemav2field.html
//
type CfnTable_SchemaV2FieldProperty struct {
	// The unique identifier for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-schemav2field.html#cfn-s3tables-table-schemav2field-id
	//
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// The name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-schemav2field.html#cfn-s3tables-table-schemav2field-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A Boolean value that specifies whether values are required for each row in this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-schemav2field.html#cfn-s3tables-table-schemav2field-required
	//
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-schemav2field.html#cfn-s3tables-table-schemav2field-type
	//
	Type interface{} `field:"required" json:"type" yaml:"type"`
	// Optional documentation for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-schemav2field.html#cfn-s3tables-table-schemav2field-doc
	//
	Doc *string `field:"optional" json:"doc" yaml:"doc"`
}

