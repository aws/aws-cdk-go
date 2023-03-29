// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/internal"
)

// Represents a Gamelift matchmaking ruleset.
// Experimental.
type IMatchmakingRuleSet interface {
	awscdk.IResource
	// Return the given named metric for this matchmaking ruleSet.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Rule evaluations during matchmaking that failed since the last report.
	//
	// This metric is limited to the top 50 rules.
	// Experimental.
	MetricRuleEvaluationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Rule evaluations during the matchmaking process that passed since the last report.
	//
	// This metric is limited to the top 50 rules.
	// Experimental.
	MetricRuleEvaluationsPassed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the ruleSet.
	// Experimental.
	MatchmakingRuleSetArn() *string
	// The unique name of the ruleSet.
	// Experimental.
	MatchmakingRuleSetName() *string
}

// The jsii proxy for IMatchmakingRuleSet
type jsiiProxy_IMatchmakingRuleSet struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IMatchmakingRuleSet) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMatchmakingRuleSet) MetricRuleEvaluationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricRuleEvaluationsFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricRuleEvaluationsFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMatchmakingRuleSet) MetricRuleEvaluationsPassed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricRuleEvaluationsPassedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricRuleEvaluationsPassed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IMatchmakingRuleSet) MatchmakingRuleSetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchmakingRuleSetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMatchmakingRuleSet) MatchmakingRuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchmakingRuleSetName",
		&returns,
	)
	return returns
}

