package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergPartitionSpecProperty := &IcebergPartitionSpecProperty{
//   	Fields: []interface{}{
//   		&IcebergPartitionFieldProperty{
//   			Name: jsii.String("name"),
//   			SourceId: jsii.Number(123),
//   			Transform: jsii.String("transform"),
//
//   			// the properties below are optional
//   			FieldId: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	SpecId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionspec.html
//
type CfnTable_IcebergPartitionSpecProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionspec.html#cfn-glue-table-icebergpartitionspec-fields
	//
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergpartitionspec.html#cfn-glue-table-icebergpartitionspec-specid
	//
	SpecId *float64 `field:"optional" json:"specId" yaml:"specId"`
}

