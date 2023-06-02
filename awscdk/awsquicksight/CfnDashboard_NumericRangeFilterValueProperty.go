package awsquicksight


// The value input pf the numeric range filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericRangeFilterValueProperty := &NumericRangeFilterValueProperty{
//   	Parameter: jsii.String("parameter"),
//   	StaticValue: jsii.Number(123),
//   }
//
type CfnDashboard_NumericRangeFilterValueProperty struct {
	// The parameter that is used in the numeric range.
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// The static value of the numeric range filter.
	StaticValue *float64 `field:"optional" json:"staticValue" yaml:"staticValue"`
}

