package awsopensearchserverless


// An OpenSearch Serverless index resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   indexProperty := &IndexProperty{
//   	Knn: jsii.Boolean(false),
//   	KnnAlgoParamEfSearch: jsii.Number(123),
//   	RefreshInterval: jsii.String("refreshInterval"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-index.html
//
type CfnIndex_IndexProperty struct {
	// Enable or disable k-nearest neighbor search capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-index.html#cfn-opensearchserverless-index-index-knn
	//
	Knn interface{} `field:"optional" json:"knn" yaml:"knn"`
	// The size of the dynamic list for the nearest neighbors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-index.html#cfn-opensearchserverless-index-index-knnalgoparamefsearch
	//
	KnnAlgoParamEfSearch *float64 `field:"optional" json:"knnAlgoParamEfSearch" yaml:"knnAlgoParamEfSearch"`
	// How often to perform a refresh operation.
	//
	// For example, 1s or 5s.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-index.html#cfn-opensearchserverless-index-index-refreshinterval
	//
	RefreshInterval *string `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
}

