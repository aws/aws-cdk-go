package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   rtmpPushNamingConfigurationProperty := &RtmpPushNamingConfigurationProperty{
//   	Automatic: automatic,
//   	Manual: &ManualRtmpPushNamingConfigurationProperty{
//   		ApplicationName: jsii.String("applicationName"),
//   		StreamName: jsii.String("streamName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-rtmppushnamingconfiguration.html
//
type CfnFlowSource_RtmpPushNamingConfigurationProperty struct {
	// The automatic RTMP naming configuration, where MediaConnect generates the application and stream names.
	//
	// This type does not require any additional settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-rtmppushnamingconfiguration.html#cfn-mediaconnect-flowsource-rtmppushnamingconfiguration-automatic
	//
	Automatic interface{} `field:"optional" json:"automatic" yaml:"automatic"`
	// The manual RTMP naming configuration, where you specify the application and stream names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-rtmppushnamingconfiguration.html#cfn-mediaconnect-flowsource-rtmppushnamingconfiguration-manual
	//
	Manual interface{} `field:"optional" json:"manual" yaml:"manual"`
}

