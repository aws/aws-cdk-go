package awslogs


// The CSV processor parses comma-separated values (CSV) from the log events into columns.
//
// For more information about this processor including examples, see csv in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvProperty := &CsvProperty{
//   	Columns: []*string{
//   		jsii.String("columns"),
//   	},
//   	Delimiter: awscdk.Aws_logs.DelimiterCharacter_COMMA,
//   	QuoteCharacter: awscdk.*Aws_logs.QuoteCharacter_DOUBLE_QUOTE,
//   	Source: jsii.String("source"),
//   }
//
type CsvProperty struct {
	// An array of names to use for the columns in the transformed log event.
	// Default: - Column names ([column_1, column_2 ...]) are used
	//
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Character used to separate each column in the original comma-separated value log event.
	// Default: DelimiterCharacter.COMMA
	//
	Delimiter DelimiterCharacter `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Character used as a text qualifier for a single column of data.
	// Default: QuoteCharacter.DOUBLE_QUOTE
	//
	QuoteCharacter QuoteCharacter `field:"optional" json:"quoteCharacter" yaml:"quoteCharacter"`
	// The path to the field in the log event that has the comma separated values to be parsed.
	// Default: '@message'.
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

