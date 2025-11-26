package previewawsapplicationsignalsmixins


// This structure specifies the information about the service and the performance metric that an SLO is to monitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sliProperty := &SliProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	MetricThreshold: jsii.Number(123),
//   	SliMetric: &SliMetricProperty{
//   		DependencyConfig: &DependencyConfigProperty{
//   			DependencyKeyAttributes: map[string]*string{
//   				"dependencyKeyAttributesKey": jsii.String("dependencyKeyAttributes"),
//   			},
//   			DependencyOperationName: jsii.String("dependencyOperationName"),
//   		},
//   		KeyAttributes: map[string]*string{
//   			"keyAttributesKey": jsii.String("keyAttributes"),
//   		},
//   		MetricDataQueries: []interface{}{
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
//   		MetricType: jsii.String("metricType"),
//   		OperationName: jsii.String("operationName"),
//   		PeriodSeconds: jsii.Number(123),
//   		Statistic: jsii.String("statistic"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-sli.html
//
type CfnServiceLevelObjectivePropsMixin_SliProperty struct {
	// The arithmetic operation to use when comparing the specified metric to the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-sli.html#cfn-applicationsignals-servicelevelobjective-sli-comparisonoperator
	//
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value that the SLI metric is compared to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-sli.html#cfn-applicationsignals-servicelevelobjective-sli-metricthreshold
	//
	MetricThreshold *float64 `field:"optional" json:"metricThreshold" yaml:"metricThreshold"`
	// Use this structure to specify the metric to be used for the SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-sli.html#cfn-applicationsignals-servicelevelobjective-sli-slimetric
	//
	SliMetric interface{} `field:"optional" json:"sliMetric" yaml:"sliMetric"`
}

