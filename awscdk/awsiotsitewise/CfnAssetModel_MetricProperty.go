package awsiotsitewise


// Contains an asset metric property.
//
// With metrics, you can calculate aggregate functions, such as an average, maximum, or minimum, as specified through an expression. A metric maps several values to a single value (such as a sum).
//
// The maximum number of dependent/cascading variables used in any one metric calculation is 10. Therefore, a *root* metric can have up to 10 cascading metrics in its computational dependency tree. Additionally, a metric can only have a data type of `DOUBLE` and consume properties with data types of `INTEGER` or `DOUBLE` .
//
// For more information, see [Defining data properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#metrics) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricProperty := &MetricProperty{
//   	Expression: jsii.String("expression"),
//   	Variables: []interface{}{
//   		&ExpressionVariableProperty{
//   			Name: jsii.String("name"),
//   			Value: &VariableValueProperty{
//   				PropertyLogicalId: jsii.String("propertyLogicalId"),
//
//   				// the properties below are optional
//   				HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   			},
//   		},
//   	},
//   	Window: &MetricWindowProperty{
//   		Tumbling: &TumblingWindowProperty{
//   			Interval: jsii.String("interval"),
//
//   			// the properties below are optional
//   			Offset: jsii.String("offset"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-metric.html
//
type CfnAssetModel_MetricProperty struct {
	// The mathematical expression that defines the metric aggregation function.
	//
	// You can specify up to 10 variables per expression. You can specify up to 10 functions per expression.
	//
	// For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-metric.html#cfn-iotsitewise-assetmodel-metric-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The list of variables used in the expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-metric.html#cfn-iotsitewise-assetmodel-metric-variables
	//
	Variables interface{} `field:"required" json:"variables" yaml:"variables"`
	// The window (time interval) over which AWS IoT SiteWise computes the metric's aggregation expression.
	//
	// AWS IoT SiteWise computes one data point per `window` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-metric.html#cfn-iotsitewise-assetmodel-metric-window
	//
	Window interface{} `field:"required" json:"window" yaml:"window"`
}

