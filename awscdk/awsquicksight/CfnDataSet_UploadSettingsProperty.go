package awsquicksight


// <p>Information about the format for a source file or files.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uploadSettingsProperty := &UploadSettingsProperty{
//   	ContainsHeader: jsii.Boolean(false),
//   	Delimiter: jsii.String("delimiter"),
//   	Format: jsii.String("format"),
//   	StartFromRow: jsii.Number(123),
//   	TextQualifier: jsii.String("textQualifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html
//
type CfnDataSet_UploadSettingsProperty struct {
	// <p>Whether the file has a header row, or the files each have a header row.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html#cfn-quicksight-dataset-uploadsettings-containsheader
	//
	ContainsHeader interface{} `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// <p>The delimiter between values in the file.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html#cfn-quicksight-dataset-uploadsettings-delimiter
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html#cfn-quicksight-dataset-uploadsettings-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// <p>A row number to start reading data from.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html#cfn-quicksight-dataset-uploadsettings-startfromrow
	//
	StartFromRow *float64 `field:"optional" json:"startFromRow" yaml:"startFromRow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html#cfn-quicksight-dataset-uploadsettings-textqualifier
	//
	TextQualifier *string `field:"optional" json:"textQualifier" yaml:"textQualifier"`
}

