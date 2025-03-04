package awsquicksight


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html#cfn-quicksight-dataset-uploadsettings-containsheader
	//
	ContainsHeader interface{} `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html#cfn-quicksight-dataset-uploadsettings-delimiter
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html#cfn-quicksight-dataset-uploadsettings-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html#cfn-quicksight-dataset-uploadsettings-startfromrow
	//
	StartFromRow *float64 `field:"optional" json:"startFromRow" yaml:"startFromRow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadsettings.html#cfn-quicksight-dataset-uploadsettings-textqualifier
	//
	TextQualifier *string `field:"optional" json:"textQualifier" yaml:"textQualifier"`
}

