package awsmedialive


// The settings for a blackout slate.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blackoutSlateProperty := &blackoutSlateProperty{
//   	blackoutSlateImage: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	networkEndBlackout: jsii.String("networkEndBlackout"),
//   	networkEndBlackoutImage: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	networkId: jsii.String("networkId"),
//   	state: jsii.String("state"),
//   }
//
type CfnChannel_BlackoutSlateProperty struct {
	// The blackout slate image to be used.
	//
	// Keep empty for solid black. Only .bmp and .png images are supported.
	BlackoutSlateImage interface{} `field:"optional" json:"blackoutSlateImage" yaml:"blackoutSlateImage"`
	// Setting to enabled causes MediaLive to blackout the video, audio, and captions, and raise the "Network Blackout Image" slate when an SCTE104/35 Network End Segmentation Descriptor is encountered.
	//
	// The blackout is lifted when the Network Start Segmentation Descriptor is encountered. The Network End and Network Start descriptors must contain a network ID that matches the value entered in Network ID.
	NetworkEndBlackout *string `field:"optional" json:"networkEndBlackout" yaml:"networkEndBlackout"`
	// The path to the local file to use as the Network End Blackout image.
	//
	// The image is scaled to fill the entire output raster.
	NetworkEndBlackoutImage interface{} `field:"optional" json:"networkEndBlackoutImage" yaml:"networkEndBlackoutImage"`
	// Provides a Network ID that matches EIDR ID format (for example, "10.XXXX/XXXX-XXXX-XXXX-XXXX-XXXX-C").
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// When set to enabled, this causes video, audio, and captions to be blanked when indicated by program metadata.
	State *string `field:"optional" json:"state" yaml:"state"`
}

