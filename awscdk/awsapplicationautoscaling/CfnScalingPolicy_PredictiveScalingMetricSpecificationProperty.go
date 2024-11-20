package awsapplicationautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveScalingMetricSpecificationProperty := &PredictiveScalingMetricSpecificationProperty{
//   	TargetValue: jsii.Number(123),
//
//   	// the properties below are optional
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
//
//   		// the properties below are optional
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	PredefinedMetricPairSpecification: &PredictiveScalingPredefinedMetricPairProperty{
//   		PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	PredefinedScalingMetricSpecification: &PredictiveScalingPredefinedScalingMetricProperty{
//   		PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html
//
type CfnScalingPolicy_PredictiveScalingMetricSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-targetvalue
	//
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-customizedcapacitymetricspecification
	//
	CustomizedCapacityMetricSpecification interface{} `field:"optional" json:"customizedCapacityMetricSpecification" yaml:"customizedCapacityMetricSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-customizedloadmetricspecification
	//
	CustomizedLoadMetricSpecification interface{} `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-customizedscalingmetricspecification
	//
	CustomizedScalingMetricSpecification interface{} `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-predefinedloadmetricspecification
	//
	PredefinedLoadMetricSpecification interface{} `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-predefinedmetricpairspecification
	//
	PredefinedMetricPairSpecification interface{} `field:"optional" json:"predefinedMetricPairSpecification" yaml:"predefinedMetricPairSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification-predefinedscalingmetricspecification
	//
	PredefinedScalingMetricSpecification interface{} `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
}

