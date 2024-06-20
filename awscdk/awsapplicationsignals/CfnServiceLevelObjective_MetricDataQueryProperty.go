package awsapplicationsignals


// Use this structure to define a metric or metric math expression that you want to use as for a service level objective.
//
// Each `MetricDataQuery` in the `MetricDataQueries` array specifies either a metric to retrieve, or a metric math expression to be performed on retrieved metrics. A single `MetricDataQueries` array can include as many as 20 `MetricDataQuery` structures in the array. The 20 structures can include as many as 10 structures that contain a `MetricStat` parameter to retrieve a metric, and as many as 10 structures that contain the `Expression` parameter to perform a math expression. Of those Expression structures, exactly one must have true as the value for `ReturnData`. The result of this expression used for the SLO.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDataQueryProperty := &MetricDataQueryProperty{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	AccountId: jsii.String("accountId"),
//   	Expression: jsii.String("expression"),
//   	MetricStat: &MetricStatProperty{
//   		Metric: &MetricProperty{
//   			Dimensions: []interface{}{
//   				&DimensionProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MetricName: jsii.String("metricName"),
//   			Namespace: jsii.String("namespace"),
//   		},
//   		Period: jsii.Number(123),
//   		Stat: jsii.String("stat"),
//
//   		// the properties below are optional
//   		Unit: jsii.String("unit"),
//   	},
//   	ReturnData: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html
//
type CfnServiceLevelObjective_MetricDataQueryProperty struct {
	// A short name used to tie this object to the results in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html#cfn-applicationsignals-servicelevelobjective-metricdataquery-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of the account where the metrics are located, if this is a cross-account alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html#cfn-applicationsignals-servicelevelobjective-metricdataquery-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The math expression to be performed on the returned data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html#cfn-applicationsignals-servicelevelobjective-metricdataquery-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// A metric to be used directly for the SLO, or to be used in the math expression that will be used for the SLO.
	//
	// Within one MetricDataQuery object, you must specify either Expression or MetricStat but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html#cfn-applicationsignals-servicelevelobjective-metricdataquery-metricstat
	//
	MetricStat interface{} `field:"optional" json:"metricStat" yaml:"metricStat"`
	// This option indicates whether to return the timestamps and raw data values of this metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html#cfn-applicationsignals-servicelevelobjective-metricdataquery-returndata
	//
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

