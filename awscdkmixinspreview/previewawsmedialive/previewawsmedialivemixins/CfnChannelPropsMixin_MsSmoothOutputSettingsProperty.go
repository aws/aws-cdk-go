package previewawsmedialivemixins


// Configuration of a Microsoft Smooth output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   msSmoothOutputSettingsProperty := &MsSmoothOutputSettingsProperty{
//   	H265PackagingType: jsii.String("h265PackagingType"),
//   	NameModifier: jsii.String("nameModifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothoutputsettings.html
//
type CfnChannelPropsMixin_MsSmoothOutputSettingsProperty struct {
	// Only applicable when this output is referencing an H.265 video description. Specifies whether MP4 segments should be packaged as HEV1 or HVC1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothoutputsettings.html#cfn-medialive-channel-mssmoothoutputsettings-h265packagingtype
	//
	H265PackagingType *string `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// A string that is concatenated to the end of the destination file name.
	//
	// This is required for multiple outputs of the same type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothoutputsettings.html#cfn-medialive-channel-mssmoothoutputsettings-namemodifier
	//
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

