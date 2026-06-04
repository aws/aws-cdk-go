package awsmediapackagev2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dashAvailabilityStartTimeConfigurationProperty := &DashAvailabilityStartTimeConfigurationProperty{
//   	FixedAvailabilityStartTime: jsii.String("fixedAvailabilityStartTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashavailabilitystarttimeconfiguration.html
//
type CfnOriginEndpointPropsMixin_DashAvailabilityStartTimeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashavailabilitystarttimeconfiguration.html#cfn-mediapackagev2-originendpoint-dashavailabilitystarttimeconfiguration-fixedavailabilitystarttime
	//
	FixedAvailabilityStartTime *string `field:"optional" json:"fixedAvailabilityStartTime" yaml:"fixedAvailabilityStartTime"`
}

