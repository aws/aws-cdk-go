package awsmedialive


// The configuration of an audio-only HLS output.
//
// The parent of this entity is HlsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioOnlyHlsSettingsProperty := &AudioOnlyHlsSettingsProperty{
//   	AudioGroupId: jsii.String("audioGroupId"),
//   	AudioOnlyImage: &InputLocationProperty{
//   		PasswordParam: jsii.String("passwordParam"),
//   		Uri: jsii.String("uri"),
//   		Username: jsii.String("username"),
//   	},
//   	AudioTrackType: jsii.String("audioTrackType"),
//   	SegmentType: jsii.String("segmentType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioonlyhlssettings.html
//
type CfnChannel_AudioOnlyHlsSettingsProperty struct {
	// Specifies the group that the audio rendition belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioonlyhlssettings.html#cfn-medialive-channel-audioonlyhlssettings-audiogroupid
	//
	AudioGroupId *string `field:"optional" json:"audioGroupId" yaml:"audioGroupId"`
	// Used with an audio-only stream.
	//
	// It must be a .jpg or .png file. If given, this image is used as the cover art for the audio-only output. Ideally, it should be formatted for an iPhone screen for two reasons. The iPhone does not resize the image; instead, it crops a centered image on the top/bottom and left/right. Additionally, this image file gets saved bit-for-bit into every 10-second segment file, so it increases bandwidth by {image file size} * {segment count} * {user count.}.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioonlyhlssettings.html#cfn-medialive-channel-audioonlyhlssettings-audioonlyimage
	//
	AudioOnlyImage interface{} `field:"optional" json:"audioOnlyImage" yaml:"audioOnlyImage"`
	// Four types of audio-only tracks are supported: Audio-Only Variant Stream The client can play back this audio-only stream instead of video in low-bandwidth scenarios.
	//
	// Represented as an EXT-X-STREAM-INF in the HLS manifest. Alternate Audio, Auto Select, Default Alternate rendition that the client should try to play back by default. Represented as an EXT-X-MEDIA in the HLS manifest with DEFAULT=YES, AUTOSELECT=YES Alternate Audio, Auto Select, Not Default Alternate rendition that the client might try to play back by default. Represented as an EXT-X-MEDIA in the HLS manifest with DEFAULT=NO, AUTOSELECT=YES Alternate Audio, not Auto Select Alternate rendition that the client will not try to play back by default. Represented as an EXT-X-MEDIA in the HLS manifest with DEFAULT=NO, AUTOSELECT=NO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioonlyhlssettings.html#cfn-medialive-channel-audioonlyhlssettings-audiotracktype
	//
	AudioTrackType *string `field:"optional" json:"audioTrackType" yaml:"audioTrackType"`
	// Specifies the segment type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioonlyhlssettings.html#cfn-medialive-channel-audioonlyhlssettings-segmenttype
	//
	SegmentType *string `field:"optional" json:"segmentType" yaml:"segmentType"`
}

