package awsmediapackagev2alpha


// For use with DVB-DASH profiles only.
//
// The settings for font downloads that you want AWS Elemental MediaPackage to pass through to the manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   dashDvbFontDownload := &DashDvbFontDownload{
//   	FontFamily: jsii.String("fontFamily"),
//   	MimeType: jsii.String("mimeType"),
//   	Url: jsii.String("url"),
//   }
//
// Experimental.
type DashDvbFontDownload struct {
	// The fontFamily name for subtitles, as described in EBU-TT-D Subtitling Distribution Format.
	// Default: - No font family specified.
	//
	// Experimental.
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// The mimeType of the resource that's at the font download URL.
	//
	// For information about font MIME types, see the MPEG-DASH Profile for Transport of ISO BMFF Based DVB Services over IP Based Networks document.
	// Default: - No MIME type specified.
	//
	// Experimental.
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
	// The URL for downloading fonts for subtitles.
	// Default: - No font download URL specified.
	//
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

