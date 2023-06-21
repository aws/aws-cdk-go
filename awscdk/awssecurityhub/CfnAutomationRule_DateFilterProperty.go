package awssecurityhub


// A date filter for querying findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateFilterProperty := &DateFilterProperty{
//   	DateRange: &DateRangeProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   	End: jsii.String("end"),
//   	Start: jsii.String("start"),
//   }
//
type CfnAutomationRule_DateFilterProperty struct {
	// A date range for the date filter.
	DateRange interface{} `field:"optional" json:"dateRange" yaml:"dateRange"`
	// A timestamp that provides the end date for the date filter.
	//
	// A correctly formatted example is `2020-05-21T20:16:34.724Z` . The value cannot contain spaces, and date and time should be separated by `T` . For more information, see [RFC 3339 section 5.6, Internet Date/Time Format](https://docs.aws.amazon.com/https://www.rfc-editor.org/rfc/rfc3339#section-5.6) .
	End *string `field:"optional" json:"end" yaml:"end"`
	// A timestamp that provides the start date for the date filter.
	//
	// A correctly formatted example is `2020-05-21T20:16:34.724Z` . The value cannot contain spaces, and date and time should be separated by `T` . For more information, see [RFC 3339 section 5.6, Internet Date/Time Format](https://docs.aws.amazon.com/https://www.rfc-editor.org/rfc/rfc3339#section-5.6) .
	Start *string `field:"optional" json:"start" yaml:"start"`
}

