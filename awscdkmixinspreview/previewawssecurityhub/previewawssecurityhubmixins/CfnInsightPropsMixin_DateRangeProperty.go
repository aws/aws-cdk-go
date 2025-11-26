package previewawssecurityhubmixins


// A date range for the date filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dateRangeProperty := &DateRangeProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-daterange.html
//
type CfnInsightPropsMixin_DateRangeProperty struct {
	// A date range unit for the date filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-daterange.html#cfn-securityhub-insight-daterange-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// A date range value for the date filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-daterange.html#cfn-securityhub-insight-daterange-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

