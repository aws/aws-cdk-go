package awscdk


// Example:
//   awscdk.NewCfnParameter(this, jsii.String("MyParameter"), &CfnParameterProps{
//   	Type: jsii.String("Number"),
//   	Default: jsii.Number(1337),
//   })
//
type CfnParameterProps struct {
	// A regular expression that represents the patterns to allow for String types.
	// Default: - No constraints on patterns allowed for parameter.
	//
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// An array containing the list of values allowed for the parameter.
	// Default: - No constraints on values allowed for parameter.
	//
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// A string that explains a constraint when the constraint is violated.
	//
	// For example, without a constraint description, a parameter that has an allowed
	// pattern of [A-Za-z0-9]+ displays the following error message when the user specifies
	// an invalid value:.
	// Default: - No description with customized error message when user specifies invalid values.
	//
	ConstraintDescription *string `field:"optional" json:"constraintDescription" yaml:"constraintDescription"`
	// A value of the appropriate type for the template to use if no value is specified when a stack is created.
	//
	// If you define constraints for the parameter, you must specify
	// a value that adheres to those constraints.
	// Default: - No default value for parameter.
	//
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// A string of up to 4000 characters that describes the parameter.
	// Default: - No description for the parameter.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An integer value that determines the largest number of characters you want to allow for String types.
	// Default: - None.
	//
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// A numeric value that determines the largest numeric value you want to allow for Number types.
	// Default: - None.
	//
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// An integer value that determines the smallest number of characters you want to allow for String types.
	// Default: - None.
	//
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// A numeric value that determines the smallest numeric value you want to allow for Number types.
	// Default: - None.
	//
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// Whether to mask the parameter value when anyone makes a call that describes the stack.
	//
	// If you set the value to ``true``, the parameter value is masked with asterisks (``*****``).
	// Default: - Parameter values are not masked.
	//
	NoEcho *bool `field:"optional" json:"noEcho" yaml:"noEcho"`
	// The data type for the parameter (DataType).
	// Default: String.
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

