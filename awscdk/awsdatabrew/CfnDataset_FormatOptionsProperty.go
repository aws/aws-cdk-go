package awsdatabrew


// Represents a set of options that define the structure of either comma-separated value (CSV), Excel, or JSON input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formatOptionsProperty := &FormatOptionsProperty{
//   	Csv: &CsvOptionsProperty{
//   		Delimiter: jsii.String("delimiter"),
//   		HeaderRow: jsii.Boolean(false),
//   	},
//   	Excel: &ExcelOptionsProperty{
//   		HeaderRow: jsii.Boolean(false),
//   		SheetIndexes: []interface{}{
//   			jsii.Number(123),
//   		},
//   		SheetNames: []*string{
//   			jsii.String("sheetNames"),
//   		},
//   	},
//   	Json: &JsonOptionsProperty{
//   		MultiLine: jsii.Boolean(false),
//   	},
//   }
//
type CfnDataset_FormatOptionsProperty struct {
	// Options that define how CSV input is to be interpreted by DataBrew.
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
	// Options that define how Excel input is to be interpreted by DataBrew.
	Excel interface{} `field:"optional" json:"excel" yaml:"excel"`
	// Options that define how JSON input is to be interpreted by DataBrew.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
}

