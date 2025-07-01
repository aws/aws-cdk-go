package awsglue


// A classifier for custom `CSV` content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvClassifierProperty := &CsvClassifierProperty{
//   	AllowSingleColumn: jsii.Boolean(false),
//   	ContainsCustomDatatype: []*string{
//   		jsii.String("containsCustomDatatype"),
//   	},
//   	ContainsHeader: jsii.String("containsHeader"),
//   	CustomDatatypeConfigured: jsii.Boolean(false),
//   	Delimiter: jsii.String("delimiter"),
//   	DisableValueTrimming: jsii.Boolean(false),
//   	Header: []*string{
//   		jsii.String("header"),
//   	},
//   	Name: jsii.String("name"),
//   	QuoteSymbol: jsii.String("quoteSymbol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html
//
type CfnClassifier_CsvClassifierProperty struct {
	// Enables the processing of files that contain only one column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html#cfn-glue-classifier-csvclassifier-allowsinglecolumn
	//
	AllowSingleColumn interface{} `field:"optional" json:"allowSingleColumn" yaml:"allowSingleColumn"`
	// Indicates whether the CSV file contains custom data types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html#cfn-glue-classifier-csvclassifier-containscustomdatatype
	//
	ContainsCustomDatatype *[]*string `field:"optional" json:"containsCustomDatatype" yaml:"containsCustomDatatype"`
	// Indicates whether the CSV file contains a header.
	//
	// A value of `UNKNOWN` specifies that the classifier will detect whether the CSV file contains headings.
	//
	// A value of `PRESENT` specifies that the CSV file contains headings.
	//
	// A value of `ABSENT` specifies that the CSV file does not contain headings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html#cfn-glue-classifier-csvclassifier-containsheader
	//
	ContainsHeader *string `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// Enables the configuration of custom data types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html#cfn-glue-classifier-csvclassifier-customdatatypeconfigured
	//
	CustomDatatypeConfigured interface{} `field:"optional" json:"customDatatypeConfigured" yaml:"customDatatypeConfigured"`
	// A custom symbol to denote what separates each column entry in the row.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html#cfn-glue-classifier-csvclassifier-delimiter
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Specifies not to trim values before identifying the type of column values.
	//
	// The default value is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html#cfn-glue-classifier-csvclassifier-disablevaluetrimming
	//
	DisableValueTrimming interface{} `field:"optional" json:"disableValueTrimming" yaml:"disableValueTrimming"`
	// A list of strings representing column names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html#cfn-glue-classifier-csvclassifier-header
	//
	Header *[]*string `field:"optional" json:"header" yaml:"header"`
	// The name of the classifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html#cfn-glue-classifier-csvclassifier-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A custom symbol to denote what combines content into a single column value.
	//
	// It must be different from the column delimiter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-csvclassifier.html#cfn-glue-classifier-csvclassifier-quotesymbol
	//
	QuoteSymbol *string `field:"optional" json:"quoteSymbol" yaml:"quoteSymbol"`
}

