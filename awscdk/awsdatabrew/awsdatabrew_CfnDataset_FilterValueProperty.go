package awsdatabrew


// Represents a single entry in the `ValuesMap` of a `FilterExpression` .
//
// A `FilterValue` associates the name of a substitution variable in an expression to its value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterValueProperty := &filterValueProperty{
//   	value: jsii.String("value"),
//   	valueReference: jsii.String("valueReference"),
//   }
//
type CfnDataset_FilterValueProperty struct {
	// The value to be associated with the substitution variable.
	Value *string `field:"required" json:"value" yaml:"value"`
	// The substitution variable reference.
	ValueReference *string `field:"required" json:"valueReference" yaml:"valueReference"`
}

