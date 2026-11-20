package awsmediaconnect


// Configuration details for a standard router input stream type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   standardRouterInputStreamDetailsProperty := &StandardRouterInputStreamDetailsProperty{
//   	RtmpPush: &RtmpPushStreamDetailsProperty{
//   		ApplicationName: jsii.String("applicationName"),
//   		StreamName: jsii.String("streamName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-standardrouterinputstreamdetails.html
//
type CfnRouterInputPropsMixin_StandardRouterInputStreamDetailsProperty struct {
	// Configuration details for an RTMP (Real-Time Messaging Protocol) router input stream, including the application and stream names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-standardrouterinputstreamdetails.html#cfn-mediaconnect-routerinput-standardrouterinputstreamdetails-rtmppush
	//
	RtmpPush interface{} `field:"optional" json:"rtmpPush" yaml:"rtmpPush"`
}

