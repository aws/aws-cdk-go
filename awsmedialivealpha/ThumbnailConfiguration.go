package awsmedialivealpha


// Thumbnail configuration for the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var thumbnailState ThumbnailState
//
//   thumbnailConfiguration := &ThumbnailConfiguration{
//   	State: thumbnailState,
//   }
//
// Experimental.
type ThumbnailConfiguration struct {
	// Whether to enable thumbnail generation.
	// Default: ThumbnailState.AUTO
	//
	// Experimental.
	State ThumbnailState `field:"optional" json:"state" yaml:"state"`
}

