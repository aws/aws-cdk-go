package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   icebergSortFieldProperty := &IcebergSortFieldProperty{
//   	Direction: jsii.String("direction"),
//   	NullOrder: jsii.String("nullOrder"),
//   	SourceId: jsii.Number(123),
//   	Transform: jsii.String("transform"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergsortfield.html
//
type CfnTablePropsMixin_IcebergSortFieldProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergsortfield.html#cfn-glue-table-icebergsortfield-direction
	//
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergsortfield.html#cfn-glue-table-icebergsortfield-nullorder
	//
	NullOrder *string `field:"optional" json:"nullOrder" yaml:"nullOrder"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergsortfield.html#cfn-glue-table-icebergsortfield-sourceid
	//
	SourceId *float64 `field:"optional" json:"sourceId" yaml:"sourceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergsortfield.html#cfn-glue-table-icebergsortfield-transform
	//
	Transform *string `field:"optional" json:"transform" yaml:"transform"`
}

