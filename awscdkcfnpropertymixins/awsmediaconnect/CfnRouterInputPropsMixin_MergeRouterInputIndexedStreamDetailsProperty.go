package awsmediaconnect


// Configuration details for an indexed stream in a merge router input setup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mergeRouterInputIndexedStreamDetailsProperty := &MergeRouterInputIndexedStreamDetailsProperty{
//   	SourceIndex: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputindexedstreamdetails.html
//
type CfnRouterInputPropsMixin_MergeRouterInputIndexedStreamDetailsProperty struct {
	// The index number (0 or 1) assigned to this source in the merge configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputindexedstreamdetails.html#cfn-mediaconnect-routerinput-mergerouterinputindexedstreamdetails-sourceindex
	//
	SourceIndex *float64 `field:"optional" json:"sourceIndex" yaml:"sourceIndex"`
}

