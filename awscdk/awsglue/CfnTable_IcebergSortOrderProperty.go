package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergSortOrderProperty := &IcebergSortOrderProperty{
//   	Fields: []interface{}{
//   		&IcebergSortFieldProperty{
//   			Direction: jsii.String("direction"),
//   			NullOrder: jsii.String("nullOrder"),
//   			SourceId: jsii.Number(123),
//   			Transform: jsii.String("transform"),
//   		},
//   	},
//   	OrderId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergsortorder.html
//
type CfnTable_IcebergSortOrderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergsortorder.html#cfn-glue-table-icebergsortorder-fields
	//
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergsortorder.html#cfn-glue-table-icebergsortorder-orderid
	//
	OrderId *float64 `field:"required" json:"orderId" yaml:"orderId"`
}

