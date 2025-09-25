package awsmediapackagev2


// For use with DVB-DASH profiles only.
//
// The settings for error reporting from the playback device that you want AWS Elemental MediaPackage to pass through to the manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashDvbMetricsReportingProperty := &DashDvbMetricsReportingProperty{
//   	ReportingUrl: jsii.String("reportingUrl"),
//
//   	// the properties below are optional
//   	Probability: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashdvbmetricsreporting.html
//
type CfnOriginEndpoint_DashDvbMetricsReportingProperty struct {
	// The URL where playback devices send error reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashdvbmetricsreporting.html#cfn-mediapackagev2-originendpoint-dashdvbmetricsreporting-reportingurl
	//
	ReportingUrl *string `field:"required" json:"reportingUrl" yaml:"reportingUrl"`
	// The number of playback devices per 1000 that will send error reports to the reporting URL.
	//
	// This represents the probability that a playback device will be a reporting player for this session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashdvbmetricsreporting.html#cfn-mediapackagev2-originendpoint-dashdvbmetricsreporting-probability
	//
	Probability *float64 `field:"optional" json:"probability" yaml:"probability"`
}

