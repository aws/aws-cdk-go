package awsmediapackagev2alpha


// The font download and error reporting information that you want MediaPackage to pass through to the manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   dashDvbSettings := &DashDvbSettings{
//   	ErrorMetrics: []DashDvbMetricsReporting{
//   		&DashDvbMetricsReporting{
//   			ReportingUrl: jsii.String("reportingUrl"),
//
//   			// the properties below are optional
//   			Probability: jsii.Number(123),
//   		},
//   	},
//   	FontDownload: &DashDvbFontDownload{
//   		FontFamily: jsii.String("fontFamily"),
//   		MimeType: jsii.String("mimeType"),
//   		Url: jsii.String("url"),
//   	},
//   }
//
// Experimental.
type DashDvbSettings struct {
	// Playback device error reporting settings.
	// Default: - No error metrics configured.
	//
	// Experimental.
	ErrorMetrics *[]*DashDvbMetricsReporting `field:"optional" json:"errorMetrics" yaml:"errorMetrics"`
	// Subtitle font settings.
	// Default: - No font download configured.
	//
	// Experimental.
	FontDownload *DashDvbFontDownload `field:"optional" json:"fontDownload" yaml:"fontDownload"`
}

