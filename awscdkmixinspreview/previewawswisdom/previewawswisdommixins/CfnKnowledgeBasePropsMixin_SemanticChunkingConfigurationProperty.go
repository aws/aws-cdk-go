package previewawswisdommixins


// Settings for semantic document chunking for a data source.
//
// Semantic chunking splits a document into smaller documents based on groups of similar content derived from the text with natural language processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   semanticChunkingConfigurationProperty := &SemanticChunkingConfigurationProperty{
//   	BreakpointPercentileThreshold: jsii.Number(123),
//   	BufferSize: jsii.Number(123),
//   	MaxTokens: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-semanticchunkingconfiguration.html
//
type CfnKnowledgeBasePropsMixin_SemanticChunkingConfigurationProperty struct {
	// The dissimilarity threshold for splitting chunks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-semanticchunkingconfiguration.html#cfn-wisdom-knowledgebase-semanticchunkingconfiguration-breakpointpercentilethreshold
	//
	BreakpointPercentileThreshold *float64 `field:"optional" json:"breakpointPercentileThreshold" yaml:"breakpointPercentileThreshold"`
	// The buffer size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-semanticchunkingconfiguration.html#cfn-wisdom-knowledgebase-semanticchunkingconfiguration-buffersize
	//
	BufferSize *float64 `field:"optional" json:"bufferSize" yaml:"bufferSize"`
	// The maximum number of tokens that a chunk can contain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-semanticchunkingconfiguration.html#cfn-wisdom-knowledgebase-semanticchunkingconfiguration-maxtokens
	//
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
}

