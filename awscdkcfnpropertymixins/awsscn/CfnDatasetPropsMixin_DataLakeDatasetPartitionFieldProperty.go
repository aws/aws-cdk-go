package awsscn


// The partition field details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataLakeDatasetPartitionFieldProperty := &DataLakeDatasetPartitionFieldProperty{
//   	Name: jsii.String("name"),
//   	Transform: &TransformProperty{
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-datalakedatasetpartitionfield.html
//
type CfnDatasetPropsMixin_DataLakeDatasetPartitionFieldProperty struct {
	// The name of the partition field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-datalakedatasetpartitionfield.html#cfn-scn-dataset-datalakedatasetpartitionfield-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The transformation of the partition field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-datalakedatasetpartitionfield.html#cfn-scn-dataset-datalakedatasetpartitionfield-transform
	//
	Transform interface{} `field:"optional" json:"transform" yaml:"transform"`
}

