package awsapplicationsignals


// This structure defines the metric that is used as the "good request" or "bad request" value for a request-based SLO.
//
// This value observed for the metric defined in `TotalRequestCountMetric` is divided by the number found for `MonitoredRequestCountMetric` to determine the percentage of successful requests that this SLO tracks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoredRequestCountMetricProperty := &MonitoredRequestCountMetricProperty{
//   	BadCountMetric: []interface{}{
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
//   	GoodCountMetric: []interface{}{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-monitoredrequestcountmetric.html
//
type CfnServiceLevelObjective_MonitoredRequestCountMetricProperty struct {
	// If you want to count "bad requests" to determine the percentage of successful requests for this request-based SLO, specify the metric to use as "bad requests" in this structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-monitoredrequestcountmetric.html#cfn-applicationsignals-servicelevelobjective-monitoredrequestcountmetric-badcountmetric
	//
	BadCountMetric interface{} `field:"optional" json:"badCountMetric" yaml:"badCountMetric"`
	// If you want to count "good requests" to determine the percentage of successful requests for this request-based SLO, specify the metric to use as "good requests" in this structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-monitoredrequestcountmetric.html#cfn-applicationsignals-servicelevelobjective-monitoredrequestcountmetric-goodcountmetric
	//
	GoodCountMetric interface{} `field:"optional" json:"goodCountMetric" yaml:"goodCountMetric"`
}

