package awsquicksight


// Metadata for a column that is used as the input of a transform operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputColumnProperty := &InputColumnProperty{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inputcolumn.html
//
type CfnDataSet_InputColumnProperty struct {
	// The name of this column in the underlying data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inputcolumn.html#cfn-quicksight-dataset-inputcolumn-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inputcolumn.html#cfn-quicksight-dataset-inputcolumn-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

