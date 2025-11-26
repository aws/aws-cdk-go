package previewawsapplicationsignalsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapplicationsignals/previewawsapplicationsignalsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates a service level objective (SLO), which can help you ensure that your critical business operations are meeting customer expectations.
//
// Use SLOs to set and track specific target levels for the reliability and availability of your applications and services. SLOs use service level indicators (SLIs) to calculate whether the application is performing at the level that you want.
//
// Create an SLO to set a target for a service operation, or service dependency's availability or latency. CloudWatch measures this target frequently you can find whether it has been breached.
//
// The target performance quality that is defined for an SLO is the *attainment goal* . An attainment goal is the percentage of time or requests that the SLI is expected to meet the threshold over each time interval. For example, an attainment goal of 99.9% means that within your interval, you are targeting 99.9% of the periods to be in healthy state.
//
// When you create an SLO, you specify whether it is a *period-based SLO* or a *request-based SLO* . Each type of SLO has a different way of evaluating your application's performance against its attainment goal.
//
// - A *period-based SLO* uses defined *periods* of time within a specified total time interval. For each period of time, Application Signals determines whether the application met its goal. The attainment rate is calculated as the `number of good periods/number of total periods` .
//
// For example, for a period-based SLO, meeting an attainment goal of 99.9% means that within your interval, your application must meet its performance goal during at least 99.9% of the time periods.
// - A *request-based SLO* doesn't use pre-defined periods of time. Instead, the SLO measures `number of good requests/number of total requests` during the interval. At any time, you can find the ratio of good requests to total requests for the interval up to the time stamp that you specify, and measure that ratio against the goal set in your SLO.
//
// After you have created an SLO, you can retrieve error budget reports for it. An *error budget* is the amount of time or amount of requests that your application can be non-compliant with the SLO's goal, and still have your application meet the goal.
//
// - For a period-based SLO, the error budget starts at a number defined by the highest number of periods that can fail to meet the threshold, while still meeting the overall goal. The *remaining error budget* decreases with every failed period that is recorded. The error budget within one interval can never increase.
//
// For example, an SLO with a threshold that 99.95% of requests must be completed under 2000ms every month translates to an error budget of 21.9 minutes of downtime per month.
// - For a request-based SLO, the remaining error budget is dynamic and can increase or decrease, depending on the ratio of good requests to total requests.
//
// When you call this operation, Application Signals creates the *AWSServiceRoleForCloudWatchApplicationSignals* service-linked role, if it doesn't already exist in your account. This service- linked role has the following permissions:
//
// - `xray:GetServiceGraph`
// - `logs:StartQuery`
// - `logs:GetQueryResults`
// - `cloudwatch:GetMetricData`
// - `cloudwatch:ListMetrics`
// - `tag:GetResources`
// - `autoscaling:DescribeAutoScalingGroups`
//
// You can easily set SLO targets for your applications, and their dependencies, that are discovered by Application Signals, using critical metrics such as latency and availability. You can also set SLOs against any CloudWatch metric or math expression that produces a time series.
//
// > You can't create an SLO for a service operation that was discovered by Application Signals until after that operation has reported standard metrics to Application Signals.
//
// You cannot change from a period-based SLO to a request-based SLO, or change from a request-based SLO to a period-based SLO.
//
// For more information about SLOs, see [Service level objectives (SLOs)](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-ServiceLevelObjectives.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceLevelObjectivePropsMixin := awscdkmixinspreview.Mixins.NewCfnServiceLevelObjectivePropsMixin(&CfnServiceLevelObjectiveMixinProps{
//   	BurnRateConfigurations: []interface{}{
//   		&BurnRateConfigurationProperty{
//   			LookBackWindowMinutes: jsii.Number(123),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ExclusionWindows: []interface{}{
//   		&ExclusionWindowProperty{
//   			Reason: jsii.String("reason"),
//   			RecurrenceRule: &RecurrenceRuleProperty{
//   				Expression: jsii.String("expression"),
//   			},
//   			StartTime: jsii.String("startTime"),
//   			Window: &WindowProperty{
//   				Duration: jsii.Number(123),
//   				DurationUnit: jsii.String("durationUnit"),
//   			},
//   		},
//   	},
//   	Goal: &GoalProperty{
//   		AttainmentGoal: jsii.Number(123),
//   		Interval: &IntervalProperty{
//   			CalendarInterval: &CalendarIntervalProperty{
//   				Duration: jsii.Number(123),
//   				DurationUnit: jsii.String("durationUnit"),
//   				StartTime: jsii.Number(123),
//   			},
//   			RollingInterval: &RollingIntervalProperty{
//   				Duration: jsii.Number(123),
//   				DurationUnit: jsii.String("durationUnit"),
//   			},
//   		},
//   		WarningThreshold: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	RequestBasedSli: &RequestBasedSliProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		MetricThreshold: jsii.Number(123),
//   		RequestBasedSliMetric: &RequestBasedSliMetricProperty{
//   			DependencyConfig: &DependencyConfigProperty{
//   				DependencyKeyAttributes: map[string]*string{
//   					"dependencyKeyAttributesKey": jsii.String("dependencyKeyAttributes"),
//   				},
//   				DependencyOperationName: jsii.String("dependencyOperationName"),
//   			},
//   			KeyAttributes: map[string]*string{
//   				"keyAttributesKey": jsii.String("keyAttributes"),
//   			},
//   			MetricType: jsii.String("metricType"),
//   			MonitoredRequestCountMetric: &MonitoredRequestCountMetricProperty{
//   				BadCountMetric: []interface{}{
//   					&MetricDataQueryProperty{
//   						AccountId: jsii.String("accountId"),
//   						Expression: jsii.String("expression"),
//   						Id: jsii.String("id"),
//   						MetricStat: &MetricStatProperty{
//   							Metric: &MetricProperty{
//   								Dimensions: []interface{}{
//   									&DimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   							},
//   							Period: jsii.Number(123),
//   							Stat: jsii.String("stat"),
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   				GoodCountMetric: []interface{}{
//   					&MetricDataQueryProperty{
//   						AccountId: jsii.String("accountId"),
//   						Expression: jsii.String("expression"),
//   						Id: jsii.String("id"),
//   						MetricStat: &MetricStatProperty{
//   							Metric: &MetricProperty{
//   								Dimensions: []interface{}{
//   									&DimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   							},
//   							Period: jsii.Number(123),
//   							Stat: jsii.String("stat"),
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			OperationName: jsii.String("operationName"),
//   			TotalRequestCountMetric: []interface{}{
//   				&MetricDataQueryProperty{
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					Id: jsii.String("id"),
//   					MetricStat: &MetricStatProperty{
//   						Metric: &MetricProperty{
//   							Dimensions: []interface{}{
//   								&DimensionProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MetricName: jsii.String("metricName"),
//   							Namespace: jsii.String("namespace"),
//   						},
//   						Period: jsii.Number(123),
//   						Stat: jsii.String("stat"),
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	Sli: &SliProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		MetricThreshold: jsii.Number(123),
//   		SliMetric: &SliMetricProperty{
//   			DependencyConfig: &DependencyConfigProperty{
//   				DependencyKeyAttributes: map[string]*string{
//   					"dependencyKeyAttributesKey": jsii.String("dependencyKeyAttributes"),
//   				},
//   				DependencyOperationName: jsii.String("dependencyOperationName"),
//   			},
//   			KeyAttributes: map[string]*string{
//   				"keyAttributesKey": jsii.String("keyAttributes"),
//   			},
//   			MetricDataQueries: []interface{}{
//   				&MetricDataQueryProperty{
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					Id: jsii.String("id"),
//   					MetricStat: &MetricStatProperty{
//   						Metric: &MetricProperty{
//   							Dimensions: []interface{}{
//   								&DimensionProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MetricName: jsii.String("metricName"),
//   							Namespace: jsii.String("namespace"),
//   						},
//   						Period: jsii.Number(123),
//   						Stat: jsii.String("stat"),
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   			MetricType: jsii.String("metricType"),
//   			OperationName: jsii.String("operationName"),
//   			PeriodSeconds: jsii.Number(123),
//   			Statistic: jsii.String("statistic"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html
//
type CfnServiceLevelObjectivePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnServiceLevelObjectiveMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServiceLevelObjectivePropsMixin
type jsiiProxy_CfnServiceLevelObjectivePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnServiceLevelObjectivePropsMixin) Props() *CfnServiceLevelObjectiveMixinProps {
	var returns *CfnServiceLevelObjectiveMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjectivePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApplicationSignals::ServiceLevelObjective`.
func NewCfnServiceLevelObjectivePropsMixin(props *CfnServiceLevelObjectiveMixinProps, options *mixins.CfnPropertyMixinOptions) CfnServiceLevelObjectivePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServiceLevelObjectivePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServiceLevelObjectivePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApplicationSignals::ServiceLevelObjective`.
func NewCfnServiceLevelObjectivePropsMixin_Override(c CfnServiceLevelObjectivePropsMixin, props *CfnServiceLevelObjectiveMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnServiceLevelObjectivePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceLevelObjectivePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServiceLevelObjectivePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServiceLevelObjectivePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceLevelObjectivePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

