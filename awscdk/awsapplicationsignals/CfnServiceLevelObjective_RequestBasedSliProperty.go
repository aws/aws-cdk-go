package awsapplicationsignals


// This structure contains information about the performance metric that a request-based SLO monitors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestBasedSliProperty := &RequestBasedSliProperty{
//   	RequestBasedSliMetric: &RequestBasedSliMetricProperty{
//   		KeyAttributes: map[string]*string{
//   			"keyAttributesKey": jsii.String("keyAttributes"),
//   		},
//   		MetricType: jsii.String("metricType"),
//   		MonitoredRequestCountMetric: &MonitoredRequestCountMetricProperty{
//   			BadCountMetric: []interface{}{
//   				&MetricDataQueryProperty{
//   					Id: jsii.String("id"),
//
//   					// the properties below are optional
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					MetricStat: &MetricStatProperty{
//   						Metric: &MetricProperty{
//   							Dimensions: []interface{}{
//   								&DimensionProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MetricName: jsii.String("metricName"),
//   							Namespace: jsii.String("namespace"),
//   						},
//   						Period: jsii.Number(123),
//   						Stat: jsii.String("stat"),
//
//   						// the properties below are optional
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   			GoodCountMetric: []interface{}{
//   				&MetricDataQueryProperty{
//   					Id: jsii.String("id"),
//
//   					// the properties below are optional
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					MetricStat: &MetricStatProperty{
//   						Metric: &MetricProperty{
//   							Dimensions: []interface{}{
//   								&DimensionProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MetricName: jsii.String("metricName"),
//   							Namespace: jsii.String("namespace"),
//   						},
//   						Period: jsii.Number(123),
//   						Stat: jsii.String("stat"),
//
//   						// the properties below are optional
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		OperationName: jsii.String("operationName"),
//   		TotalRequestCountMetric: []interface{}{
//   			&MetricDataQueryProperty{
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				AccountId: jsii.String("accountId"),
//   				Expression: jsii.String("expression"),
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
//
//   					// the properties below are optional
//   					Unit: jsii.String("unit"),
//   				},
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	MetricThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-requestbasedsli.html
//
type CfnServiceLevelObjective_RequestBasedSliProperty struct {
	// This structure contains the information about the metric that is used for a request-based SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-requestbasedsli.html#cfn-applicationsignals-servicelevelobjective-requestbasedsli-requestbasedslimetric
	//
	RequestBasedSliMetric interface{} `field:"required" json:"requestBasedSliMetric" yaml:"requestBasedSliMetric"`
	// The arithmetic operation used when comparing the specified metric to the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-requestbasedsli.html#cfn-applicationsignals-servicelevelobjective-requestbasedsli-comparisonoperator
	//
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value that the SLI metric is compared to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-requestbasedsli.html#cfn-applicationsignals-servicelevelobjective-requestbasedsli-metricthreshold
	//
	MetricThreshold *float64 `field:"optional" json:"metricThreshold" yaml:"metricThreshold"`
}

