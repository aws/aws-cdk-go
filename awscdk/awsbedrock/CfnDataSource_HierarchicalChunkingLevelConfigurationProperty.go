package awsbedrock


// Token settings for a layer in a hierarchical chunking configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hierarchicalChunkingLevelConfigurationProperty := &HierarchicalChunkingLevelConfigurationProperty{
//   	MaxTokens: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-hierarchicalchunkinglevelconfiguration.html
//
type CfnDataSource_HierarchicalChunkingLevelConfigurationProperty struct {
	// The maximum number of tokens that a chunk can contain in this layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-hierarchicalchunkinglevelconfiguration.html#cfn-bedrock-datasource-hierarchicalchunkinglevelconfiguration-maxtokens
	//
	MaxTokens *float64 `field:"required" json:"maxTokens" yaml:"maxTokens"`
}

