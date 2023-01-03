package awsdynamodb


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
	// `CfnTable.InputFormatOptionsProperty.Csv`.
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
}

