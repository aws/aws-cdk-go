package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessAgentCoreMemoryConfigurationProperty := &HarnessAgentCoreMemoryConfigurationProperty{
//   	ActorId: jsii.String("actorId"),
//   	Arn: jsii.String("arn"),
//   	MessagesCount: jsii.Number(123),
//   	RetrievalConfig: map[string]interface{}{
//   		"retrievalConfigKey": &HarnessAgentCoreMemoryRetrievalConfigProperty{
//   			"relevanceScore": jsii.Number(123),
//   			"strategyId": jsii.String("strategyId"),
//   			"topK": jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorememoryconfiguration.html
//
type CfnHarnessPropsMixin_HarnessAgentCoreMemoryConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorememoryconfiguration.html#cfn-bedrockagentcore-harness-harnessagentcorememoryconfiguration-actorid
	//
	ActorId *string `field:"optional" json:"actorId" yaml:"actorId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorememoryconfiguration.html#cfn-bedrockagentcore-harness-harnessagentcorememoryconfiguration-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorememoryconfiguration.html#cfn-bedrockagentcore-harness-harnessagentcorememoryconfiguration-messagescount
	//
	MessagesCount *float64 `field:"optional" json:"messagesCount" yaml:"messagesCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorememoryconfiguration.html#cfn-bedrockagentcore-harness-harnessagentcorememoryconfiguration-retrievalconfig
	//
	RetrievalConfig interface{} `field:"optional" json:"retrievalConfig" yaml:"retrievalConfig"`
}

