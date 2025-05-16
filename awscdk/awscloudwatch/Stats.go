package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory functions for standard statistics strings.
//
// Example:
//   var dashboard dashboard
//   var executionCountMetric metric
//   var errorCountMetric metric
//
//
//   dashboard.AddWidgets(cloudwatch.NewGraphWidget(&GraphWidgetProps{
//   	Title: jsii.String("Executions vs error rate"),
//
//   	Left: []iMetric{
//   		executionCountMetric,
//   	},
//
//   	Right: []*iMetric{
//   		errorCountMetric.With(&MetricOptions{
//   			Statistic: cloudwatch.Stats_AVERAGE(),
//   			Label: jsii.String("Error rate"),
//   			Color: cloudwatch.Color_GREEN(),
//   		}),
//   	},
//   }))
//
type Stats interface {
}

// The jsii proxy struct for Stats
type jsiiProxy_Stats struct {
	_ byte // padding
}

func NewStats_Override(s Stats) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		nil, // no parameters
		s,
	)
}

// A shorter alias for `percentile()`.
func Stats_P(percentile *float64) *string {
	_init_.Initialize()

	if err := validateStats_PParameters(percentile); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"p",
		[]interface{}{percentile},
		&returns,
	)

	return returns
}

// Percentile indicates the relative standing of a value in a dataset.
//
// Percentiles help you get a better understanding of the distribution of your metric data.
//
// For example, `p(90)` is the 90th percentile and means that 90% of the data
// within the period is lower than this value and 10% of the data is higher
// than this value.
func Stats_Percentile(percentile *float64) *string {
	_init_.Initialize()

	if err := validateStats_PercentileParameters(percentile); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"percentile",
		[]interface{}{percentile},
		&returns,
	)

	return returns
}

// Percentile rank (PR) is the percentage of values that meet a fixed threshold.
//
// - If two numbers are given, they define the lower and upper bounds in absolute values,
//   respectively.
// - If one number is given, it defines the upper bound (the lower bound is assumed to
//   be 0).
//
// For example, `percentileRank(300)` returns the percentage of data points that have a value of 300 or less.
// `percentileRank(100, 2000)` returns the percentage of data points that have a value between 100 and 2000.
func Stats_PercentileRank(v1 *float64, v2 *float64) *string {
	_init_.Initialize()

	if err := validateStats_PercentileRankParameters(v1); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"percentileRank",
		[]interface{}{v1, v2},
		&returns,
	)

	return returns
}

// Shorter alias for `percentileRank()`.
func Stats_Pr(v1 *float64, v2 *float64) *string {
	_init_.Initialize()

	if err := validateStats_PrParameters(v1); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"pr",
		[]interface{}{v1, v2},
		&returns,
	)

	return returns
}

// Shorter alias for `trimmedCount()`.
func Stats_Tc(p1 *float64, p2 *float64) *string {
	_init_.Initialize()

	if err := validateStats_TcParameters(p1); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"tc",
		[]interface{}{p1, p2},
		&returns,
	)

	return returns
}

// A shorter alias for `trimmedMean()`.
func Stats_Tm(p1 *float64, p2 *float64) *string {
	_init_.Initialize()

	if err := validateStats_TmParameters(p1); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"tm",
		[]interface{}{p1, p2},
		&returns,
	)

	return returns
}

// Trimmed count (TC) is the number of data points in the chosen range for a trimmed mean statistic.
//
// - If two numbers are given, they define the lower and upper bounds in percentages,
//   respectively.
// - If one number is given, it defines the upper bound (the lower bound is assumed to
//   be 0).
//
// For example, `tc(90)` returns the number of data points not including any
// data points that fall in the highest 10% of the values. `tc(10, 90)`
// returns the number of data points not including any data points that fall
// in the lowest 10% of the values and the highest 90% of the values.
func Stats_TrimmedCount(p1 *float64, p2 *float64) *string {
	_init_.Initialize()

	if err := validateStats_TrimmedCountParameters(p1); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"trimmedCount",
		[]interface{}{p1, p2},
		&returns,
	)

	return returns
}

