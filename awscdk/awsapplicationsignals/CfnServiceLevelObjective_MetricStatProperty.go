package awsapplicationsignals


// A metric to be used directly for the SLO, or to be used in the math expression that will be used for the SLO.
//
// Within one MetricDataQuery object, you must specify either Expression or MetricStat but not both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricStatProperty := &MetricStatProperty{
//   	Metric: &MetricProperty{
//   		Dimensions: []interface{}{
//   			&DimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   	},
//   	Period: jsii.Number(123),
//   	Stat: jsii.String("stat"),
//
//   	// the properties below are optional
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricstat.html
//
type CfnServiceLevelObjective_MetricStatProperty struct {
	// This structure defines the metric used for a service level indicator, including the metric name, namespace, and dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricstat.html#cfn-applicationsignals-servicelevelobjective-metricstat-metric
	//
	Metric interface{} `field:"required" json:"metric" yaml:"metric"`
	// The granularity, in seconds, to be used for the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricstat.html#cfn-applicationsignals-servicelevelobjective-metricstat-period
	//
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The statistic to use for comparison to the threshold.
	//
	// It can be any CloudWatch statistic or extended statistic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricstat.html#cfn-applicationsignals-servicelevelobjective-metricstat-stat
	//
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// If you omit Unit then all data that was collected with any unit is returned, along with the corresponding units that were specified when the data was reported to CloudWatch.
	//
	// If you specify a unit, the operation returns only data that was collected with that unit specified. If you specify a unit that does not match the data collected, the results of the operation are null. CloudWatch does not perform unit conversions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricstat.html#cfn-applicationsignals-servicelevelobjective-metricstat-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

