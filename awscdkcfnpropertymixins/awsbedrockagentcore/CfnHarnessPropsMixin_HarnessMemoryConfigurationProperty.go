package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessMemoryConfigurationProperty := &HarnessMemoryConfigurationProperty{
//   	AgentCoreMemoryConfiguration: &HarnessAgentCoreMemoryConfigurationProperty{
//   		ActorId: jsii.String("actorId"),
//   		Arn: jsii.String("arn"),
//   		MessagesCount: jsii.Number(123),
//   		RetrievalConfig: map[string]interface{}{
//   			"retrievalConfigKey": &HarnessAgentCoreMemoryRetrievalConfigProperty{
//   				"relevanceScore": jsii.Number(123),
//   				"strategyId": jsii.String("strategyId"),
//   				"topK": jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmemoryconfiguration.html
//
type CfnHarnessPropsMixin_HarnessMemoryConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmemoryconfiguration.html#cfn-bedrockagentcore-harness-harnessmemoryconfiguration-agentcorememoryconfiguration
	//
	AgentCoreMemoryConfiguration interface{} `field:"optional" json:"agentCoreMemoryConfiguration" yaml:"agentCoreMemoryConfiguration"`
}