// Trimmed mean (TM) is the mean of all values that are between two specified boundaries.
//
// Values outside of the boundaries are ignored when the mean is calculated.
// You define the boundaries as one or two numbers between 0 and 100, up to 10
// decimal places. The numbers are percentages.
//
// - If two numbers are given, they define the lower and upper bounds in percentages,
//   respectively.
// - If one number is given, it defines the upper bound (the lower bound is assumed to
//   be 0).
//
// For example, `tm(90)` calculates the average after removing the 10% of data
// points with the highest values; `tm(10, 90)` calculates the average after removing the
// 10% with the lowest and 10% with the highest values.
func Stats_TrimmedMean(p1 *float64, p2 *float64) *string {
	_init_.Initialize()

	if err := validateStats_TrimmedMeanParameters(p1); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"trimmedMean",
		[]interface{}{p1, p2},
		&returns,
	)

	return returns
}

// Trimmed sum (TS) is the sum of the values of data points in a chosen range for a trimmed mean statistic.
//
// It is equivalent to `(Trimmed Mean) * (Trimmed count)`.
//
// - If two numbers are given, they define the lower and upper bounds in percentages,
//   respectively.
// - If one number is given, it defines the upper bound (the lower bound is assumed to
//   be 0).
//
// For example, `ts(90)` returns the sum of the data points not including any
// data points that fall in the highest 10% of the values.  `ts(10, 90)`
// returns the sum of the data points not including any data points that fall
// in the lowest 10% of the values and the highest 90% of the values.
func Stats_TrimmedSum(p1 *float64, p2 *float64) *string {
	_init_.Initialize()

	if err := validateStats_TrimmedSumParameters(p1); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"trimmedSum",
		[]interface{}{p1, p2},
		&returns,
	)

	return returns
}

// Shorter alias for `trimmedSum()`.
func Stats_Ts(p1 *float64, p2 *float64) *string {
	_init_.Initialize()

	if err := validateStats_TsParameters(p1); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"ts",
		[]interface{}{p1, p2},
		&returns,
	)

	return returns
}

// Winsorized mean (WM) is similar to trimmed mean.
//
// However, with winsorized mean, the values that are outside the boundary are
// not ignored, but instead are considered to be equal to the value at the
// edge of the appropriate boundary.  After this normalization, the average is
// calculated. You define the boundaries as one or two numbers between 0 and
// 100, up to 10 decimal places.
//
// - If two numbers are given, they define the lower and upper bounds in percentages,
//   respectively.
// - If one number is given, it defines the upper bound (the lower bound is assumed to
//   be 0).
//
// For example, `tm(90)` calculates the average after removing the 10% of data
// points with the highest values; `tm(10, 90)` calculates the average after removing the
// 10% with the lowest and 10% with the highest values.
//
// For example, `wm(90)` calculates the average while treating the 10% of the
// highest values to be equal to the value at the 90th percentile.
// `wm(10, 90)` calculates the average while treaing the bottom 10% and the
// top 10% of values to be equal to the boundary values.
func Stats_WinsorizedMean(p1 *float64, p2 *float64) *string {
	_init_.Initialize()

	if err := validateStats_WinsorizedMeanParameters(p1); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"winsorizedMean",
		[]interface{}{p1, p2},
		&returns,
	)

	return returns
}

// A shorter alias for `winsorizedMean()`.
func Stats_Wm(p1 *float64, p2 *float64) *string {
	_init_.Initialize()

	if err := validateStats_WmParameters(p1); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"wm",
		[]interface{}{p1, p2},
		&returns,
	)

	return returns
}

func Stats_AVERAGE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"AVERAGE",
		&returns,
	)
	return returns
}

func Stats_IQM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"IQM",
		&returns,
	)
	return returns
}

func Stats_MAXIMUM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"MAXIMUM",
		&returns,
	)
	return returns
}

func Stats_MINIMUM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"MINIMUM",
		&returns,
	)
	return returns
}

func Stats_SAMPLE_COUNT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"SAMPLE_COUNT",
		&returns,
	)
	return returns
}

func Stats_SUM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		"SUM",
		&returns,
	)
	return returns
}

