package awsapplicationsignals


// A structure that contains information about the metric that the SLO monitors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sliMetricProperty := &SliMetricProperty{
//   	KeyAttributes: map[string]*string{
//   		"keyAttributesKey": jsii.String("keyAttributes"),
//   	},
//   	MetricDataQueries: []interface{}{
//   		&MetricDataQueryProperty{
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			AccountId: jsii.String("accountId"),
//   			Expression: jsii.String("expression"),
//   			MetricStat: &MetricStatProperty{
//   				Metric: &MetricProperty{
//   					Dimensions: []interface{}{
//   						&DimensionProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					MetricName: jsii.String("metricName"),
//   					Namespace: jsii.String("namespace"),
//   				},
//   				Period: jsii.Number(123),
//   				Stat: jsii.String("stat"),
//
//   				// the properties below are optional
//   				Unit: jsii.String("unit"),
//   			},
//   			ReturnData: jsii.Boolean(false),
//   		},
//   	},
//   	MetricType: jsii.String("metricType"),
//   	OperationName: jsii.String("operationName"),
//   	PeriodSeconds: jsii.Number(123),
//   	Statistic: jsii.String("statistic"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html
//
type CfnServiceLevelObjective_SliMetricProperty struct {
	// This is a string-to-string map that contains information about the type of object that this SLO is related to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-keyattributes
	//
	KeyAttributes interface{} `field:"optional" json:"keyAttributes" yaml:"keyAttributes"`
	// If this SLO monitors a CloudWatch metric or the result of a CloudWatch metric math expression, this structure includes the information about that metric or expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-metricdataqueries
	//
	MetricDataQueries interface{} `field:"optional" json:"metricDataQueries" yaml:"metricDataQueries"`
	// If the SLO monitors either the LATENCY or AVAILABILITY metric that Application Signals collects, this field displays which of those metrics is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-metrictype
	//
	MetricType *string `field:"optional" json:"metricType" yaml:"metricType"`
	// If the SLO monitors a specific operation of the service, this field displays that operation name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-operationname
	//
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// The number of seconds to use as the period for SLO evaluation.
	//
	// Your application's performance is compared to the SLI during each period. For each period, the application is determined to have either achieved or not achieved the necessary performance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-periodseconds
	//
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// The statistic to use for comparison to the threshold.
	//
	// It can be any CloudWatch statistic or extended statistic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-statistic
	//
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
}

