package awsimagebuilderalpha


// The phase to run in a specific workflow in an image build, which define the steps to execute to customize or test the instance.
//
// Example:
//   phase := &ComponentDocumentPhase{
//   	Name: imagebuilder.ComponentPhaseName_BUILD,
//   	Steps: []ComponentDocumentStep{
//   		&ComponentDocumentStep{
//   			Name: jsii.String("configure-app"),
//   			Action: imagebuilder.ComponentAction_CREATE_FILE,
//   			Inputs: imagebuilder.ComponentStepInputs_FromObject(map[string]interface{}{
//   				"path": jsii.String("/etc/myapp/config.json"),
//   				"content": jsii.String("{\"env\": \"production\"}"),
//   			}),
//   		},
//   	},
//   }
//
// Experimental.
type ComponentDocumentPhase struct {
	// The name of the phase.
	// Experimental.
	Name ComponentPhaseName `field:"required" json:"name" yaml:"name"`
	// The list of steps to execute to modify or test the build/test instance.
	// Experimental.
	Steps *[]*ComponentDocumentStep `field:"required" json:"steps" yaml:"steps"`
}

