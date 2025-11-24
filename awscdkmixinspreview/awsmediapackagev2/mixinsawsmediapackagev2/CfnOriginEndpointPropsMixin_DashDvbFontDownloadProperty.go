package mixinsawsmediapackagev2


// For use with DVB-DASH profiles only.
//
// The settings for font downloads that you want AWS Elemental MediaPackage to pass through to the manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dashDvbFontDownloadProperty := &DashDvbFontDownloadProperty{
//   	FontFamily: jsii.String("fontFamily"),
//   	MimeType: jsii.String("mimeType"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashdvbfontdownload.html
//
type CfnOriginEndpointPropsMixin_DashDvbFontDownloadProperty struct {
	// The `fontFamily` name for subtitles, as described in [EBU-TT-D Subtitling Distribution Format](https://docs.aws.amazon.com/https://tech.ebu.ch/publications/tech3380) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashdvbfontdownload.html#cfn-mediapackagev2-originendpoint-dashdvbfontdownload-fontfamily
	//
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// The `mimeType` of the resource that's at the font download URL.
	//
	// For information about font MIME types, see the [MPEG-DASH Profile for Transport of ISO BMFF Based DVB Services over IP Based Networks](https://docs.aws.amazon.com/https://dvb.org/wp-content/uploads/2021/06/A168r4_MPEG-DASH-Profile-for-Transport-of-ISO-BMFF-Based-DVB-Services_Draft-ts_103-285-v140_November_2021.pdf) document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashdvbfontdownload.html#cfn-mediapackagev2-originendpoint-dashdvbfontdownload-mimetype
	//
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
	// The URL for downloading fonts for subtitles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashdvbfontdownload.html#cfn-mediapackagev2-originendpoint-dashdvbfontdownload-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

