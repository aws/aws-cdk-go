package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Represents a Stage.
type IStage interface {
	awscdk.IResource
	// Adds a stage variable to this stage.
	AddStageVariable(name *string, value *string)
	// Return the given named metric for this HTTP Api Gateway Stage.
	// Default: - average over 5 minutes.
	//
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The name of the stage;
	//
	// its primary identifier.
	StageName() *string
	// The URL to this stage.
	Url() *string
}

// The jsii proxy for IStage
type jsiiProxy_IStage struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IStage) AddStageVariable(name *string, value *string) {
	if err := i.validateAddStageVariableParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addStageVariable",
		[]interface{}{name, value},
	)
}

func (i *jsiiProxy_IStage) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (j *jsiiProxy_IStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

