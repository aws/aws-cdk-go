package awsopensearchserverless


// Custom analysis configuration including analyzers, tokenizers, and filters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisProperty := &AnalysisProperty{
//   	Analyzer: map[string]interface{}{
//   		"analyzerKey": &AnalyzerItemsProperty{
//   			"charFilter": []*string{
//   				jsii.String("charFilter"),
//   			},
//   			"filter": []*string{
//   				jsii.String("filter"),
//   			},
//   			"tokenizer": jsii.String("tokenizer"),
//   			"type": jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-analysis.html
//
type CfnIndex_AnalysisProperty struct {
	// Custom analyzer definitions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-analysis.html#cfn-opensearchserverless-index-analysis-analyzer
	//
	Analyzer interface{} `field:"optional" json:"analyzer" yaml:"analyzer"`
}

