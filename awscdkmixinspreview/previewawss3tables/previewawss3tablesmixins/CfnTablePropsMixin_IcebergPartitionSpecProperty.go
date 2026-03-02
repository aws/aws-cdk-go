package previewawss3tablesmixins


// Partition specification for an Iceberg table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   icebergPartitionSpecProperty := &IcebergPartitionSpecProperty{
//   	Fields: []interface{}{
//   		&IcebergPartitionFieldProperty{
//   			FieldId: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			SourceId: jsii.Number(123),
//   			Transform: jsii.String("transform"),
//   		},
//   	},
//   	SpecId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergpartitionspec.html
//
type CfnTablePropsMixin_IcebergPartitionSpecProperty struct {
	// List of partition fields for an Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergpartitionspec.html#cfn-s3tables-table-icebergpartitionspec-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The partition spec ID (defaults to 0 if not specified).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-icebergpartitionspec.html#cfn-s3tables-table-icebergpartitionspec-specid
	//
	SpecId *float64 `field:"optional" json:"specId" yaml:"specId"`
}

