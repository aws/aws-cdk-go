package awsmedialive


// Configuration of a Microsoft Smooth output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   msSmoothOutputSettingsProperty := &msSmoothOutputSettingsProperty{
//   	h265PackagingType: jsii.String("h265PackagingType"),
//   	nameModifier: jsii.String("nameModifier"),
//   }
//
type CfnChannel_MsSmoothOutputSettingsProperty struct {
	// Only applicable when this output is referencing an H.265 video description. Specifies whether MP4 segments should be packaged as HEV1 or HVC1.
	H265PackagingType *string `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// A string that is concatenated to the end of the destination file name.
	//
	// This is required for multiple outputs of the same type.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

