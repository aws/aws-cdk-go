package awsapplicationsignals


// This structure specifies the information about the service and the performance metric that an SLO is to monitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sliProperty := &SliProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	MetricThreshold: jsii.Number(123),
//   	SliMetric: &SliMetricProperty{
//   		KeyAttributes: map[string]*string{
//   			"keyAttributesKey": jsii.String("keyAttributes"),
//   		},
//   		MetricDataQueries: []interface{}{
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
//   		MetricType: jsii.String("metricType"),
//   		OperationName: jsii.String("operationName"),
//   		PeriodSeconds: jsii.Number(123),
//   		Statistic: jsii.String("statistic"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-sli.html
//
type CfnServiceLevelObjective_SliProperty struct {
	// The arithmetic operation to use when comparing the specified metric to the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-sli.html#cfn-applicationsignals-servicelevelobjective-sli-comparisonoperator
	//
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value that the SLI metric is compared to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-sli.html#cfn-applicationsignals-servicelevelobjective-sli-metricthreshold
	//
	MetricThreshold *float64 `field:"required" json:"metricThreshold" yaml:"metricThreshold"`
	// Use this structure to specify the metric to be used for the SLO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-sli.html#cfn-applicationsignals-servicelevelobjective-sli-slimetric
	//
	SliMetric interface{} `field:"required" json:"sliMetric" yaml:"sliMetric"`
}

