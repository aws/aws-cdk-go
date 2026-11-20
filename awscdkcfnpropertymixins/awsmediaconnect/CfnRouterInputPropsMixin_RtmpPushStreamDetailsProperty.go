package awsmediaconnect


// Configuration details for an RTMP (Real-Time Messaging Protocol) router input stream, including the application and stream names.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   rtmpPushStreamDetailsProperty := &RtmpPushStreamDetailsProperty{
//   	ApplicationName: jsii.String("applicationName"),
//   	StreamName: jsii.String("streamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtmppushstreamdetails.html
//
type CfnRouterInputPropsMixin_RtmpPushStreamDetailsProperty struct {
	// The name of the RTMP application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtmppushstreamdetails.html#cfn-mediaconnect-routerinput-rtmppushstreamdetails-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The name of the RTMP stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtmppushstreamdetails.html#cfn-mediaconnect-routerinput-rtmppushstreamdetails-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

