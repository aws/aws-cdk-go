package awssecurityhub


// A date range for the date filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateRangeProperty := &DateRangeProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
type CfnAutomationRule_DateRangeProperty struct {
	// A date range unit for the date filter.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// A date range value for the date filter.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

