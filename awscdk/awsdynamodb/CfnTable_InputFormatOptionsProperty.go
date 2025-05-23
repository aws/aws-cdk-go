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
//   inputFormatOptionsProperty := &InputFormatOptionsProperty{
//   	Csv: &CsvProperty{
//   		Delimiter: jsii.String("delimiter"),
//   		HeaderList: []*string{
//   			jsii.String("headerList"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-inputformatoptions.html
//
type CfnTable_InputFormatOptionsProperty struct {
	// The options for imported source files in CSV format.
	//
	// The values are Delimiter and HeaderList.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-inputformatoptions.html#cfn-dynamodb-table-inputformatoptions-csv
	//
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
}

