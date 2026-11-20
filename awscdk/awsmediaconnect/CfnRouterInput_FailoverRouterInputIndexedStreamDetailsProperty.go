package awsmediaconnect


// Configuration details for an indexed stream in a failover router input setup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverRouterInputIndexedStreamDetailsProperty := &FailoverRouterInputIndexedStreamDetailsProperty{
//   	SourceIndex: jsii.Number(123),
//
//   	// the properties below are optional
//   	RtmpPush: &RtmpPushStreamDetailsProperty{
//   		ApplicationName: jsii.String("applicationName"),
//   		StreamName: jsii.String("streamName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputindexedstreamdetails.html
//
type CfnRouterInput_FailoverRouterInputIndexedStreamDetailsProperty struct {
	// The index number (0 or 1) assigned to this source in the failover configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputindexedstreamdetails.html#cfn-mediaconnect-routerinput-failoverrouterinputindexedstreamdetails-sourceindex
	//
	SourceIndex *float64 `field:"required" json:"sourceIndex" yaml:"sourceIndex"`
	// Configuration details for an RTMP (Real-Time Messaging Protocol) router input stream, including the application and stream names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputindexedstreamdetails.html#cfn-mediaconnect-routerinput-failoverrouterinputindexedstreamdetails-rtmppush
	//
	RtmpPush interface{} `field:"optional" json:"rtmpPush" yaml:"rtmpPush"`
}

