package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Component resource.
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
type ComponentProps struct {
	// The component document content that defines the build, validation, or test steps to be executed during the image building process.
	// Experimental.
	Data ComponentData `field:"required" json:"data" yaml:"data"`
	// The operating system platform of the component.
	// Experimental.
	Platform Platform `field:"required" json:"platform" yaml:"platform"`
	// The change description of the component.
	//
	// Describes what change has been made in this version of the component, or
	// what makes this version different from other versions.
	// Default: None.
	//
	// Experimental.
	ChangeDescription *string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// The name of the component.
	// Default: - a name is generated.
	//
	// Experimental.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// The version of the component.
	// Default: 1.0.0
	//
	// Experimental.
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
	// The description of the component.
	// Default: None.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key used to encrypt this component.
	// Default: - an Image Builder owned key will be used to encrypt the component.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The operating system versions supported by the component.
	// Default: None.
	//
	// Experimental.
	SupportedOsVersions *[]OSVersion `field:"optional" json:"supportedOsVersions" yaml:"supportedOsVersions"`
	// The tags to apply to the component.
	// Default: None.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

