package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   semanticChunkingConfigurationProperty := &SemanticChunkingConfigurationProperty{
//   	BreakpointPercentileThreshold: jsii.Number(123),
//   	BufferSize: jsii.Number(123),
//   	MaxTokens: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-semanticchunkingconfiguration.html
//
type CfnKnowledgeBase_SemanticChunkingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-semanticchunkingconfiguration.html#cfn-wisdom-knowledgebase-semanticchunkingconfiguration-breakpointpercentilethreshold
	//
	BreakpointPercentileThreshold *float64 `field:"required" json:"breakpointPercentileThreshold" yaml:"breakpointPercentileThreshold"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-semanticchunkingconfiguration.html#cfn-wisdom-knowledgebase-semanticchunkingconfiguration-buffersize
	//
	BufferSize *float64 `field:"required" json:"bufferSize" yaml:"bufferSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-semanticchunkingconfiguration.html#cfn-wisdom-knowledgebase-semanticchunkingconfiguration-maxtokens
	//
	MaxTokens *float64 `field:"required" json:"maxTokens" yaml:"maxTokens"`
}

