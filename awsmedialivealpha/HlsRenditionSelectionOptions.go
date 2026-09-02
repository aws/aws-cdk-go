package awsmedialivealpha


// Options for selecting an HLS audio rendition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsRenditionSelectionOptions := &HlsRenditionSelectionOptions{
//   	GroupId: jsii.String("groupId"),
//   	RenditionName: jsii.String("renditionName"),
//   }
//
// Experimental.
type HlsRenditionSelectionOptions struct {
	// The `GROUP-ID` in the `#EXT-X-MEDIA` tag of the target HLS audio rendition.
	// Experimental.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The `NAME` in the `#EXT-X-MEDIA` tag of the target HLS audio rendition.
	// Experimental.
	RenditionName *string `field:"required" json:"renditionName" yaml:"renditionName"`
}

