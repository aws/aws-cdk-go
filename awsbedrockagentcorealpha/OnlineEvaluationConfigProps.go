package awsbedrockagentcorealpha

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
//   	Evaluators: []EvaluatorReference{
//   		agentcore.EvaluatorReference_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS()),
//   		agentcore.EvaluatorReference_*Builtin(agentcore.BuiltinEvaluator_CORRECTNESS()),
//   		agentcore.EvaluatorReference_Custom(customEvaluator),
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
type OnlineEvaluationConfigProps struct {
	// The name of the online evaluation configuration.
	//
	// Must be unique within your account. Valid characters are a-z, A-Z, 0-9, _ (underscore).
	// Must start with a letter and can be up to 48 characters long.
	// Experimental.
	OnlineEvaluationConfigName *string `field:"required" json:"onlineEvaluationConfigName" yaml:"onlineEvaluationConfigName"`
	// The description of the online evaluation configuration.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM role that provides permissions for the evaluation to access AWS services.
	//
	// If not provided, a role will be created automatically with the required permissions
	// including cross-region Bedrock model invocation (to support cross-region inference
	// profiles). For strict cost controls or data residency compliance, provide a custom
	// role with region-scoped permissions.
	// Default: - A new role will be created.
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The execution status of the online evaluation configuration.
	//
	// Controls whether the evaluation actively processes agent traces.
	// Default: ExecutionStatus.ENABLED
	//
	// Experimental.
	ExecutionStatus ExecutionStatus `field:"optional" json:"executionStatus" yaml:"executionStatus"`
	// The list of filters that determine which agent traces should be evaluated.
	// Default: - No filters (evaluate all sampled traces).
	//
	// Experimental.
	Filters *[]*FilterConfig `field:"optional" json:"filters" yaml:"filters"`
	// The percentage of agent traces to sample for evaluation.
	// Default: 10.
	//
	// Experimental.
	SamplingPercentage *float64 `field:"optional" json:"samplingPercentage" yaml:"samplingPercentage"`
	// The duration of inactivity after which an agent session is considered complete and ready for evaluation.
	//
	// Must be between 1 minute and 1440 minutes (24 hours).
	// Default: Duration.minutes(15)
	//
	// Experimental.
	SessionTimeout awscdk.Duration `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// The data source configuration that specifies where to read agent traces from.
	// Experimental.
	DataSource DataSourceConfig `field:"required" json:"dataSource" yaml:"dataSource"`
	// The list of evaluators to apply during online evaluation.
	//
	// Can include both built-in evaluators and custom evaluators.
	// Experimental.
	Evaluators *[]EvaluatorReference `field:"required" json:"evaluators" yaml:"evaluators"`
}

