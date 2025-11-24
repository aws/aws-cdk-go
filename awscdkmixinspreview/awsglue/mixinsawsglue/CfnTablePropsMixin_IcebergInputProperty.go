package mixinsawsglue


// Specifies an input structure that defines an Apache Iceberg metadata table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   icebergInputProperty := &IcebergInputProperty{
//   	MetadataOperation: jsii.String("metadataOperation"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-iceberginput.html
//
type CfnTablePropsMixin_IcebergInputProperty struct {
	// A required metadata operation.
	//
	// Can only be set to CREATE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-iceberginput.html#cfn-glue-table-iceberginput-metadataoperation
	//
	MetadataOperation *string `field:"optional" json:"metadataOperation" yaml:"metadataOperation"`
	// The table version for the Iceberg table.
	//
	// Defaults to 2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-iceberginput.html#cfn-glue-table-iceberginput-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

