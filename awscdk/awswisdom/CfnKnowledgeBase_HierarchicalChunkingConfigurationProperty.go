package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hierarchicalChunkingConfigurationProperty := &HierarchicalChunkingConfigurationProperty{
//   	LevelConfigurations: []interface{}{
//   		&HierarchicalChunkingLevelConfigurationProperty{
//   			MaxTokens: jsii.Number(123),
//   		},
//   	},
//   	OverlapTokens: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-hierarchicalchunkingconfiguration.html
//
type CfnKnowledgeBase_HierarchicalChunkingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-hierarchicalchunkingconfiguration.html#cfn-wisdom-knowledgebase-hierarchicalchunkingconfiguration-levelconfigurations
	//
	LevelConfigurations interface{} `field:"required" json:"levelConfigurations" yaml:"levelConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-hierarchicalchunkingconfiguration.html#cfn-wisdom-knowledgebase-hierarchicalchunkingconfiguration-overlaptokens
	//
	OverlapTokens *float64 `field:"required" json:"overlapTokens" yaml:"overlapTokens"`
}

