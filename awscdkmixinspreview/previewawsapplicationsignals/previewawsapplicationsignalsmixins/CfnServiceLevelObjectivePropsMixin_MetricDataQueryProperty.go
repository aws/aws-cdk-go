package previewawsapplicationsignalsmixins


// Use this structure to define a metric or metric math expression that you want to use as for a service level objective.
//
// Each `MetricDataQuery` in the `MetricDataQueries` array specifies either a metric to retrieve, or a metric math expression to be performed on retrieved metrics. A single `MetricDataQueries` array can include as many as 20 `MetricDataQuery` structures in the array. The 20 structures can include as many as 10 structures that contain a `MetricStat` parameter to retrieve a metric, and as many as 10 structures that contain the `Expression` parameter to perform a math expression. Of those `Expression` structures, exactly one must have true as the value for `ReturnData` . The result of this expression used for the SLO.
//
// For more information about metric math expressions, see [Use metric math](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html) .
//
// Within each `MetricDataQuery` object, you must specify either `Expression` or `MetricStat` but not both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricDataQueryProperty := &MetricDataQueryProperty{
//   	AccountId: jsii.String("accountId"),
//   	Expression: jsii.String("expression"),
//   	Id: jsii.String("id"),
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
//   		Unit: jsii.String("unit"),
//   	},
//   	ReturnData: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html
//
type CfnServiceLevelObjectivePropsMixin_MetricDataQueryProperty struct {
	// The ID of the account where this metric is located.
	//
	// If you are performing this operation in a monitoring account, use this to specify which source account to retrieve this metric from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html#cfn-applicationsignals-servicelevelobjective-metricdataquery-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// This field can contain a metric math expression to be performed on the other metrics that you are retrieving within this `MetricDataQueries` structure.
	//
	// A math expression can use the `Id` of the other metrics or queries to refer to those metrics, and can also use the `Id` of other expressions to use the result of those expressions. For more information about metric math expressions, see [Metric Math Syntax and Functions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax) in the *Amazon CloudWatch User Guide* .
	//
	// Within each `MetricDataQuery` object, you must specify either `Expression` or `MetricStat` but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html#cfn-applicationsignals-servicelevelobjective-metricdataquery-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// A short name used to tie this object to the results in the response.
	//
	// This `Id` must be unique within a `MetricDataQueries` array. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the metric math expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html#cfn-applicationsignals-servicelevelobjective-metricdataquery-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A metric to be used directly for the SLO, or to be used in the math expression that will be used for the SLO.
	//
	// Within one `MetricDataQuery` object, you must specify either `Expression` or `MetricStat` but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html#cfn-applicationsignals-servicelevelobjective-metricdataquery-metricstat
	//
	MetricStat interface{} `field:"optional" json:"metricStat" yaml:"metricStat"`
	// Use this only if you are using a metric math expression for the SLO.
	//
	// Specify `true` for `ReturnData` for only the one expression result to use as the alarm. For all other metrics and expressions in the same `CreateServiceLevelObjective` operation, specify `ReturnData` as `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricdataquery.html#cfn-applicationsignals-servicelevelobjective-metricdataquery-returndata
	//
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

