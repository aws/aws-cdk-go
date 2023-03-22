package awsdynamodb


// The format options for the data that was imported into the target table.
//
// There is one value, CsvOption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputFormatOptionsProperty := &inputFormatOptionsProperty{
//   	csv: &csvProperty{
//   		delimiter: jsii.String("delimiter"),
//   		headerList: []*string{
//   			jsii.String("headerList"),
//   		},
//   	},
//   }
//
type CfnTable_InputFormatOptionsProperty struct {
	// The options for imported source files in CSV format.
	//
	// The values are Delimiter and HeaderList.
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
}

