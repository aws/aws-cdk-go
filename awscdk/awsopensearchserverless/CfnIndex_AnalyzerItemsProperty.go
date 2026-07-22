package awsopensearchserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analyzerItemsProperty := &AnalyzerItemsProperty{
//   	CharFilter: []*string{
//   		jsii.String("charFilter"),
//   	},
//   	Filter: []*string{
//   		jsii.String("filter"),
//   	},
//   	Tokenizer: jsii.String("tokenizer"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-analyzeritems.html
//
type CfnIndex_AnalyzerItemsProperty struct {
	// Character filters to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-analyzeritems.html#cfn-opensearchserverless-index-analyzeritems-charfilter
	//
	CharFilter *[]*string `field:"optional" json:"charFilter" yaml:"charFilter"`
	// Token filters to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-analyzeritems.html#cfn-opensearchserverless-index-analyzeritems-filter
	//
	Filter *[]*string `field:"optional" json:"filter" yaml:"filter"`
	// The tokenizer to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-analyzeritems.html#cfn-opensearchserverless-index-analyzeritems-tokenizer
	//
	Tokenizer *string `field:"optional" json:"tokenizer" yaml:"tokenizer"`
	// The analyzer type (e.g. custom, standard, simple).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-analyzeritems.html#cfn-opensearchserverless-index-analyzeritems-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

