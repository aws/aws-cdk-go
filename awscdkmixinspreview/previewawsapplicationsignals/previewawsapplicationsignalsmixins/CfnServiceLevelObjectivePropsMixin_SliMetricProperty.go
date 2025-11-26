package previewawsapplicationsignalsmixins


// Use this structure to specify the metric to be used for the SLO.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sliMetricProperty := &SliMetricProperty{
//   	DependencyConfig: &DependencyConfigProperty{
//   		DependencyKeyAttributes: map[string]*string{
//   			"dependencyKeyAttributesKey": jsii.String("dependencyKeyAttributes"),
//   		},
//   		DependencyOperationName: jsii.String("dependencyOperationName"),
//   	},
//   	KeyAttributes: map[string]*string{
//   		"keyAttributesKey": jsii.String("keyAttributes"),
//   	},
//   	MetricDataQueries: []interface{}{
//   		&MetricDataQueryProperty{
//   			AccountId: jsii.String("accountId"),
//   			Expression: jsii.String("expression"),
//   			Id: jsii.String("id"),
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
type CfnServiceLevelObjectivePropsMixin_SliMetricProperty struct {
	// Identifies the dependency using the `DependencyKeyAttributes` and `DependencyOperationName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-dependencyconfig
	//
	DependencyConfig interface{} `field:"optional" json:"dependencyConfig" yaml:"dependencyConfig"`
	// If this SLO is related to a metric collected by Application Signals, you must use this field to specify which service the SLO metric is related to.
	//
	// To do so, you must specify at least the `Type` , `Name` , and `Environment` attributes.
	//
	// This is a string-to-string map. It can include the following fields.
	//
	// - `Type` designates the type of object this is.
	// - `ResourceType` specifies the type of the resource. This field is used only when the value of the `Type` field is `Resource` or `AWS::Resource` .
	// - `Name` specifies the name of the object. This is used only if the value of the `Type` field is `Service` , `RemoteService` , or `AWS::Service` .
	// - `Identifier` identifies the resource objects of this resource. This is used only if the value of the `Type` field is `Resource` or `AWS::Resource` .
	// - `Environment` specifies the location where this object is hosted, or what it belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-keyattributes
	//
	KeyAttributes interface{} `field:"optional" json:"keyAttributes" yaml:"keyAttributes"`
	// If this SLO monitors a CloudWatch metric or the result of a CloudWatch metric math expression, use this structure to specify that metric or expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-metricdataqueries
	//
	MetricDataQueries interface{} `field:"optional" json:"metricDataQueries" yaml:"metricDataQueries"`
	// If the SLO is to monitor either the `LATENCY` or `AVAILABILITY` metric that Application Signals collects, use this field to specify which of those metrics is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-metrictype
	//
	MetricType *string `field:"optional" json:"metricType" yaml:"metricType"`
	// If the SLO is to monitor a specific operation of the service, use this field to specify the name of that operation.
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
	// It can be any CloudWatch statistic or extended statistic. For more information about statistics, see [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-slimetric.html#cfn-applicationsignals-servicelevelobjective-slimetric-statistic
	//
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
}

