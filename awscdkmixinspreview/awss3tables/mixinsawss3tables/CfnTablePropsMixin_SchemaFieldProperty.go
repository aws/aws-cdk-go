package mixinsawss3tables


// Contains details about a schema field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaFieldProperty := &SchemaFieldProperty{
//   	Name: jsii.String("name"),
//   	Required: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-schemafield.html
//
type CfnTablePropsMixin_SchemaFieldProperty struct {
	// The name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-schemafield.html#cfn-s3tables-table-schemafield-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A Boolean value that specifies whether values are required for each row in this field.
	//
	// By default, this is `false` and null values are allowed in the field. If this is `true` the field does not allow null values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-schemafield.html#cfn-s3tables-table-schemafield-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The field type.
	//
	// S3 Tables supports all Apache Iceberg primitive types. For more information, see the [Apache Iceberg documentation](https://docs.aws.amazon.com/https://iceberg.apache.org/spec/#primitive-types) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-schemafield.html#cfn-s3tables-table-schemafield-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

