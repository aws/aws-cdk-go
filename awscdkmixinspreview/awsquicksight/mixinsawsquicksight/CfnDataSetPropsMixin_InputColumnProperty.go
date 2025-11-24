package mixinsawsquicksight


// Metadata for a column that is used as the input of a transform operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputColumnProperty := &InputColumnProperty{
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   	SubType: jsii.String("subType"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inputcolumn.html
//
type CfnDataSetPropsMixin_InputColumnProperty struct {
	// A unique identifier for the input column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inputcolumn.html#cfn-quicksight-dataset-inputcolumn-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of this column in the underlying data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inputcolumn.html#cfn-quicksight-dataset-inputcolumn-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sub data type of the column.
	//
	// Sub types are only available for decimal columns that are part of a SPICE dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inputcolumn.html#cfn-quicksight-dataset-inputcolumn-subtype
	//
	SubType *string `field:"optional" json:"subType" yaml:"subType"`
	// The data type of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inputcolumn.html#cfn-quicksight-dataset-inputcolumn-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

