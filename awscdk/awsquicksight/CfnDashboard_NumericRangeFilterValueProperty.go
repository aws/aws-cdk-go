package awsquicksight


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
	// `CfnDashboard.NumericRangeFilterValueProperty.Parameter`.
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// `CfnDashboard.NumericRangeFilterValueProperty.StaticValue`.
	StaticValue *float64 `field:"optional" json:"staticValue" yaml:"staticValue"`
}

