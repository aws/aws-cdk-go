package awsquicksight


// A parameter declaration for the `Integer` data type.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerValueWhenUnsetConfigurationProperty := &IntegerValueWhenUnsetConfigurationProperty{
//   	CustomValue: jsii.Number(123),
//   	ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   }
//
type CfnAnalysis_IntegerValueWhenUnsetConfigurationProperty struct {
	// A custom value that's used when the value of a parameter isn't set.
	CustomValue *float64 `field:"optional" json:"customValue" yaml:"customValue"`
	// The built-in options for default values. The value can be one of the following:.
	//
	// - `RECOMMENDED` : The recommended value.
	// - `NULL` : The `NULL` value.
	ValueWhenUnsetOption *string `field:"optional" json:"valueWhenUnsetOption" yaml:"valueWhenUnsetOption"`
}

