package awsmediaconnect


// Configuration details for a failover router input that can automatically switch between two sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   failoverRouterInputStreamDetailsProperty := &FailoverRouterInputStreamDetailsProperty{
//   	SourceIndexOneStreamDetails: &FailoverRouterInputIndexedStreamDetailsProperty{
//   		RtmpPush: &RtmpPushStreamDetailsProperty{
//   			ApplicationName: jsii.String("applicationName"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   		SourceIndex: jsii.Number(123),
//   	},
//   	SourceIndexZeroStreamDetails: &FailoverRouterInputIndexedStreamDetailsProperty{
//   		RtmpPush: &RtmpPushStreamDetailsProperty{
//   			ApplicationName: jsii.String("applicationName"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   		SourceIndex: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputstreamdetails.html
//
type CfnRouterInputPropsMixin_FailoverRouterInputStreamDetailsProperty struct {
	// Configuration details for an indexed stream in a failover router input setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputstreamdetails.html#cfn-mediaconnect-routerinput-failoverrouterinputstreamdetails-sourceindexonestreamdetails
	//
	SourceIndexOneStreamDetails interface{} `field:"optional" json:"sourceIndexOneStreamDetails" yaml:"sourceIndexOneStreamDetails"`
	// Configuration details for an indexed stream in a failover router input setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputstreamdetails.html#cfn-mediaconnect-routerinput-failoverrouterinputstreamdetails-sourceindexzerostreamdetails
	//
	SourceIndexZeroStreamDetails interface{} `field:"optional" json:"sourceIndexZeroStreamDetails" yaml:"sourceIndexZeroStreamDetails"`
}

