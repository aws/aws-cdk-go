package awsmediaconnect


// The configuration settings for a router input using the RTMP (Real-Time Messaging Protocol) protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   rtmpPushRouterInputConfigurationProperty := &RtmpPushRouterInputConfigurationProperty{
//   	NamingConfiguration: &RtmpPushNamingConfigurationProperty{
//   		Automatic: automatic,
//   		Manual: &ManualRtmpPushNamingConfigurationProperty{
//   			ApplicationName: jsii.String("applicationName"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   	},
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	NamingType: jsii.String("namingType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtmppushrouterinputconfiguration.html
//
type CfnRouterInput_RtmpPushRouterInputConfigurationProperty struct {
	// The configuration settings for the RTMP (Real-Time Messaging Protocol) application and stream names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtmppushrouterinputconfiguration.html#cfn-mediaconnect-routerinput-rtmppushrouterinputconfiguration-namingconfiguration
	//
	NamingConfiguration interface{} `field:"required" json:"namingConfiguration" yaml:"namingConfiguration"`
	// The port number for the RTMP protocol.
	//
	// The only valid value is 1935.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtmppushrouterinputconfiguration.html#cfn-mediaconnect-routerinput-rtmppushrouterinputconfiguration-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtmppushrouterinputconfiguration.html#cfn-mediaconnect-routerinput-rtmppushrouterinputconfiguration-namingtype
	//
	NamingType *string `field:"optional" json:"namingType" yaml:"namingType"`
}

