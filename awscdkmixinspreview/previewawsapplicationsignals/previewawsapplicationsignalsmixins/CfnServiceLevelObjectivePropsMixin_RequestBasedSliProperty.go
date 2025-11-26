package previewawsapplicationsignalsmixins


// This structure contains information about the performance metric that a request-based SLO monitors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestBasedSliProperty := &RequestBasedSliProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	MetricThreshold: jsii.Number(123),
//   	RequestBasedSliMetric: &RequestBasedSliMetricProperty{
//   		DependencyConfig: &DependencyConfigProperty{
//   			DependencyKeyAttributes: map[string]*string{
//   				"dependencyKeyAttributesKey": jsii.String("dependencyKeyAttributes"),
//   			},
//   			DependencyOperationName: jsii.String("dependencyOperationName"),
//   		},
//   		KeyAttributes: map[string]*string{
//   			"keyAttributesKey": jsii.String("keyAttributes"),
//   		},
//   		MetricType: jsii.String("metricType"),
//   		MonitoredRequestCountMetric: &MonitoredRequestCountMetricProperty{
//   			BadCountMetric: []interface{}{
//   				&MetricDataQueryProperty{
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					Id: jsii.String("id"),
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
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   			GoodCountMetric: []interface{}{
//   				&MetricDataQueryProperty{
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					Id: jsii.String("id"),
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
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		OperationName: jsii.String("operationName"),
//   		TotalRequestCountMetric: []interface{}{
//   			&MetricDataQueryProperty{
//   				AccountId: jsii.String("accountId"),
//   				Expression: jsii.String("expression"),
//   				Id: jsii.String("id"),
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
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-requestbasedsli.html
//
type CfnServiceLevelObjectivePropsMixin_RequestBasedSliProperty struct {
	// The arithmetic operation used when comparing the specified metric to the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-requestbasedsli.html#cfn-applicationsignals-servicelevelobjective-requestbasedsli-comparisonoperator
	//
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// This value is the threshold that the observed metric values of the SLI metric are compared to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-requestbasedsli.html#cfn-applicationsignals-servicelevelobjective-requestbasedsli-metricthreshold
	//
	MetricThreshold *float64 `field:"optional" json:"metricThreshold" yaml:"metricThreshold"`
	// A structure that contains information about the metric that the SLO monitors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-requestbasedsli.html#cfn-applicationsignals-servicelevelobjective-requestbasedsli-requestbasedslimetric
	//
	RequestBasedSliMetric interface{} `field:"optional" json:"requestBasedSliMetric" yaml:"requestBasedSliMetric"`
}

