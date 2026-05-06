package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   icebergSchemaProperty := &IcebergSchemaProperty{
//   	Fields: []interface{}{
//   		&IcebergStructFieldProperty{
//   			Doc: jsii.String("doc"),
//   			Id: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			Required: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	IdentifierFieldIds: []interface{}{
//   		jsii.Number(123),
//   	},
//   	SchemaId: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergschema.html
//
type CfnTablePropsMixin_IcebergSchemaProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergschema.html#cfn-glue-table-icebergschema-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergschema.html#cfn-glue-table-icebergschema-identifierfieldids
	//
	IdentifierFieldIds interface{} `field:"optional" json:"identifierFieldIds" yaml:"identifierFieldIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergschema.html#cfn-glue-table-icebergschema-schemaid
	//
	SchemaId *float64 `field:"optional" json:"schemaId" yaml:"schemaId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergschema.html#cfn-glue-table-icebergschema-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

