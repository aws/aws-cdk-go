package awsmediaconnectalpha


// Options for adding an output to a flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   var outputConfiguration OutputConfiguration
//
//   addFlowOutputOptions := &AddFlowOutputOptions{
//   	Output: outputConfiguration,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	FlowOutputName: jsii.String("flowOutputName"),
//   	OutputStatus: mediaconnect_alpha.State_ENABLED,
//   }
//
// Experimental.
type AddFlowOutputOptions struct {
	// The output configuration.
	// Experimental.
	Output OutputConfiguration `field:"required" json:"output" yaml:"output"`
	// A description of the output.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the flow output.
	// Default: - auto-generated.
	//
	// Experimental.
	FlowOutputName *string `field:"optional" json:"flowOutputName" yaml:"flowOutputName"`
	// Whether the output is enabled.
	// Default: State.ENABLED
	//
	// Experimental.
	OutputStatus State `field:"optional" json:"outputStatus" yaml:"outputStatus"`
}

