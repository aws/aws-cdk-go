package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var disabled interface{}
//
//   harnessMemoryConfigurationProperty := &HarnessMemoryConfigurationProperty{
//   	AgentCoreMemoryConfiguration: &HarnessAgentCoreMemoryConfigurationProperty{
//   		ActorId: jsii.String("actorId"),
//   		Arn: jsii.String("arn"),
//   		MessagesCount: jsii.Number(123),
//   		RetrievalConfig: map[string]interface{}{
//   			"retrievalConfigKey": &HarnessAgentCoreMemoryRetrievalConfigProperty{
//   				"relevanceScore": jsii.String("relevanceScore"),
//   				"strategyId": jsii.String("strategyId"),
//   				"topK": jsii.String("topK"),
//   			},
//   		},
//   	},
//   	Disabled: disabled,
//   	ManagedMemoryConfiguration: &HarnessManagedMemoryConfigurationProperty{
//   		Arn: jsii.String("arn"),
//   		EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   		EventExpiryDuration: jsii.Number(123),
//   		Strategies: []*string{
//   			jsii.String("strategies"),
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
	// Explicitly opt out of memory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmemoryconfiguration.html#cfn-bedrockagentcore-harness-harnessmemoryconfiguration-disabled
	//
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Configuration for managed memory.
	//
	// The harness creates and manages a memory resource in the customer's account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmemoryconfiguration.html#cfn-bedrockagentcore-harness-harnessmemoryconfiguration-managedmemoryconfiguration
	//
	ManagedMemoryConfiguration interface{} `field:"optional" json:"managedMemoryConfiguration" yaml:"managedMemoryConfiguration"`
}

