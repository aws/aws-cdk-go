package awsapplicationautoscaling


// This structure defines the CloudWatch metric to return, along with the statistic, period, and unit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingMetricStatProperty := &TargetTrackingMetricStatProperty{
//   	Metric: &TargetTrackingMetricProperty{
//   		Dimensions: []interface{}{
//   			&TargetTrackingMetricDimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   	},
//   	Stat: jsii.String("stat"),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricstat.html
//
type CfnScalingPolicy_TargetTrackingMetricStatProperty struct {
	// Represents a specific metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-targettrackingmetricstat-metric
	//
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// The statistic to return.
	//
	// It can include any CloudWatch statistic or extended statistic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-targettrackingmetricstat-stat
	//
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
	// The unit to use for the returned data points.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-targettrackingmetricstat-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

