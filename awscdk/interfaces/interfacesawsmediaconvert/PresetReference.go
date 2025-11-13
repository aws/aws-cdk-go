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
//   	PresetId: jsii.String("presetId"),
//   }
//
type PresetReference struct {
	// The ARN of the Preset resource.
	PresetArn *string `field:"required" json:"presetArn" yaml:"presetArn"`
	// The Id of the Preset resource.
	PresetId *string `field:"required" json:"presetId" yaml:"presetId"`
}

