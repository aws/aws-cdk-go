package awsimagebuilderalpha


// Specifies what the step should do in case of failure.
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
type ComponentOnFailure string

const (
	// Fails the step and document execution.
	// Experimental.
	ComponentOnFailure_ABORT ComponentOnFailure = "ABORT"
	// Fails the step and proceeds to execute the next step in the document.
	// Experimental.
	ComponentOnFailure_CONTINUE ComponentOnFailure = "CONTINUE"
	// Ignores the step and proceeds to execute the next step in the document.
	// Experimental.
	ComponentOnFailure_IGNORE ComponentOnFailure = "IGNORE"
)

