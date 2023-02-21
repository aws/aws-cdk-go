package awsquicksight


// Information about the format for a source file or files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uploadSettingsProperty := &uploadSettingsProperty{
//   	containsHeader: jsii.Boolean(false),
//   	delimiter: jsii.String("delimiter"),
//   	format: jsii.String("format"),
//   	startFromRow: jsii.Number(123),
//   	textQualifier: jsii.String("textQualifier"),
//   }
//
type CfnDataSet_UploadSettingsProperty struct {
	// Whether the file has a header row, or the files each have a header row.
	ContainsHeader interface{} `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// The delimiter between values in the file.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// File format.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// A row number to start reading data from.
	StartFromRow *float64 `field:"optional" json:"startFromRow" yaml:"startFromRow"`
	// Text qualifier.
	TextQualifier *string `field:"optional" json:"textQualifier" yaml:"textQualifier"`
}

