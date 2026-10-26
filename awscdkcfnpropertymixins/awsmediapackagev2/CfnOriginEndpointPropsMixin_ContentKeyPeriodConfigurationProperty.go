package awsmediapackagev2


// <p>The configuration that controls the content key period timing information that MediaPackage signals to your DRM key provider.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   contentKeyPeriodConfigurationProperty := &ContentKeyPeriodConfigurationProperty{
//   	ContentKeyPeriodTiming: jsii.String("contentKeyPeriodTiming"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-contentkeyperiodconfiguration.html
//
type CfnOriginEndpointPropsMixin_ContentKeyPeriodConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-contentkeyperiodconfiguration.html#cfn-mediapackagev2-originendpoint-contentkeyperiodconfiguration-contentkeyperiodtiming
	//
	ContentKeyPeriodTiming *string `field:"optional" json:"contentKeyPeriodTiming" yaml:"contentKeyPeriodTiming"`
}

