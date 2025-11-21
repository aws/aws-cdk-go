package awsimagebuilderalpha


// The schema version of the component.
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
type ComponentSchemaVersion string

const (
	// Schema version 1.0 for the component document.
	// Experimental.
	ComponentSchemaVersion_V1_0 ComponentSchemaVersion = "V1_0"
)

