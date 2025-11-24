package mixinsawsquicksight


// Specifies a mapping to override the name of an output column from a transform operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputColumnNameOverrideProperty := &OutputColumnNameOverrideProperty{
//   	OutputColumnName: jsii.String("outputColumnName"),
//   	SourceColumnName: jsii.String("sourceColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-outputcolumnnameoverride.html
//
type CfnDataSetPropsMixin_OutputColumnNameOverrideProperty struct {
	// The new name to assign to the column in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-outputcolumnnameoverride.html#cfn-quicksight-dataset-outputcolumnnameoverride-outputcolumnname
	//
	OutputColumnName *string `field:"optional" json:"outputColumnName" yaml:"outputColumnName"`
	// The original name of the column from the source transform operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-outputcolumnnameoverride.html#cfn-quicksight-dataset-outputcolumnnameoverride-sourcecolumnname
	//
	SourceColumnName *string `field:"optional" json:"sourceColumnName" yaml:"sourceColumnName"`
}

