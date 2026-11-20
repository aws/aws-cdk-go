package awsmediaconnect


// The manual RTMP naming configuration, where you specify the application and stream names.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   manualRtmpPushNamingConfigurationProperty := &ManualRtmpPushNamingConfigurationProperty{
//   	ApplicationName: jsii.String("applicationName"),
//   	StreamName: jsii.String("streamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-manualrtmppushnamingconfiguration.html
//
type CfnRouterInputPropsMixin_ManualRtmpPushNamingConfigurationProperty struct {
	// The name of the RTMP application.
	//
	// Together with the stream name, the application name forms the RTMP URL path, in the pattern rtmp://ingestAddress/applicationName/streamName.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-manualrtmppushnamingconfiguration.html#cfn-mediaconnect-routerinput-manualrtmppushnamingconfiguration-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The name of the RTMP stream.
	//
	// The stream name forms the final segment of the RTMP URL path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-manualrtmppushnamingconfiguration.html#cfn-mediaconnect-routerinput-manualrtmppushnamingconfiguration-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

