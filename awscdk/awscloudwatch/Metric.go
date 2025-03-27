package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A metric emitted by a service.
//
// The metric is a combination of a metric identifier (namespace, name and dimensions)
// and an aggregation function (statistic, period and unit).
//
// It also contains metadata which is used only in graphs, such as color and label.
// It makes sense to embed this in here, so that compound constructs can attach
// that metadata to metrics they expose.
//
// This class does not represent a resource, so hence is not a construct. Instead,
// Metric is an abstraction that makes it easy to specify metrics for use in both
// alarms and graphs.
//
// Example:
//   var fn function
//
//
//   minuteErrorRate := fn.metricErrors(&MetricOptions{
//   	Statistic: cloudwatch.Stats_AVERAGE(),
//   	Period: awscdk.Duration_Minutes(jsii.Number(1)),
//   	Label: jsii.String("Lambda failure rate"),
//   })
//
type Metric interface {
	IMetric
	// Account which this metric comes from.
	Account() *string
	// The hex color code used when this metric is rendered on a graph.
	Color() *string
	// Dimensions of this metric.
	Dimensions() *map[string]interface{}
	// Label for this metric when added to a Graph in a Dashboard.
	Label() *string
	// Name of this metric.
	MetricName() *string
	// Namespace of this metric.
	Namespace() *string
	// Period of this metric.
	Period() awscdk.Duration
	// Region which this metric comes from.
	Region() *string
	// Statistic of this metric.
	Statistic() *string
	// Unit of the metric.
	Unit() Unit
	// Warnings attached to this metric.
	// Deprecated: - use warningsV2.
	Warnings() *[]*string
	// Warnings attached to this metric.
	WarningsV2() *map[string]*string
	// Attach the metric object to the given construct scope.
	//
	// Returns a Metric object that uses the account and region from the Stack
	// the given construct is defined in. If the metric is subsequently used
	// in a Dashboard or Alarm in a different Stack defined in a different
	// account or region, the appropriate 'region' and 'account' fields
	// will be added to it.
	//
	// If the scope we attach to is in an environment-agnostic stack,
	// nothing is done and the same Metric object is returned.
	AttachTo(scope constructs.IConstruct) Metric
	// Make a new Alarm for this metric.
	//
	// Combines both properties that may adjust the metric (aggregation) as well
	// as alarm properties.
	CreateAlarm(scope constructs.Construct, id *string, props *CreateAlarmOptions) Alarm
	// Inspect the details of the metric object.
	ToMetricConfig() *MetricConfig
	// Returns a string representation of an object.
	ToString() *string
	// Return a copy of Metric `with` properties changed.
	//
	// All properties except namespace and metricName can be changed.
	With(props *MetricOptions) Metric
}

// The jsii proxy struct for Metric
type jsiiProxy_Metric struct {
	jsiiProxy_IMetric
}

func (j *jsiiProxy_Metric) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Dimensions() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Period() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Unit() Unit {
	var returns Unit
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}


func NewMetric(props *MetricProps) Metric {
	_init_.Initialize()

	if err := validateNewMetricParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Metric{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.Metric",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewMetric_Override(m Metric, props *MetricProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.Metric",
		[]interface{}{props},
		m,
	)
}

// Grant permissions to the given identity to write metrics.
func Metric_GrantPutMetricData(grantee awsiam.IGrantable) awsiam.Grant {
	_init_.Initialize()

	if err := validateMetric_GrantPutMetricDataParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Metric",
		"grantPutMetricData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) AttachTo(scope constructs.IConstruct) Metric {
	if err := m.validateAttachToParameters(scope); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.Invoke(
		m,
		"attachTo",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) CreateAlarm(scope constructs.Construct, id *string, props *CreateAlarmOptions) Alarm {
	if err := m.validateCreateAlarmParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns Alarm

	_jsii_.Invoke(
		m,
		"createAlarm",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) ToMetricConfig() *MetricConfig {
	var returns *MetricConfig

	_jsii_.Invoke(
		m,
		"toMetricConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) With(props *MetricOptions) Metric {
	if err := m.validateWithParameters(props); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.Invoke(
		m,
		"with",
		[]interface{}{props},
		&returns,
	)

	return returns
}

