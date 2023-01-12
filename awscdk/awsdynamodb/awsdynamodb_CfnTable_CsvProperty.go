package awsdynamodb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvProperty := &csvProperty{
//   	delimiter: jsii.String("delimiter"),
//   	headerList: []*string{
//   		jsii.String("headerList"),
//   	},
//   }
//
type CfnTable_CsvProperty struct {
	// `CfnTable.CsvProperty.Delimiter`.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// `CfnTable.CsvProperty.HeaderList`.
	HeaderList *[]*string `field:"optional" json:"headerList" yaml:"headerList"`
}

