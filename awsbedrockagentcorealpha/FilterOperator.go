package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Filter operators for online evaluation filtering.
//
// Example:
//   evaluation := agentcore.NewOnlineEvaluationConfig(this, jsii.String("FilteredEval"), &OnlineEvaluationConfigProps{
//   	OnlineEvaluationConfigName: jsii.String("filtered_evaluation"),
//   	Evaluators: []EvaluatorReference{
//   		agentcore.EvaluatorReference_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS()),
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
// Experimental.
type FilterOperator interface {
	// The string value of the filter operator.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for FilterOperator
type jsiiProxy_FilterOperator struct {
	_ byte // padding
}

func (j *jsiiProxy_FilterOperator) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewFilterOperator(value *string) FilterOperator {
	_init_.Initialize()

	if err := validateNewFilterOperatorParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_FilterOperator{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewFilterOperator_Override(f FilterOperator, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		[]interface{}{value},
		f,
	)
}

func FilterOperator_CONTAINS() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"CONTAINS",
		&returns,
	)
	return returns
}

func FilterOperator_EQUAL() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"EQUAL",
		&returns,
	)
	return returns
}

func FilterOperator_GREATER_THAN() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"GREATER_THAN",
		&returns,
	)
	return returns
}

func FilterOperator_GREATER_THAN_OR_EQUAL() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"GREATER_THAN_OR_EQUAL",
		&returns,
	)
	return returns
}

func FilterOperator_LESS_THAN() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"LESS_THAN",
		&returns,
	)
	return returns
}

func FilterOperator_LESS_THAN_OR_EQUAL() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"LESS_THAN_OR_EQUAL",
		&returns,
	)
	return returns
}

func FilterOperator_NOT_CONTAINS() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"NOT_CONTAINS",
		&returns,
	)
	return returns
}

func FilterOperator_NOT_EQUAL() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"NOT_EQUAL",
		&returns,
	)
	return returns
}

