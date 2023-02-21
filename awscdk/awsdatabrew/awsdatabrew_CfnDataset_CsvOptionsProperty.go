package awsdatabrew


// Represents a set of options that define how DataBrew will read a comma-separated value (CSV) file when creating a dataset from that file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvOptionsProperty := &csvOptionsProperty{
//   	delimiter: jsii.String("delimiter"),
//   	headerRow: jsii.Boolean(false),
//   }
//
type CfnDataset_CsvOptionsProperty struct {
	// A single character that specifies the delimiter being used in the CSV file.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// A variable that specifies whether the first row in the file is parsed as the header.
	//
	// If this value is false, column names are auto-generated.
	HeaderRow interface{} `field:"optional" json:"headerRow" yaml:"headerRow"`
}

