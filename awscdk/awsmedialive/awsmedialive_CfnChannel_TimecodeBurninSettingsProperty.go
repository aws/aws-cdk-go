package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timecodeBurninSettingsProperty := &TimecodeBurninSettingsProperty{
//   	FontSize: jsii.String("fontSize"),
//   	Position: jsii.String("position"),
//   	Prefix: jsii.String("prefix"),
//   }
//
type CfnChannel_TimecodeBurninSettingsProperty struct {
	// `CfnChannel.TimecodeBurninSettingsProperty.FontSize`.
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// `CfnChannel.TimecodeBurninSettingsProperty.Position`.
	Position *string `field:"optional" json:"position" yaml:"position"`
	// `CfnChannel.TimecodeBurninSettingsProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

