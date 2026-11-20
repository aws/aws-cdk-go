package awsmediaconnect


// Configuration details for a merge router input that combines two input sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mergeRouterInputStreamDetailsProperty := &MergeRouterInputStreamDetailsProperty{
//   	SourceIndexOneStreamDetails: &MergeRouterInputIndexedStreamDetailsProperty{
//   		SourceIndex: jsii.Number(123),
//   	},
//   	SourceIndexZeroStreamDetails: &MergeRouterInputIndexedStreamDetailsProperty{
//   		SourceIndex: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputstreamdetails.html
//
type CfnRouterInputPropsMixin_MergeRouterInputStreamDetailsProperty struct {
	// Configuration details for an indexed stream in a merge router input setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputstreamdetails.html#cfn-mediaconnect-routerinput-mergerouterinputstreamdetails-sourceindexonestreamdetails
	//
	SourceIndexOneStreamDetails interface{} `field:"optional" json:"sourceIndexOneStreamDetails" yaml:"sourceIndexOneStreamDetails"`
	// Configuration details for an indexed stream in a merge router input setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputstreamdetails.html#cfn-mediaconnect-routerinput-mergerouterinputstreamdetails-sourceindexzerostreamdetails
	//
	SourceIndexZeroStreamDetails interface{} `field:"optional" json:"sourceIndexZeroStreamDetails" yaml:"sourceIndexZeroStreamDetails"`
}

