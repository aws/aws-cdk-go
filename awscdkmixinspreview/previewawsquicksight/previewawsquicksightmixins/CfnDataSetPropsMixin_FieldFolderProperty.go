package previewawsquicksightmixins


// A FieldFolder element is a folder that contains fields and nested subfolders.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fieldFolderProperty := &FieldFolderProperty{
//   	Columns: []*string{
//   		jsii.String("columns"),
//   	},
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-fieldfolder.html
//
type CfnDataSetPropsMixin_FieldFolderProperty struct {
	// A folder has a list of columns.
	//
	// A column can only be in one folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-fieldfolder.html#cfn-quicksight-dataset-fieldfolder-columns
	//
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// The description for a field folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-fieldfolder.html#cfn-quicksight-dataset-fieldfolder-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

