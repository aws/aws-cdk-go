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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-datefilter.html
//
type CfnInsight_DateFilterProperty struct {
	// A date range for the date filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-datefilter.html#cfn-securityhub-insight-datefilter-daterange
	//
	DateRange interface{} `field:"optional" json:"dateRange" yaml:"dateRange"`
	// A timestamp that provides the end date for the date filter.
	//
	// This field accepts only the specified formats. Timestamps can end with `Z` or `("+" / "-") time-hour [":" time-minute]` . The time-secfrac after seconds is limited to a maximum of 9 digits. The offset is bounded by +/-18:00. Here are valid timestamp formats that you can send to Security Hub:
	//
	// - `YYYY-MM-DDTHH:MM:SSZ` (for example, `2019-01-31T23:00:00Z` )
	// - `YYYY-MM-DDTHH:MM:SS.mmmmmmmmmZ` (for example, `2019-01-31T23:00:00.123456789Z` )
	// - `YYYY-MM-DDTHH:MM:SS+HH:MM` (for example, `2024-01-04T15:25:10+17:59` )
	// - `YYYY-MM-DDTHH:MM:SS-HHMM` (for example, `2024-01-04T15:25:10-1759` )
	// - `YYYY-MM-DDTHH:MM:SS.mmmmmmmmm+HH:MM` (for example, `2024-01-04T15:25:10.123456789+17:59` )
	//
	// If a finding provider sends a finding to Security Hub that contains a timestamp in nanoseconds, we round it to milliseconds. For example, we round `2024-10-31T23:00:00.123456789Z` to `2024-10-31T23:00:00.123Z` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-datefilter.html#cfn-securityhub-insight-datefilter-end
	//
	End *string `field:"optional" json:"end" yaml:"end"`
	// A timestamp that provides the start date for the date filter.
	//
	// This field accepts only the specified formats. Timestamps can end with `Z` or `("+" / "-") time-hour [":" time-minute]` . The time-secfrac after seconds is limited to a maximum of 9 digits. The offset is bounded by +/-18:00. Here are valid timestamp formats that you can send to Security Hub:
	//
	// - `YYYY-MM-DDTHH:MM:SSZ` (for example, `2019-01-31T23:00:00Z` )
	// - `YYYY-MM-DDTHH:MM:SS.mmmmmmmmmZ` (for example, `2019-01-31T23:00:00.123456789Z` )
	// - `YYYY-MM-DDTHH:MM:SS+HH:MM` (for example, `2024-01-04T15:25:10+17:59` )
	// - `YYYY-MM-DDTHH:MM:SS-HHMM` (for example, `2024-01-04T15:25:10-1759` )
	// - `YYYY-MM-DDTHH:MM:SS.mmmmmmmmm+HH:MM` (for example, `2024-01-04T15:25:10.123456789+17:59` )
	//
	// If a finding provider sends a finding to Security Hub that contains a timestamp in nanoseconds, we round it to milliseconds. For example, we round `2024-10-31T23:00:00.123456789Z` to `2024-10-31T23:00:00.123Z` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-datefilter.html#cfn-securityhub-insight-datefilter-start
	//
	Start *string `field:"optional" json:"start" yaml:"start"`
}

