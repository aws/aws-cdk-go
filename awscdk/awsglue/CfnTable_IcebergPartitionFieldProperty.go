package awsglue


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionfield.html
//
type CfnTable_IcebergPartitionFieldProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionfield.html#cfn-glue-table-icebergpartitionfield-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionfield.html#cfn-glue-table-icebergpartitionfield-sourceid
	//
	SourceId *float64 `field:"required" json:"sourceId" yaml:"sourceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionfield.html#cfn-glue-table-icebergpartitionfield-transform
	//
	Transform *string `field:"required" json:"transform" yaml:"transform"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionfield.html#cfn-glue-table-icebergpartitionfield-fieldid
	//
	FieldId *float64 `field:"optional" json:"fieldId" yaml:"fieldId"`
}

