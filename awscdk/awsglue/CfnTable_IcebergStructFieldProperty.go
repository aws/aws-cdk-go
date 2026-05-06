package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergStructFieldProperty := &IcebergStructFieldProperty{
//   	Id: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Required: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Doc: jsii.String("doc"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html
//
type CfnTable_IcebergStructFieldProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html#cfn-glue-table-icebergstructfield-id
	//
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html#cfn-glue-table-icebergstructfield-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html#cfn-glue-table-icebergstructfield-required
	//
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html#cfn-glue-table-icebergstructfield-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergstructfield.html#cfn-glue-table-icebergstructfield-doc
	//
	Doc *string `field:"optional" json:"doc" yaml:"doc"`
}

