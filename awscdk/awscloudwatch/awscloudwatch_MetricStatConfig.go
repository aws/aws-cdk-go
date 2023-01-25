package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for a concrete metric.
//
// NOTE: `unit` is no longer on this object since it is only used for `Alarms`, and doesn't mean what one
// would expect it to mean there anyway. It is most likely to be misused.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var value interface{}
//
//   metricStatConfig := &metricStatConfig{
//   	metricName: jsii.String("metricName"),
//   	namespace: jsii.String("namespace"),
//   	period: duration,
//   	statistic: jsii.String("statistic"),
//
//   	// the properties below are optional
//   	account: jsii.String("account"),
//   	dimensions: []dimension{
//   		&dimension{
//   			name: jsii.String("name"),
//   			value: value,
//   		},
//   	},
//   	region: jsii.String("region"),
//   	unitFilter: awscdk.Aws_cloudwatch.unit_SECONDS,
//   }
//
// Experimental.
type MetricStatConfig struct {
	// Name of the metric.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Namespace of the metric.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// How many seconds to aggregate over.
	// Experimental.
	Period awscdk.Duration `field:"required" json:"period" yaml:"period"`
	// Aggregation function to use (can be either simple or a percentile).
	// Experimental.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
	// Account which this metric comes from.
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The dimensions to apply to the alarm.
	// Experimental.
	Dimensions *[]*Dimension `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Region which this metric comes from.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Unit used to filter the metric stream.
	//
	// Only refer to datums emitted to the metric stream with the given unit and
	// ignore all others. Only useful when datums are being emitted to the same
	// metric stream under different units.
	//
	// This field has been renamed from plain `unit` to clearly communicate
	// its purpose.
	// Experimental.
	UnitFilter Unit `field:"optional" json:"unitFilter" yaml:"unitFilter"`
}

