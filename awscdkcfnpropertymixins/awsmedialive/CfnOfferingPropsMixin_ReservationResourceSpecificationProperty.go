package awsmedialive


// Resource configuration (codec, resolution, bitrate, ...).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   reservationResourceSpecificationProperty := &ReservationResourceSpecificationProperty{
//   	ChannelClass: jsii.String("channelClass"),
//   	Codec: jsii.String("codec"),
//   	MaximumBitrate: jsii.String("maximumBitrate"),
//   	MaximumFramerate: jsii.String("maximumFramerate"),
//   	Resolution: jsii.String("resolution"),
//   	ResourceType: jsii.String("resourceType"),
//   	SpecialFeature: jsii.String("specialFeature"),
//   	VideoQuality: jsii.String("videoQuality"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-offering-reservationresourcespecification.html
//
type CfnOfferingPropsMixin_ReservationResourceSpecificationProperty struct {
	// Channel class, e.g. 'STANDARD'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-offering-reservationresourcespecification.html#cfn-medialive-offering-reservationresourcespecification-channelclass
	//
	ChannelClass *string `field:"optional" json:"channelClass" yaml:"channelClass"`
	// Codec, e.g. 'AVC'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-offering-reservationresourcespecification.html#cfn-medialive-offering-reservationresourcespecification-codec
	//
	Codec *string `field:"optional" json:"codec" yaml:"codec"`
	// Maximum bitrate, e.g. 'MAX_20_MBPS'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-offering-reservationresourcespecification.html#cfn-medialive-offering-reservationresourcespecification-maximumbitrate
	//
	MaximumBitrate *string `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// Maximum framerate, e.g. 'MAX_30_FPS' (Outputs only).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-offering-reservationresourcespecification.html#cfn-medialive-offering-reservationresourcespecification-maximumframerate
	//
	MaximumFramerate *string `field:"optional" json:"maximumFramerate" yaml:"maximumFramerate"`
	// Resolution, e.g. 'HD'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-offering-reservationresourcespecification.html#cfn-medialive-offering-reservationresourcespecification-resolution
	//
	Resolution *string `field:"optional" json:"resolution" yaml:"resolution"`
	// Resource type, 'INPUT', 'OUTPUT', 'MULTIPLEX', or 'CHANNEL'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-offering-reservationresourcespecification.html#cfn-medialive-offering-reservationresourcespecification-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Special feature, e.g. 'AUDIO_NORMALIZATION' (Channels only).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-offering-reservationresourcespecification.html#cfn-medialive-offering-reservationresourcespecification-specialfeature
	//
	SpecialFeature *string `field:"optional" json:"specialFeature" yaml:"specialFeature"`
	// Video quality, e.g. 'STANDARD' (Outputs only).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-offering-reservationresourcespecification.html#cfn-medialive-offering-reservationresourcespecification-videoquality
	//
	VideoQuality *string `field:"optional" json:"videoQuality" yaml:"videoQuality"`
}

