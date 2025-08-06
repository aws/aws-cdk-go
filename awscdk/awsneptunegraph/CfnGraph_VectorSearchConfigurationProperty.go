package awsneptunegraph


// The vector-search configuration for the graph, which specifies the vector dimension to use in the vector index, if any.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vectorSearchConfigurationProperty := &VectorSearchConfigurationProperty{
//   	VectorSearchDimension: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-vectorsearchconfiguration.html
//
type CfnGraph_VectorSearchConfigurationProperty struct {
	// The number of dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-vectorsearchconfiguration.html#cfn-neptunegraph-graph-vectorsearchconfiguration-vectorsearchdimension
	//
	VectorSearchDimension *float64 `field:"required" json:"vectorSearchDimension" yaml:"vectorSearchDimension"`
}

