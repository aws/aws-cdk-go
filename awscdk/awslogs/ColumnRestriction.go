package awslogs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnRestriction := &ColumnRestriction{
//   	Comparison: jsii.String("comparison"),
//
//   	// the properties below are optional
//   	NumberValue: jsii.Number(123),
//   	StringValue: jsii.String("stringValue"),
//   }
//
// Experimental.
type ColumnRestriction struct {
	// Comparison operator to use.
	// Experimental.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Number value to compare to.
	//
	// Exactly one of 'stringValue' and 'numberValue' must be set.
	// Experimental.
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// String value to compare to.
	//
	// Exactly one of 'stringValue' and 'numberValue' must be set.
	// Experimental.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

