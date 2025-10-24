package awsbedrockalpha


// Properties for a function in a function schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   functionProps := &FunctionProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Parameters: map[string]FunctionParameterProps{
//   		"parametersKey": &FunctionParameterProps{
//   			"type": bedrock_alpha.ParameterType_STRING,
//
//   			// the properties below are optional
//   			"description": jsii.String("description"),
//   			"required": jsii.Boolean(false),
//   		},
//   	},
//   	RequireConfirmation: bedrock_alpha.RequireConfirmation_ENABLED,
//   }
//
// Experimental.
type FunctionProps struct {
	// Description of the function.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the function.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Parameters for the function as a record of parameter name to parameter properties.
	// Default: {}.
	//
	// Experimental.
	Parameters *map[string]*FunctionParameterProps `field:"optional" json:"parameters" yaml:"parameters"`
	// Whether to require confirmation before executing the function.
	// Default: RequireConfirmation.DISABLED
	//
	// Experimental.
	RequireConfirmation RequireConfirmation `field:"optional" json:"requireConfirmation" yaml:"requireConfirmation"`
}

