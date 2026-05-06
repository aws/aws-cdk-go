package awsbedrockagentcorealpha


// Configuration for CloudWatch Logs data source.
//
// Example:
//   evaluation := agentcore.NewOnlineEvaluationConfig(this, jsii.String("MyEvaluation"), &OnlineEvaluationConfigProps{
//   	OnlineEvaluationConfigName: jsii.String("my_evaluation"),
//   	Evaluators: []EvaluatorReference{
//   		agentcore.EvaluatorReference_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS()),
//   		agentcore.EvaluatorReference_*Builtin(agentcore.BuiltinEvaluator_CORRECTNESS()),
//   	},
//   	DataSource: agentcore.DataSourceConfig_FromCloudWatchLogs(&CloudWatchLogsDataSourceConfig{
//   		LogGroupNames: []*string{
//   			jsii.String("/aws/bedrock-agentcore/my-agent"),
//   		},
//   		ServiceNames: []*string{
//   			jsii.String("my-agent.default"),
//   		},
//   	}),
//   })
//
// Experimental.
type CloudWatchLogsDataSourceConfig struct {
	// The list of CloudWatch log group names to monitor for agent traces.
	// Experimental.
	LogGroupNames *[]*string `field:"required" json:"logGroupNames" yaml:"logGroupNames"`
	// The list of service names to filter traces within the specified log groups. Used to identify relevant agent sessions.
	//
	// For agents hosted on AgentCore Runtime, service name follows the format:
	// `<agent-runtime-name>.<agent-runtime-endpoint-name>`
	// Experimental.
	ServiceNames *[]*string `field:"required" json:"serviceNames" yaml:"serviceNames"`
}

