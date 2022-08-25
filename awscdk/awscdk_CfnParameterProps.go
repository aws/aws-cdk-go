// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   awscdk.NewCfnParameter(this, jsii.String("MyParameter"), &cfnParameterProps{
//   	type: jsii.String("Number"),
//   	default: jsii.Number(1337),
//   })
//
type CfnParameterProps struct {
	// A regular expression that represents the patterns to allow for String types.
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// An array containing the list of values allowed for the parameter.
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// A string that explains a constraint when the constraint is violated.
	//
	// For example, without a constraint description, a parameter that has an allowed
	// pattern of [A-Za-z0-9]+ displays the following error message when the user specifies
	// an invalid value:.
	ConstraintDescription *string `field:"optional" json:"constraintDescription" yaml:"constraintDescription"`
	// A value of the appropriate type for the template to use if no value is specified when a stack is created.
	//
	// If you define constraints for the parameter, you must specify
	// a value that adheres to those constraints.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// A string of up to 4000 characters that describes the parameter.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An integer value that determines the largest number of characters you want to allow for String types.
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// A numeric value that determines the largest numeric value you want to allow for Number types.
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// An integer value that determines the smallest number of characters you want to allow for String types.
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// A numeric value that determines the smallest numeric value you want to allow for Number types.
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// Whether to mask the parameter value when anyone makes a call that describes the stack.
	//
	// If you set the value to ``true``, the parameter value is masked with asterisks (``*****``).
	NoEcho *bool `field:"optional" json:"noEcho" yaml:"noEcho"`
	// The data type for the parameter (DataType).
	Type *string `field:"optional" json:"type" yaml:"type"`
}

