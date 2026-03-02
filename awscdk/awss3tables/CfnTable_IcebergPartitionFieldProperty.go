package awss3tables


// A partition field specification for an Iceberg table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergPartitionFieldProperty := &IcebergPartitionFieldProperty{
//   	Name: jsii.String("name"),
//   	SourceId: jsii.Number(123),
//   	Transform: jsii.String("transform"),
//
//   	// the properties below are optional
//   	FieldId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergpartitionfield.html
//
type CfnTable_IcebergPartitionFieldProperty struct {
	// The name of the partition field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergpartitionfield.html#cfn-s3tables-table-icebergpartitionfield-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source column ID to partition on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergpartitionfield.html#cfn-s3tables-table-icebergpartitionfield-sourceid
	//
	SourceId *float64 `field:"required" json:"sourceId" yaml:"sourceId"`
	// The partition transform function (identity, bucket[N], truncate[N], year, month, day, hour).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergpartitionfield.html#cfn-s3tables-table-icebergpartitionfield-transform
	//
	Transform *string `field:"required" json:"transform" yaml:"transform"`
	// The partition field ID (auto-assigned starting from 1000 if not specified).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergpartitionfield.html#cfn-s3tables-table-icebergpartitionfield-fieldid
	//
	FieldId *float64 `field:"optional" json:"fieldId" yaml:"fieldId"`
}

