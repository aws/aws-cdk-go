package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergInputProperty := &IcebergInputProperty{
//   	MetadataOperation: jsii.String("metadataOperation"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-iceberginput.html
//
type CfnTable_IcebergInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-iceberginput.html#cfn-glue-table-iceberginput-metadataoperation
	//
	MetadataOperation *string `field:"optional" json:"metadataOperation" yaml:"metadataOperation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-iceberginput.html#cfn-glue-table-iceberginput-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

