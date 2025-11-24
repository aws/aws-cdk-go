package mixinsawsbedrock


// Details of the memory configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   memoryConfigurationProperty := &MemoryConfigurationProperty{
//   	EnabledMemoryTypes: []*string{
//   		jsii.String("enabledMemoryTypes"),
//   	},
//   	SessionSummaryConfiguration: &SessionSummaryConfigurationProperty{
//   		MaxRecentSessions: jsii.Number(123),
//   	},
//   	StorageDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-memoryconfiguration.html
//
type CfnAgentPropsMixin_MemoryConfigurationProperty struct {
	// The type of memory that is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-memoryconfiguration.html#cfn-bedrock-agent-memoryconfiguration-enabledmemorytypes
	//
	EnabledMemoryTypes *[]*string `field:"optional" json:"enabledMemoryTypes" yaml:"enabledMemoryTypes"`
	// Contains the configuration for SESSION_SUMMARY memory type enabled for the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-memoryconfiguration.html#cfn-bedrock-agent-memoryconfiguration-sessionsummaryconfiguration
	//
	SessionSummaryConfiguration interface{} `field:"optional" json:"sessionSummaryConfiguration" yaml:"sessionSummaryConfiguration"`
	// The number of days the agent is configured to retain the conversational context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-memoryconfiguration.html#cfn-bedrock-agent-memoryconfiguration-storagedays
	//
	StorageDays *float64 `field:"optional" json:"storageDays" yaml:"storageDays"`
}

