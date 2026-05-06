package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   icebergStructFieldProperty := &IcebergStructFieldProperty{
//   	Doc: jsii.String("doc"),
//   	Id: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Required: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html
//
type CfnTablePropsMixin_IcebergStructFieldProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html#cfn-glue-table-icebergstructfield-doc
	//
	Doc *string `field:"optional" json:"doc" yaml:"doc"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html#cfn-glue-table-icebergstructfield-id
	//
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html#cfn-glue-table-icebergstructfield-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html#cfn-glue-table-icebergstructfield-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html#cfn-glue-table-icebergstructfield-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

