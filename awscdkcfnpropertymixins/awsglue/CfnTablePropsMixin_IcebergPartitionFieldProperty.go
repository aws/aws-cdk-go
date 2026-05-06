package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   icebergPartitionFieldProperty := &IcebergPartitionFieldProperty{
//   	FieldId: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	SourceId: jsii.Number(123),
//   	Transform: jsii.String("transform"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionfield.html
//
type CfnTablePropsMixin_IcebergPartitionFieldProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionfield.html#cfn-glue-table-icebergpartitionfield-fieldid
	//
	FieldId *float64 `field:"optional" json:"fieldId" yaml:"fieldId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionfield.html#cfn-glue-table-icebergpartitionfield-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionfield.html#cfn-glue-table-icebergpartitionfield-sourceid
	//
	SourceId *float64 `field:"optional" json:"sourceId" yaml:"sourceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionfield.html#cfn-glue-table-icebergpartitionfield-transform
	//
	Transform *string `field:"optional" json:"transform" yaml:"transform"`
}

