package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating an OnlineEvaluationConfig.
//
// Example:
//   var customEvaluator Evaluator
//
//
//   evaluation := agentcore.NewOnlineEvaluationConfig(this, jsii.String("MixedEvaluation"), &OnlineEvaluationConfigProps{
//   	OnlineEvaluationConfigName: jsii.String("mixed_evaluation"),
//   	Evaluators: []EvaluatorSelector{
//   		agentcore.EvaluatorSelector_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS()),
//   		agentcore.EvaluatorSelector_*Builtin(agentcore.BuiltinEvaluator_CORRECTNESS()),
//   		agentcore.EvaluatorSelector_Custom(customEvaluator),
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
type OnlineEvaluationConfigProps struct {
	// The name of the online evaluation configuration.
	//
	// Must be unique within your account. Valid characters are a-z, A-Z, 0-9, _ (underscore).
	// Must start with a letter and can be up to 48 characters long.
	OnlineEvaluationConfigName *string `field:"required" json:"onlineEvaluationConfigName" yaml:"onlineEvaluationConfigName"`
	// The description of the online evaluation configuration.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM role that provides permissions for the evaluation to access AWS services.
	//
	// If not provided, a role will be created automatically with the required permissions
	// including cross-region Bedrock model invocation (to support cross-region inference
	// profiles). For strict cost controls or data residency compliance, provide a custom
	// role with region-scoped permissions.
	// Default: - A new role will be created.
	//
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The execution status of the online evaluation configuration.
	//
	// Controls whether the evaluation actively processes agent traces.
	// Default: ExecutionStatus.ENABLED
	//
	ExecutionStatus ExecutionStatus `field:"optional" json:"executionStatus" yaml:"executionStatus"`
	// The list of filters that determine which agent traces should be evaluated.
	// Default: - No filters (evaluate all sampled traces).
	//
	Filters *[]*FilterConfig `field:"optional" json:"filters" yaml:"filters"`
	// The percentage of agent traces to sample for evaluation.
	// Default: 10.
	//
	SamplingPercentage *float64 `field:"optional" json:"samplingPercentage" yaml:"samplingPercentage"`
	// The duration of inactivity after which an agent session is considered complete and ready for evaluation.
	//
	// Must be between 1 minute and 1440 minutes (24 hours).
	// Default: Duration.minutes(15)
	//
	SessionTimeout awscdk.Duration `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// The data source configuration that specifies where to read agent traces from.
	DataSource DataSourceConfig `field:"required" json:"dataSource" yaml:"dataSource"`
	// The list of evaluators to apply during online evaluation.
	//
	// Can include both built-in evaluators and custom evaluators.
	Evaluators *[]EvaluatorSelector `field:"required" json:"evaluators" yaml:"evaluators"`
	// Tags for the online evaluation configuration.
	//
	// A list of key:value pairs of tags to apply to this OnlineEvaluationConfig resource.
	// Default: - No tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

