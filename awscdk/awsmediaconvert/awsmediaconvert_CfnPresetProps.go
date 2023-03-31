package awsmediaconvert


// Properties for defining a `CfnPreset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var settingsJson interface{}
//   var tags interface{}
//
//   cfnPresetProps := &cfnPresetProps{
//   	settingsJson: settingsJson,
//
//   	// the properties below are optional
//   	category: jsii.String("category"),
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	tags: tags,
//   }
//
type CfnPresetProps struct {
	// Specify, in JSON format, the transcoding job settings for this output preset.
	//
	// This specification must conform to the AWS Elemental MediaConvert job validation. For information about forming this specification, see the Remarks section later in this topic.
	//
	// For more information about MediaConvert output presets, see [Working with AWS Elemental MediaConvert Output Presets](https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-presets.html) in the ** .
	SettingsJson interface{} `field:"required" json:"settingsJson" yaml:"settingsJson"`
	// The new category for the preset, if you are changing it.
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The new description for the preset, if you are changing it.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the preset that you are modifying.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

