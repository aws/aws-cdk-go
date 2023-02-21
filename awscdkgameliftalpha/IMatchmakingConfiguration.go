// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/internal"
)

// Represents a Gamelift matchmaking configuration.
// Experimental.
type IMatchmakingConfiguration interface {
	awscdk.IResource
	// Return the given named metric for this matchmaking configuration.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Matchmaking requests currently being processed or waiting to be processed.
	// Experimental.
	MetricCurrentTickets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// For matchmaking configurations that require acceptance, the potential matches that were accepted since the last report.
	// Experimental.
	MetricMatchesAccepted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Potential matches that were created since the last report.
	// Experimental.
	MetricMatchesCreated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Matches that were successfully placed into a game session since the last report.
	// Experimental.
	MetricMatchesPlaced(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// For matchmaking configurations that require acceptance, the potential matches that were rejected by at least one player since the last report.
	// Experimental.
	MetricMatchesRejected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Players in matchmaking tickets that were added since the last report.
	// Experimental.
	MetricPlayersStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// For matchmaking requests that were put into a potential match before the last report, the amount of time between ticket creation and potential match creation.
	//
	// Units: seconds.
	// Experimental.
	MetricTimeToMatch(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the matchmaking configuration.
	// Experimental.
	MatchmakingConfigurationArn() *string
	// The name of the matchmaking configuration.
	// Experimental.
	MatchmakingConfigurationName() *string
	// The notification target for matchmaking events.
	// Experimental.
	NotificationTarget() awssns.ITopic
}

// The jsii proxy for IMatchmakingConfiguration
type jsiiProxy_IMatchmakingConfiguration struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IMatchmakingConfiguration) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IMatchmakingConfiguration) MetricCurrentTickets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCurrentTicketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCurrentTickets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMatchmakingConfiguration) MetricMatchesAccepted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMatchesAcceptedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMatchesAccepted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMatchmakingConfiguration) MetricMatchesCreated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMatchesCreatedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMatchesCreated",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMatchmakingConfiguration) MetricMatchesPlaced(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMatchesPlacedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMatchesPlaced",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMatchmakingConfiguration) MetricMatchesRejected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMatchesRejectedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMatchesRejected",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMatchmakingConfiguration) MetricPlayersStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPlayersStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPlayersStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMatchmakingConfiguration) MetricTimeToMatch(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTimeToMatchParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTimeToMatch",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IMatchmakingConfiguration) MatchmakingConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchmakingConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMatchmakingConfiguration) MatchmakingConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchmakingConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMatchmakingConfiguration) NotificationTarget() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"notificationTarget",
		&returns,
	)
	return returns
}

