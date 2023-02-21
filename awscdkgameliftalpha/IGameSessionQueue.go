// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/internal"
)

// Represents a Gamelift GameSessionQueue for a Gamelift fleet destination.
// Experimental.
type IGameSessionQueue interface {
	awscdk.IResource
	// Return the given named metric for this fleet.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Average amount of time that game session placement requests in the queue with status PENDING have been waiting to be fulfilled.
	// Experimental.
	MetricAverageWaitTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Game session placement requests that were canceled before timing out since the last report.
	// Experimental.
	MetricPlacementsCanceled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Game session placement requests that failed for any reason since the last report.
	// Experimental.
	MetricPlacementsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// New game session placement requests that were added to the queue since the last report.
	// Experimental.
	MetricPlacementsStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Game session placement requests that resulted in a new game session since the last report.
	// Experimental.
	MetricPlacementsSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Game session placement requests that reached the queue's timeout limit without being fulfilled since the last report.
	// Experimental.
	MetricPlacementsTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the gameSessionQueue.
	// Experimental.
	GameSessionQueueArn() *string
	// The Name of the gameSessionQueue.
	// Experimental.
	GameSessionQueueName() *string
}

// The jsii proxy for IGameSessionQueue
type jsiiProxy_IGameSessionQueue struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGameSessionQueue) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGameSessionQueue) MetricAverageWaitTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricAverageWaitTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricAverageWaitTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGameSessionQueue) MetricPlacementsCanceled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPlacementsCanceledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPlacementsCanceled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGameSessionQueue) MetricPlacementsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPlacementsFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPlacementsFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGameSessionQueue) MetricPlacementsStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPlacementsStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPlacementsStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGameSessionQueue) MetricPlacementsSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPlacementsSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPlacementsSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGameSessionQueue) MetricPlacementsTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPlacementsTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPlacementsTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IGameSessionQueue) GameSessionQueueArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameSessionQueueArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGameSessionQueue) GameSessionQueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameSessionQueueName",
		&returns,
	)
	return returns
}

