package awsmediaconnect


// The settings for the RTMP (Real-Time Messaging Protocol) source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   rtmpPushSettingsProperty := &RtmpPushSettingsProperty{
//   	NamingConfiguration: &RtmpPushNamingConfigurationProperty{
//   		Automatic: automatic,
//   		Manual: &ManualRtmpPushNamingConfigurationProperty{
//   			ApplicationName: jsii.String("applicationName"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-rtmppushsettings.html
//
type CfnFlowSource_RtmpPushSettingsProperty struct {
	// The configuration settings for the RTMP (Real-Time Messaging Protocol) application and stream names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-rtmppushsettings.html#cfn-mediaconnect-flowsource-rtmppushsettings-namingconfiguration
	//
	NamingConfiguration interface{} `field:"required" json:"namingConfiguration" yaml:"namingConfiguration"`
}

