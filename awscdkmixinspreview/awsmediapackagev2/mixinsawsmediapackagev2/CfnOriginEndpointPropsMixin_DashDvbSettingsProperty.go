package mixinsawsmediapackagev2


// For endpoints that use the DVB-DASH profile only.
//
// The font download and error reporting information that you want MediaPackage to pass through to the manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dashDvbSettingsProperty := &DashDvbSettingsProperty{
//   	ErrorMetrics: []interface{}{
//   		&DashDvbMetricsReportingProperty{
//   			Probability: jsii.Number(123),
//   			ReportingUrl: jsii.String("reportingUrl"),
//   		},
//   	},
//   	FontDownload: &DashDvbFontDownloadProperty{
//   		FontFamily: jsii.String("fontFamily"),
//   		MimeType: jsii.String("mimeType"),
//   		Url: jsii.String("url"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashdvbsettings.html
//
type CfnOriginEndpointPropsMixin_DashDvbSettingsProperty struct {
	// Playback device error reporting settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashdvbsettings.html#cfn-mediapackagev2-originendpoint-dashdvbsettings-errormetrics
	//
	ErrorMetrics interface{} `field:"optional" json:"errorMetrics" yaml:"errorMetrics"`
	// Subtitle font settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashdvbsettings.html#cfn-mediapackagev2-originendpoint-dashdvbsettings-fontdownload
	//
	FontDownload interface{} `field:"optional" json:"fontDownload" yaml:"fontDownload"`
}

