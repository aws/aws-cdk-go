package previewawsmediatailormixins


// The configuration for time-shifted viewing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeShiftConfigurationProperty := &TimeShiftConfigurationProperty{
//   	MaxTimeDelaySeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-timeshiftconfiguration.html
//
type CfnChannelPropsMixin_TimeShiftConfigurationProperty struct {
	// The maximum time delay for time-shifted viewing.
	//
	// The minimum allowed maximum time delay is 0 seconds, and the maximum allowed maximum time delay is 21600 seconds (6 hours).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-timeshiftconfiguration.html#cfn-mediatailor-channel-timeshiftconfiguration-maxtimedelayseconds
	//
	MaxTimeDelaySeconds *float64 `field:"optional" json:"maxTimeDelaySeconds" yaml:"maxTimeDelaySeconds"`
}

