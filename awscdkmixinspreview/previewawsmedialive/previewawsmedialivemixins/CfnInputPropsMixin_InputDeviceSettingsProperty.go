package previewawsmedialivemixins


// Settings that apply only if the input is an Elemental Link input.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputDeviceSettingsProperty := &InputDeviceSettingsProperty{
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdevicesettings.html
//
type CfnInputPropsMixin_InputDeviceSettingsProperty struct {
	// The unique ID for the device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdevicesettings.html#cfn-medialive-input-inputdevicesettings-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

