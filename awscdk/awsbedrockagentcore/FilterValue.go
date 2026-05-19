package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A typed filter value for online evaluation filtering.
//
// Use the static factory methods to create filter values:
// - `FilterValue.string()` for string comparisons
// - `FilterValue.number()` for numeric comparisons
// - `FilterValue.boolean()` for boolean comparisons
//
// Example:
//   evaluation := agentcore.NewOnlineEvaluationConfig(this, jsii.String("FilteredEval"), &OnlineEvaluationConfigProps{
//   	OnlineEvaluationConfigName: jsii.String("filtered_evaluation"),
//   	Evaluators: []EvaluatorSelector{
//   		agentcore.EvaluatorSelector_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS()),
//   	},
//   	DataSource: agentcore.DataSourceConfig_FromCloudWatchLogs(&CloudWatchLogsDataSourceConfig{
//   		LogGroupNames: []*string{
//   			jsii.String("/aws/bedrock-agentcore/my-agent"),
//   		},
//   		ServiceNames: []*string{
//   			jsii.String("my-agent.default"),
//   		},
//   	}),
//   	// Sample 25% of traces
//   	SamplingPercentage: jsii.Number(25),
//   	// Only evaluate traces matching these filters
//   	Filters: []FilterConfig{
//   		&FilterConfig{
//   			Key: jsii.String("user.region"),
//   			Operator: agentcore.FilterOperator_EQUAL(),
//   			Value: agentcore.FilterValue_String(jsii.String("us-east-1")),
//   		},
//   		&FilterConfig{
//   			Key: jsii.String("session.duration"),
//   			Operator: agentcore.FilterOperator_GREATER_THAN(),
//   			Value: agentcore.FilterValue_Number(jsii.Number(60)),
//   		},
//   	},
//   	// Consider sessions complete after 30 minutes of inactivity
//   	SessionTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   })
//
type FilterValue interface {
}

// The jsii proxy struct for FilterValue
type jsiiProxy_FilterValue struct {
	_ byte // padding
}

// Creates a boolean filter value.
func FilterValue_Boolean(value *bool) FilterValue {
	_init_.Initialize()

	if err := validateFilterValue_BooleanParameters(value); err != nil {
		panic(err)
	}
	var returns FilterValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.FilterValue",
		"boolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Creates a numeric filter value.
func FilterValue_Number(value *float64) FilterValue {
	_init_.Initialize()

	if err := validateFilterValue_NumberParameters(value); err != nil {
		panic(err)
	}
	var returns FilterValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.FilterValue",
		"number",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Creates a string filter value.
func FilterValue_String(value *string) FilterValue {
	_init_.Initialize()

	if err := validateFilterValue_StringParameters(value); err != nil {
		panic(err)
	}
	var returns FilterValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.FilterValue",
		"string",
		[]interface{}{value},
		&returns,
	)

	return returns
}

