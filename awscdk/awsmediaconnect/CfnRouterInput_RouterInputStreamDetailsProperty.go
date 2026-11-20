package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var destinationEndpoint interface{}
//   var mediaConnectFlow interface{}
//   var mediaLiveChannel interface{}
//
//   routerInputStreamDetailsProperty := &RouterInputStreamDetailsProperty{
//   	DestinationEndpoint: destinationEndpoint,
//   	Failover: &FailoverRouterInputStreamDetailsProperty{
//   		SourceIndexOneStreamDetails: &FailoverRouterInputIndexedStreamDetailsProperty{
//   			SourceIndex: jsii.Number(123),
//
//   			// the properties below are optional
//   			RtmpPush: &RtmpPushStreamDetailsProperty{
//   				ApplicationName: jsii.String("applicationName"),
//   				StreamName: jsii.String("streamName"),
//   			},
//   		},
//   		SourceIndexZeroStreamDetails: &FailoverRouterInputIndexedStreamDetailsProperty{
//   			SourceIndex: jsii.Number(123),
//
//   			// the properties below are optional
//   			RtmpPush: &RtmpPushStreamDetailsProperty{
//   				ApplicationName: jsii.String("applicationName"),
//   				StreamName: jsii.String("streamName"),
//   			},
//   		},
//   	},
//   	MediaConnectFlow: mediaConnectFlow,
//   	MediaLiveChannel: mediaLiveChannel,
//   	Merge: &MergeRouterInputStreamDetailsProperty{
//   		SourceIndexOneStreamDetails: &MergeRouterInputIndexedStreamDetailsProperty{
//   			SourceIndex: jsii.Number(123),
//   		},
//   		SourceIndexZeroStreamDetails: &MergeRouterInputIndexedStreamDetailsProperty{
//   			SourceIndex: jsii.Number(123),
//   		},
//   	},
//   	Standard: &StandardRouterInputStreamDetailsProperty{
//   		RtmpPush: &RtmpPushStreamDetailsProperty{
//   			ApplicationName: jsii.String("applicationName"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputstreamdetails.html
//
type CfnRouterInput_RouterInputStreamDetailsProperty struct {
	// Configuration details for a router input stream that is represented by a destination endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputstreamdetails.html#cfn-mediaconnect-routerinput-routerinputstreamdetails-destinationendpoint
	//
	DestinationEndpoint interface{} `field:"optional" json:"destinationEndpoint" yaml:"destinationEndpoint"`
	// Configuration details for a failover router input that can automatically switch between two sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputstreamdetails.html#cfn-mediaconnect-routerinput-routerinputstreamdetails-failover
	//
	Failover interface{} `field:"optional" json:"failover" yaml:"failover"`
	// Configuration details for a MediaConnect flow when used as a router input source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputstreamdetails.html#cfn-mediaconnect-routerinput-routerinputstreamdetails-mediaconnectflow
	//
	MediaConnectFlow interface{} `field:"optional" json:"mediaConnectFlow" yaml:"mediaConnectFlow"`
	// Configuration details for a MediaLive channel when used as a router input source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputstreamdetails.html#cfn-mediaconnect-routerinput-routerinputstreamdetails-medialivechannel
	//
	MediaLiveChannel interface{} `field:"optional" json:"mediaLiveChannel" yaml:"mediaLiveChannel"`
	// Configuration details for a merge router input that combines two input sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputstreamdetails.html#cfn-mediaconnect-routerinput-routerinputstreamdetails-merge
	//
	Merge interface{} `field:"optional" json:"merge" yaml:"merge"`
	// Configuration details for a standard router input stream type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputstreamdetails.html#cfn-mediaconnect-routerinput-routerinputstreamdetails-standard
	//
	Standard interface{} `field:"optional" json:"standard" yaml:"standard"`
}

