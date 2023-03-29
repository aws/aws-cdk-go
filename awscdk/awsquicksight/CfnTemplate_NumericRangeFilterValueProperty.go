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
type CfnTemplate_NumericRangeFilterValueProperty struct {
	// `CfnTemplate.NumericRangeFilterValueProperty.Parameter`.
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// `CfnTemplate.NumericRangeFilterValueProperty.StaticValue`.
	StaticValue *float64 `field:"optional" json:"staticValue" yaml:"staticValue"`
}

