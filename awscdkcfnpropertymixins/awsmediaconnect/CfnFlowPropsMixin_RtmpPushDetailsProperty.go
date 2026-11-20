package awsmediaconnect


// The read-only details for the RTMP (Real-Time Messaging Protocol) source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   rtmpPushDetailsProperty := &RtmpPushDetailsProperty{
//   	ApplicationName: jsii.String("applicationName"),
//   	StreamName: jsii.String("streamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-rtmppushdetails.html
//
type CfnFlowPropsMixin_RtmpPushDetailsProperty struct {
	// The name of the RTMP application.
	//
	// Together with the stream name, the application name forms the RTMP URL path, in the pattern rtmp://ingestAddress/applicationName/streamName.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-rtmppushdetails.html#cfn-mediaconnect-flow-rtmppushdetails-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The name of the RTMP stream.
	//
	// The stream name forms the final segment of the RTMP URL path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-rtmppushdetails.html#cfn-mediaconnect-flow-rtmppushdetails-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

