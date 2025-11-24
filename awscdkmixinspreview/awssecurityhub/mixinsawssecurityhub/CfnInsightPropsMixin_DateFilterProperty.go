package mixinsawssecurityhub


// A date filter for querying findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnInsightPropsMixin_DateFilterProperty struct {
	// A date range for the date filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-datefilter.html#cfn-securityhub-insight-datefilter-daterange
	//
	DateRange interface{} `field:"optional" json:"dateRange" yaml:"dateRange"`
	// A timestamp that provides the end date for the date filter.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-datefilter.html#cfn-securityhub-insight-datefilter-end
	//
	End *string `field:"optional" json:"end" yaml:"end"`
	// A timestamp that provides the start date for the date filter.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-datefilter.html#cfn-securityhub-insight-datefilter-start
	//
	Start *string `field:"optional" json:"start" yaml:"start"`
}

