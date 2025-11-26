package previewawsquicksightmixins


// Output column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputColumnProperty := &OutputColumnProperty{
//   	Description: jsii.String("description"),
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   	SubType: jsii.String("subType"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-outputcolumn.html
//
type CfnDataSetPropsMixin_OutputColumnProperty struct {
	// A description for a column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-outputcolumn.html#cfn-quicksight-dataset-outputcolumn-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A unique identifier for the output column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-outputcolumn.html#cfn-quicksight-dataset-outputcolumn-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The display name of the column..
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-outputcolumn.html#cfn-quicksight-dataset-outputcolumn-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sub data type of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-outputcolumn.html#cfn-quicksight-dataset-outputcolumn-subtype
	//
	SubType *string `field:"optional" json:"subType" yaml:"subType"`
	// The data type of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-outputcolumn.html#cfn-quicksight-dataset-outputcolumn-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

