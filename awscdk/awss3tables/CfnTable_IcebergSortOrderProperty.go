package awss3tables


// Sort order specification for an Iceberg table.
//
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
//
//   	// the properties below are optional
//   	OrderId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergsortorder.html
//
type CfnTable_IcebergSortOrderProperty struct {
	// List of sort fields for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergsortorder.html#cfn-s3tables-table-icebergsortorder-fields
	//
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The sort order ID (defaults to 1 if not specified, 0 is reserved for unsorted).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergsortorder.html#cfn-s3tables-table-icebergsortorder-orderid
	//
	OrderId *float64 `field:"optional" json:"orderId" yaml:"orderId"`
}

