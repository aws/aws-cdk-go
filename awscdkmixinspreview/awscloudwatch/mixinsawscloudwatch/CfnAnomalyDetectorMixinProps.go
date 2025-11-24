package mixinsawscloudwatch


// Properties for CfnAnomalyDetectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnomalyDetectorMixinProps := &CfnAnomalyDetectorMixinProps{
//   	Configuration: &ConfigurationProperty{
//   		ExcludedTimeRanges: []interface{}{
//   			&RangeProperty{
//   				EndTime: jsii.String("endTime"),
//   				StartTime: jsii.String("startTime"),
//   			},
//   		},
//   		MetricTimeZone: jsii.String("metricTimeZone"),
//   	},
//   	Dimensions: []interface{}{
//   		&DimensionProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MetricCharacteristics: &MetricCharacteristicsProperty{
//   		PeriodicSpikes: jsii.Boolean(false),
//   	},
//   	MetricMathAnomalyDetector: &MetricMathAnomalyDetectorProperty{
//   		MetricDataQueries: []interface{}{
//   			&MetricDataQueryProperty{
//   				AccountId: jsii.String("accountId"),
//   				Expression: jsii.String("expression"),
//   				Id: jsii.String("id"),
//   				Label: jsii.String("label"),
//   				MetricStat: &MetricStatProperty{
//   					Metric: &MetricProperty{
//   						Dimensions: []interface{}{
//   							&DimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//   					},
//   					Period: jsii.Number(123),
//   					Stat: jsii.String("stat"),
//   					Unit: jsii.String("unit"),
//   				},
//   				Period: jsii.Number(123),
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Namespace: jsii.String("namespace"),
//   	SingleMetricAnomalyDetector: &SingleMetricAnomalyDetectorProperty{
//   		AccountId: jsii.String("accountId"),
//   		Dimensions: []interface{}{
//   			&DimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   		Stat: jsii.String("stat"),
//   	},
//   	Stat: jsii.String("stat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html
//
type CfnAnomalyDetectorMixinProps struct {
	// Specifies details about how the anomaly detection model is to be trained, including time ranges to exclude when training and updating the model.
	//
	// The configuration can also include the time zone to use for the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html#cfn-cloudwatch-anomalydetector-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The dimensions of the metric associated with the anomaly detection band.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html#cfn-cloudwatch-anomalydetector-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Use this object to include parameters to provide information about your metric to CloudWatch to help it build more accurate anomaly detection models.
	//
	// Currently, it includes the `PeriodicSpikes` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html#cfn-cloudwatch-anomalydetector-metriccharacteristics
	//
	MetricCharacteristics interface{} `field:"optional" json:"metricCharacteristics" yaml:"metricCharacteristics"`
	// The CloudWatch metric math expression for this anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html#cfn-cloudwatch-anomalydetector-metricmathanomalydetector
	//
	MetricMathAnomalyDetector interface{} `field:"optional" json:"metricMathAnomalyDetector" yaml:"metricMathAnomalyDetector"`
	// The name of the metric associated with the anomaly detection band.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html#cfn-cloudwatch-anomalydetector-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric associated with the anomaly detection band.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html#cfn-cloudwatch-anomalydetector-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The CloudWatch metric and statistic for this anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html#cfn-cloudwatch-anomalydetector-singlemetricanomalydetector
	//
	SingleMetricAnomalyDetector interface{} `field:"optional" json:"singleMetricAnomalyDetector" yaml:"singleMetricAnomalyDetector"`
	// The statistic of the metric associated with the anomaly detection band.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html#cfn-cloudwatch-anomalydetector-stat
	//
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
}

