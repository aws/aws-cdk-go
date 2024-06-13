package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingMetricDataQueryProperty := &TargetTrackingMetricDataQueryProperty{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	Expression: jsii.String("expression"),
//   	Label: jsii.String("label"),
//   	MetricStat: &TargetTrackingMetricStatProperty{
//   		Metric: &MetricProperty{
//   			MetricName: jsii.String("metricName"),
//   			Namespace: jsii.String("namespace"),
//
//   			// the properties below are optional
//   			Dimensions: []interface{}{
//   				&MetricDimensionProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		Stat: jsii.String("stat"),
//
//   		// the properties below are optional
//   		Unit: jsii.String("unit"),
//   	},
//   	ReturnData: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingmetricdataquery.html
//
type CfnScalingPolicy_TargetTrackingMetricDataQueryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingmetricdataquery.html#cfn-autoscaling-scalingpolicy-targettrackingmetricdataquery-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingmetricdataquery.html#cfn-autoscaling-scalingpolicy-targettrackingmetricdataquery-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingmetricdataquery.html#cfn-autoscaling-scalingpolicy-targettrackingmetricdataquery-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingmetricdataquery.html#cfn-autoscaling-scalingpolicy-targettrackingmetricdataquery-metricstat
	//
	MetricStat interface{} `field:"optional" json:"metricStat" yaml:"metricStat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingmetricdataquery.html#cfn-autoscaling-scalingpolicy-targettrackingmetricdataquery-returndata
	//
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

