package mixinsawsdatabrew


// Represents a set of options that define how DataBrew will interpret a Microsoft Excel file when creating a dataset from that file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   excelOptionsProperty := &ExcelOptionsProperty{
//   	HeaderRow: jsii.Boolean(false),
//   	SheetIndexes: []interface{}{
//   		jsii.Number(123),
//   	},
//   	SheetNames: []*string{
//   		jsii.String("sheetNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html
//
type CfnDatasetPropsMixin_ExcelOptionsProperty struct {
	// A variable that specifies whether the first row in the file is parsed as the header.
	//
	// If this value is false, column names are auto-generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html#cfn-databrew-dataset-exceloptions-headerrow
	//
	HeaderRow interface{} `field:"optional" json:"headerRow" yaml:"headerRow"`
	// One or more sheet numbers in the Excel file that will be included in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html#cfn-databrew-dataset-exceloptions-sheetindexes
	//
	SheetIndexes interface{} `field:"optional" json:"sheetIndexes" yaml:"sheetIndexes"`
	// One or more named sheets in the Excel file that will be included in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html#cfn-databrew-dataset-exceloptions-sheetnames
	//
	SheetNames *[]*string `field:"optional" json:"sheetNames" yaml:"sheetNames"`
}

