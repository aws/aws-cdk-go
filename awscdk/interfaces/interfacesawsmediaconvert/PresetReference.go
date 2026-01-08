package interfacesawsmediaconvert


// A reference to a Preset resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   presetReference := &PresetReference{
//   	PresetArn: jsii.String("presetArn"),
//   	PresetName: jsii.String("presetName"),
//   }
//
type PresetReference struct {
	// The ARN of the Preset resource.
	PresetArn *string `field:"required" json:"presetArn" yaml:"presetArn"`
	// The Name of the Preset resource.
	PresetName *string `field:"required" json:"presetName" yaml:"presetName"`
}

