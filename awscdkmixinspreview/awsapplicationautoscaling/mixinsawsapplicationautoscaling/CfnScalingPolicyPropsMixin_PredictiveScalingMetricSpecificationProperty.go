package mixinsawsapplicationautoscaling


// This structure specifies the metrics and target utilization settings for a predictive scaling policy.
//
// You must specify either a metric pair, or a load metric and a scaling metric individually. Specifying a metric pair instead of individual metrics provides a simpler way to configure metrics for a scaling policy. You choose the metric pair, and the policy automatically knows the correct sum and average statistics to use for the load metric and the scaling metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   predictiveScalingMetricSpecificationProperty := &PredictiveScalingMetricSpecificationProperty{
//   	CustomizedCapacityMetricSpecification: &PredictiveScalingCustomizedCapacityMetricProperty{
//   		MetricDataQueries: []interface{}{
//   			&PredictiveScalingMetricDataQueryProperty{
//   				Expression: jsii.String("expression"),
//   				Id: jsii.String("id"),
//   				Label: jsii.String("label"),
//   				MetricStat: &PredictiveScalingMetricStatProperty{
//   					Metric: &PredictiveScalingMetricProperty{
//   						Dimensions: []interface{}{
//   							&PredictiveScalingMetricDimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//   					},
//   					Stat: jsii.String("stat"),
//   					Unit: jsii.String("unit"),
//   				},
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	CustomizedLoadMetricSpecification: &PredictiveScalingCustomizedLoadMetricProperty{
//   		MetricDataQueries: []interface{}{
//   			&PredictiveScalingMetricDataQueryProperty{
//   				Expression: jsii.String("expression"),
//   				Id: jsii.String("id"),
//   				Label: jsii.String("label"),
//   				MetricStat: &PredictiveScalingMetricStatProperty{
//   					Metric: &PredictiveScalingMetricProperty{
//   						Dimensions: []interface{}{
//   							&PredictiveScalingMetricDimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//   					},
//   					Stat: jsii.String("stat"),
//   					Unit: jsii.String("unit"),
//   				},
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	CustomizedScalingMetricSpecification: &PredictiveScalingCustomizedScalingMetricProperty{
//   		MetricDataQueries: []interface{}{
//   			&PredictiveScalingMetricDataQueryProperty{
//   				Expression: jsii.String("expression"),
//   				Id: jsii.String("id"),
//   				Label: jsii.String("label"),
//   				MetricStat: &PredictiveScalingMetricStatProperty{
//   					Metric: &PredictiveScalingMetricProperty{
//   						Dimensions: []interface{}{
//   							&PredictiveScalingMetricDimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//   					},
//   					Stat: jsii.String("stat"),
//   					Unit: jsii.String("unit"),
//   				},
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	PredefinedLoadMetricSpecification: &PredictiveScalingPredefinedLoadMetricProperty{
//   		PredefinedMetricType: jsii.String("predefinedMetricType"),
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	PredefinedMetricPairSpecification: &PredictiveScalingPredefinedMetricPairProperty{
//   		PredefinedMetricType: jsii.String("predefinedMetricType"),
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	PredefinedScalingMetricSpecification: &PredictiveScalingPredefinedScalingMetricProperty{
//   		PredefinedMetricType: jsii.String("predefinedMetricType"),
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	TargetValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html
//
type CfnScalingPolicyPropsMixin_PredictiveScalingMetricSpecificationProperty struct {
	// The customized capacity metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-customizedcapacitymetricspecification
	//
	CustomizedCapacityMetricSpecification interface{} `field:"optional" json:"customizedCapacityMetricSpecification" yaml:"customizedCapacityMetricSpecification"`
	// The customized load metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-customizedloadmetricspecification
	//
	CustomizedLoadMetricSpecification interface{} `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// The customized scaling metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-customizedscalingmetricspecification
	//
	CustomizedScalingMetricSpecification interface{} `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// The predefined load metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-predefinedloadmetricspecification
	//
	PredefinedLoadMetricSpecification interface{} `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// The predefined metric pair specification that determines the appropriate scaling metric and load metric to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-predefinedmetricpairspecification
	//
	PredefinedMetricPairSpecification interface{} `field:"optional" json:"predefinedMetricPairSpecification" yaml:"predefinedMetricPairSpecification"`
	// The predefined scaling metric specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-predefinedscalingmetricspecification
	//
	PredefinedScalingMetricSpecification interface{} `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
	// Specifies the target utilization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-targetvalue
	//
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

