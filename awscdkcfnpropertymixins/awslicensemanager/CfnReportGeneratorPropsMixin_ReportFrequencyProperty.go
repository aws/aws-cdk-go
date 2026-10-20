package awslicensemanager


// Details about how frequently reports are generated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   reportFrequencyProperty := &ReportFrequencyProperty{
//   	Period: jsii.String("period"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-reportgenerator-reportfrequency.html
//
type CfnReportGeneratorPropsMixin_ReportFrequencyProperty struct {
	// Time period between each report.
	//
	// The period can be daily, weekly, or monthly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-reportgenerator-reportfrequency.html#cfn-licensemanager-reportgenerator-reportfrequency-period
	//
	Period *string `field:"optional" json:"period" yaml:"period"`
	// Number of times within the frequency period that a report is generated.
	//
	// The only supported value is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-reportgenerator-reportfrequency.html#cfn-licensemanager-reportgenerator-reportfrequency-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

