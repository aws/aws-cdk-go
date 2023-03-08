package awsglue


// A classifier for custom `CSV` content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvClassifierProperty := &csvClassifierProperty{
//   	allowSingleColumn: jsii.Boolean(false),
//   	containsHeader: jsii.String("containsHeader"),
//   	delimiter: jsii.String("delimiter"),
//   	disableValueTrimming: jsii.Boolean(false),
//   	header: []*string{
//   		jsii.String("header"),
//   	},
//   	name: jsii.String("name"),
//   	quoteSymbol: jsii.String("quoteSymbol"),
//   }
//
type CfnClassifier_CsvClassifierProperty struct {
	// Enables the processing of files that contain only one column.
	AllowSingleColumn interface{} `field:"optional" json:"allowSingleColumn" yaml:"allowSingleColumn"`
	// Indicates whether the CSV file contains a header.
	//
	// A value of `UNKNOWN` specifies that the classifier will detect whether the CSV file contains headings.
	//
	// A value of `PRESENT` specifies that the CSV file contains headings.
	//
	// A value of `ABSENT` specifies that the CSV file does not contain headings.
	ContainsHeader *string `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// A custom symbol to denote what separates each column entry in the row.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Specifies not to trim values before identifying the type of column values.
	//
	// The default value is `true` .
	DisableValueTrimming interface{} `field:"optional" json:"disableValueTrimming" yaml:"disableValueTrimming"`
	// A list of strings representing column names.
	Header *[]*string `field:"optional" json:"header" yaml:"header"`
	// The name of the classifier.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A custom symbol to denote what combines content into a single column value.
	//
	// It must be different from the column delimiter.
	QuoteSymbol *string `field:"optional" json:"quoteSymbol" yaml:"quoteSymbol"`
}

