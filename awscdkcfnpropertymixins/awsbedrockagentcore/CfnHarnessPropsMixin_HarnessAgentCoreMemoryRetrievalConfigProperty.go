package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessAgentCoreMemoryRetrievalConfigProperty := &HarnessAgentCoreMemoryRetrievalConfigProperty{
//   	RelevanceScore: jsii.Number(123),
//   	StrategyId: jsii.String("strategyId"),
//   	TopK: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorememoryretrievalconfig.html
//
type CfnHarnessPropsMixin_HarnessAgentCoreMemoryRetrievalConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorememoryretrievalconfig.html#cfn-bedrockagentcore-harness-harnessagentcorememoryretrievalconfig-relevancescore
	//
	RelevanceScore *float64 `field:"optional" json:"relevanceScore" yaml:"relevanceScore"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorememoryretrievalconfig.html#cfn-bedrockagentcore-harness-harnessagentcorememoryretrievalconfig-strategyid
	//
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorememoryretrievalconfig.html#cfn-bedrockagentcore-harness-harnessagentcorememoryretrievalconfig-topk
	//
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
}

