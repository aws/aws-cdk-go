package awsautoscaling


// Contains customized metric specification information for a target tracking scaling policy for Amazon EC2 Auto Scaling.
//
// To create your customized metric specification:
//
// - Add values for each required property from CloudWatch. You can use an existing metric, or a new metric that you create. To use your own metric, you must first publish the metric to CloudWatch. For more information, see [Publish Custom Metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html) in the *Amazon CloudWatch User Guide* .
// - Choose a metric that changes proportionally with capacity. The value of the metric should increase or decrease in inverse proportion to the number of capacity units. That is, the value of the metric should decrease when capacity increases.
//
// For more information about CloudWatch, see [Amazon CloudWatch Concepts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html) .
//
// `CustomizedMetricSpecification` is a property of the [AWS::AutoScaling::ScalingPolicy TargetTrackingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customizedMetricSpecificationProperty := &CustomizedMetricSpecificationProperty{
//   	Dimensions: []interface{}{
//   		&MetricDimensionProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Metrics: []interface{}{
//   		&TargetTrackingMetricDataQueryProperty{
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			Expression: jsii.String("expression"),
//   			Label: jsii.String("label"),
//   			MetricStat: &TargetTrackingMetricStatProperty{
//   				Metric: &MetricProperty{
//   					MetricName: jsii.String("metricName"),
//   					Namespace: jsii.String("namespace"),
//
//   					// the properties below are optional
//   					Dimensions: []interface{}{
//   						&MetricDimensionProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				Stat: jsii.String("stat"),
//
//   				// the properties below are optional
//   				Period: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   			Period: jsii.Number(123),
//   			ReturnData: jsii.Boolean(false),
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	Period: jsii.Number(123),
//   	Statistic: jsii.String("statistic"),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html
//
type CfnScalingPolicy_CustomizedMetricSpecificationProperty struct {
	// The dimensions of the metric.
	//
	// Conditional: If you published your metric with dimensions, you must specify the same dimensions in your scaling policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The name of the metric.
	//
	// To get the exact metric name, namespace, and dimensions, inspect the [Metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html) object that is returned by a call to [ListMetrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The metrics to include in the target tracking scaling policy, as a metric data query.
	//
	// This can include both raw metric and metric math expressions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-metrics
	//
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// The namespace of the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-period
	//
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The statistic of the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-statistic
	//
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The unit of the metric.
	//
	// For a complete list of the units that CloudWatch supports, see the [MetricDatum](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html) data type in the *Amazon CloudWatch API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

