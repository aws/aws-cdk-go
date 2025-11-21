package awsimagebuilderalpha


// Properties of an EC2 Image Builder Component Document.
//
// Example:
//   component := imagebuilder.NewComponent(this, jsii.String("StructuredComponent"), &ComponentProps{
//   	Platform: imagebuilder.Platform_LINUX,
//   	Data: imagebuilder.ComponentData_FromComponentDocumentJsonObject(&ComponentDocument{
//   		SchemaVersion: imagebuilder.ComponentSchemaVersion_V1_0,
//   		Phases: []ComponentDocumentPhase{
//   			&ComponentDocumentPhase{
//   				Name: imagebuilder.ComponentPhaseName_BUILD,
//   				Steps: []ComponentDocumentStep{
//   					&ComponentDocumentStep{
//   						Name: jsii.String("install-with-timeout"),
//   						Action: imagebuilder.ComponentAction_EXECUTE_BASH,
//   						Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
//   						OnFailure: imagebuilder.ComponentOnFailure_CONTINUE,
//   						Inputs: imagebuilder.ComponentStepInputs_FromObject(map[string]interface{}{
//   							"commands": []interface{}{
//   								jsii.String("./install-script.sh"),
//   							},
//   						}),
//   					},
//   				},
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type ComponentDocument struct {
	// The phases which define the grouping of steps to run in the build and test workflows of the image build.
	// Experimental.
	Phases *[]*ComponentDocumentPhase `field:"required" json:"phases" yaml:"phases"`
	// The schema version of the component.
	// Experimental.
	SchemaVersion ComponentSchemaVersion `field:"required" json:"schemaVersion" yaml:"schemaVersion"`
	// The constants to define in the document.
	// Default: None.
	//
	// Experimental.
	Constants *map[string]ComponentConstantValue `field:"optional" json:"constants" yaml:"constants"`
	// The description of the document.
	// Default: None.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the document.
	// Default: None.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameters to define in the document.
	// Default: None.
	//
	// Experimental.
	Parameters *map[string]*ComponentDocumentParameterDefinition `field:"optional" json:"parameters" yaml:"parameters"`
}

