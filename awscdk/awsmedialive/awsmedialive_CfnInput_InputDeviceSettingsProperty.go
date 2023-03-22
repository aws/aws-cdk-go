package awsmedialive


// Settings that apply only if the input is an Elemental Link input.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputDeviceSettingsProperty := &inputDeviceSettingsProperty{
//   	id: jsii.String("id"),
//   }
//
type CfnInput_InputDeviceSettingsProperty struct {
	// The unique ID for the device.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

