package awsscn


// The partition specification of the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   partitionSpecProperty := &PartitionSpecProperty{
//   	Fields: []interface{}{
//   		&DataLakeDatasetPartitionFieldProperty{
//   			Name: jsii.String("name"),
//   			Transform: &TransformProperty{
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-partitionspec.html
//
type CfnDatasetPropsMixin_PartitionSpecProperty struct {
	// The partition fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-partitionspec.html#cfn-scn-dataset-partitionspec-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
}

