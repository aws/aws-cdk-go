package awsquicksight


// A FieldFolder element is a folder that contains fields and nested subfolders.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldFolderProperty := &fieldFolderProperty{
//   	columns: []*string{
//   		jsii.String("columns"),
//   	},
//   	description: jsii.String("description"),
//   }
//
type CfnDataSet_FieldFolderProperty struct {
	// A folder has a list of columns.
	//
	// A column can only be in one folder.
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// The description for a field folder.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

