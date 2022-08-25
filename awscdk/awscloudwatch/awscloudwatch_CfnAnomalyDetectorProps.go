package awscloudwatch


// Properties for defining a `CfnAnomalyDetector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalyDetectorProps := &cfnAnomalyDetectorProps{
//   	configuration: &configurationProperty{
//   		excludedTimeRanges: []interface{}{
//   			&rangeProperty{
//   				endTime: jsii.String("endTime"),
//   				startTime: jsii.String("startTime"),
//   			},
//   		},
//   		metricTimeZone: jsii.String("metricTimeZone"),
//   	},
//   	dimensions: []interface{}{
//   		&dimensionProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	metricMathAnomalyDetector: &metricMathAnomalyDetectorProperty{
//   		metricDataQueries: []interface{}{
//   			&metricDataQueryProperty{
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				accountId: jsii.String("accountId"),
//   				expression: jsii.String("expression"),
//   				label: jsii.String("label"),
//   				metricStat: &metricStatProperty{
//   					metric: &metricProperty{
//   						metricName: jsii.String("metricName"),
//   						namespace: jsii.String("namespace"),
//
//   						// the properties below are optional
//   						dimensions: []interface{}{
//   							&dimensionProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					period: jsii.Number(123),
//   					stat: jsii.String("stat"),
//
//   					// the properties below are optional
//   					unit: jsii.String("unit"),
//   				},
//   				period: jsii.Number(123),
//   				returnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	metricName: jsii.String("metricName"),
//   	namespace: jsii.String("namespace"),
//   	singleMetricAnomalyDetector: &singleMetricAnomalyDetectorProperty{
//   		dimensions: []interface{}{
//   			&dimensionProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		metricName: jsii.String("metricName"),
//   		namespace: jsii.String("namespace"),
//   		stat: jsii.String("stat"),
//   	},
//   	stat: jsii.String("stat"),
//   }
//
type CfnAnomalyDetectorProps struct {
	// Specifies details about how the anomaly detection model is to be trained, including time ranges to exclude when training and updating the model.
	//
	// The configuration can also include the time zone to use for the metric.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The dimensions of the metric associated with the anomaly detection band.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The CloudWatch metric math expression for this anomaly detector.
	MetricMathAnomalyDetector interface{} `field:"optional" json:"metricMathAnomalyDetector" yaml:"metricMathAnomalyDetector"`
	// The name of the metric associated with the anomaly detection band.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric associated with the anomaly detection band.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The CloudWatch metric and statistic for this anomaly detector.
	SingleMetricAnomalyDetector interface{} `field:"optional" json:"singleMetricAnomalyDetector" yaml:"singleMetricAnomalyDetector"`
	// The statistic of the metric associated with the anomaly detection band.
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
}

