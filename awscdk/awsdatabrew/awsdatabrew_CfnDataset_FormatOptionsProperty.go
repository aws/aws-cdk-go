package awsdatabrew


// Represents a set of options that define the structure of either comma-separated value (CSV), Excel, or JSON input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formatOptionsProperty := &formatOptionsProperty{
//   	csv: &csvOptionsProperty{
//   		delimiter: jsii.String("delimiter"),
//   		headerRow: jsii.Boolean(false),
//   	},
//   	excel: &excelOptionsProperty{
//   		headerRow: jsii.Boolean(false),
//   		sheetIndexes: []interface{}{
//   			jsii.Number(123),
//   		},
//   		sheetNames: []*string{
//   			jsii.String("sheetNames"),
//   		},
//   	},
//   	json: &jsonOptionsProperty{
//   		multiLine: jsii.Boolean(false),
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

