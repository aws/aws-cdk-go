package awsimagebuilderalpha


// Represents a platform for an EC2 Image Builder image.
//
// Example:
//   component := imagebuilder.NewComponent(this, jsii.String("JsonComponent"), &ComponentProps{
//   	Platform: imagebuilder.Platform_LINUX,
//   	Data: imagebuilder.ComponentData_FromJsonObject(map[string]interface{}{
//   		"schemaVersion": imagebuilder.ComponentSchemaVersion_V1_0,
//   		"phases": []interface{}{
//   			map[string]interface{}{
//   				"name": imagebuilder.ComponentPhaseName_BUILD,
//   				"steps": []map[string]interface{}{
//   					map[string]interface{}{
//   						"name": jsii.String("configure-app"),
//   						"action": imagebuilder.ComponentAction_CREATE_FILE,
//   						"inputs": map[string]*string{
//   							"path": jsii.String("/etc/myapp/config.json"),
//   							"content": jsii.String("{\"env\": \"production\"}"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type Platform string

const (
	// Platform for Linux.
	// Experimental.
	Platform_LINUX Platform = "LINUX"
	// Platform for Windows.
	// Experimental.
	Platform_WINDOWS Platform = "WINDOWS"
	// Platform for macOS.
	// Experimental.
	Platform_MAC_OS Platform = "MAC_OS"
)

