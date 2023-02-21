package awsdatabrew


// A key-value pair to associate an expression's substitution variable names with their values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   substitutionValueProperty := &substitutionValueProperty{
//   	value: jsii.String("value"),
//   	valueReference: jsii.String("valueReference"),
//   }
//
type CfnRuleset_SubstitutionValueProperty struct {
	// Value or column name.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Variable name.
	ValueReference *string `field:"required" json:"valueReference" yaml:"valueReference"`
}

