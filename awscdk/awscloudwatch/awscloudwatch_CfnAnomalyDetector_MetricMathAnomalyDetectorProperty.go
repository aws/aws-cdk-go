package awscloudwatch


// Indicates the CloudWatch math expression that provides the time series the anomaly detector uses as input.
//
// The designated math expression must return a single time series.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricMathAnomalyDetectorProperty := &metricMathAnomalyDetectorProperty{
//   	metricDataQueries: []interface{}{
//   		&metricDataQueryProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			accountId: jsii.String("accountId"),
//   			expression: jsii.String("expression"),
//   			label: jsii.String("label"),
//   			metricStat: &metricStatProperty{
//   				metric: &metricProperty{
//   					metricName: jsii.String("metricName"),
//   					namespace: jsii.String("namespace"),
//
//   					// the properties below are optional
//   					dimensions: []interface{}{
//   						&dimensionProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				period: jsii.Number(123),
//   				stat: jsii.String("stat"),
//
//   				// the properties below are optional
//   				unit: jsii.String("unit"),
//   			},
//   			period: jsii.Number(123),
//   			returnData: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnAnomalyDetector_MetricMathAnomalyDetectorProperty struct {
	// An array of metric data query structures that enables you to create an anomaly detector based on the result of a metric math expression.
	//
	// Each item in `MetricDataQueries` gets a metric or performs a math expression. One item in `MetricDataQueries` is the expression that provides the time series that the anomaly detector uses as input. Designate the expression by setting `ReturnData` to `True` for this object in the array. For all other expressions and metrics, set `ReturnData` to `False` . The designated expression must return a single time series.
	MetricDataQueries interface{} `field:"optional" json:"metricDataQueries" yaml:"metricDataQueries"`
}

