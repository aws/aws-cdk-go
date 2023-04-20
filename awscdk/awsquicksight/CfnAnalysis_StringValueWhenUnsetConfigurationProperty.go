package awsquicksight


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
	// `CfnAnalysis.StringValueWhenUnsetConfigurationProperty.CustomValue`.
	CustomValue *string `field:"optional" json:"customValue" yaml:"customValue"`
	// `CfnAnalysis.StringValueWhenUnsetConfigurationProperty.ValueWhenUnsetOption`.
	ValueWhenUnsetOption *string `field:"optional" json:"valueWhenUnsetOption" yaml:"valueWhenUnsetOption"`
}

