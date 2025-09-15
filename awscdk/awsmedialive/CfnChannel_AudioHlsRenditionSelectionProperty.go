package awsmedialive


// Selector for HLS audio rendition.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioHlsRenditionSelectionProperty := &AudioHlsRenditionSelectionProperty{
//   	GroupId: jsii.String("groupId"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiohlsrenditionselection.html
//
type CfnChannel_AudioHlsRenditionSelectionProperty struct {
	// Specifies the GROUP-ID in the #EXT-X-MEDIA tag of the target HLS audio rendition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiohlsrenditionselection.html#cfn-medialive-channel-audiohlsrenditionselection-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Specifies the NAME in the #EXT-X-MEDIA tag of the target HLS audio rendition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiohlsrenditionselection.html#cfn-medialive-channel-audiohlsrenditionselection-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

