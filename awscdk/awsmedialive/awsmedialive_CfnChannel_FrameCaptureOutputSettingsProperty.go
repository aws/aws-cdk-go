package awsmedialive


// The frame capture output settings.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureOutputSettingsProperty := &frameCaptureOutputSettingsProperty{
//   	nameModifier: jsii.String("nameModifier"),
//   }
//
type CfnChannel_FrameCaptureOutputSettingsProperty struct {
	// Required if the output group contains more than one output.
	//
	// This modifier forms part of the output file name.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

