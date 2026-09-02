package awsmediapackagev2


// <p>The configuration for the DASH <code>availabilityStartTime</code> attribute of the Media Presentation Description (MPD).
//
// Use this configuration to set a custom availability start time for your DASH manifest.</p>
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashAvailabilityStartTimeConfigurationProperty := &DashAvailabilityStartTimeConfigurationProperty{
//   	FixedAvailabilityStartTime: jsii.String("fixedAvailabilityStartTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashavailabilitystarttimeconfiguration.html
//
type CfnOriginEndpoint_DashAvailabilityStartTimeConfigurationProperty struct {
	// <p>The fixed availability start time for the DASH manifest, in ISO 8601 date-time format.
	//
	// The value must have hourly granularity, meaning that the minutes, seconds, and fractional seconds must be zero. The value must be on or after <code>2024-01-01T00:00:00Z</code> and must be at least 14 days before the current time.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashavailabilitystarttimeconfiguration.html#cfn-mediapackagev2-originendpoint-dashavailabilitystarttimeconfiguration-fixedavailabilitystarttime
	//
	FixedAvailabilityStartTime *string `field:"required" json:"fixedAvailabilityStartTime" yaml:"fixedAvailabilityStartTime"`
}

