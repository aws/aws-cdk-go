package awsimagebuilderalpha


// The phases of a component document.
//
// Example:
//   component := imagebuilder.NewComponent(this, jsii.String("EncryptedComponent"), &ComponentProps{
//   	Platform: imagebuilder.Platform_LINUX,
//   	KmsKey: kms.NewKey(this, jsii.String("ComponentKey")),
//   	Data: imagebuilder.ComponentData_FromJsonObject(map[string]interface{}{
//   		"schemaVersion": imagebuilder.ComponentSchemaVersion_V1_0,
//   		"phases": []interface{}{
//   			map[string]interface{}{
//   				"name": imagebuilder.ComponentPhaseName_BUILD,
//   				"steps": []map[string]interface{}{
//   					map[string]interface{}{
//   						"name": jsii.String("secure-setup"),
//   						"action": imagebuilder.ComponentAction_EXECUTE_BASH,
//   						"inputs": map[string][]*string{
//   							"commands": []*string{
//   								jsii.String("echo \"This component data is encrypted with KMS\""),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type ComponentPhaseName string

const (
	// Build phase of the component.
	//
	// This phase is run during the BUILDING phase of the image build.
	// Experimental.
	ComponentPhaseName_BUILD ComponentPhaseName = "BUILD"
	// Test phase of the component, executed directly on the instance which is used to build the container image.
	//
	// This
	// phase is run during the TESTING phase of the image build.
	// Experimental.
	ComponentPhaseName_CONTAINER_HOST_TEST ComponentPhaseName = "CONTAINER_HOST_TEST"
	// Test phase of the component.
	//
	// This phase is run during the TESTING phase of the image build.
	// Experimental.
	ComponentPhaseName_TEST ComponentPhaseName = "TEST"
	// Validate phase of the component.
	//
	// This phase is run during the BUILDING phase of the image build, after the build
	// step of the component is executed.
	// Experimental.
	ComponentPhaseName_VALIDATE ComponentPhaseName = "VALIDATE"
)

