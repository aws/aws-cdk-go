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
//   audioHlsRenditionSelectionProperty := &audioHlsRenditionSelectionProperty{
//   	groupId: jsii.String("groupId"),
//   	name: jsii.String("name"),
//   }
//
type CfnChannel_AudioHlsRenditionSelectionProperty struct {
	// Specifies the GROUP-ID in the #EXT-X-MEDIA tag of the target HLS audio rendition.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Specifies the NAME in the #EXT-X-MEDIA tag of the target HLS audio rendition.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

