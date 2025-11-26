package awsimagebuilderalpha


// The schema version of the component.
//
// Example:
//   customComponent := imagebuilder.NewComponent(this, jsii.String("MyComponent"), &ComponentProps{
//   	Platform: imagebuilder.Platform_LINUX,
//   	Data: imagebuilder.ComponentData_FromJsonObject(map[string]interface{}{
//   		"schemaVersion": imagebuilder.ComponentSchemaVersion_V1_0,
//   		"phases": []interface{}{
//   			map[string]interface{}{
//   				"name": imagebuilder.ComponentPhaseName_BUILD,
//   				"steps": []map[string]interface{}{
//   					map[string]interface{}{
//   						"name": jsii.String("install-app"),
//   						"action": imagebuilder.ComponentAction_EXECUTE_BASH,
//   						"inputs": map[string][]*string{
//   							"commands": []*string{
//   								jsii.String("yum install -y my-application"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	}),
//   })
//
//   imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("ComponentImageRecipe"), &ImageRecipeProps{
//   	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
//   	Components: []ComponentConfiguration{
//   		&ComponentConfiguration{
//   			Component: customComponent,
//   		},
//   	},
//   })
//
// Experimental.
type ComponentSchemaVersion string

const (
	// Schema version 1.0 for the component document.
	// Experimental.
	ComponentSchemaVersion_V1_0 ComponentSchemaVersion = "V1_0"
)

