package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingMetricStatProperty := &TargetTrackingMetricStatProperty{
//   	Metric: &MetricProperty{
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//
//   		// the properties below are optional
//   		Dimensions: []interface{}{
//   			&MetricDimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Stat: jsii.String("stat"),
//
//   	// the properties below are optional
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingmetricstat.html
//
type CfnScalingPolicy_TargetTrackingMetricStatProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingmetricstat.html#cfn-autoscaling-scalingpolicy-targettrackingmetricstat-metric
	//
	Metric interface{} `field:"required" json:"metric" yaml:"metric"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingmetricstat.html#cfn-autoscaling-scalingpolicy-targettrackingmetricstat-stat
	//
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingmetricstat.html#cfn-autoscaling-scalingpolicy-targettrackingmetricstat-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

