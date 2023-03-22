package awsmedialive


// Settings to configure color space settings in the incoming video.
//
// The parent of this entity is VideoSelector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoSelectorColorSpaceSettingsProperty := &videoSelectorColorSpaceSettingsProperty{
//   	hdr10Settings: &hdr10SettingsProperty{
//   		maxCll: jsii.Number(123),
//   		maxFall: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_VideoSelectorColorSpaceSettingsProperty struct {
	// Settings to configure color space settings in the incoming video.
	Hdr10Settings interface{} `field:"optional" json:"hdr10Settings" yaml:"hdr10Settings"`
}

