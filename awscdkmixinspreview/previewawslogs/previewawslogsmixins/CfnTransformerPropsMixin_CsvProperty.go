package previewawslogsmixins


// The `CSV` processor parses comma-separated values (CSV) from the log events into columns.
//
// For more information about this processor including examples, see [csv](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-csv) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   csvProperty := &CsvProperty{
//   	Columns: []*string{
//   		jsii.String("columns"),
//   	},
//   	Delimiter: jsii.String("delimiter"),
//   	QuoteCharacter: jsii.String("quoteCharacter"),
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-csv.html
//
type CfnTransformerPropsMixin_CsvProperty struct {
	// An array of names to use for the columns in the transformed log event.
	//
	// If you omit this, default column names ( `[column_1, column_2 ...]` ) are used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-csv.html#cfn-logs-transformer-csv-columns
	//
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// The character used to separate each column in the original comma-separated value log event.
	//
	// If you omit this, the processor looks for the comma `,` character as the delimiter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-csv.html#cfn-logs-transformer-csv-delimiter
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The character used used as a text qualifier for a single column of data.
	//
	// If you omit this, the double quotation mark `"` character is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-csv.html#cfn-logs-transformer-csv-quotecharacter
	//
	QuoteCharacter *string `field:"optional" json:"quoteCharacter" yaml:"quoteCharacter"`
	// The path to the field in the log event that has the comma separated values to be parsed.
	//
	// If you omit this value, the whole log message is processed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-csv.html#cfn-logs-transformer-csv-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

