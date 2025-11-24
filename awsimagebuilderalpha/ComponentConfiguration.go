package awsimagebuilderalpha


// Configuration details for a component, to include in a recipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   var component Component
//   var componentParameterValue ComponentParameterValue
//
//   componentConfiguration := &ComponentConfiguration{
//   	Component: component,
//
//   	// the properties below are optional
//   	Parameters: map[string]ComponentParameterValue{
//   		"parametersKey": componentParameterValue,
//   	},
//   }
//
// Experimental.
type ComponentConfiguration struct {
	// The component to execute as part of the image build.
	// Experimental.
	Component IComponent `field:"required" json:"component" yaml:"component"`
	// The parameters to use when executing the component.
	// Default: - no parameters. if the component contains parameters, their default values will be used. otherwise, any
	// required parameters that are not included will result in a build failure.
	//
	// Experimental.
	Parameters *map[string]ComponentParameterValue `field:"optional" json:"parameters" yaml:"parameters"`
}

