package previewawsbedrockmixins


// Configuration for SESSION_SUMMARY memory type enabled for the agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sessionSummaryConfigurationProperty := &SessionSummaryConfigurationProperty{
//   	MaxRecentSessions: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-sessionsummaryconfiguration.html
//
type CfnAgentPropsMixin_SessionSummaryConfigurationProperty struct {
	// Maximum number of recent session summaries to include in the agent's prompt context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-sessionsummaryconfiguration.html#cfn-bedrock-agent-sessionsummaryconfiguration-maxrecentsessions
	//
	MaxRecentSessions *float64 `field:"optional" json:"maxRecentSessions" yaml:"maxRecentSessions"`
}

