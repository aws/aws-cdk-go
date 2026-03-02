package awss3tables


// A sort field specification for an Iceberg table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergSortFieldProperty := &IcebergSortFieldProperty{
//   	Direction: jsii.String("direction"),
//   	NullOrder: jsii.String("nullOrder"),
//   	SourceId: jsii.Number(123),
//   	Transform: jsii.String("transform"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergsortfield.html
//
type CfnTable_IcebergSortFieldProperty struct {
	// Sort direction (asc or desc).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergsortfield.html#cfn-s3tables-table-icebergsortfield-direction
	//
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Null value ordering (nulls-first or nulls-last).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergsortfield.html#cfn-s3tables-table-icebergsortfield-nullorder
	//
	NullOrder *string `field:"required" json:"nullOrder" yaml:"nullOrder"`
	// The source column ID to sort on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergsortfield.html#cfn-s3tables-table-icebergsortfield-sourceid
	//
	SourceId *float64 `field:"required" json:"sourceId" yaml:"sourceId"`
	// The sort transform function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergsortfield.html#cfn-s3tables-table-icebergsortfield-transform
	//
	Transform *string `field:"required" json:"transform" yaml:"transform"`
}

