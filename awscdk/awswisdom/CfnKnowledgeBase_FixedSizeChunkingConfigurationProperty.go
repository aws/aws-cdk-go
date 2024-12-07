package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fixedSizeChunkingConfigurationProperty := &FixedSizeChunkingConfigurationProperty{
//   	MaxTokens: jsii.Number(123),
//   	OverlapPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-fixedsizechunkingconfiguration.html
//
type CfnKnowledgeBase_FixedSizeChunkingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-fixedsizechunkingconfiguration.html#cfn-wisdom-knowledgebase-fixedsizechunkingconfiguration-maxtokens
	//
	MaxTokens *float64 `field:"required" json:"maxTokens" yaml:"maxTokens"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-fixedsizechunkingconfiguration.html#cfn-wisdom-knowledgebase-fixedsizechunkingconfiguration-overlappercentage
	//
	OverlapPercentage *float64 `field:"required" json:"overlapPercentage" yaml:"overlapPercentage"`
}

