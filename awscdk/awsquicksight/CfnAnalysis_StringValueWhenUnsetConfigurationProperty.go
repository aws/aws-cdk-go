package awsquicksight


// The configuration that defines the default value of a `String` parameter when a value has not been set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringValueWhenUnsetConfigurationProperty := &StringValueWhenUnsetConfigurationProperty{
//   	CustomValue: jsii.String("customValue"),
//   	ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   }
//
type CfnAnalysis_StringValueWhenUnsetConfigurationProperty struct {
	// A custom value that's used when the value of a parameter isn't set.
	CustomValue *string `field:"optional" json:"customValue" yaml:"customValue"`
	// The built-in options for default values. The value can be one of the following:.
	//
	// - `RECOMMENDED` : The recommended value.
	// - `NULL` : The `NULL` value.
	ValueWhenUnsetOption *string `field:"optional" json:"valueWhenUnsetOption" yaml:"valueWhenUnsetOption"`
}

