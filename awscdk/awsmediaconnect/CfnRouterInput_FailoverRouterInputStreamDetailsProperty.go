package awsmediaconnect


// Configuration details for a failover router input that can automatically switch between two sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverRouterInputStreamDetailsProperty := &FailoverRouterInputStreamDetailsProperty{
//   	SourceIndexOneStreamDetails: &FailoverRouterInputIndexedStreamDetailsProperty{
//   		SourceIndex: jsii.Number(123),
//
//   		// the properties below are optional
//   		RtmpPush: &RtmpPushStreamDetailsProperty{
//   			ApplicationName: jsii.String("applicationName"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   	},
//   	SourceIndexZeroStreamDetails: &FailoverRouterInputIndexedStreamDetailsProperty{
//   		SourceIndex: jsii.Number(123),
//
//   		// the properties below are optional
//   		RtmpPush: &RtmpPushStreamDetailsProperty{
//   			ApplicationName: jsii.String("applicationName"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputstreamdetails.html
//
type CfnRouterInput_FailoverRouterInputStreamDetailsProperty struct {
	// Configuration details for an indexed stream in a failover router input setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputstreamdetails.html#cfn-mediaconnect-routerinput-failoverrouterinputstreamdetails-sourceindexonestreamdetails
	//
	SourceIndexOneStreamDetails interface{} `field:"required" json:"sourceIndexOneStreamDetails" yaml:"sourceIndexOneStreamDetails"`
	// Configuration details for an indexed stream in a failover router input setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputstreamdetails.html#cfn-mediaconnect-routerinput-failoverrouterinputstreamdetails-sourceindexzerostreamdetails
	//
	SourceIndexZeroStreamDetails interface{} `field:"required" json:"sourceIndexZeroStreamDetails" yaml:"sourceIndexZeroStreamDetails"`
}

