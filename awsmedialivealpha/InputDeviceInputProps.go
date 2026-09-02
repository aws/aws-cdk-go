package awsmedialivealpha


// Properties for an Elemental Link input device input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   inputDeviceInputProps := &InputDeviceInputProps{
//   	DeviceIds: []*string{
//   		jsii.String("deviceIds"),
//   	},
//   }
//
// Experimental.
type InputDeviceInputProps struct {
	// The IDs of one or two registered Elemental Link devices.
	//
	// Two provides pipeline redundancy.
	// Experimental.
	DeviceIds *[]*string `field:"required" json:"deviceIds" yaml:"deviceIds"`
}

